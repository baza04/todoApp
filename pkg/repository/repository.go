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
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
