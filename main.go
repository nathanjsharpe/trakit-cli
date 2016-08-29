package main

import (
	"flag"
	"fmt"
	"github.com/nathanjsharpe/trakit/config"
	"github.com/nathanjsharpe/trakit/handlers"
	"github.com/nathanjsharpe/trakit/output"
	"os/exec"
	"regexp"
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
			fmt.Println(config.CurrentEnvironment())
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
		case "logs":
			fmt.Println("Searching for error with ID", args[1])
			env := config.CurrentEnvironment()
			dest := env.SshUser + "@trakit." + env.Key
			currentCheck := "App Api"
			out, err := exec.Command("ssh", dest, "cat", env.AppApiLogPath).Output()
			if err != nil {
				panic(err)
			}
			rx := `ERROR.+` + args[1] + `\n(?:!.+\n)+`
			regx, _ := regexp.Compile(rx)
			result := regx.FindString(string(out))
			if result == "" {
				fmt.Println("No error found with ID", args[1])
			} else {
				fmt.Printf("Found error with ID %s on %s:\n %s", args[1], currentCheck, result)
			}
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
