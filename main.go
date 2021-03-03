package main

import (
	"absensi/config"
	"absensi/tools"
	"fmt"
	"log"
	"net/url"
)

func init() {
	config.LoadEnv()
}

type TokenResult struct {
	token       string
	permissions map[string]string
}

func main() {
	config := config.NewConfig()
	// var params = url.Values{}
	// username := config.Username
	// password := config.Password
	// loginUrl := config.BaseURL + "login"
	//fmt.Println("pointer:", &config)
	var params = url.Values{}
	params.Add("username", config.Username)
	params.Add("password", config.Password)
	requestResult := tools.PostForm(config.BaseURL+"login", params)
	if requestResult == "" {
		log.Fatal("Error on Request Response")
	}
	// fmt.Println("pointer:", config.Username)
	fmt.Println(requestResult)

}
