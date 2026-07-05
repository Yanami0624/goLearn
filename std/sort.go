package std

import (
	"fmt"
	"sort"
	"strconv"
)

type Person struct {
	Id   int
	name string
}

func NewPerson(n int) []Person {
	const N int = 16
	persons := make([]Person, N)
	for i := range N {
		persons[i] = Person{
			Id:   N - i,
			name: strconv.Itoa(100 + i),
		}
	}
	return persons
}

type personslice []Person

func (a personslice) Len() int {
	return len(a)
}

func (a personslice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a personslice) Less(i, j int) bool {
	return a[i].Id < a[j].Id
}

func FunSort() {
	arr := NewPerson(16)
	fmt.Println(arr)
	sort.Sort(personslice(arr))
	fmt.Println(arr)

}
