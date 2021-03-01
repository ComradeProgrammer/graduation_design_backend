package model

import (
	"encoding/json"
	"fmt"
	"graduation_design/internal/app/config"
	"graduation_design/internal/pkg/logs"
	"graduation_design/internal/pkg/request"
	"strconv"
)

type MileStone struct {
	Id          int         `json:"id"`
	IId         int         `json:"iid"`
	ProjectId   int         `json:"project_id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	DueDate     string      `json:"due_date"`
	StartDate   string      `json:"start_date"`
	State       string      `json:"state"`
	UpdateAt    string      `json"update_at"`
	CreateAt    string      `json:"create_at"`
	Expired     bool        `json:"expired"`
	Issues      interface{} `json:"issues"`
	MRs         interface{} `json:"mrs"`
}

func CreateMileStone(token string, projectId int, title, description string, startDate, dueDate string) error {
	data := map[string]string{
		"title":       title,
		"description": description,
		"start_date":  startDate,
		"due_date":    dueDate,
	}
	status, resp, err := request.FormForJson(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectId)+"/milestones",
		"POST",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		data,
		5,
	)
	if err != nil {
		logs.Error("CreateMileStone:Request Failed,%s", err)
		return err
	}
	if status != 201 {
		logs.Error("CreateMileStone Failed,Code %d", status)
		return fmt.Errorf("CreateMileStone Request Failed,Code %d", status)
	}
	logs.Info("CreateMileStone success,response %s", resp)
	return nil
}

func GetAllProjectMilestones(token string, projectId int) (string, error) {
	status, resp, err := request.StringForString(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectId)+"/milestones",
		"GET",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		"",
		5,
	)
	if err != nil {
		logs.Error("GetProjectMilestone:Request Failed,%s", err)
		return "", err
	}
	if status != 200 {
		logs.Error("GetProjectMilestone Failed,Code %d", status)
		return "", fmt.Errorf("GetProjectMilestone Request Failed,Code %d", status)
	}

	var milestones = make([]MileStone, 0)
	err = json.Unmarshal([]byte(resp), &milestones)
	if err != nil {
		logs.Error("GetProjectMilestone:unmarshalFailed,%s", err)
		return "", err
	}
	var issueChans = make([]chan string, len(milestones))
	var mrChans = make([]chan string, len(milestones))

	for i_ := 0; i_ < len(milestones); i_++ {
		var i = i_
		issueChan := make(chan string)
		mrChan := make(chan string)
		issueChans[i] = issueChan
		mrChans[i] = mrChan
		go func() {
			issues, err1 := getAllIssuesOfMilestone(token, projectId, milestones[i].Id)
			if err1 != nil {
				issueChan <- ""
				logs.Error("getAllIssuesOfMilestone:%s", err1)
				return
			}
			issueChan <- issues
		}()
		go func() {
			mrs, err1 := getAllMROfMilestone(token, projectId, milestones[i].Id)
			if err1 != nil {
				mrChan <- ""
				logs.Error("getAllMROfMilestone:%s", err1)
				return
			}
			mrChan <- mrs
		}()
	}
	for i := 0; i < len(milestones); i++ {
		issuesStr := <-issueChans[i]
		mrsStr := <-mrChans[i]
		var issueObj = make([]map[string]interface{}, 0)
		var mrObj = make([]map[string]interface{}, 0)
		err = json.Unmarshal([]byte(issuesStr), &issueObj)
		if err != nil {
			logs.Error("GetProjectMilestone:unmarshalFailed,%s", err)
			return "", err
		}
		err = json.Unmarshal([]byte(mrsStr), &mrObj)
		if err != nil {
			logs.Error("GetProjectMilestone:unmarshalFailed,%s", err)
			return "", err
		}
		milestones[i].Issues = issueObj
		milestones[i].MRs = mrObj
	}
	res, err := json.Marshal(milestones)
	return string(res), nil
}
func GetProjectMilestone(token string, projectId int, milestoneId int) (string, error) {
	status, resp, err := request.StringForString(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectId)+"/milestones/"+strconv.Itoa(milestoneId),
		"GET",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		"",
		5,
	)
	if err != nil {
		logs.Error("GetProjectMilestone:Request Failed,%s", err)
		return "", err
	}
	if status != 200 {
		logs.Error("GetProjectMilestone Failed,Code %d", status)
		return "", fmt.Errorf("GetProjectMilestone Request Failed,Code %d", status)
	}
	var issueChan = make(chan string)
	var mrChan = make(chan string)
	var milestone = MileStone{}
	err = json.Unmarshal([]byte(resp), &milestone)
	if err != nil {
		logs.Error("GetProjectMilestone:unmarshalFailed,%s", err)
		return "", err
	}
	go func() {
		issues, err1 := getAllIssuesOfMilestone(token, projectId, milestoneId)
		if err1 != nil {
			issueChan <- ""
			logs.Error("getAllIssuesOfMilestone:%s", err1)
			return
		}
		issueChan <- issues
	}()
	go func() {
		mrs, err1 := getAllMROfMilestone(token, projectId, milestoneId)
		if err1 != nil {
			mrChan <- ""
			logs.Error("getAllMROfMilestone:%s", err1)
			return
		}
		mrChan <- mrs
	}()
	issuesStr := <-issueChan
	mrsStr := <-mrChan
	var issueObj = make([]map[string]interface{}, 0)
	var mrObj = make([]map[string]interface{}, 0)
	err = json.Unmarshal([]byte(issuesStr), &issueObj)
	if err != nil {
		logs.Error("GetProjectMilestone:unmarshalFailed,%s", err)
		return "", err
	}
	err = json.Unmarshal([]byte(mrsStr), &mrObj)
	if err != nil {
		logs.Error("GetProjectMilestone:unmarshalFailed,%s", err)
		return "", err
	}
	milestone.Issues = issueObj
	milestone.MRs = mrObj
	res, err := json.Marshal(milestone)
	if err != nil {
		logs.Error("GetProjectMilestone:marshalFailed,%s", err)
		return "", err
	}

	return string(res), nil
}
func EditMileStone(token string, projectId int, milestoneId int, title, description string, startDate, dueDate string) error {
	data := map[string]string{
		"title":       title,
		"description": description,
		"start_date":  startDate,
		"due_date":    dueDate,
	}
	status, resp, err := request.FormForJson(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectId)+"/milestones/"+strconv.Itoa(milestoneId),
		"PUT",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		data,
		5,
	)
	if err != nil {
		logs.Error("CreateMileStone:Request Failed,%s", err)
		return err
	}
	if status != 200 {
		logs.Error("CreateMileStone Failed,Code %d", status)
		return fmt.Errorf("CreateMileStone Request Failed,Code %d", status)
	}
	logs.Info("CreateMileStone success,response %s", resp)
	return nil
}

func DeleteMilestone(token string, projectId int, milestoneId int) error {
	status, _, err := request.StringForString(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectId)+"/milestones/"+strconv.Itoa(milestoneId),
		"DELETE",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		"",
		5,
	)
	if err != nil {
		logs.Error(" DeleteMileStone:Request Failed,%s", err)
		return err
	}
	if status != 204 {
		logs.Error("DeleteMileStone Failed,Code %d", status)
		return fmt.Errorf(" DeleteMileStone Request Failed,Code %d", status)
	}
	return nil
}

func getAllIssuesOfMilestone(token string, projectId int, milestoneId int) (string, error) {
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
		logs.Error("getAllIssuesOfMilestone:Request Failed,%s", err)
		return "", err
	}
	if status != 200 {
		logs.Error("getAllIssuesOfMilestone Failed,Code %d", status)
		return "", fmt.Errorf("getAllIssuesOfMilestone Request Failed,Code %d", status)
	}
	logs.Info("getAllIssuesOfMilestone success,response %s", resp)
	return resp, nil
}

func getAllMROfMilestone(token string, projectId int, milestoneId int) (string, error) {
	status, resp, err := request.StringForString(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectId)+"/milestones/"+strconv.Itoa(milestoneId)+"/merge_requests",
		"GET",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		"",
		5,
	)
	if err != nil {
		logs.Error("getAllMROfMilestone:Request Failed,%s", err)
		return "", err
	}
	if status != 200 {
		logs.Error("getAllMROfMilestone Failed,Code %d", status)
		return "", fmt.Errorf("getAllMROfMilestone Request Failed,Code %d", status)
	}
	logs.Info("getAllMROfMilestone success,response %s", resp)
	return resp, nil
}
