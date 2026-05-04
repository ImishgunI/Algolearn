package storage

import (
	"Algolearn/internal/execution/state"
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Storage struct {
	rdb *redis.Client
}

func New(rdb *redis.Client) *Storage {
	return &Storage{rdb: rdb}
}

func (s *Storage) SaveSteps(ctx context.Context, execID string, steps []state.Step) error {
	key := fmt.Sprintf("execution:%s", execID)

	data, err := json.Marshal(steps)
	if err != nil {
		return err
	}

	return s.rdb.Set(ctx, key, data, 0).Err()
}

func (s *Storage) GetSteps(ctx context.Context, execID string) ([]state.Step, error) {
	key := fmt.Sprintf("execution:%s", execID)

	data, err := s.rdb.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	var steps []state.Step
	if err := json.Unmarshal(data, &steps); err != nil {
		return nil, err
	}

	return steps, nil
}
