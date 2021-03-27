package model

import (
	"graduation_design/internal/app/config"
	"graduation_design/internal/pkg/logs"
	"graduation_design/internal/pkg/request"
	"strconv"
)

func CheckProjectAuthorization(token string, id int) (bool, error) {

	status, _, err := request.StringForString(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(id),
		"GET",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		"",
		5,
	)
	if err != nil {
		logs.Error("getProjectFromGitlab:Request Failed,%s", err)
		return false, err
	}
	if status != 200 {
		logs.Error("getProjectFromGitlab Request Failed,Code %d", status)
		return false, nil
	}
	return true, nil
}
