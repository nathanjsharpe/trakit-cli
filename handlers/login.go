package handlers

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/nathanjsharpe/trakit/config"
	"github.com/nathanjsharpe/trakit/trakitapi"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"syscall"
)

func Login(args []string) {
	if len(args) < 1 {
		showLoginHelp()
		os.Exit(1)
	}

	environment, err := config.GetEnvironment(args[0])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	username, password := getCredentials()
	loginResponse := trakitapi.Login(username, password)
	session := config.Session{
		Environment: environment.Key,
		User:        loginResponse.User,
		Token:       loginResponse.Key,
	}
	session.Save()
}

func getCredentials() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	color.Set(color.FgWhite, color.Bold)
	fmt.Print("Username: ")
	color.Unset()

	username, _ := reader.ReadString('\n')

	color.Set(color.FgWhite, color.Bold)
	fmt.Print("Password: ")
	color.Unset()

	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}
	password := string(bytePassword)

	return strings.TrimSpace(username), strings.TrimSpace(password)
}

func showLoginHelp() {
	fmt.Println("Usage: trakit login <environment>")
	fmt.Println("Environment must be one of:", config.GetConfig().EnvironmentKeys)
}
