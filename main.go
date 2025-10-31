package main

import (
	"fmt"
	"os"

	"github.com/albertsko/zed-everforest/scripts"
)

var usage string = `Usage:
go run . script_name

Available scripts:
- generate
`

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%s \nError:\n%+v", usage, r)
		}
	}()

	arg1 := os.Args[1]
	switch arg1 {
	case "generate":
		scripts.Generate()
	default:
		panic("provided incorrect first arg\n")
	}
}
