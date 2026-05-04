package algorithms

import "Algolearn/internal/execution/state"

type Algorithm interface {
	Run(input []int) []state.Step
}
