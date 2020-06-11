package http_client

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"net/http"
)

type CallRestAPIOpt func(*http.Request)

func CallRestAPI(httpClient *http.Client, url, method string, body []byte, opt ...CallRestAPIOpt) (int, http.Header, []byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return 0, map[string][]string{}, []byte{}, err
	}

	for _, o := range opt {
		o(req)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return 0, map[string][]string{}, []byte{}, err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, map[string][]string{}, []byte{}, err
	}
	resHeader := res.Header
	return res.StatusCode, resHeader, resBody, nil
}

func WithBaiscTokenHeader(username, password string) CallRestAPIOpt {
	return func(req *http.Request) {
		req.Header.Add("Authorization", "Basic "+basicAuth(username, password))
	}
}

func WithBearTokenHeader(key string) CallRestAPIOpt {
	return func(req *http.Request) {
		req.Header.Add("Authorization", "Bearer " + key)
	}
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
