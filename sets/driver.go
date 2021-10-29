package sets

import "fmt"

const (
	LOG_SEPARATOR = "============================"
)

func Run() {

	printSeparator()
}

func printSeparator() {
	fmt.Printf("%s\n", LOG_SEPARATOR)
}
