package config

import (
	"encoding/json"
	"errors"
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
)

var configDir string
var ConfigPath string
var SessionPath string

type Environment struct {
	Key    string `json:"key"`
	AppApi string `json:"appApi"`
	WebApp string `json:"webApp"`
}

type Configuration struct {
	Environments    []Environment `json:"environments"`
	EnvironmentKeys []string
}

func init() {
	configDir, err := homedir.Expand("~/.config/trakit")
	if err != nil {
		fmt.Println("Could not initialize configuration")
		os.Exit(1)
	}

	err = os.MkdirAll(configDir, 0644)
	if err != nil {
		fmt.Println("Could not create config directory:", configDir)
		os.Exit(1)
	}

	ConfigPath = configDir + "/config.json"
	SessionPath = configDir + "/.session.json"
}

func GetConfig() Configuration {
	bytes, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		fmt.Println("No configuration file found. Run 'trakit config generate' to create one")
		os.Exit(1)
	}

	var c Configuration
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		panic(err)
	}

	keys := make([]string, len(c.Environments))
	for i, env := range c.Environments {
		keys[i] = env.Key
	}
	c.EnvironmentKeys = keys
	return c
}

func GetEnvironment(key string) (Environment, error) {
	for _, env := range GetConfig().Environments {
		if env.Key == key {
			return env, nil
		}
	}

	return Environment{}, errors.New("No environment with key " + key + ". Make sure your configuration specified an environment with that key.")
}

func GenerateConfigFile() {
	conf := []byte("{\n  \"environments\": []\n}")

	ioutil.WriteFile(ConfigPath, conf, 0644)
}

func AddEnvironment(env Environment) {
	config := GetConfig()
	config.Environments = append(config.Environments, env)
	configJson, _ := json.Marshal(config)
	ioutil.WriteFile(ConfigPath, configJson, 0644)
}
