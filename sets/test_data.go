package sets

var Knapsack01TinyTestData = []Subject{
	{Value: 1, Weight: 2},
	{Value: 2, Weight: 3},
	{Value: 5, Weight: 4},
	{Value: 6, Weight: 5},
}

func GetKnapsackDataCopy(data []Subject) []Subject {
	tmp := make([]Subject, len(data))

	copy(tmp, data)
	return tmp
}
