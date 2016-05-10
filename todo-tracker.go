package main

import (
	"runtime"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	// _ "github.com/mattn/go-sqlite3"
)

type TodoItem struct {
	ID          int
	Description string
	Deadline    string
	Progress    int
}

func ListView(c *gin.Context) {

	TodoList := make([]TodoItem, 1)

	TodoList[0] = TodoItem{
		ID:          1,
		Description: "Poke Tom!",
		Deadline:    "11/05/2016",
		Progress:    40,
	}

	// Make a database call for all todo items.

	// Forward todo list to template parser.
	c.HTML(200, "index.html", gin.H{
		"TodoList": TodoList,
	})
}

func ImprintView(c *gin.Context) {
	c.HTML(200, "imprint.html", gin.H{})
}

func AddView(c *gin.Context) {
	c.HTML(200, "add.html", gin.H{})
}

func AddHandler(c *gin.Context) {

	// Retrieve data from formular.

	// Create todo item based on input data.

	// Save model to database.

	// On success - redirect to list view.
}

func EditView(c *gin.Context) {

	// Database query.

	// Forward data to template parser.
	c.HTML(200, "edit.html", gin.H{})
}

func EditHandler(c *gin.Context) {}

func DeleteHandler(c *gin.Context) {}

func main() {

	// Set maximum available CPUs to be used by go.
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	// Instantiate new gin router with default middleware.
	app := gin.Default()

	// Load HTML template files.
	app.LoadHTMLGlob("views/*")
	app.Static("/css", "./css")
	app.Static("/js", "./js")
	app.Static("/fonts", "./fonts")

	// Define routes to end points.
	app.GET("/", ListView)
	app.GET("/add", AddView)
	app.POST("/add", AddHandler)
	app.GET("/edit/:todoID", EditView)
	app.POST("/edit/:todoID", EditHandler)
	app.GET("/delete/:todoID", DeleteHandler)
	app.GET("/imprint", ImprintView)

	// Start the web application.
	app.Run(":8080")
}
