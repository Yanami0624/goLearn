package std

import (
	"fmt"
	"strings"
)

func FunString() {}

func FunString_copy() {
	a := "hello"
	fmt.Println(strings.Clone(a))
	b := a
	b = b[:2]
	fmt.Println(a)

	c := []byte("hello")
	d := c
	d[0] = 'H'
	fmt.Println(string(c))
}

func FunString_compare() {
	fmt.Println(strings.Compare("abc", "abe"))
	fmt.Println(strings.Compare("abcd", "abe"))
	fmt.Println(strings.Compare("abijk", "abe"))
	fmt.Println(strings.Compare("abe", "abe"))
}

func FunString_contain() {
	fmt.Println(strings.Contains("abcdefg", "a"))
	fmt.Println(strings.Contains("abcdefg", "abc"))
	fmt.Println(strings.Contains("abcdefg", "ba"))
}

func FunString_fieldspilt() {
	fmt.Printf("%q\n", strings.Fields(" a b c d e f g "))
	fmt.Printf("%q\n", strings.FieldsFunc("a,b,c,d,e,f,g", func(r rune) bool {
		return r == ','
	}))
}
