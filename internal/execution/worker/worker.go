package worker

import (
	"Algolearn/internal/execution/algorithms"
	"Algolearn/internal/execution/manager"
	"Algolearn/internal/execution/storage"
	"context"
	"encoding/json"
)

type Worker struct {
	storage    *storage.Storage
	algorithms map[string]algorithms.Algorithm
}

func New(storage *storage.Storage) *Worker {
	return &Worker{
		storage: storage,
		algorithms: map[string]algorithms.Algorithm{
			"bubble_sort": &algorithms.BubbleSort{},
		},
	}
}

func (w *Worker) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			w.process(ctx)
		}
	}
}

func (w *Worker) process(ctx context.Context) {
	payload, err := w.storage.Dequeue(ctx)
	if err != nil {
		return
	}

	var job manager.Job
	if err := json.Unmarshal([]byte(payload), &job); err != nil {
		return
	}

	algo, ok := w.algorithms[job.Algorithm]
	if !ok {
		return
	}

	steps := algo.Run(job.Data)

	_ = w.storage.SaveSteps(ctx, job.ExecutionID, steps)
}
