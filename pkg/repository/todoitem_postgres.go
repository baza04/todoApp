package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	todoapp "github.com/baza04/todoApp"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{
		db: db,
	}
}

func (r *TodoItemPostgres) Create(userID, listID int, item todoapp.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemID int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", todoItemsTable)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	if err = row.Scan(&itemID); err != nil {
		return 0, fmt.Errorf("scan ID after insert todo item failed: %w", tx.Rollback())
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listsItemsTable)
	if _, err = tx.Exec(createListItemsQuery, listID, itemID); err != nil {
		return 0, fmt.Errorf("todo item list insert executing failed: %w", tx.Rollback())
	}

	return itemID, tx.Commit()
}

func (r *TodoItemPostgres) GetAll(userID, listID int) ([]todoapp.TodoItem, error) {
	var items []todoapp.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id
		INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2`, todoItemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Select(&items, query, listID, userID); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TodoItemPostgres) GetByID(userID, itemID int) (todoapp.TodoItem, error) {
	var item todoapp.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id 
		INNER JOIN %s ul ON ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2`, todoItemsTable, listsItemsTable, usersListsTable)
	err := r.db.Get(&item, query, itemID, userID)

	return item, err
}

func (r *TodoItemPostgres) Update(userID, itemID int, input todoapp.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, *input.Title)
		argID++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argID))
		args = append(args, *input.Description)
		argID++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argID))
		args = append(args, *input.Done)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		todoItemsTable, setQuery, listsItemsTable, usersListsTable, argID, argID+1)
	args = append(args, userID, itemID)

	logrus.Debugf("itemUpdateQuery %s\n", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)

	return err
}

func (r *TodoItemPostgres) Delete(userID, itemID int) error {
	query := fmt.Sprintf("DELETE FROM %s ti USING %s li, %s ul WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2",
		todoItemsTable, listsItemsTable, usersListsTable)
	_, err := r.db.Exec(query, userID, itemID)

	return err
}
