package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"graduation_design/internal/pkg/logs"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// request body is json,response is also json
func JsonForJson(urlStr string, method string, header map[string]string, content map[string]interface{}, timeout int) (int, map[string]interface{}, error) {

	bodyJson, err := json.Marshal(content)
	if err != nil {
		logs.Fatal("JsonForJson:marshal body failure,%v", err)
		return -1, nil, err
	}
	req, err := http.NewRequest(method, urlStr, bytes.NewBuffer(bodyJson))
	if err != nil {
		logs.Fatal("JsonForJson:NewRequest failure,%v", err)
		return -1, nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type","application/json")
	logs.Info("Send Request url:%s,method %s,header %v,body %v", urlStr, method, header,string(bodyJson))
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logs.Error("JsonForJson:get response failure,%v", err)
		return -1, nil, err
	}

	var data map[string]interface{}
	dataJsonRecv, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	logs.Info("JsonForJson:response received %s", string(dataJsonRecv))

	err = json.Unmarshal(dataJsonRecv, &data)
	if err != nil {
		logs.Error("FormForJson:response not json,%s", string(dataJsonRecv))
		return -1,data,fmt.Errorf("FormForJson:response not json,%s", string(dataJsonRecv))
	}
	return resp.StatusCode, data, nil
}

// request body is string,response is json
func StringForJson(urlStr string, method string, header map[string]string, content string, timeout int) (int, map[string]interface{}, error) {
	logs.Info("Send Request url:%s,method %s,header %v,body %s", urlStr, method, header, content)

	req, err := http.NewRequest(method, urlStr, bytes.NewBuffer([]byte(content)))
	if err != nil {
		logs.Fatal("StringForJson:NewRequest failure,%v", err)
		return -1, nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logs.Error("StringForJson:get response failure,%v", err)
		return -1, nil, err
	}

	var data map[string]interface{}
	dataJsonRecv, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	logs.Info("StringForJson:response received %s", string(dataJsonRecv))

	err = json.Unmarshal(dataJsonRecv, &data)
	if err != nil {
		logs.Error("StringForJson:response not json,%s", string(dataJsonRecv))
		return -1,data,fmt.Errorf("StringForJson:response not json,%s", string(dataJsonRecv))
	}
	return resp.StatusCode, data, nil
}

// request body is string,response is also string
func StringForString(urlStr string, method string, header map[string]string, content string, timeout int) (int, string, error) {
	logs.Info("Send Request url:%s,method %s,header %v,body %s", urlStr, method, header, content)

	req, err := http.NewRequest(method, urlStr, bytes.NewBuffer([]byte(content)))
	if err != nil {
		logs.Fatal("StringForString:NewRequest failure,%v", err)
		return -1, "", err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logs.Error("StringForString:get response failure,%v", err)
		return -1, "", err
	}
	
	dataJsonRecv, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	logs.Info("StringForString:response received %s", string(dataJsonRecv))

	return resp.StatusCode, string(dataJsonRecv), nil
}

// request body is post for,response is also json
func FormForJson(urlStr string, method string, header map[string]string, content map[string]string, timeout int) (int, map[string]interface{}, error) {

	body:=url.Values{}
	for k,v:=range content{
		body.Set(k,v)
	}
	req, err := http.NewRequest(method, urlStr, bytes.NewBuffer([]byte(body.Encode())))
	if err != nil {
		logs.Fatal("JsonForJson:NewRequest failure,%v", err)
		return -1, nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	logs.Info("Send Request url:%s,method %s,header %v,body %v", urlStr, method, header,body.Encode())
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logs.Error("FormForJson:get response failure,%v", err)
		return -1, nil, err
	}

	var data map[string]interface{}
	dataJsonRecv, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	logs.Info("FormForJson:response received %s", string(dataJsonRecv))

	err = json.Unmarshal(dataJsonRecv, &data)
	if err != nil {
		logs.Error("FormForJson:response not json,%s", string(dataJsonRecv))
		return -1,data,fmt.Errorf("FormForJson:response not json,%s", string(dataJsonRecv))
	}
	return resp.StatusCode, data, nil
}
// request body is string,response is also string
func StringForStringWithHeader(urlStr string, method string, header map[string]string, content string, timeout int) (http.Header, string, error) {
	logs.Info("Send Request url:%s,method %s,header %v,body %s", urlStr, method, header, content)

	req, err := http.NewRequest(method, urlStr, bytes.NewBuffer([]byte(content)))
	if err != nil {
		logs.Fatal("StringForString:NewRequest failure,%v", err)
		return nil, "", err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logs.Error("StringForString:get response failure,%v", err)
		return nil, "", err
	}
	
	dataJsonRecv, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	logs.Info("StringForString:response received %s", string(dataJsonRecv))

	return resp.Header, string(dataJsonRecv), nil
}
