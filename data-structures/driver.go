package datastructures

import (
	"fmt"

	"github.com/el-Mike/algs/data-structures/hashtable"
)

const (
	LOG_SEPARATOR = "--------------------------------------------"
)

func Run() {
	// array.RunOpeations()

	printSeparator()

	// list.RunLinkedList()

	printSeparator()

	hashtable.RunChainingHashtable()
}

func printSeparator() {
	fmt.Printf("%s\n", LOG_SEPARATOR)
}
