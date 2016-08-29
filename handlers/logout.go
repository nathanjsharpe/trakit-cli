package handlers

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/nathanjsharpe/trakit/config"
)

func Logout(args []string) {
	err := config.DeleteSession()
	if err != nil {
		panic(err)
	}

	color.Set(color.FgGreen)
	fmt.Println("Logged out")
	color.Unset()
}
