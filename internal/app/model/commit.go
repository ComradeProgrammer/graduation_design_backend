package model

import (
	"fmt"
	"graduation_design/internal/app/config"
	"graduation_design/internal/pkg/logs"
	"graduation_design/internal/pkg/request"
	"strconv"
)

func GetCommit(token string, projectID int, sha string) (map[string]interface{}, error) {
	status, resp, err := request.StringForJson(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectID)+"/repository/commits/"+sha,
		"GET",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		"",
		5,
	)
	if err != nil {
		logs.Error("GetCommit:Request Failed,%s", err)
		return nil, err
	}
	if status != 200 {
		logs.Error("GetCommit Failed,Code %d", status)
		return nil, fmt.Errorf("GetCommit Request Failed,Code %d", status)
	}
	logs.Info("GetCommit success,response %s", resp)
	return resp, nil
}
