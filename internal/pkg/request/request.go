package request

import (
	"bytes"
	"encoding/json"
	"graduation_design/internal/pkg/logs"
	"io/ioutil"
	"net/http"
	"time"
)
// request body is json,response is also json
func JsonForJson(url string, method string, header map[string]string, content map[string]interface{}, timeout int) (int, map[string]interface{}, error) {
	logs.Info("Send Request url:%s,method %s,header %v,body %v", url, method, header, content)

	bodyJson, err := json.Marshal(content)
	if err != nil {
		logs.Fatal("JsonForJson:marshal body failure,%v", err)
		return -1, nil, err
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyJson))
	if err != nil {
		logs.Fatal("JsonForJson:NewRequest failure,%v", err)
		return -1, nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logs.Error("JsonForJson:get response failure,%v", err)
		return -1, nil, err
	}
	if resp.StatusCode != 200 {
		logs.Error("JsonForJson:response code %d", resp.StatusCode)
		return resp.StatusCode, nil, nil
	}
	var data map[string]interface{}
	dataJsonRecv, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	logs.Info("JsonForJson:response received %s", string(dataJsonRecv))

	err = json.Unmarshal(dataJsonRecv, &data)
	if err != nil {
		logs.Error("JsonForJson:response not json,%s", string(dataJsonRecv))
	}
	return 200, data, nil
}

// request body is string,response is json
func StringForJson(url string, method string, header map[string]string, content string, timeout int) (int, map[string]interface{}, error) {
	logs.Info("Send Request url:%s,method %s,header %v,body %s", url, method, header, content)

	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(content)))
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
	if resp.StatusCode != 200 {
		logs.Error("StringForJson:response code %d", resp.StatusCode)
		return resp.StatusCode, nil, nil
	}
	var data map[string]interface{}
	dataJsonRecv, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	logs.Info("StringForJson:response received %s", string(dataJsonRecv))

	err = json.Unmarshal(dataJsonRecv, &data)
	if err != nil {
		logs.Error("StringForJson:response not json,%s", string(dataJsonRecv))
	}
	return 200, data, nil
}

// request body is string,response is also json
func StringForString(url string, method string, header map[string]string, content string, timeout int) (int, string, error) {
	logs.Info("Send Request url:%s,method %s,header %v,body %s", url, method, header, content)

	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(content)))
	if err != nil {
		logs.Fatal("StringForJson:NewRequest failure,%v", err)
		return -1, "", err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logs.Error("StringForJson:get response failure,%v", err)
		return -1, "", err
	}
	if resp.StatusCode != 200 {
		logs.Error("StringForJson:response code %d", resp.StatusCode)
		return resp.StatusCode, "", nil
	}
	
	dataJsonRecv, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	logs.Info("StringForJson:response received %s", string(dataJsonRecv))

	return 200, string(dataJsonRecv), nil
}