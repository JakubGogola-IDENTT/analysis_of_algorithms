package main

import (
	"crypto/md5"
	"fmt"
	"lab2/mincount"
)

func main() {
	mc := mincount.New(md5.New, 400)
	v := 2137
	hash := mc.XD(v, 16)
	fmt.Println(hash)

	// analytics.TestAll()
}
