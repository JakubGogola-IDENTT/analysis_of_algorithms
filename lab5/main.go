package main

import (
	"fmt"
	"lab5/mis"
)

func main() {
	// me.Simulate(4)
	g := mis.New(10)
	g.Simulate()
	fmt.Println(g)
}
