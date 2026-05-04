package manager

import (
	"Algolearn/internal/execution/algorithms"
	"Algolearn/internal/execution/state"
	"Algolearn/internal/execution/storage"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Manager struct {
	algorithms map[string]algorithms.Algorithm
	storage    *storage.Storage
}

func New(storage *storage.Storage) *Manager {
	return &Manager{
		storage: storage,
		algorithms: map[string]algorithms.Algorithm{
			"bubble_sort": &algorithms.BubbleSort{},
		},
	}
}

func (m *Manager) Execute(ctx context.Context, algorithm string, data []int) (string, error) {
	algo, ok := m.algorithms[algorithm]
	if !ok {
		return "", fmt.Errorf("unknown algorithm")
	}

	steps := algo.Run(data)

	execID := uuid.New().String()

	err := m.storage.SaveSteps(ctx, execID, steps)
	if err != nil {
		return "", err
	}

	return execID, nil
}

func (m *Manager) Get(ctx context.Context, execID string) ([]state.Step, error) {
	return m.storage.GetSteps(ctx, execID)
}
