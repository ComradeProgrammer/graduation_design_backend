package model

import (
	"encoding/json"
	"fmt"
	"graduation_design/internal/app/config"
	"graduation_design/internal/app/db"
	"graduation_design/internal/pkg/logs"
	"graduation_design/internal/pkg/request"
	"strconv"
)

type Project struct {
	db.ProjectDB
	SSHURL string `json:"ssh_url_to_repo"`
}

// return (tracked projects,untracked projects)
func GetUserProjects(token string) ([]Project, []Project, error) {
	logs.Info("GetTrackingProjects")
	userProjects, err := getAllProjectsFromGitlab(token)
	if err != nil {
		logs.Error("GetTrackingProjects: read gitlab projects failed,%s", err)
		return nil, nil, err
	}
	var tracked = make([]Project, 0)
	var untracked = make([]Project, 0)
	for i := 0; i < len(userProjects); i++ {
		if _, err := db.FindProjectByID(userProjects[i].ID); err != nil {
			//not found
			untracked = append(untracked, userProjects[i])
		} else {
			tracked = append(tracked, userProjects[i])
		}
	}
	return tracked, untracked, nil

}

func TrackProject(token string, id int) error {
	logs.Info("TrackProject id %d", id)
	prj, err := getProjectFromGitlab(token, id)
	if err != nil {
		logs.Error("TrackProject: read gitlab projects failed,%s", err)
		return err
	}
	return prj.ProjectDB.SaveProject()

}

func UntrackProject(token string, id int) error {
	logs.Info("UnTrackProject id %d", id)
	var tmp = db.ProjectDB{ID: id}
	return tmp.DeleteProject()
}

func getProjectFromGitlab(token string, id int) (Project, error) {
	logs.Info("getProjectFromGitlab")
	var res = Project{}
	status, resp, err := request.StringForString(
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
		return res, err
	}
	if status != 200 {
		logs.Error("getProjectFromGitlab Request Failed,Code %d", status)
		return res, fmt.Errorf("getProjectFromGitlab Request Failed,Code %d", status)
	}
	err = json.Unmarshal([]byte(resp), &res)
	if err != nil {
		logs.Error("getAllProjectsFromGitlab:Unmarshal Failed,%s", err)
		return res, err
	}
	return res, nil
}

func getAllProjectsFromGitlab(token string) ([]Project, error) {
	logs.Info("getAllProjectsFromGitlab")
	resp, err := request.StringForStringWithPagination(
		config.GITLABAPIURL+"/projects?per_page=100",
		"GET",
		map[string]string{
			"Authorization": "Bearer " + token,
		},
		"",
		5,
	)
	if err != nil {
		logs.Error("getAllProjectsFromGitlab:Request Failed,%s", err)
		return nil, err
	}
	// if status != 200 {
	// 	logs.Error("getAllProjectsFromGitlab Request Failed,Code %d", status)
	// 	return nil, fmt.Errorf("getAllProjectsFromGitlab Request Failed,Code %d", status)
	// }
	var res = make([]Project, 0)
	err = json.Unmarshal([]byte(resp), &res)
	if err != nil {
		logs.Error("getAllProjectsFromGitlab:Unmarshal Failed,%s", err)
		return nil, err
	}
	return res, nil

}
