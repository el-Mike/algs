package sets

type Subject struct {
	Value  int
	Weight int
}

type KnapsackResult struct {
	maxValue int
	subjects *[]Subject
}

var EMPTY_RESULT = &KnapsackResult{
	maxValue: 0,
	subjects: &[]Subject{},
}

// Knapsack01 - implementation of 0-1 Knapsack problem solution, based on
// Dynamic Programming approach.
func Knapsack01(baseSubjects []Subject, capacity int) *KnapsackResult {
	if capacity <= 0 {
		return EMPTY_RESULT
	}

	var subjects []Subject

	// There is no point to include Subjects that alone exceed the capacity,
	// therefore we can filter them out.
	for _, subject := range baseSubjects {
		if subject.Weight <= capacity {
			subjects = append(subjects, subject)
		}
	}

	subjectsLength := len(subjects)

	if subjectsLength == 0 {
		return EMPTY_RESULT
	}

	matrix := make([][]int, subjectsLength)

	// Fill matrix with nested slices.
	for i := range matrix {
		// We need to test the capacity inclusively [0..capacity], therefore we're
		// adding 1 to the number of columns.
		matrix[i] = make([]int, capacity+1)
	}

	for i := 0; i < subjectsLength; i += 1 {
		for j := 0; j <= capacity; j += 1 {
			// No Subject can fit in 0 capacity, therefore the solution will always be 0.
			if j == 0 {
				matrix[i][j] = 0

				continue
			}

			currentSubject := subjects[i]
			currentCapacity := j

			currentWeight := currentSubject.Weight
			currentValue := currentSubject.Value

			// For the first row (and therefore first Subject) we just want to check
			// if it can fit in given capacity. This is the first "level" of solved
			// subproblems.
			if i == 0 {
				if currentWeight <= currentCapacity {
					matrix[i][j] = currentValue
				} else {
					matrix[i][j] = 0
				}

				continue
			}

			// For any other row, check if it fits the limit - if yes, check if sum
			// of its value and related result for the capacity reduced by its weight
			// is bigger than solution for previous row.
			// If yes - save the new amount.
			if currentWeight <= currentCapacity {
				matrix[i][j] = max(matrix[i-1][currentCapacity],
					currentValue+matrix[i-1][currentCapacity-currentWeight],
				)
				// Otherwise, copy the value from previous row.
			} else {
				matrix[i][j] = matrix[i-1][currentCapacity]
			}
		}
	}

	maxValue := matrix[subjectsLength-1][capacity]

	result := []Subject{}
	currentValue := maxValue
	currentCapacity := capacity

	// In order to retrieve the Subjects that belong to the solution,
	// we need to backtrack the matrix and pick correct ones based on related values.
	for i := (subjectsLength - 1); i > 0 && currentValue > 0; i -= 1 {
		// If currentValue is different than value from previous row, that means
		// that given Subject should be added.
		if currentValue != matrix[i-1][currentCapacity] {
			currentSubject := subjects[i]

			result = append(result, currentSubject)

			// Subtract capacity's "index".
			currentCapacity -= currentSubject.Weight

			// As long as we are dealing with integers, we could simply subtract
			// currentSubject's value, but since subtraction can be inconsistent
			// when using floating point numbers, it's safer to rely on matrix
			// values.
			currentValue = matrix[i-1][currentCapacity]
		}
	}

	return &KnapsackResult{
		maxValue: maxValue,
		subjects: &result,
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
