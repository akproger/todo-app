package service

import "github.com/akproger/todo-app/pkg/repository"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService() *Service {
	return &Service{}
}

func NewRepository(repos *repository.Repository) *Service {
	return &Service{}
}
