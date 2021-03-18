package main

import (
	"absensi/source/config"
	"absensi/source/tools"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/prprprus/scheduler"
	log "github.com/sirupsen/logrus"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	config.LoadEnv()
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

type TokenResult struct {
	Token       string   `json:"token"`
	Permissions []string `json:"permissions"`
}

type Daftar struct {
	gorm.Model
	WorkingDate       string    `gorm:"column:working_date;uniqueIndex"`
	LoginDateTime     time.Time `gorm:"column:login_time"`
	LogoutDateTime    time.Time `gorm:"column:logout_time"`
	LoginExecuteTime  time.Time `gorm:"column:login_execute_time"`
	LogoutExecuteTime time.Time `gorm:"column:logout_execute_time"`
}

func insertIntoDB(filename string, daftar *Daftar) error {
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		log.Panic("failed to connect database")
		return err
	}
	result := db.Create(daftar)
	if err := result.Error; err != nil {
		log.Panic("record couldn't be saved")
		return err
	}
	if count := result.RowsAffected; count == 0 {
		log.Panic("record couldn't be saved")
		return err
	}
	return nil
}

func doLoginOrLogoutDB(filename string, isLogin bool) error {
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		log.Panic("failed to connect database")
		return err
	}
	daftar := &Daftar{}
	result := db.Where("working_date=?", time.Now().Format("2006-01-02")).First(&daftar)
	if err = result.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if isLogin {
		daftar.LoginExecuteTime = time.Now()
	} else {
		daftar.LogoutExecuteTime = time.Now()
	}
	db.Save(&daftar)
	return nil
}

func checkAndCreateSchedule(filename string) (*Daftar, error) {
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		log.Panic("failed to connect database")
		return nil, err
	}
	daftar := &Daftar{}
	//db.First(&daftar, "workingDate=?", time.Now().Format("2006-01-02")).Error
	log.Printf("Saved for Data: %s", time.Now().Format("2006-01-02"))
	result := db.Where("working_date=?", time.Now().Format("2006-01-02")).First(&daftar)
	if err = result.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		// Create another record
		//log.Panicf("Create new record. %s", err.Error())
		loginTime, logoutTime := getDateLoginLogout()
		daftar = &Daftar{
			WorkingDate:    time.Now().Format("2006-01-02"),
			LoginDateTime:  loginTime,
			LogoutDateTime: logoutTime,
		}
		insertIntoDB(filename, daftar)
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
func doCron(delay int) {
	log.Printf("do Absensi at %s", time.Now())
	s, err := scheduler.NewScheduler(1000)
	if err != nil {
		log.Println("Error when doing DoCron. Message: " + err.Error())
	}
	s.Delay().Second(delay).Do(doAbsensi)
}

func copyToAbsensiDirectory(config *config.Config) error {
	home, err := tools.GetHomeDirectory()
	if err != nil {
		log.Panic("Create File Error. Panic and abort the apps")
	}
	folderPath := home + "/" + ".absensi"
	if tools.CheckFileExists(".env") {
		return err
	}
	if tools.CheckFileExists(config.Picture) {
		return err
	}
	length, err := tools.Copy(".env", folderPath+"/.env")
	if length == 0 || err != nil {
		return errors.New("Cant copy file .env")
	}
	length, err = tools.Copy(config.Picture, folderPath+"/"+config.Picture)
	if length == 0 || err != nil {
		return errors.New("Cant copy file " + config.Picture)
	}
	return nil
}

func main() {
	config := config.NewConfig()
	err := copyToAbsensiDirectory(config)
	if err != nil {
		log.Panicf("%s", err.Error())
	}
	doAbsensi(config)
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

func doAbsensi(config *config.Config) {
	format := "2006-01-02 15:04"

	home, err := tools.GetHomeDirectory()
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
	daftar, err := checkAndCreateSchedule(fileName)
	if err != nil {
		log.Panic("Create File Error. Panic and abort the apps")
	}
	log.Printf("Login: %s,Logout: %s,Now: %s",
		daftar.LoginDateTime.Format(format),
		daftar.LogoutDateTime.Format(format),
		time.Now().Format(format))
	if daftar.LoginDateTime.Format(format) == time.Now().Format(format) {
		token := Login(config)
		doLoginOrLogout(config, token, true)
		doLoginOrLogoutDB(fileName, true)
	}
	if daftar.LogoutDateTime.Format(format) == time.Now().Format(format) {
		token := Login(config)
		doLoginOrLogout(config, token, false)
		doLoginOrLogoutDB(fileName, false)
	}
	doCron(59)
}

func doLoginOrLogout(config *config.Config, token string, isLogin bool) {
	url := ""
	if isLogin {
		url = config.BaseURL + "presences/punchIn"
	} else {
		url = config.BaseURL + "presences/punchOut"
	}
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("location_latitude", config.Lattitude)
	_ = writer.WriteField("location_longitude", config.Longitude)
	_ = writer.WriteField("location_description", config.Description)
	_ = writer.WriteField("location_timezone", config.Region)
	home, err := tools.GetHomeDirectory()
	if err != nil {
		log.Panic("Create File Error. Panic and abort the apps")
	}
	folderPath := home + "/" + ".absensi"
	file, errFile5 := os.Open(folderPath + "/" + config.Picture)
	defer file.Close()
	part5, errFile5 := writer.CreateFormFile("photo", filepath.Base(folderPath+"/"+config.Picture))
	_, errFile5 = io.Copy(part5, file)
	if errFile5 != nil {
		fmt.Println(errFile5)
		return
	}
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+token)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func getDateLoginLogout() (time.Time, time.Time) {
	year, month, day := time.Now().Date()
	loginTime := time.Date(year, month, day, 7, 0, 0, 0, time.Now().Location())
	logoutTime := time.Date(year, month, day, 16, 30, 0, 0, time.Now().Location())
	randomLoginMinutes := returnRandom(0, 29)
	randomLogoutMinutes := returnRandom(0, 29)
	randomLoginShift := returnRandom(0, 2)
	startRandomLogoutShift := 0 + randomLoginShift
	if startRandomLogoutShift >= 2 {
		startRandomLogoutShift = 2
	}
	randomLogoutShift := returnRandom(startRandomLogoutShift, 2)
	fullLoginTime := loginTime.Add(time.Minute * time.Duration(randomLoginMinutes+(randomLoginShift*30)))
	fullLogoutTime := logoutTime.Add(time.Minute * time.Duration(randomLogoutMinutes+(randomLogoutShift*30)))
	log.Printf("Login: %s,Logout: %s", fullLoginTime, fullLogoutTime)
	return fullLoginTime, fullLogoutTime
}

func returnRandom(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(max-min+1) + min
	// fmt.Println("Random %i", x)
	log.Debugf("Random %d", x)
	return x
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
