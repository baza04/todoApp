package service

import (
	todoapp "github.com/baza04/todoApp"
	"github.com/baza04/todoApp/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{
		repo:     repo,
		listRepo: listRepo,
	}
}

func (s *TodoItemService) Create(userId, listId int, input todoapp.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil { // list doesn't exist or not belong to user
		return 0, nil
	}

	return s.repo.Create(userId, listId, input)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]todoapp.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (todoapp.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemService) Update(userId, itemId int, input todoapp.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, itemId, input)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}
