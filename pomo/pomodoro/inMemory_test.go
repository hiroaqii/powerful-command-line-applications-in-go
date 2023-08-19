package pomodoro_test

import (
	"testing"

	"github.com/hiroaqii/powerful-command-line-applications-in-go/pomo/pomodoro"
	"github.com/hiroaqii/powerful-command-line-applications-in-go/pomo/pomodoro/repository"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()

	return repository.NewInMemoryRepo(), func() {}
}
