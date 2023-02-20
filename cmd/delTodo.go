/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// delTodoCmd represents the delTodo command
var delTodoCmd = &cobra.Command{
	Use:   "delTodo",
	Short: "A brief description of your command",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delTodo called")
		DelTodo(args)
	},
}

func DelTodo(args []string) {
	str := args[0]
	url := "http://localhost:8080/todos/" + str

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Printf("Could not make a request. Error %v", err)
	}

	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "CLI TOOL")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Could not make a request. Error %v", err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println(string(body))
}

func init() {
	rootCmd.AddCommand(delTodoCmd)
}
