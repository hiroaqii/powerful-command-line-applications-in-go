package repository

import (
	"fmt"
	"sync"

	"github.com/hiroaqii/powerful-command-line-applications-in-go/pomo/pomodoro"
)

type InMemoryRepo struct {
	sync.RWMutex
	intervals []pomodoro.Interval
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		intervals: []pomodoro.Interval{},
	}
}

func (r *InMemoryRepo) Create(i pomodoro.Interval) (int64, error) {
	r.Lock()
	defer r.Unlock()

	i.ID = int64(len(r.intervals) + 1)

	r.intervals = append(r.intervals, i)

	return i.ID, nil
}

func (r *InMemoryRepo) Update(i pomodoro.Interval) error {
	r.Lock()
	defer r.Unlock()

	if i.ID == 0 {
		return fmt.Errorf("%w: %d", pomodoro.ErrInvalidID, i.ID)
	}

	r.intervals[i.ID-1] = i

	return nil
}

func (r *InMemoryRepo) ById(id int64) (pomodoro.Interval, error) {
	r.RLock()
	defer r.RUnlock()

	i := pomodoro.Interval{}
	if id == 0 {
		return i, fmt.Errorf("%w: %d", pomodoro.ErrInvalidID, id)
	}

	i = r.intervals[id-1]

	return i, nil
}

func (r *InMemoryRepo) Breaks(n int) ([]pomodoro.Interval, error) {
	r.RLock()
	defer r.RUnlock()

	data := []pomodoro.Interval{}
	for k := len(r.intervals) - 1; k >= 0; k-- {
		if r.intervals[k].Category == pomodoro.CategoryPomodoro {
			continue
		}

		data = append(data, r.intervals[k])

		if len(data) == n {
			return data, nil
		}
	}

	return data, nil
}
