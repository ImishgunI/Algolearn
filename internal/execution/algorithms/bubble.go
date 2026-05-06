package algorithms

import "Algolearn/internal/execution/state"

type BubbleSort struct{}

func (b *BubbleSort) Run(input []int) []state.Step {
	arr := make([]int, len(input))
	copy(arr, input)

	var steps []state.Step

	n := len(arr)

	for i := range n {
		for j := 0; j < n-i-1; j++ {

			if arr[j] > arr[j+1] {
				active := []int{j, j + 1}
				swap := []int{j, j + 1}
				arr[j], arr[j+1] = arr[j+1], arr[j]

				snapshot := make([]int, len(arr))
				copy(snapshot, arr)

				steps = append(steps, state.Step{
					Array:    snapshot,
					Active:   active,
					Swapping: swap,
				})
			}
		}
	}

	return steps
}
