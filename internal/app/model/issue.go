package model

import (
	"fmt"
	"graduation_design/internal/app/config"
	"graduation_design/internal/pkg/logs"
	"graduation_design/internal/pkg/request"
	"strconv"
)

// todo: remove restriction of pagination in gitlab source code

func CreateIssue(token string,projectID int,milestoneId int,
	title string,description string,startDate,dueDate string,
	typeTag,priorityTag string)error{
	status,resp,err:=request.FormForJson(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectID)+"/issues",
		"POST",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		map[string]string{
			"title":title,
			"description":description,
			"start_date":startDate,
			"due_date":dueDate,
			"labels":typeTag+","+priorityTag,
			"milestone_id":strconv.Itoa(milestoneId),
		},
		5,
	)	
	if err != nil {
		logs.Error(" CreateIssue:Request Failed,%s", err)
		return err
	}
	if status != 201 {
		logs.Error(" CreateIssue Failed,Code %d", status)
		return  fmt.Errorf(" CreateIssue Request Failed,Code %d", status)
	}
	logs.Info(" CreateIssue success,response %s",resp)
	return nil
}

func GetAllIssues(token string,projectId int,milestoneId int)(string,error){
	status,resp,err:=request.StringForString(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectId)+"/milestones/"+strconv.Itoa(milestoneId)+"/issues",
		"GET",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		"",
		5,
	
	)
	if err != nil {
		logs.Error("GetAllIssue:Request Failed,%s", err)
		return "",err
	}
	if status != 200 {
		logs.Error("GetAllIssue Failed,Code %d", status)
		return "",fmt.Errorf("GetAllIssue Request Failed,Code %d", status)
	}
	return resp,nil
}

