/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"sync"

	"github.com/Dee-6777/scheduler/cmd"
	"github.com/Dee-6777/scheduler/ui"
)

func main() {
	var waitgroup sync.WaitGroup
	waitgroup.Add(2)
	go func() {
		cmd.Execute()
		waitgroup.Done()
	}()
	go func() {
		ui.Greet()
		waitgroup.Done()
	}()
	waitgroup.Wait()
}
