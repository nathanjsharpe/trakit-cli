package handlers

import (
	"fmt"
)

var userHeaders = []string{
	"Id",
	"Name",
	"Address",
	"Description",
	"Icon",
	"ShowOnMap",
	"Lat",
	"Lng",
	"Status",
}

func User(args []string) {
	if len(args) == 0 {
		showAllUsers()
	} else {
		first, args := args[0], args[1:]
		switch first {
		case "activate":
			handleActivate(args)
		}
	}
}

func showAllUsers() {
	fmt.Println("fetching users")
}

func handleActivate(args []string) {
	if len(args) == 0 {
		fmt.Println("Activates a user (sets inactive to false)")
		fmt.Println("Usage: trakit user activate <username>")
	} else {
		fmt.Println("activating", args[0])
	}
}
