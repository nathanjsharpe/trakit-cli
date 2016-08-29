package main

import (
	"flag"
	"fmt"
	"github.com/nathanjsharpe/trakit/config"
	"github.com/nathanjsharpe/trakit/handlers"
	"github.com/nathanjsharpe/trakit/output"
)

const VERSION = "0.1.0"

func main() {
	flag.Parse()

	handle(flag.Args())

	fmt.Println()
}

func handle(args []string) {
	if len(args) == 0 {
		Help(args)
	} else {
		first, args := args[0], args[1:]
		switch first {
		case "config":
			handlers.Config(args)
		case "environments", "env":
			fmt.Println(config.GetConfig().EnvironmentKeys)
		case "landmark", "landmarks", "lm":
			config.ValidateSession()
			handlers.Landmark(args)
		case "login":
			handlers.Login(args)
		case "logout":
			handlers.Logout(args)
		case "session":
			handlers.Session(args)
		case "user":
			handlers.User(args)
		case "version", "versions":
			config.ValidateSession()
			handlers.Version(args, VERSION)
		default:
			Help(args)
		}
	}
}

func Help(args []string) {
	data := [][]string{
		[]string{"trakit config", "Manage trakit-cli configuration"},
		[]string{"trakit environments", "List possible trakit environments"},
		[]string{"trakit landmark", "Landmarks resource"},
		[]string{"trakit login", "Log into Trakit"},
		[]string{"trakit logout", "Log out of Trakit"},
		[]string{"trakit session", "Display details of current session"},
		[]string{"trakit user", "Users resource"},
		[]string{"trakit version", "Print version of Trakit CLI and services for selected environment"},
	}

	output.PrintBorderlessTable(data)
}
