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
	Create(userId int, list todoapp.TodoList) (int, error)
	GetAll(userId int) ([]todoapp.TodoList, error)
	GetById(userId, id int) (todoapp.TodoList, error)
	Update(userId, id int, input *todoapp.UpdateListInput) error
	Delete(userId, id int) error
}

type TodoItem interface{}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
		TodoItem:      nil,
	}
}

/* type Authorization interface{
	CreateUser(user todoapp.User) (int, error)
}


func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		//TodoList:      repo.TodoList,
		//TodoItem:      repo.TodoItem,
	}
} */
