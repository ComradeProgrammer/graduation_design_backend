package model

import (
	"encoding/json"
	"fmt"
	"graduation_design/internal/app/config"
	"graduation_design/internal/pkg/logs"
	"graduation_design/internal/pkg/request"
	"strconv"
)

func CreateIssue(token string, projectID int, milestoneId int,
	title string, description string, startDate, dueDate string,
	typeTag, priorityTag string) error {
	status, resp, err := request.FormForJson(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectID)+"/issues",
		"POST",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		map[string]string{
			"title":        title,
			"description":  description,
			"start_date":   startDate,
			"due_date":     dueDate,
			"labels":       typeTag + "," + priorityTag,
			"milestone_id": strconv.Itoa(milestoneId),
		},
		5,
	)
	if err != nil {
		logs.Error(" CreateIssue:Request Failed,%s", err)
		return err
	}
	if status != 201 {
		logs.Error(" CreateIssue Failed,Code %d", status)
		return fmt.Errorf(" CreateIssue Request Failed,Code %d", status)
	}
	logs.Info(" CreateIssue success,response %s", resp)
	return nil
}

func GetAllIssuesForMilestone(token string, projectId int, milestoneId int) (string, error) {
	status, resp, err := request.StringForString(
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
		return "", err
	}
	if status != 200 {
		logs.Error("GetAllIssue Failed,Code %d", status)
		return "", fmt.Errorf("GetAllIssue Request Failed,Code %d", status)
	}
	return resp, nil
}

func GetIssue(token string, projectId int, IssueIid int) (string, error) {
	status, resp, err := request.StringForString(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectId)+"/issues/"+strconv.Itoa(IssueIid),
		"GET",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		"",
		5,
	)
	if err != nil {
		logs.Error("GetIssue:Request Failed,%s", err)
		return "", err
	}
	if status != 200 {
		logs.Error("GetIssue Failed,Code %d", status)
		return "", fmt.Errorf("GetIssue Request Failed,Code %d", status)
	}
	return resp, nil
}

func EditIssue(token string, projectID int, milestoneId int, issueIid int,
	title string, description string, startDate, dueDate string,
	typeTag, priorityTag string) error {
	status, resp, err := request.FormForJson(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectID)+"/issues/"+strconv.Itoa(issueIid),
		"PUT",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		map[string]string{
			"title":        title,
			"description":  description,
			"start_date":   startDate,
			"due_date":     dueDate,
			"labels":       typeTag + "," + priorityTag,
			"milestone_id": strconv.Itoa(milestoneId),
		},
		5,
	)
	if err != nil {
		logs.Error("EditIssue:Request Failed,%s", err)
		return err
	}
	if status != 200 {
		logs.Error(" EditIssue Failed,Code %d", status)
		return fmt.Errorf(" EditIssue Request Failed,Code %d", status)
	}
	logs.Info(" EditIssue success,response %s", resp)
	return nil
}

func ChangeIssueState(token string, projectID int, issueIid int, stateEvent string) error {
	status, resp, err := request.FormForJson(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectID)+"/issues/"+strconv.Itoa(issueIid),
		"PUT",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		map[string]string{
			"state_event": stateEvent,
		},
		5,
	)
	if err != nil {
		logs.Error("ChangeIssueState:Request Failed,%s", err)
		return err
	}
	if status != 200 {
		logs.Error(" ChangeIssueState Failed,Code %d", status)
		return fmt.Errorf(" EditIssue Request Failed,Code %d", status)
	}
	logs.Info("ChangeIssueState success,response %s", resp)
	return nil
}

func GetAllProjectIssue(token string, projectId int) (string, error) {
	resp, err := request.StringForStringWithPagination(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectId)+"/issues?per_page=100",
		"GET",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		"",
		5,
	)
	if err != nil {
		logs.Error("GetIssue:Request Failed,%s", err)
		return "", err
	}
	// if status != 200 {
	// 	logs.Error("GetIssue Failed,Code %d", status)
	// 	return "",fmt.Errorf("GetIssue Request Failed,Code %d", status)
	// }
	return resp, nil
}

func GetAllProjectIssueInObject(token string, projectId int) ([]map[string]interface{}, error) {
	resp, err := request.StringForStringWithPagination(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectId)+"/issues?per_page=100",
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
	var ret = make([]map[string]interface{}, 0)
	json.Unmarshal([]byte(resp), &ret)
	return ret, nil
}

func GetProjectIssueNum(token string, projectId int) (int, error) {
	header, _, err := request.StringForStringWithHeader(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectId)+"/issues?per_page=100",
		"GET",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		"",
		5,
	)

	if err != nil {
		logs.Error("GetIssueStatistic failed,%s", err)
		return -1, err
	}
	total, err := strconv.Atoi(header.Get("x-total"))
	return total, nil

}
