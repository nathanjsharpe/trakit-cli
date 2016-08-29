package handlers

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/nathanjsharpe/trakit/config"
	"os"
)

func Logout(args []string) {
	err := os.Remove(config.SessionPath)
	if err != nil {
		panic(err)
	}
	color.Set(color.FgGreen)
	fmt.Println("Logged out")
	color.Unset()
}
