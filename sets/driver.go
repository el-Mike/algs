package sets

import "fmt"

const (
	LOG_SEPARATOR = "============================"
)

func Run() {
	result := Knapsack01(GetKnapsackDataCopy(Knapsack01TinyTestData), 8)
	printKnapsackResult(result)
	printSeparator()
}

func printSeparator() {
	fmt.Printf("%s\n", LOG_SEPARATOR)
}

func printKnapsackResult(result *KnapsackResult) {
	fmt.Printf("Subjects: \n")
	fmt.Printf("%v\n", result.subjects)
	fmt.Printf("Max value: \n")
	fmt.Printf("%d\n\n", result.maxValue)
}
