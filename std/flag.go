package std

import (
	"flag"
	"fmt"
)

func FunFlag() {
	name := flag.String("name", "Alice", "name")
	age := flag.Int("age", 18, "age")
	sex := flag.String("sex", "male", "sex")

	flag.Parse()

	fmt.Printf("%s, age %d, %s\n", *name, *age, *sex)
}
