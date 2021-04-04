package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"graduation_design/internal/pkg/logs"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func StringForStringWithPagination(urlStr string, method string, header map[string]string, content string, timeout int) (string, error) {
	var allObj = make([]map[string]interface{}, 0)
	res, err := RequestForGitlabPagination(urlStr, method, header, content, timeout)
	if err != nil {
		return "", err
	}
	for _, item := range res {
		var objList = make([]map[string]interface{}, 0)
		json.Unmarshal([]byte(item), &objList)
		allObj = append(allObj, objList...)
	}
	ret, _ := json.Marshal(allObj)
	return string(ret), nil
}

func RequestForGitlabPagination(url string, method string, header map[string]string, content string, timeout int) ([]string, error) {
	var ret = make([]string, 0)
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	var first = true
	var pageNum = 0
	var currentPage = 1
	var nextUrl = ""
	for {
		if first {
			first = false
			logs.Info("Send Request url:%s,method %s,header %v,body %s", url, method, header, content)
			req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(content)))
			if err != nil {
				logs.Fatal("RequestForGitlabPagination:NewRequest failure,%v", err)
				return nil, err
			}
			for k, v := range header {
				req.Header.Set(k, v)
			}
			resp, err := client.Do(req)
			if err != nil {
				logs.Fatal("RequestForGitlabPagination:Requestfailure,%v", err)
				return nil, err
			}

			dataJsonRecv, err := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			logs.Info("RequestForGitlabPagination:turn 0,response received %s", string(dataJsonRecv))

			pageNum, err = strconv.Atoi(resp.Header.Get("x-total-pages"))
			if err != nil {
				return nil, fmt.Errorf("x-total-pages not found:%s ", err.Error())
			}
			ret = append(ret, string(dataJsonRecv))
			logs.Info("link in header:%s", resp.Header.Get("link"))
			nextUrl = ""
			links := strings.Split(resp.Header.Get("link"), ",")
			regex := `\s*\<(.*)\>;\s*rel="next"`
			reg := regexp.MustCompile(regex)
			for _, link := range links {
				groups := reg.FindStringSubmatch(link)
				if groups == nil {
					continue
				}
				nextUrl = groups[1]

			}
			if nextUrl == "" && currentPage < pageNum {
				logs.Error("next url not match,%s", links)
				break
			}
			currentPage++

		} else {
			if currentPage > pageNum {
				break
			}
			urlStr := nextUrl
			logs.Info("Send Request url:%s,method %s,header %v,body %s", urlStr, method, header, content)
			req, err := http.NewRequest(method, urlStr, bytes.NewBuffer([]byte(content)))
			if err != nil {
				logs.Fatal("RequestForGitlabPagination:NewRequest failure,%v", err)
				return nil, err
			}
			for k, v := range header {
				req.Header.Set(k, v)
			}
			resp, err := client.Do(req)
			if err != nil {
				logs.Fatal("RequestForGitlabPagination:Requestfailure,%v", err)
				return nil, err
			}

			dataJsonRecv, err := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			logs.Info("RequestForGitlabPagination:turn %d,response received %s", currentPage, string(dataJsonRecv))
			ret = append(ret, string(dataJsonRecv))

			logs.Info("link in header:%s", resp.Header.Get("link"))
			nextUrl = ""
			links := strings.Split(resp.Header.Get("link"), ",")
			regex := `\s*\<(.*)\>;\s*rel="next"`
			reg := regexp.MustCompile(regex)
			for _, link := range links {
				groups := reg.FindStringSubmatch(link)
				if groups == nil {
					continue
				}
				nextUrl = groups[1]

			}
			if nextUrl == "" && currentPage < pageNum {
				logs.Error("next url not match,%s", links)
				break
			}
			currentPage++

		}

	}
	return ret, nil
}
