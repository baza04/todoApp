package repository

import (
	todoapp "github.com/baza04/todoApp"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todoapp.User) (int, error)
	GetUser(username, password string) (todoapp.User, error)
}

type TodoList interface {
	Create(userId int, list todoapp.TodoList) (int, error)
	GetAll(userId int) ([]todoapp.TodoList, error)
	GetById(userId, id int) (todoapp.TodoList, error)
	Update(userId, id int, input *todoapp.UpdateListInput) error
	Delete(userId, id int) error
}

type TodoItem interface {
}
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      nil,
	}
}
