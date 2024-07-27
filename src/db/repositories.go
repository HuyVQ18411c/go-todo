package db

import (
	"gorm.io/gorm"
)

type ITodoRepository interface {
	GetAllTodos() []Todo
	GetTodoById(todoId int) *Todo
	CreateTodo(data CreateTodoDto) *Todo
	DeleteTodo(todoId int)
}

type TodoRepository struct {
	database *gorm.DB
}

func NewTodoRepository(database *gorm.DB) *TodoRepository {
	newService := TodoRepository{database: database}
	return &newService
}

func (service *TodoRepository) GetAllTodos() []Todo {
	var todos []Todo
	service.database.Raw("SELECT * FROM todos").Scan(&todos)
	return todos
}

func (service *TodoRepository) CreateTodo(data CreateTodoDto) *Todo {
	newTodo := Todo{Title: data.Title, Description: data.Description, Deadline: data.Deadline}
	service.database.Create(&newTodo)
	return &newTodo
}

func (service *TodoRepository) DeleteTodo(todoId int) {
	service.database.Exec("DELETE FROM todos WHERE id = ?", todoId)
}

func (service *TodoRepository) GetTodoById(todoId int) *Todo {
	var todo Todo
	service.database.Raw("SELECT id FROM todos WHERE id = ?", todoId).Scan(&todo)

	if todo.ID == 0 {
		return nil
	}

	return &todo
}
