/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json" // To encode it to json
	"fmt"           // This package allows to format basic strings, values, or anything and print
	"log"           // To logout errors
	"net/http"      // allows to create a server

	"github.com/gorilla/mux" // gorilla/mux is a package which adapts to Go’s default HTTP router and establish different routes
	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("api called")
		Schedule()
	},
}

type Todo struct { // struct and slices has been used as database
	ID        string `json:"id"`
	Item      string `json:"title"`
	Time      string `json:"time"`
	Completed string `json:"completed"`
}

var Todos = []Todo{ // slice of type Todo has been declared and initialised with some json values
	{ID: "1", Item: "Make your bed", Time: "06:00", Completed: "false"},
	{ID: "2", Item: "Brush your Teeth", Time: "06:20", Completed: "false"},
	{ID: "3", Item: "Go For a Walk", Time: "06:30", Completed: "false"},
	{ID: "4", Item: "It's Breakfast time!", Time: "07:00", Completed: "false"},
	{ID: "5", Item: "Take a shower", Time: "07:30", Completed: "false"},
	{ID: "6", Item: "Standup", Time: "09:00", Completed: "false"},
	{ID: "7", Item: "Take a Break", Time: "09:30", Completed: "false"},
}

func Schedule() {
	r := mux.NewRouter()
	// Registering Request Handlers
	r.HandleFunc("/todos", GetTodos).Methods("GET")           // Registering get request handler for fetching all the elements present in our database
	r.HandleFunc("/todos/{id}", GetTodo).Methods("GET")       // Registering get request handler which will fetch a single element using it's id in our database
	r.HandleFunc("/todos", CreateTodo).Methods("POST")        // Registering post request handler to create and add a new element in our database
	r.HandleFunc("/todos/{id}", UpdateTodo).Methods("PUT")    // Registering put request handler to update a todo
	r.HandleFunc("/todos/{id}", DeleteTodo).Methods("DELETE") // Registering delete request handler to delete a todo

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r)) // to start a server at port:8080
}

// @GET ('/todos') ...to fetch data
func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Todos)
}

// @DELETE ('/todos/{id}') ...to delete a value by using id
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Todos {
		if item.ID == params["id"] {
			Todos = append(Todos[:index], Todos[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Todos)
}

// @GET method ('/todos/{id}') ...to fetch data by id
func GetTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range Todos {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// @POST ('/todos') ...to create data
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	Todos = append(Todos, todo)
	json.NewEncoder(w).Encode(todo)
}

// @PUT ('/todos/{id}') ...update data
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Todos {
		if item.ID == params["id"] {
			Todos = append(Todos[:index], Todos[index+1:]...)
			var todo Todo
			_ = json.NewDecoder(r.Body).Decode(&todo)
			todo.ID = params["id"]
			Todos = append(Todos, todo)
			json.NewEncoder(w).Encode(todo)
		}
	}
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
