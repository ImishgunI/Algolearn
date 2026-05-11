package algorithms

import "Algolearn/internal/execution/state"

type MergeSort struct{}

func (m *MergeSort) Run(input []int) []state.Step {
	arr := make([]int, len(input))
	copy(arr, input)

	steps := make([]state.Step, 0)

	mergeSortHelper(arr, 0, len(arr)-1, &steps)

	return steps
}

func mergeSortHelper(
	arr []int,
	left, right int,
	steps *[]state.Step,
) {
	if left >= right {
		return
	}

	mid := left + (right-left)/2

	// -------- SPLIT STEP --------
	*steps = append(*steps, state.Step{
		Array: copyArray(arr),

		LeftRange:  []int{left, mid},
		RightRange: []int{mid + 1, right},

		Lines: []int{2, 3, 4},
	})

	mergeSortHelper(arr, left, mid, steps)
	mergeSortHelper(arr, mid+1, right, steps)

	merge(arr, left, mid, right, steps)
}

func merge(
	arr []int,
	left, mid, right int,
	steps *[]state.Step,
) {
	n1 := mid - left + 1
	n2 := right - mid

	L := make([]int, n1)
	R := make([]int, n2)

	for i := 0; i < n1; i++ {
		L[i] = arr[left+i]
	}

	for j := 0; j < n2; j++ {
		R[j] = arr[mid+1+j]
	}

	i := 0
	j := 0
	k := left

	// -------- MAIN MERGE LOOP --------
	for i < n1 && j < n2 {

		// COMPARISON STEP
		*steps = append(*steps, state.Step{
			Array: copyArray(arr),

			Active: []int{
				left + i,
				mid + 1 + j,
			},

			LeftRange:  []int{left, mid},
			RightRange: []int{mid + 1, right},
			MergedRange: []int{
				left,
				right,
			},

			Lines: []int{16},
		})

		if L[i] <= R[j] {

			arr[k] = L[i]

			// WRITE LEFT VALUE
			*steps = append(*steps, state.Step{
				Array: copyArray(arr),

				Writing: []int{k},

				LeftRange:  []int{left, mid},
				RightRange: []int{mid + 1, right},
				MergedRange: []int{
					left,
					right,
				},

				Lines: []int{18},
			})

			i++

		} else {

			arr[k] = R[j]

			// WRITE RIGHT VALUE
			*steps = append(*steps, state.Step{
				Array: copyArray(arr),

				Writing: []int{k},

				LeftRange:  []int{left, mid},
				RightRange: []int{mid + 1, right},
				MergedRange: []int{
					left,
					right,
				},

				Lines: []int{22},
			})

			j++
		}

		k++
	}

	// -------- REMAINING LEFT --------
	for i < n1 {

		arr[k] = L[i]

		*steps = append(*steps, state.Step{
			Array: copyArray(arr),

			Writing: []int{k},

			LeftRange:  []int{left, mid},
			RightRange: []int{mid + 1, right},
			MergedRange: []int{
				left,
				right,
			},

			Lines: []int{31},
		})

		i++
		k++
	}

	// -------- REMAINING RIGHT --------
	for j < n2 {

		arr[k] = R[j]

		*steps = append(*steps, state.Step{
			Array: copyArray(arr),

			Writing: []int{k},

			LeftRange:  []int{left, mid},
			RightRange: []int{mid + 1, right},
			MergedRange: []int{
				left,
				right,
			},

			Lines: []int{42},
		})

		j++
		k++
	}

	// -------- FINAL MERGED STEP --------
	*steps = append(*steps, state.Step{
		Array: copyArray(arr),

		MergedRange: []int{
			left,
			right,
		},

		Lines: []int{47},
	})
}
