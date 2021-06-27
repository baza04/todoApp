package service

import "github.com/baza04/todoApp/pkg/repository"

type Authorization interface{}

type TodoList interface{}

type TodoItem interface{}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: repo.Authorization,
		TodoList:      repo.TodoList,
		TodoItem:      repo.TodoItem,
	}
}
