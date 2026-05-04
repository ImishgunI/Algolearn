package manager

import (
	"Algolearn/internal/execution/algorithms"
	"Algolearn/internal/execution/state"
	"fmt"
)

type Manager struct {
	algorithms map[string]algorithms.Algorithm
}

func New() *Manager {
	return &Manager{
		algorithms: map[string]algorithms.Algorithm{
			"bubble_sort": &algorithms.BubbleSort{},
		},
	}
}

func (m *Manager) Execute(algorithm string, data []int) ([]state.Step, error) {
	algo, ok := m.algorithms[algorithm]
	if !ok {
		return nil, fmt.Errorf("unknown algorithm")
	}

	steps := algo.Run(data)

	return steps, nil
}
