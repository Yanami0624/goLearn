package std

import "log"

func FunLog() {
	log.SetPrefix("LOG ")
	log.SetFlags(log.Lshortfile | log.Ltime)
	log.Println("start")
	log.Panicln("panic")

}
