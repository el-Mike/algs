package main

import (
	"fmt"

	"github.com/el-Mike/algs/data-structures/array"
)

const (
	LOG_SEPARATOR = "--------------------------------------------"
)

func main() {
	array.RunOpeations()

	printSeparator()
}

func printSeparator() {
	fmt.Printf("%s\n", LOG_SEPARATOR)
}
