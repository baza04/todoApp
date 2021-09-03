package service

import (
	todoapp "github.com/baza04/todoApp"
	"github.com/baza04/todoApp/pkg/repository"
)

type Authorization interface {
	CreateUser(user todoapp.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type TodoList interface {
	Create(userID int, list todoapp.TodoList) (int, error)
	GetAll(userID int) ([]todoapp.TodoList, error)
	GetByID(userID, id int) (todoapp.TodoList, error)
	Update(userID, id int, input *todoapp.UpdateListInput) error
	Delete(userID, id int) error
}

type TodoItem interface {
	Create(userID, listID int, item todoapp.TodoItem) (int, error)
	GetAll(userID, listID int) ([]todoapp.TodoItem, error)
	GetByID(userID, itemID int) (todoapp.TodoItem, error)
	Update(userID, itemID int, input todoapp.UpdateItemInput) error
	Delete(userID, itemID int) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
		TodoItem:      NewTodoItemService(repo.TodoItem, repo.TodoList),
	}
}
