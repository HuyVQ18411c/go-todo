package main

import (
	"net/http"

	"github.com/HuyVQ18411c/go-todo/src/config"
	"github.com/HuyVQ18411c/go-todo/src/controllers"
	"github.com/HuyVQ18411c/go-todo/src/db"
	"github.com/HuyVQ18411c/go-todo/src/middlewares"
	"gorm.io/gorm"
)

type App struct {
	router *http.ServeMux
	config *config.Config
	db     *gorm.DB
}

func (app *App) AddRouter(path string, router *http.ServeMux) {
	app.router.Handle(path, router)
}

func NewApp() *App {
	app := App{}

	// Add config
	app.config = config.NewConfig()

	// Add database
	// Basically, there is a connection pool
	app.db = db.NewDatabase(app.config)

	// Add router
	app.router = http.NewServeMux()
	app.AddRouter("/", controllers.CreateHomeRouter(app.db))

	return &app
}

func main() {
	app := NewApp()
	server := http.Server{
		Addr:    app.config.Address,
		Handler: middlewares.Logging(app.router),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
