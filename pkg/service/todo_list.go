package service

import (
	todoapp "github.com/baza04/todoApp"
	"github.com/baza04/todoApp/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todoapp.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]todoapp.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, id int) (todoapp.TodoList, error) {
	return s.repo.GetById(userId, id)
}

func (s *TodoListService) Update(userId, id int, input *todoapp.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, id, input)
}

func (s *TodoListService) Delete(userId, id int) error {
	return s.repo.Delete(userId, id)
}
