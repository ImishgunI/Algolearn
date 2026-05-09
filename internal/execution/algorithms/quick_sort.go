package algorithms

import "Algolearn/internal/execution/state"

type QuickSort struct{}

func (q *QuickSort) Run(input []int) []state.Step {
	arr := make([]int, len(input))
	copy(arr, input)
	steps := []state.Step{}
	quickSortHelper(arr, 0, len(arr)-1, &steps)
	return steps
}

func quickSortHelper(arr []int, low, high int, steps *[]state.Step) {
	if low < high {
		pi := partition(arr, low, high, steps)
		quickSortHelper(arr, low, pi-1, steps)
		quickSortHelper(arr, pi+1, high, steps)
	}
}

func partition(arr []int, low, high int, steps *[]state.Step) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		// Сравнение arr[j] с pivot
		*steps = append(*steps, state.Step{
			Array: copyArray(arr),

			Active:   []int{j, high},
			Swapping: []int{},

			Pivot:     []int{high},
			Partition: []int{low, high},

			Lines: []int{9},
		})

		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
			*steps = append(*steps, state.Step{
				Array: copyArray(arr),

				Active:   []int{},
				Swapping: []int{i, j},

				Pivot:     []int{high},
				Partition: []int{low, high},

				Lines: []int{10, 11},
			})
		}
	}

	// Ставим pivot на место
	arr[i+1], arr[high] = arr[high], arr[i+1]
	*steps = append(*steps, state.Step{
		Array: copyArray(arr),

		Active:   []int{},
		Swapping: []int{i + 1, high},

		Pivot:     []int{i + 1},
		Partition: []int{low, high},

		Lines: []int{13},
	})
	return i + 1
}
