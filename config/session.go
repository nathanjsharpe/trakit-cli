package config

import (
	"encoding/json"
	"fmt"
	"github.com/nathanjsharpe/trakit/trakitapi"
	"io/ioutil"
	"os"
)

type Session struct {
	Environment string                 `json:"environment"`
	Token       trakitapi.SessionToken `json:"token"`
	User        trakitapi.User         `json:"user"`
}

func (session *Session) validate() (bool, string) {
	if session.Token.Token == "" {
		return false, "No current session. Log in to create a session."
	}

	if session.Token.IsExpired() {
		return false, "Session is expired."
	}

	return true, ""
}

func (session *Session) Save() bool {
	sessionJson, _ := json.Marshal(session)
	ioutil.WriteFile(SessionPath, sessionJson, 0644)
	return true
}

func LoadSession() Session {
	raw, err := ioutil.ReadFile(SessionPath)
	if err != nil {
		fmt.Println("No current session. ")
		os.Exit(1)
	}

	var s Session
	json.Unmarshal(raw, &s)

	currentEnv, err := GetEnvironment(s.Environment)
	if err != nil {
		exitForRelogin("Environment not found: " + s.Environment)
	}
	trakitapi.SetAppApiUrl(currentEnv.AppApi)
	trakitapi.SetSessionToken(s.Token.Token)
	return s
}

func ValidateSession() {
	session := LoadSession()
	valid, message := session.validate()
	if !valid {
		exitForRelogin(message)
	}
}

func DeleteSession() error {
	return os.Remove(SessionPath)
}

func exitForRelogin(message string) {
	fmt.Println(message, "Log in and try again.")
	os.Exit(1)
}
