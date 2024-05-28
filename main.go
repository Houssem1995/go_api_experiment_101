package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Learn Go", Completed: false},
	{ID: "2", Item: "Build a RESTful API", Completed: true},
	{ID: "3", Item: "Build a React App", Completed: false},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func addTodos(c *gin.Context) {
	var newTodo todo
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func getToddoByID(c *gin.Context) {
	id := c.Param("id")
	for _, t := range todos {
		if t.ID == id {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func deleteTodos(c *gin.Context) {
	id := c.Param("id")
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "todo deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func updateTodos(c *gin.Context) {
	id := c.Param("id")
	for i, t := range todos {
		if t.ID == id {
			var updatedTodo todo
			if err := c.BindJSON(&updatedTodo); err != nil {
				return
			}
			todos[i] = updatedTodo
			c.IndentedJSON(http.StatusOK, updatedTodo)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getToddoByID)
	router.POST("/todos", addTodos)
	router.DELETE("/todos/:id", deleteTodos)
	router.PUT("/todos/:id", updateTodos)
	router.Run("localhost:9999")
}
