package main

import (
	"fmt"

	"github.com/nus/trygofuzz/lib"
)

func main() {
	trygofuzz.Checker("aa")
	fmt.Println("aaa")
}
