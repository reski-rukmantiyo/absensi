package main

import (
	"absensi/config"
	"absensi/tools"
	"log"
	"net/url"
	"time"

	"gorm.io/gorm"
)

func init() {
	config.LoadEnv()
}

type TokenResult struct {
	Token       string   `json:"token"`
	Permissions []string `json:"permissions"`
}

type Daftar struct {
	gorm.Model
	WorkData          time.Time `gorm:"uniqueIndex"`
	LoginDateTime     time.Time
	LogoutDateTime    time.Time
	LoginExecuteTime  time.Time
	LogoutExecuteTime time.Time
}

func main() {
	// config := config.NewConfig()
	// token := Login(config)
	// log.Println(token)

	// db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	// if err != nil {
	// 	log.Panic("failed to connect database")
	// }
	// db.AutoMigrate(&Daftar{})

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
