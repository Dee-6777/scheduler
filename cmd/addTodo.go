/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// addTodoCmd represents the addTodo command
var addTodoCmd = &cobra.Command{
	Use:   "addTodo",
	Short: "A brief description of your command",
	Long:  ``,
	Args:  cobra.MaximumNArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addTodo called")
		AddTodo(args)
	},
}

func AddTodo(args []string) {
	url := "http://localhost:8080/todos"
	payload := strings.NewReader("{\n  \"id\": \"" + args[0] + "\",\n  \"title\": \"" + args[1] + "\",\n  \"time\": \"" + args[2] + "\",\n  \"completed\": \"" + args[3] + "\"\n} \n\n\n")
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Printf("Could not make a request. Error %v", err)
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "CLI Tool")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Could not make a request. Error %v", err)
	}
	defer res.Body.Close()
}

func init() {
	rootCmd.AddCommand(addTodoCmd)
}
