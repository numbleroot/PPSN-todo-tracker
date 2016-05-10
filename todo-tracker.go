package main

import (
	"log"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/numbleroot/PPSN-todo-tracker/db"
)

func DatabaseMiddleware(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func ListView(c *gin.Context) {

	var TodoList []db.TodoItem

	// Retrieve database connection instance from context.
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		log.Fatal("[ListView] Could not retrieve database connection from gin context.")
	}

	// Make a database call for all todo items.
	db.Find(&TodoList)

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
	todoDescription := c.PostForm("todoDescription")
	todoDeadline := c.PostForm("todoDeadline")
	todoProgress, _ := strconv.Atoi(c.PostForm("todoProgress"))

	// Create todo item based on input data.
	Todo := db.TodoItem{
		Description: todoDescription,
		Deadline:    todoDeadline,
		Progress:    todoProgress,
	}

	// Retrieve database connection instance from context.
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		log.Fatal("[AddHandler] Could not retrieve database connection from gin context.")
	}

	// Save model to database.
	db.Create(&Todo)

	// On success - redirect to list view.
	c.Redirect(301, "/")
}

func EditView(c *gin.Context) {

	var Todo db.TodoItem

	// Retrieve database connection instance from context.
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		log.Fatal("[EditView] Could not retrieve database connection from gin context.")
	}

	// Get ID of todo from context.
	todoID := c.Params.ByName("todoID")

	// Database query.
	db.Find(&Todo, "id = ?", todoID)

	// Forward data to template parser.
	c.HTML(200, "edit.html", gin.H{
		"Todo": Todo,
	})
}

func EditHandler(c *gin.Context) {}

func DeleteHandler(c *gin.Context) {}

func main() {

	// Set maximum available CPUs to be used by go.
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	// Instantiate new gin router with default middleware.
	app := gin.Default()

	// Open up database connection and add it as middleware.
	db := db.InitDB("todos.sqlite3")
	app.Use(DatabaseMiddleware(db))

	// Load HTML template files.db
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
