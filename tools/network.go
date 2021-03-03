package tools

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Get : Get values thru HTTP Get
func Get(url string) string {
	log.Debugf("Request: %s\n", url)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Debugf("Response: " + string(content))
	return string(content)
}

// PostForm : PostForm values thru HTTP PostForm
func PostForm(url string, values url.Values) string {
	log.Debugf("Request: %s\n", url)
	client := &http.Client{}
	r, err := http.NewRequest("POST", url, strings.NewReader(values.Encode())) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(values.Encode())))

	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Debugf("Response: " + string(body))
	return string(body)
}
