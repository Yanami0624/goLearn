package std

import (
	"encoding/json"
	"fmt"
	"sync"
)

type StrJson struct {
	Id   int
	Name string
}

func FunJson() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan string, 1)
	go func() {
		defer wg.Done()
		str := StrJson{Id: 1, Name: "Alice"}
		msg, _ := json.Marshal(str)
		ch <- string(msg)
	}()
	go func() {
		defer wg.Done()
		msg := <-ch
		fmt.Println(msg)
		var str StrJson
		json.Unmarshal([]byte(msg), &str)
		fmt.Printf("%d, %s\n", str.Id, str.Name)
	}()
	wg.Wait()
}
