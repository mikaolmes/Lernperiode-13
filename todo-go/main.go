package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type todo struct {
	ID 			string 	`json:"id"`
	Item		string 	`json:"item"`
	Completed	bool 	`json:"completed"`
}

var todos = []todo{
	{ID:"1", Item: "Clean Room", Completed: false},
	{ID:"2", Item: "Read Book", Completed: false},
	{ID:"3", Item: "Record Video", Completed: false},
}

func getTodos(context * gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context * gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil { // context.BindJSON binds the received JSON to newTodo
		return // return if there is an error
	}

	todos = append(todos, newTodo) // append newTodo to todos slice

	context.IndentedJSON(http.StatusCreated, newTodo) // return the new todo with status created
}


func getTodo(context * gin.Context) {
	id := context.Param("id")
	todo, err := getTodoByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func toggleTodoStatus(context * gin.Context) {
	id := context.Param("id")
	todo, err := getTodoByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	todo.Completed = !todo.Completed // flip the bool status

	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoByID(id string) (*todo, error) {
	for i, t := range todos { 
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func main() {
	router := gin.Default() // router = server
	router.GET("/todos", getTodos) // endpoint to get all todos
	router.GET("/todos/:id", getTodo) // endpoint to get a todo by id
	router.PATCH("/todos/:id", toggleTodoStatus) // endpoint to update a todo by id
	router.POST("/todos", addTodo) // endpoint to add a new todo
	router.Run("localhost:9090") // run server on port 9090
}