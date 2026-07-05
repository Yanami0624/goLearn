package std

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func FunHttpExample() {
	address := flag.String("a", "https://baidu.com", "address")
	method := flag.String("m", "Get", "method")
	flag.Parse()

	switch {
	case *method == "Get":
		resp, _ := http.Get(*address)
		// defer resp.Body.Close()
		content, _ := io.ReadAll(resp.Body)
		fmt.Println(string(content[:10]))
	}
}

func FunHttp() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 100)
		client := &http.Client{}
		// request, _ := http.NewRequest("GET", "http://localhost:8080/hello", nil)
		// request.Header.Add("name", "alice")
		// resp, err := client.Do(request)
		resp, err := client.Get("http://127.0.0.1:8080/hello")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer resp.Body.Close()
		content, _ := io.ReadAll(resp.Body)
		fmt.Printf("client recieve: %s\n", string(content))
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "home page")
	})
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, %s\n", r.Host)
	})
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	go func() {
		defer wg.Done()

		server.ListenAndServe()
	}()

	time.Sleep(time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1000)
	defer cancel()
	server.Shutdown(ctx)

	wg.Wait()
}
