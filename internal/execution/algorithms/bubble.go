package algorithms

import "Algolearn/internal/execution/state"

type BubbleSort struct{}

func (b *BubbleSort) Run(input []int) []state.Step {
	arr := make([]int, len(input))
	copy(arr, input)

	var steps []state.Step

	n := len(arr)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {

			if arr[j] > arr[j+1] {

				arr[j], arr[j+1] = arr[j+1], arr[j]

				snapshot := make([]int, len(arr))
				copy(snapshot, arr)

				steps = append(steps, state.Step{
					Array: snapshot,
				})
			}
		}
	}

	return steps
}
