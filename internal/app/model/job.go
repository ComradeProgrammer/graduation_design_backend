package model

import (
	"encoding/json"
	"fmt"
	"graduation_design/internal/app/config"
	"graduation_design/internal/pkg/logs"
	"graduation_design/internal/pkg/request"
	"strconv"
)

type Job struct {
	Name   string `json:"name"`
	Id     int    `json:"id"`
	Status string `json:"status"`
	WebUrl string `json:"web_url"`
}

func GetJobLog(token string, projectID int, jobID int) (string, error) {
	status, resp, err := request.StringForString(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectID)+"/jobs/"+strconv.Itoa(jobID)+"/trace",
		"GET",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		"",
		5,
	)
	if err != nil {
		logs.Error("getJobLog:Request Failed,%s", err)
		return "", err
	}
	if status != 200 {
		logs.Error("getJobLog Failed,Code %d", status)
		return "", fmt.Errorf("getJobLog Request Failed,Code %d", status)
	}
	return resp, nil
}

func getPipelineJob(token string, projectID int, pipelineID int) ([]Job, error) {
	status, resp, err := request.StringForString(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectID)+"/pipelines/"+strconv.Itoa(pipelineID)+"/jobs",
		"GET",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		"",
		5,
	)
	if err != nil {
		logs.Error("getPipelineJob:Request Failed,%s", err)
		return nil, err
	}
	if status != 200 {
		logs.Error("getPipelineJob Failed,Code %d", status)
		return nil, fmt.Errorf("getPipelineJob Request Failed,Code %d", status)
	}
	var res = make([]Job, 0)
	err = json.Unmarshal([]byte(resp), &res)
	if err != nil {
		return res, err
	}
	return res, nil
}
