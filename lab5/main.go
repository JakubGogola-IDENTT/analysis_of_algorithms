package main

import (
	"lab5/mis"
)

func main() {
	g := mis.New(10)
	g.Simulate()
	g.PrintMIS()
}
