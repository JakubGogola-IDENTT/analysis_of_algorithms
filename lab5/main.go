package main

import (
	"lab5/mis"
)

func main() {
	g := mis.New(4)
	g.Simulate()
	g.PrintMIS()
}
