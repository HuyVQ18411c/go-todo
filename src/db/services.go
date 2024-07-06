package db

import (
	"gorm.io/gorm"
)

type ITodoService interface {
	GetAllTodos() []Todo
	GetTodoById(todoId int) *Todo
	CreateTodo(data CreateTodoDto) *Todo
	DeleteTodo(todoId int)
}

type TodoService struct {
	database *gorm.DB
}

func NewTodoService(database *gorm.DB) *TodoService {
	newService := TodoService{database: database}
	return &newService
}

func (service *TodoService) GetAllTodos() []Todo {
	var todos []Todo
	service.database.Raw("SELECT * FROM todos").Scan(&todos)
	return todos
}

func (service *TodoService) CreateTodo(data CreateTodoDto) *Todo {
	newTodo := Todo{Title: data.Title, Description: data.Description, Deadline: data.Deadline}
	service.database.Create(&newTodo)
	return &newTodo
}

func (service *TodoService) DeleteTodo(todoId int) {
	service.database.Exec("DELETE FROM todos WHERE id = ?", todoId)
}

func (service *TodoService) GetTodoById(todoId int) *Todo {
	var todo Todo
	service.database.Raw("SELECT id FROM todos WHERE id = ?", todoId).Scan(&todo)

	if todo.ID == 0 {
		return nil
	}

	return &todo
}
