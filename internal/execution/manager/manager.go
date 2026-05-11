package manager

import (
	"Algolearn/internal/execution/algorithms"
	"Algolearn/internal/execution/state"
	"Algolearn/internal/execution/storage"
	"context"
	"encoding/json"

	"github.com/google/uuid"
)

type Job struct {
	ExecutionID string `json:"execution_id"`
	Algorithm   string `json:"algorithm"`
	Data        []int  `json:"data"`
}

type Manager struct {
	algorithms map[string]algorithms.Algorithm
	storage    *storage.Storage
}

func New(storage *storage.Storage) *Manager {
	return &Manager{
		storage: storage,
		algorithms: map[string]algorithms.Algorithm{
			"bubble_sort": &algorithms.BubbleSort{},
			"quick_sort":  &algorithms.QuickSort{},
			"merge_sort":  &algorithms.MergeSort{},
		},
	}
}

func (m *Manager) Execute(ctx context.Context, algorithm string, data []int) (string, error) {
	execID := uuid.New().String()

	job := Job{
		ExecutionID: execID,
		Algorithm:   algorithm,
		Data:        data,
	}

	payload, err := json.Marshal(job)
	if err != nil {
		return "", err
	}

	err = m.storage.Enqueue(ctx, string(payload))
	if err != nil {
		return "", err
	}

	return execID, nil
}

func (m *Manager) Get(ctx context.Context, execID string) ([]state.Step, error) {
	return m.storage.GetSteps(ctx, execID)
}
