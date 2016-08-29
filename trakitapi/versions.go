package trakitapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

const AppApiRegex = `trakitng_app_api-([^-]+)`
const DataApiRegex = `trakitng_data_api-([^-]+)`
const AuthApiRegex = `trakitng_auth_api-([^-]+)`
const EventApiRegex = `trakitng_event_api-([^-]+)`
const WebAppRegex = `Version:\D+(\d+\.\d+\.\d+)`

type versionResponse struct {
	App   string `json:"app-api"`
	Data  string `json:"data-api"`
	Auth  string `json:"auth-api"`
	Event string `json:"event-api"`
}

type Versions struct {
	App   string
	Data  string
	Auth  string
	Event string
	Web   string
}

func GetVersions() Versions {
	res := Get("version/all")

	defer res.Body.Close()

	var apiVersions versionResponse
	err := json.NewDecoder(res.Body).Decode(&apiVersions)
	if err != nil {
		panic(err)
	}

	sanitizedVersions := sanitizeApiVersions(apiVersions)
	webAppVersion := getWebAppVersion()

	versions := Versions{
		App:   sanitizedVersions.App,
		Data:  sanitizedVersions.Data,
		Auth:  sanitizedVersions.Auth,
		Event: sanitizedVersions.Event,
		Web:   webAppVersion,
	}

	return versions
}

func sanitizeApiVersions(rawVersions versionResponse) versionResponse {
	appApiRegex, err := regexp.Compile(AppApiRegex)
	if err != nil {
		panic(err)
	}

	dataApiRegex, err := regexp.Compile(DataApiRegex)
	if err != nil {
		panic(err)
	}

	authApiRegex, err := regexp.Compile(AuthApiRegex)
	if err != nil {
		panic(err)
	}

	eventApiRegex, err := regexp.Compile(EventApiRegex)
	if err != nil {
		panic(err)
	}

	return versionResponse{
		App:   appApiRegex.FindStringSubmatch(rawVersions.App)[1],
		Data:  dataApiRegex.FindStringSubmatch(rawVersions.Data)[1],
		Auth:  authApiRegex.FindStringSubmatch(rawVersions.Auth)[1],
		Event: eventApiRegex.FindStringSubmatch(rawVersions.Event)[1],
	}
}

func getWebAppVersion() string {
	res, err := http.Get("http://app.trakit.party/version")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	webAppRegex, err := regexp.Compile(WebAppRegex)
	if err != nil {
		panic(err)
	}

	return webAppRegex.FindStringSubmatch(string(body))[1]
}
