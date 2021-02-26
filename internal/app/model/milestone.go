package model

import (
	"fmt"
	"graduation_design/internal/app/config"
	"graduation_design/internal/pkg/logs"
	"graduation_design/internal/pkg/request"
	"strconv"
)


func CreateMileStone(token string,projectId int,title,description string,startDate,dueDate string)error{
	data:=map[string]string{
		"title":title,
		"description":description,
		"start_date":startDate,
		"due_date":dueDate,
	}
	status,resp,err:=request.FormForJson(
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
		return  fmt.Errorf("CreateMileStone Request Failed,Code %d", status)
	}
	logs.Info("CreateMileStone success,response %s",resp)
	return nil
}

func GetAllProjectMilestones(token string,projectId int)(string,error){
	status,resp,err:=request.StringForString(
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
		return "",err
	}
	if status != 200 {
		logs.Error("GetProjectMilestone Failed,Code %d", status)
		return "",fmt.Errorf("GetProjectMilestone Request Failed,Code %d", status)
	}
	return resp,nil
}
func GetProjectMilestone(token string,projectId int,milestoneId int)(string,error){
	status,resp,err:=request.StringForString(
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
		return "",err
	}
	if status != 200 {
		logs.Error("GetProjectMilestone Failed,Code %d", status)
		return "",fmt.Errorf("GetProjectMilestone Request Failed,Code %d", status)
	}
	return resp,nil
}
func EditMileStone(token string,projectId int,milestoneId int,title,description string,startDate,dueDate string)error{
	data:=map[string]string{
		"title":title,
		"description":description,
		"start_date":startDate,
		"due_date":dueDate,
	}
	status,resp,err:=request.FormForJson(
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
		return  fmt.Errorf("CreateMileStone Request Failed,Code %d", status)
	}
	logs.Info("CreateMileStone success,response %s",resp)
	return nil
}

func DeleteMilestone(token string,projectId int,milestoneId int)error{
	status,_,err:=request.StringForString(
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
		logs.Error(" DeleteMileStone Failed,Code %d", status)
		return fmt.Errorf(" DeleteMileStone Request Failed,Code %d", status)
	}
	return nil
}