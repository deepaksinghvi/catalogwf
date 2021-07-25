package utils

import (
	"bytes"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

const (
	postAction = "POST"
	getAction = "GET"
	contentType = "Content-Type"
)

func PostRequest(url string, payload *bytes.Buffer, contentTypeString string, client http.Client) (*http.Response, error) {
	method := postAction
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		log.Error(err)
	}
	req.Header.Set(contentType, contentTypeString)
	res, err := client.Do(req)
	if nil != err {
		return nil, errors.New("request error")
	}
	return res,nil;
}

func GetRequest(url string, client http.Client) ([]byte, error) {
	method := getAction
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		log.Error(err)
	}
	res, err := client.Do(req)
	if nil != err {
		return nil,err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}