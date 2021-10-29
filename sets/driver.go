package sets

import "fmt"

const (
	LOG_SEPARATOR = "============================"
)

func Run() {
	Knapsack01(GetKnapsackDataCopy(Knapsack01TinyTestData), 8)
	printSeparator()
}

func printSeparator() {
	fmt.Printf("%s\n", LOG_SEPARATOR)
}
