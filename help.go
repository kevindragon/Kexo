package main

import (
	"fmt"
)

func Help() {
	fmt.Println("Help")
}

func HelpCmd(cmd Command) {
	switch cmd {
	case NewCmd:
		fmt.Println("help for new command")
	default:
		Help()
	}
}
