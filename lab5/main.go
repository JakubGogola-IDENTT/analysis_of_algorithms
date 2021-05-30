package main

import (
	"fmt"
	"lab5/me"
	"lab5/mis"
	"strings"
)

func printResult(itName string, itValue int, callback func()) {
	str := fmt.Sprintf("-------- %s = %d --------", itName, itValue)
	fmt.Println(str)
	callback()
	fmt.Println(strings.Repeat("-", len(str)))
}

func main() {
	// ex10
	fmt.Println("Ex10")
	for i := 1; i < 5; i++ {
		printResult("i", i, func() {
			me.Simulate(i)
		})
	}

	// ex11
	fmt.Println("Ex11")
	for i := 2; i < 15; i++ {
		printResult("i", i, func() {
			g := mis.New(i)
			g.Simulate()
			g.PrintMIS()
		})
	}
}
