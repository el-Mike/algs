package datastructures

import (
	"fmt"

	"github.com/el-Mike/algs/data-structures/array"
)

const (
	LOG_SEPARATOR = "--------------------------------------------"
)

func Run() {
	array.RunOpeations()

	printSeparator()
}

func printSeparator() {
	fmt.Printf("%s\n", LOG_SEPARATOR)
}
