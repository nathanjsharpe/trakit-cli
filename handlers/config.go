package handlers

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/nathanjsharpe/trakit/config"
	"github.com/nathanjsharpe/trakit/output"
	"os"
	"strconv"
	"strings"
)

func Config(args []string) {
	if len(args) == 0 {
		configHelp()
	} else {
		first, args := args[0], args[1:]
		switch first {
		case "generate", "g":
			handleGenerateConfig(args)
		case "environments", "environment", "env":
			handleEnvironments(args)
		case "show":
			fmt.Println(config.GetConfig())
		default:
			configHelp()
		}
	}

	session := config.LoadSession()
	data := [][]string{
		[]string{"Environment", session.Environment},
		[]string{"User", session.User.Username},
		[]string{"User Id", strconv.Itoa(session.User.Id)},
		[]string{"Expires", session.Token.Expiration},
	}
	output.PrintSimpleTable(data)
}

func handleGenerateConfig(args []string) {
	config.GenerateConfigFile()
}

func handleEnvironments(args []string) {
	if len(args) == 0 {
		environmentsHelp()
	}

	first, args := args[0], args[1:]
	switch first {
	case "list":
		fmt.Println(config.GetConfig().EnvironmentKeys)
	case "add":
		handleAddEnvironment(args)
	default:
		environmentsHelp()
	}
}

func handleAddEnvironment(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: trakit config environments add <key>")
		fmt.Println("\nExample: trakit config environments add party")
	} else {
		newEnv := config.Environment{
			Key: args[0],
		}

		fmt.Println("Creating environment '" + newEnv.Key + "'")

		reader := bufio.NewReader(os.Stdin)

		color.Set(color.FgWhite, color.Bold)
		fmt.Print("\nApp api url: ")
		color.Unset()

		appApiUrl, _ := reader.ReadString('\n')
		newEnv.AppApi = strings.TrimSpace(appApiUrl)

		color.Set(color.FgWhite, color.Bold)
		fmt.Print("Web app url: ")
		color.Unset()

		webAppUrl, _ := reader.ReadString('\n')
		newEnv.WebApp = strings.TrimSpace(webAppUrl)

		config.AddEnvironment(newEnv)
	}
}

func configHelp() {
	data := [][]string{
		[]string{"trakit config show", "Show current configuration"},
		[]string{"trakit config generate", "Create a configuration file"},
		[]string{"trakit config environments", "Manage environments (alias: env)"},
	}

	fmt.Println("Manage trakit-cli configuration")
	fmt.Println("\nUsage:")

	output.PrintBorderlessTable(data)
}

func environmentsHelp() {
	data := [][]string{
		[]string{"trakit config environments list", "List available environments"},
		[]string{"trakit config environments add", "Add an environment"},
	}

	fmt.Println("Manage trakit-cli configuration")
	fmt.Println("\nUsage:")

	output.PrintBorderlessTable(data)

	fmt.Println("\nAliases: environments, env, environment")
}
