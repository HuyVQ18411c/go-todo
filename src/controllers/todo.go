package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/HuyVQ18411c/go-todo/src/db"
	"github.com/HuyVQ18411c/go-todo/src/utils"
	"gorm.io/gorm"
)

type TodoController struct {
	service db.ITodoService
}

func NewTodoController(database *gorm.DB) *TodoController {
	controller := TodoController{}
	controller.service = db.NewTodoService(database)
	return &controller
}

func (controller *TodoController) getTodoList(response http.ResponseWriter, request *http.Request) {
	data := controller.service.GetAllTodos()

	response.WriteHeader(http.StatusOK)
	response.Write(utils.GetJSONResponse(http.StatusOK, data))
}

func (controller *TodoController) createTodo(response http.ResponseWriter, request *http.Request) {
	data, err := io.ReadAll(request.Body)

	if err != nil {
		panic(err)
	}

	todoData := db.CreateTodoDto{}
	json.Unmarshal(data, &todoData)

	newTodo := controller.service.CreateTodo(todoData)
	response.WriteHeader(http.StatusOK)
	response.Write(utils.GetJSONResponse(http.StatusOK, newTodo))
}

func (controller *TodoController) deleteTodo(response http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.PathValue("id"))

	if err != nil {
		panic(err)
	}
	log.Print("#### asd", id)
	todo := controller.service.GetTodoById(id)
	if todo == nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write(utils.GetJSONResponse(http.StatusNotFound, todo))
		return
	}

	controller.service.DeleteTodo(id)
	response.WriteHeader(http.StatusOK)
}

func CreateHomeRouter(database *gorm.DB) *http.ServeMux {
	controller := NewTodoController(database)
	router := http.NewServeMux()
	router.HandleFunc("GET /", controller.getTodoList)
	router.HandleFunc("POST /", controller.createTodo)
	router.HandleFunc("DELETE /{id}", controller.deleteTodo)
	return router
}
