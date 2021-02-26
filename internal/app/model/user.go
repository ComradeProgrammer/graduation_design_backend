package model

import (
	"encoding/json"
	"fmt"
	"graduation_design/internal/app/config"
	"graduation_design/internal/pkg/logs"
	"graduation_design/internal/pkg/request"
)

type User struct{
	Id int `json:"id"`
	UserName string `json:"username"`
	Avatar string	`json:"avatar_url"`
	IsAdmin bool `json:"is_admin"`

}

func GetCurrentUser(token string)(*User,error){
	logs.Info("GetCurrentUser")
	status,resp,err:=request.StringForString(
		config.GITLABAPIURL+"/user",
		"GET",
		map[string]string{
			"Authorization":"Bearer "+token,
		},
		"",
		5,
	)
	if err != nil {
		logs.Error("GetCurrentUser:Request Failed,%s", err)
		return nil, err
	}
	if status!=200{
		logs.Error("GetCurrentUser Request Failed,Code %d", status)
		return nil, fmt.Errorf("GetCurrentUser Request Failed,Code %d", status)
	}
	var tmp=User{}
	err=json.Unmarshal([]byte(resp),&tmp)
	if err != nil {
		logs.Error("GetCurrentUser:marshal Failed,%s", err)
		return nil, err
	}
	return &tmp,nil
}