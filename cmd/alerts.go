/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type Todos struct { // struct and slices has been used as database
	ID        string `json:"id"`
	Item      string `json:"title"`
	Time      string `json:"time"`
	Completed string `json:"completed"`
}

// alertsCmd represents the alerts command
var alertsCmd = &cobra.Command{
	Use:   "alerts",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("alerts called")
		Buzz()
	},
}

var todo []Todos

func Conversion() []Todos {
	url := "https://scheduler-api-go.onrender.com/todos"
	responseBytes := GetTodosData(url)
	if err := json.Unmarshal(responseBytes, &todo); err != nil {
		log.Printf("Could not unmarshall response %v", err)
	}
	return todo
}

func Buzz() {
	values := Conversion()
	for i := range values {
		arg := []string{}
		arg = append(arg, values[i].Time, values[i].Item)
		ShowNotif(arg)
	}
}

func GetTodosData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)
	if err != nil {
		log.Printf("Could not get a joke. Error %v", err)
		os.Exit(1)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Joke generator CLI")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request. Error %v", err)
	}
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read the response. Error%v", err)
	}
	return responseBytes
}

func init() {
	rootCmd.AddCommand(alertsCmd)
}
