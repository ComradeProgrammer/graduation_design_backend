package model

import (
	"encoding/json"
	"graduation_design/internal/app/config"
	"graduation_design/internal/pkg/logs"
	"graduation_design/internal/pkg/pool"
	"graduation_design/internal/pkg/request"
	"strconv"
)

func GetAllIssueNotes(token string, projectId int) ([]map[string]interface{}, error) {
	resp, err := GetAllProjectIssueInObject(token, projectId)
	if err != nil {
		logs.Error("get all issue failed when GetAllIssueDiscussion")
	}
	var issueIIDList = make([]int, 0)
	for _, issue := range resp {
		iid, ok := issue["iid"].(float64)
		if !ok {
			logs.Error("iid not found for issue, content is %v", issue)
			continue
		}
		issueIIDList = append(issueIIDList, int(iid))
	}
	var res = make(chan []map[string]interface{}, len(issueIIDList))
	var p = pool.NewPool(10, len(issueIIDList))
	for i := 0; i < len(issueIIDList); i++ {
		var i2 = i
		p.AddTask(func() {
			r, err := request.StringForStringWithPagination(
				config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectId)+"/issues/"+strconv.Itoa(issueIIDList[i2])+"/notes?per_page=100",
				"GET",
				map[string]string{
					"Authorization": "Bearer " + token,
				},
				"",
				5,
			)
			if err != nil {
				logs.Error("GetIssue:Request Failed,%s", err)
			}
			var tmp = make([]map[string]interface{}, 0)
			json.Unmarshal([]byte(r), &tmp)
			res <- tmp
		})

	}
	p.Run()
	p.Wait()
	close(res)
	var ret = make([]map[string]interface{}, 0)
	for data := range res {
		ret = append(ret, data...)
	}
	return ret, nil

}

func GetAllMRNotes(token string, projectId int) ([]map[string]interface{}, error) {
	resp, err := GetAllProjectMrInObject(token, projectId)
	if err != nil {
		logs.Error("get all issue failed when GetAllIssueDiscussion")
	}
	var issueIIDList = make([]int, 0)
	for _, issue := range resp {
		iid, ok := issue["iid"].(float64)
		if !ok {
			logs.Error("iid not found for issue, content is %v", issue)
			continue
		}
		issueIIDList = append(issueIIDList, int(iid))
	}
	var res = make(chan []map[string]interface{}, len(issueIIDList))
	var p = pool.NewPool(10, len(issueIIDList))
	for i := 0; i < len(issueIIDList); i++ {
		var i2 = i
		p.AddTask(func() {
			r, err := request.StringForStringWithPagination(
				config.GITLABAPIURL+"/projects/"+strconv.Itoa(projectId)+"/merge_requests/"+strconv.Itoa(issueIIDList[i2])+"/notes?per_page=100",
				"GET",
				map[string]string{
					"Authorization": "Bearer " + token,
				},
				"",
				5,
			)
			if err != nil {
				logs.Error("GetIssue:Request Failed,%s", err)
			}
			var tmp = make([]map[string]interface{}, 0)
			json.Unmarshal([]byte(r), &tmp)
			res <- tmp
		})

	}
	p.Run()
	p.Wait()
	close(res)
	var ret = make([]map[string]interface{}, 0)
	for data := range res {
		ret = append(ret, data...)
	}
	return ret, nil

}
