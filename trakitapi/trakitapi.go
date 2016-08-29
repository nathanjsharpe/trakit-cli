package trakitapi

import (
	"fmt"
	"net/http"
)

var t *TrakitApi

func init() {
	t = New()
}

type TrakitApi struct {
	appApiUrl    string
	sessionToken string
	verbose      bool
}

func New() *TrakitApi {
	t := new(TrakitApi)
	t.verbose = true
	t.sessionToken = ""

	return t
}

func SetAppApiUrl(in string) { t.SetAppApiUrl(in) }
func (t *TrakitApi) SetAppApiUrl(in string) {
	t.appApiUrl = in
}

func SetSessionToken(in string) { t.SetSessionToken(in) }
func (t *TrakitApi) SetSessionToken(in string) {
	t.sessionToken = in
}

func SetVerbose(in bool) { t.SetVerbose(in) }
func (t *TrakitApi) SetVerbose(in bool) {
	t.verbose = in
}

func Get(path string) http.Response {
	url := t.appApiUrl + "/" + path

	req, err := http.NewRequest("GET", url, nil)

	if t.verbose {
		fmt.Println("Sending request to " + url + " with token " + t.sessionToken)
	}
	req.Header.Set("Content-Type", "application/json")
	if t.sessionToken != "" {
		req.Header.Set("X-Auth-Token", t.sessionToken)
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	if t.verbose {
		fmt.Println("response Status:", res.Status)
		fmt.Println("response Headers:", res.Header)
		// body, _ := ioutil.ReadAll(res.Body)
		// fmt.Println("response Body:", string(body))
	}

	return *res
}
