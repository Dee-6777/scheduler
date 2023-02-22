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

var updateTodoCmd = &cobra.Command{
	Use:   "updateTodo",
	Short: "A brief description of your command",
	Long:  ``,
	Args:  cobra.MaximumNArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updateTodo called")
		Update(args)
	},
}

func Update(args []string) {
	str := args[0]
	url := "https://scheduler-api-go.onrender.com/todos/" + str

	payload := strings.NewReader("{\n  \"id\": \"" + args[0] + "\",\n  \"title\": \"" + args[1] + "\",\n  \"time\": \"" + args[2] + "\",\n  \"completed\": \"" + args[3] + "\"\n} \n\n\n")

	req, _ := http.NewRequest("PUT", url, payload)

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
	rootCmd.AddCommand(updateTodoCmd)
}
