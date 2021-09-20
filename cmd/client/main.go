package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kozigh01/go_yt_SteveHook_RemindersCLIApp/client"
)

var (
	backendUriFlag = flag.String("backend", "http://swapi.com", "Backend API Url")
	helpFlag = flag.Bool("help", false, "Display a helpful message")
)

func main() {
	flag.Parse()
	s := client.NewSwitch(*backendUriFlag)

	if *helpFlag || len(os.Args) == 1 {
		s.Help()
		return
	}

	err := s.Switch()
	if err != nil {
		fmt.Printf("cmd switch error: %s\n", err)
		os.Exit(2)
	}
}