package model

import (
	"fmt"
	"graduation_design/internal/app/config"
	"graduation_design/internal/pkg/logs"
	"graduation_design/internal/pkg/request"
	"strconv"
)

func GetMasterBranch(token string, projectID int) (map[string]interface{}, error) {
	status, resp, err := request.StringForJson(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectID)+"/repository/branches/"+config.MAINBRANCH,
		"GET",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		"",
		5,
	)
	if err != nil {
		logs.Error("GetMasterBranch:Request Failed,%s", err)
		return nil, err
	}
	if status != 200 {
		logs.Error(" GetMasterBranch Failed,Code %d", status)
		return nil, fmt.Errorf(" GetMasterBranchRequest Failed,Code %d", status)
	}
	logs.Info("GetMasterBranch success,response %s", resp)
	return resp, nil
}
