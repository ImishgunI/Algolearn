package algorithms

import "Algolearn/internal/execution/state"

type BubbleSort struct{}

func (b *BubbleSort) Run(input []int) []state.Step {
	arr := make([]int, len(input))
	copy(arr, input)
	steps := []state.Step{}
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			steps = append(steps, state.Step{
				Array:    copyArray(arr),
				Active:   []int{j, j + 1},
				Swapping: []int{},
				Lines:    []int{3},
			})

			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				steps = append(steps, state.Step{
					Array:    copyArray(arr),
					Active:   []int{},
					Swapping: []int{j, j + 1},
					Lines:    []int{4},
				})
			}
		}
	}
	return steps
}

func copyArray(arr []int) []int {
	a := make([]int, len(arr))
	copy(a, arr)
	return a
}
