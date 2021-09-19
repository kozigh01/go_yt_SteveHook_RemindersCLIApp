package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no command provided");
		os.Exit(2)
	}

	cmd := os.Args[1]
	switch cmd {
	case "greet":
		greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
		msgFlag := greetCmd.String("msg", "CLI BASICS - REMINDER CLI", "the message for greet command")
		if err := greetCmd.Parse(os.Args[2:]); err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("hello and wlecome: %s\n", *msgFlag)
	case "help":
		fmt.Println("some help messge")
	default:
		fmt.Printf("unknown command: %q\n", cmd)
	}
}