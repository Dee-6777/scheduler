/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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

func Buzz() {
	for k := range Todos {
		arg := []string{}
		arg = append(arg, Todos[k].Time, Todos[k].Item)
		Alerts(arg)
		fmt.Printf("ID is '%s' Title is '%s'\n", Todos[k].ID, Todos[k].Item)
	}
}

func init() {
	rootCmd.AddCommand(alertsCmd)
}
