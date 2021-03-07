package main

import (
	"absensi/config"
	"absensi/tools"
	"errors"
	"log"
	"net/url"
	"os"
	"time"

	"gorm.io/driver/sqlite"
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
	WorkingDate       time.Time `gorm:"column:working_date;uniqueIndex"`
	LoginDateTime     time.Time `gorm:"column:login_date"`
	LogoutDateTime    time.Time `gorm:"column:logout_date"`
	LoginExecuteTime  time.Time
	LogoutExecuteTime time.Time
}

func checkAndCreateSchedule(filePath string) (*Daftar, error) {
	db, err := gorm.Open(sqlite.Open(filePath), &gorm.Config{})
	if err != nil {
		log.Panic("failed to connect database")
		return nil, err
	}
	daftar := &Daftar{}
	//db.First(&daftar, "workingDate=?", time.Now().Format("2006-01-02")).Error
	result := db.Where("working_date=?", time.Now().Format("2006-01-02")).First(&daftar)
	if err = result.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		// Create another record
		//log.Panicf("Create new record. %s", err.Error())
	}
	return daftar, nil
}

func createFile(filePath string) error {
	db, err := gorm.Open(sqlite.Open(filePath), &gorm.Config{})
	if err != nil {
		log.Panic("failed to connect database")
		return err
	}
	db.AutoMigrate(&Daftar{})
	return nil
}

func main() {
	// config := config.NewConfig()
	// token := Login(config)
	// log.Println(token)

	home, err := os.UserHomeDir()
	folderPath := home + "/" + ".absensi"
	fileName := folderPath + "/" + "absensi.db"
	_, err = os.Stat(folderPath)
	if err != nil {
		err = os.Mkdir(folderPath, os.ModePerm)
		if err != nil {
			log.Panic("Create File Error. Panic and abort the apps")
		}
	}
	_, err = os.Stat(fileName)
	if err != nil {
		err = createFile(fileName)
		if err != nil {
			log.Panic("Create File Error. Panic and abort the apps")
		}
	}
	checkAndCreateSchedule(fileName)
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
