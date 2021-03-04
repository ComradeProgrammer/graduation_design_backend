package model

import (
	"encoding/json"
	"fmt"
	"graduation_design/internal/app/config"
	"graduation_design/internal/pkg/logs"
	"graduation_design/internal/pkg/request"
	"strconv"
)

type Pipeline struct {
	ID        int    `json:"id"`
	Status    string `json:"status"`
	WebUrl    string `json:"web_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func getAllPipeline(token string, projectID int) ([]Pipeline, error) {
	status, resp, err := request.StringForString(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectID)+"/pipelines",
		"GET",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		"",
		5,
	)
	if err != nil {
		logs.Error("getAllPipeline:Request Failed,%s", err)
		return nil, err
	}
	if status != 200 {
		logs.Error("getAllPipeline Failed,Code %d", status)
		return nil, fmt.Errorf("getAllPipeline Request Failed,Code %d", status)
	}
	var res = make([]Pipeline, 0)
	err = json.Unmarshal([]byte(resp), &resp)
	if err != nil {
		return res, err
	}
	return res, nil
}
func getLatestFinishedPipelineOnMasterBranch(token string, projectID int) (int, string, error) {
	resp, err := GetMasterBranch(token, projectID)
	if err != nil {
		return -1, "", err
	}
	commit, ok := resp["commit"].(map[string]interface{})
	if !ok {
		return -1, "", fmt.Errorf("doesn't contain commit")
	}
	sha, ok := commit["short_id"].(string)
	if !ok {
		return -1, "", fmt.Errorf("doesn't contain sha")
	}
	resp, err = GetCommit(token, projectID, sha)
	if err != nil {
		return -1, "", fmt.Errorf("failed to get commit ")
	}

	lastPipeline, ok := resp["last_pipeline"].(map[string]interface{})
	if !ok {
		return -1, "", fmt.Errorf("doesn't contain last_pipeline")
	}
	pipelineIdFloat, ok := lastPipeline["id"].(float64)
	if !ok {
		return -1, "", fmt.Errorf("doesn't contain last_pipeline id")
	}
	pipelineId := int(pipelineIdFloat)
	status, ok := lastPipeline["status"].(string)
	if !ok {
		return -1, "", fmt.Errorf("doesn't contain last_pipeline status")
	}

	return pipelineId, status, nil
}
