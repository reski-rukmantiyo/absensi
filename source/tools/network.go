package tools

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/parnurzeal/gorequest"
	log "github.com/sirupsen/logrus"
)

// Get : Get values thru HTTP Get
func Get(url string) string {
	log.Debugf("Request: %s\n", url)
	res, err := http.Get(url)
	if err != nil {
		log.Printf("Error when request. Message: %s", err.Error())
		return ""
	}
	content, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Printf("Error when read body. Message: %s", err.Error())
		return ""
	}
	log.Debugf("Response: " + string(content))
	return string(content)
}

func GetJSON(url string, params ...string) (string, error) {
	var token = params[0]
	request := gorequest.New()
	_, data, errs := request.Get(url).
		Set("Authorization", "Bearer "+token).
		End()
	if len(errs) != 0 {
		return "", errs[0]
	}
	return data, nil
}

func PostFormJSON(url string, obj interface{}) string {
	requestObj, err := ToJSON(obj)
	if err != nil {
		log.Printf("Error when request. Message: %s", err.Error())
		return ""
	}
	log.Debugf("Request: %s\n", url)
	client := &http.Client{}
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(requestObj)) // URL-encoded payload
	if err != nil {
		log.Printf("Error when request. Message: %s", err.Error())
		return ""
	}
	r.Header.Add("Content-Type", "application/json")
	// r.Header.Add("Content-Length", strconv.Itoa(len(requestObj.Encode())))

	res, err := client.Do(r)
	if err != nil {
		log.Printf("Error when get response. Message: %s", err.Error())
		return ""
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error when read body. Message: %s", err.Error())
		return ""
	}
	log.Debugf("Response: " + string(body))
	return string(body)
}

// PostForm : PostForm values thru HTTP PostForm
func PostForm(url string, values url.Values, headers ...map[string]string) string {
	log.Debugf("Request: %s\n", url)
	client := &http.Client{}
	r, err := http.NewRequest("POST", url, strings.NewReader(values.Encode())) // URL-encoded payload
	if err != nil {
		log.Printf("Error when post request. Message: %s", err.Error())
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(values.Encode())))
	if len(headers) != 0 {
		for _, header := range headers {
			for name, value := range header {
				r.Header.Add(name, value)
			}
		}
	}
	res, err := client.Do(r)
	if err != nil {
		log.Printf("Error when get response. Message: %s", err.Error())
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error when read body. Message: %s", err.Error())
	}
	log.Debugf("Response: " + string(body))
	return string(body)
}
