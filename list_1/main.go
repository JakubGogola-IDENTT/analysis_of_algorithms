package main

import (
	"fmt"
	"list_1/election"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(election.WithUpperLimit(100, 100))
}
