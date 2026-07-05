package std

import (
	"log"
	"strconv"
)

func FunStrconv() {
	str := "123"
	num := 123

	log.SetFlags(log.Lshortfile)

	log.Println(strconv.Atoi(str))
	log.Println(strconv.Itoa(num))

}
