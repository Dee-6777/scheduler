/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"log"

	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	"github.com/spf13/cobra"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

const (
	markName  = "REMINDER"
	markValue = "1"
)

// notificationCmd represents the notification command
var notificationCmd = &cobra.Command{
	Use:   "notification",
	Short: "A brief description of your command",
	Long:  ``,
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("notification called")
		ShowNotif(args)
	},
}

func ShowNotif(args []string) {
	if len(args) < 2 {
		//fmt.Printf("Usage:%s <hh:mm> <text message\n>", os.Args[0])
		os.Exit(1)
	}

	now := time.Now()
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	t, err := w.Parse(args[0], now)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	if t == nil {
		fmt.Println("Unable to parse time")
		os.Exit(2)
	}

	if now.After(t.Time) {
		fmt.Println("set a future time")
	}

	diff := t.Time.Sub(now)
	if os.Getenv(markName) == markValue {
		time.Sleep(diff)
		f, err := os.Open("cmd/beep-06.mp3")
		if err != nil {
			log.Fatal(err)
		}

		streamer, format, err := mp3.Decode(f)
		if err != nil {
			log.Fatal(err)
		}
		defer streamer.Close()

		speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

		done := make(chan bool)
		speaker.Play(beep.Seq(streamer, beep.Callback(func() {
			done <- true
		})))

		<-done
		err = beeep.Alert("Reminder", strings.Join(args[1:], " "), "assets/information.png")
		if err != nil {
			fmt.Println(err)
			os.Exit(4)
		}
	} else {
		app := exec.Command(os.Args[0], os.Args[1:]...)
		app.Env = append(os.Environ(), fmt.Sprintf("%s=%s", markName, markValue))
		if err := app.Start(); err != nil {
			fmt.Println(err)
			os.Exit(5)
		}
		fmt.Println("Reminder will be displayed after", diff.Round(time.Second))
	}
}
func init() {
	rootCmd.AddCommand(notificationCmd)
}
