package todoapp

import "errors"

// TodoList entity used to working with TODO lists
type TodoList struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

// UserList entity used to connecting users with TODO lists
type UsersList struct {
	ID     int
	UserID int
	ListID int
}

// TodoItem entity used to working with TODO items
type TodoItem struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

// ListItem entity used to connecting TODO lists with TODO items
type ListsItem struct {
	ID     int
	ListID int
	ItemID int
}

// UpdateListInput entity is used to update the TODO list
type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

// Validate method checks the filling of the fields of the UpdaListInput struct
func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update list structure fields is empty")
	}
	return nil
}

// UpdateItemInput entity used to update the TODO item
type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

// Validate method checks the filling of the fields of the UpdateItemInput
func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("update item structure fields is empty")
	}
	return nil
}
