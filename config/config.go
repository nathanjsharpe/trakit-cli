package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

const configDir = "/home/nathan/.config/trakit"

const ConfigPath = configDir + "/config.json"
const SessionPath = configDir + "/.session.json"

type Environment struct {
	Key    string `json:"key"`
	AppApi string `json:"appApi"`
	WebApp string `json:"webApp"`
}

type Configuration struct {
	Environments    []Environment `json:"environments"`
	EnvironmentKeys []string
}

func GetConfig() Configuration {
	bytes, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		panic(err)
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
