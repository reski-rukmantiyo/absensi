package main

import (
	"absensi/source/config"
	"absensi/source/tools"
	"log"
	"net/url"
)

func init() {
	config.LoadEnv()
}

type TokenResult struct {
	Token       string   `json:"token"`
	Permissions []string `json:"permissions"`
}

func main() {
	config := config.NewConfig()
	token := Login(config)
	log.Println(token)
}

func Login(config *config.Config) string {
	var params = url.Values{}
	params.Add("username", config.Username)
	params.Add("password", config.Password)
	requestResult := tools.PostForm(config.BaseURL+"login", params)
	if requestResult == "" {
		log.Fatal("Error on Request Response")
	}
	tokenResult := &TokenResult{}
	tools.FromJSON(requestResult, tokenResult)
	return tokenResult.Token
}
