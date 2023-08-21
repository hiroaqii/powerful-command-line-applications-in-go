//go:build inmemory
// +build inmemory

package cmd

import (
	"github.com/hiroaqii/powerful-command-line-applications-in-go/pomo/pomodoro"
	"github.com/hiroaqii/powerful-command-line-applications-in-go/pomo/pomodoro/repository"
)

func getRepo() (pomodoro.Repository, error) {
	return repository.NewInMemoryRepo(), nil
}
