package model

import (
	"encoding/json"
	"fmt"
	"graduation_design/internal/app/config"
	"graduation_design/internal/pkg/git"
	"graduation_design/internal/pkg/logs"
	"graduation_design/internal/pkg/request"
	"strconv"
)

type GitlabCommit struct {
	ID         string `json:"id"`
	GitlabUser string `json:"committer_name"`
	WebUrl     string `json:"web_url"`
}
type FullCommit struct {
	git.GitCommit
	GitlabUser string `json:"committer_name"`
	WebUrl     string `json:"web_url"`
}

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

func GetAllCommitsInMap(token string,projectID int)(map[string]GitlabCommit,error){
	resp, err := request.StringForStringWithPagination(
		config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectID)+"/repository/commits/?per_page=100",
		"GET",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		"",
		5,
	)
	if err!=nil{
		logs.Error("GetAllCommitsInMap:%s",err.Error())
		return nil,err
	}
	var commitList=make([]GitlabCommit,0)
	err=json.Unmarshal([]byte(resp),&commitList)
	if err!=nil{
		logs.Error("GetAllCommitsInMap:%s",err.Error())
		return nil,err
	}
	var commitMap=make(map[string]GitlabCommit)
	for _,c:=range commitList{
		commitMap[c.ID]=c
	} 
	return commitMap,nil

}

func mergeCommitsFromGitAndGitlab(fromGit[]git.GitCommit,fromGitlab map[string]GitlabCommit)([]FullCommit,map[string]string){
	var res=make([]FullCommit,len(fromGit))
	var gitToGitlab=make(map[string]string)
	for i,c:=range fromGit{
		res[i].GitCommit=c
		res[i].GitlabUser=fromGitlab[c.Hash].GitlabUser
		res[i].WebUrl=fromGitlab[c.Hash].WebUrl
		gitToGitlab[res[i].Author]=fromGitlab[c.Hash].GitlabUser
	}

	return res,gitToGitlab
}