package db

import (
	"errors"
	models "todo/models"
)

// Todos eport empty map
var Todos = make(map[int]*models.Todo)

// AddTodo adds a todo to todos map
func AddTodo(t *models.Todo) {
	Todos[t.ID] = t
}

// DeleteTodo delete todo
func DeleteTodo(id int) {
	delete(Todos, id)
}

// UpdateTodo update todo
func UpdateTodo(id int, t *models.Todo) (*models.Todo, error) {
	if _, ok := Todos[id]; !ok {
		return nil, errors.New("id not found")
	}

	Todos[id] = t
	return t, nil
}
