package main

import (
	. "Gocodes/algorithm"
	. "Gocodes/struct"
)

func main() {
	g := BuildGraph[MyInt](10, 0.1, true)
	g.Print()
	println(g.ExistRing())
}
