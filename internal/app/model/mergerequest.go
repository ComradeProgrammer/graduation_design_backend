package model

import (
	"encoding/json"
	"graduation_design/internal/app/config"
	"graduation_design/internal/pkg/logs"
	"graduation_design/internal/pkg/request"
	"strconv"
)

func GetProjectMrNum(token string,projectId int)(int,error){
	header,_,err:=request.StringForStringWithHeader(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectId)+"/merge_requests?per_page=100",
		"GET",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		"",
		5,
	)
	if err != nil {
		logs.Error("GetIssueStatistic failed,%s", err)
		return -1,err
	}
	total,err:=strconv.Atoi(header.Get("x-total"))
	return total,nil
}

func GetAllProjectMrInObject(token string, projectId int) ([]map[string]interface{}, error){
	resp, err := request.StringForStringWithPagination(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectId)+"/merge_requests?per_page=100",
		"GET",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		"",
		5,
	)
	if err != nil {
		logs.Error("GetIssue:Request Failed,%s", err)
		return nil, err
	}
	var ret=make([]map[string]interface{},0)
	json.Unmarshal([]byte(resp),&ret)
	return ret, nil
}