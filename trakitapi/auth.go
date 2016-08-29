package trakitapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
)

type LoginResponse struct {
	Key  SessionToken `json:"key"`
	User User         `json:"user"`
}

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(username string, password string) LoginResponse {
	url := "http://api.trakit.party/session_token"

	payload := &LoginPayload{
		Username: username,
		Password: password,
	}
	jsonStr, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	var loginResponse LoginResponse
	err = json.NewDecoder(res.Body).Decode(&loginResponse)
	if err != nil {
		panic(err)
	}

	switch res.StatusCode {
	case 201:
		color.Set(color.FgGreen)
		fmt.Println("\nLogged in as " + username)
		color.Unset()
	case 401, 403:
		fmt.Println("\nUnauthorized")
	}

	if true {
		fmt.Println("response Status:", res.Status)
		fmt.Println("response Headers:", res.Header)
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println("response Body:", string(body))
	}

	return loginResponse
}
