package main

import (
	"flag"
	"fmt"
	"github.com/nathanjsharpe/trakit/config"
	"github.com/nathanjsharpe/trakit/handlers"
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
			environments := config.GetConfig()
			fmt.Println(environments)
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
