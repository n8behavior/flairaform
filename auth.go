package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Auth struct {
	Token   string `json:"access_token"`
	Expires int    `json:"expires_in"`
	Type    string `json:"token_type"`
	Scope   string `json:"scope"`
}

func Authorize() (*Auth, error) {
	// TODO externalize
	endpoint, err := url.Parse("https://api.flair.co/oauth/token")
	if err != nil {
		return nil, fmt.Errorf("Endpoint URL: %s", err)
	}
	q := endpoint.Query()
	q.Set("client_id", os.Getenv("CLIENT_ID"))
	q.Set("client_secret", os.Getenv("CLIENT_SECRET"))
	q.Set("scope", "thermostats.view")
	q.Set("grant_type", "client_credentials")
	endpoint.RawQuery = q.Encode()

	fmt.Println(endpoint)

	method := "POST"
	client := &http.Client{}
	req, err := http.NewRequest(method, endpoint.String(), nil)

	if err != nil {
		return nil, fmt.Errorf("Request: %s", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	auth := new(Auth)
	json.Unmarshal(body, &auth)

	return auth, nil
}
