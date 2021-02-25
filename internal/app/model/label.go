package model

import (
	"encoding/json"
	"fmt"
	"graduation_design/internal/app/config"
	"graduation_design/internal/pkg/logs"
	"graduation_design/internal/pkg/request"
	"strconv"
)

type Label struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Color string `json:"color"`
}
var presetLabels=[]Label{
	{0,"bug","#5843AD"},
	{0,"feature","#5cb85c"},
	{0,"P0","#FF0000"},
	{0,"P1","#F0AD4E"},
	{0,"P2","#428BCA"},
}
func (l1 *Label)equal(l2 *Label)bool{
	if l1.Name==l2.Name && l1.Color==l2.Color{
		return true
	}
	return false
}

func (l *Label)CreateLabel(token string,projectID int)error{
	status,resp,err:=request.FormForJson(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectID)+"/labels",
		"POST",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		map[string]string{
			"name":l.Name,
			"color":l.Color,
		},
		5,
	)
	if err != nil {
		logs.Error("CreateLabel Failed:Request Failed,%s", err)
		return err
	}
	if status != 201 {
		logs.Error("CreateLabel Failed,Code %d", status)
		return  fmt.Errorf("CreateLabel Request Failed,Code %d", status)
	}
	logs.Info("CreateLabel success,response %s",resp)
	return nil
}

func CheckAndCreateLabels(token string,projectID int)error{
	current,err:=getLabels(token,projectID)
	if err!=nil{
		return err
	}
	var add=make([]Label,0)
	for _,i:=range presetLabels{
		var override=false
		for _,j:=range current{
			if i.equal(&j){
				override=true
				break
			}
		}
		if override{
			continue
		}
		add = append(add, i)
	}
	for _,l:=range add{
		err=l.CreateLabel(token,projectID)
		if err!=nil{
			return err
		}
	}
	return nil
}

func getLabels(token string,projectID int)([]Label,error){
	status,resp,err:=request.StringForString(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectID)+"/labels",
		"GET",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		"",
		5,
	)
	if err!=nil{
		return nil,err
	}
	if status != 200 {
		logs.Error("getLabels Failed,Code %d", status)
		return  nil,fmt.Errorf("getLabels Request Failed,Code %d", status)
	}
	var res=make([]Label,0)
	err=json.Unmarshal([]byte(resp),&res)
	if err!=nil{
		return nil,err
	}
	return res,nil
}
