package main

import (
	"fmt"
	"strings"
	// "log"
	"os"
)

func main() {
	// fmt.Println(os.Args)

	if len(os.Args) < 2 {
		// log.Fatal("no command provided")
		fmt.Println("no command provided");
		os.Exit(2)
	}

	cmd := os.Args[1]
	switch cmd {
	case "greet":
		msg := "REMINDERS CLI - CLI BASICS"
		if len(os.Args) > 2 {
			f := strings.Split(os.Args[2], "=")
			if len(f) == 2 && f[0] == "--msg" {
				msg = f[1]
			}
		}
		fmt.Printf("hello and welcome: %s\n", msg)
	case "help":
		fmt.Println("some help messge")
	default:
		fmt.Printf("unknown command: %q\n", cmd)
	}
}