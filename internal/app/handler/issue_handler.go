package handler

import (
	"graduation_design/internal/app/model"
	"regexp"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CreateIssue(c *gin.Context){
	session := sessions.Default(c)
	accessToken,ok:=session.Get("access_token").(string)
	if !ok || accessToken == "" {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}
	reqData:=make(map[string]interface{})
	c.BindJSON(&reqData)
	projectIdFloat,ok:=reqData["project_id"].(float64)
	if !ok{
		c.JSON(400,gin.H{
			"error":"invalid project_id",
		})
		return
	}
	projectId:=int(projectIdFloat)
	ok, err := model.CheckProjectAuthorization(accessToken, projectId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !ok {
		c.JSON(401, gin.H{"message": "unauthorized"})
		return
	}
	milestoneIdFloat,ok:=reqData["milestone_id"].(float64)
	if !ok{
		c.JSON(400,gin.H{
			"error":"invalid milestone_id",
		})
		return
	}
	milestoneId:=int(milestoneIdFloat)
	title,ok:=reqData["title"].(string)
	if !ok{
		c.JSON(400,gin.H{
			"error":"invalid title",
		})
		return
	}
	description,ok:=reqData["description"].(string)
	if !ok{
		c.JSON(400,gin.H{
			"error":"invalid description",
		})
		return
	}
	
	startDateStr,ok:=reqData["start_date"].(string)
	if !ok{
		c.JSON(400,gin.H{
			"error":"invalid start_date",
		})
		return
	}
	match,_:=regexp.Match(`^\d\d\d\d-\d\d-\d\d$`,[]byte(startDateStr))
	if !match{
		c.JSON(400,gin.H{
			"error":"invalid format of start_date",
		})
		return
	}

	dueDateStr,ok:=reqData["due_date"].(string)
	if !ok{
		c.JSON(400,gin.H{
			"error":"invalid start_date",
		})
		return
	}
	match,_=regexp.Match(`^\d\d\d\d-\d\d-\d\d$`,[]byte(dueDateStr))
	if !match{
		c.JSON(400,gin.H{
			"error":"invalid format of start_date",
		})
		return
	}
	typeTag,ok:=reqData["type_tag"].(string)
	if !ok ||(typeTag!="feature"&&typeTag!="bug"){
		c.JSON(400,gin.H{
			"error":"invalid format of type_tag",
		})
		return
	}
	priorityTag,ok:=reqData["priority_tag"].(string)
	if !ok ||(priorityTag!="P0"&&priorityTag!="P1"&&priorityTag!="P2"){
		c.JSON(400,gin.H{
			"error":"invalid format of priority_tag",
		})
		return
	}
	err=model.CreateIssue(accessToken,projectId,milestoneId,
	title,description,startDateStr,dueDateStr,typeTag,priorityTag)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200,gin.H{
		"message":"success",
	})
}

func GetAllIssues(c *gin.Context){
	session := sessions.Default(c)
	accessToken,ok:=session.Get("access_token").(string)
	if !ok || accessToken == "" {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}
	projectIDStr:=c.Query("projectid")
	projectID,err:=strconv.Atoi(projectIDStr)
	if err!=nil{
		c.JSON(400,gin.H{
			"error":"invalid project id",
		})
		return
	}
	ok, err = model.CheckProjectAuthorization(accessToken, projectID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !ok {
		c.JSON(401, gin.H{"message": "unauthorized"})
		return
	}
	milestoneIDStr:=c.Query("milestoneid")
	milestoneID,err:=strconv.Atoi(milestoneIDStr)
	if err!=nil{
		c.JSON(400,gin.H{
			"error":"invalid milestoneid",
		})
		return
	}
	resp,err:=model.GetAllIssuesForMilestone(accessToken,projectID,milestoneID)
	if err!=nil{
		c.JSON(400,gin.H{
			"error":err.Error(),
		})
		return
	}
	c.Header("Content-Type", "application/json")
	c.String(200, resp)
}
func GetIssue(c *gin.Context){
	session := sessions.Default(c)
	accessToken,ok:=session.Get("access_token").(string)
	if !ok || accessToken == "" {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}
	projectIDStr:=c.Query("projectid")
	projectID,err:=strconv.Atoi(projectIDStr)
	if err!=nil{
		c.JSON(400,gin.H{
			"error":"invalid project id",
		})
		return
	}
	ok, err = model.CheckProjectAuthorization(accessToken, projectID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !ok {
		c.JSON(401, gin.H{"message": "unauthorized"})
		return
	}
	issueIiDStr:=c.Query("issue_iid")
	issueIiD,err:=strconv.Atoi(issueIiDStr)
	if err!=nil{
		c.JSON(400,gin.H{
			"error":"invalid issue_iid",
		})
		return
	}
	resp,err:=model.GetIssue(accessToken,projectID,issueIiD)
	if err!=nil{
		c.JSON(400,gin.H{
			"error":err.Error(),
		})
		return
	}
	c.Header("Content-Type", "application/json")
	c.String(200, resp)
}

func EditIssue(c *gin.Context){
	session := sessions.Default(c)
	accessToken,ok:=session.Get("access_token").(string)
	if !ok || accessToken == "" {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}
	reqData:=make(map[string]interface{})
	c.BindJSON(&reqData)
	projectIdFloat,ok:=reqData["project_id"].(float64)
	if !ok{
		c.JSON(400,gin.H{
			"error":"invalid project_id",
		})
		return
	}
	projectId:=int(projectIdFloat)
	ok, err := model.CheckProjectAuthorization(accessToken, projectId)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !ok {
		c.JSON(401, gin.H{"message": "unauthorized"})
		return
	}
	milestoneIDFloat,ok:=reqData["milestone_id"].(float64)
	if !ok{
		c.JSON(400,gin.H{
			"error":"invalid milestone_id",
		})
		return
	}
	milestoneID:=int(milestoneIDFloat)
	issueIiDFloat,ok:=reqData["issue_iid"].(float64)
	if !ok{
		c.JSON(400,gin.H{
			"error":"invalid issue_iid",
		})
		return
	}
	issueIiD:=int(issueIiDFloat)
	
	title,ok:=reqData["title"].(string)
	if !ok{
		c.JSON(400,gin.H{
			"error":"invalid title",
		})
		return
	}
	description,ok:=reqData["description"].(string)
	if !ok{
		c.JSON(400,gin.H{
			"error":"invalid description",
		})
		return
	}
	
	startDateStr,ok:=reqData["start_date"].(string)
	if !ok{
		c.JSON(400,gin.H{
			"error":"invalid start_date",
		})
		return
	}
	match,_:=regexp.Match(`^\d\d\d\d-\d\d-\d\d$`,[]byte(startDateStr))
	if !match{
		c.JSON(400,gin.H{
			"error":"invalid format of start_date",
		})
		return
	}

	dueDateStr,ok:=reqData["due_date"].(string)
	if !ok{
		c.JSON(400,gin.H{
			"error":"invalid start_date",
		})
		return
	}
	match,_=regexp.Match(`^\d\d\d\d-\d\d-\d\d$`,[]byte(dueDateStr))
	if !match{
		c.JSON(400,gin.H{
			"error":"invalid format of start_date",
		})
		return
	}
	typeTag,ok:=reqData["type_tag"].(string)
	if !ok ||(typeTag!="feature"&&typeTag!="bug"){
		c.JSON(400,gin.H{
			"error":"invalid format of type_tag",
		})
		return
	}
	priorityTag,ok:=reqData["priority_tag"].(string)
	if !ok ||(priorityTag!="P0"&&priorityTag!="P1"&&priorityTag!="P2"){
		c.JSON(400,gin.H{
			"error":"invalid format of priority_tag",
		})
		return
	}
	err=model.EditIssue(accessToken,projectId,milestoneID,issueIiD,
		title,description,startDateStr,dueDateStr,typeTag,priorityTag)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200,gin.H{
		"message":"success",
	})
	
}

func ChangeIssueState(c *gin.Context){
	session := sessions.Default(c)
	accessToken,ok:=session.Get("access_token").(string)
	if !ok || accessToken == "" {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}
	projectIDStr:=c.Query("projectid")
	projectID,err:=strconv.Atoi(projectIDStr)
	if err!=nil{
		c.JSON(400,gin.H{
			"error":"invalid project_id",
		})
		return
	}
	ok, err = model.CheckProjectAuthorization(accessToken, projectID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !ok {
		c.JSON(401, gin.H{"message": "unauthorized"})
		return
	}
	issueIiDStr:=c.Query("issue_iid")
	issueIiD,err:=strconv.Atoi(issueIiDStr)
	if err!=nil{
		c.JSON(400,gin.H{
			"error":"invalid issue_iid",
		})
		return
	}
	stateEvent:=c.Query("state_event")
	if stateEvent!="close" && stateEvent!="reopen"{
		c.JSON(400,gin.H{
			"error":"invalid state_event",
		})
		return
	}
	err=model.ChangeIssueState(accessToken,projectID,issueIiD,stateEvent)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200,gin.H{
		"message":"success",
	})
}	