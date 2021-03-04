package handler

import (
	"graduation_design/internal/app/db"
	"graduation_design/internal/app/model"

	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetCodeQuality(c *gin.Context){
	session := sessions.Default(c)
	accessToken, _ := session.Get("access_token").(string)
	projectIDStr := c.Query("projectid")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid projectid",
		})
		return
	}
	data,err:=model.CheckQuality(accessToken,projectID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": data,
	})
}

func GetJobLog(c *gin.Context) {
	session := sessions.Default(c)
	accessToken, _ := session.Get("access_token").(string)
	projectIDStr := c.Query("projectid")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid projectid",
		})
		return
	}
	jobIDStr := c.Query("jobid")
	jobID, err := strconv.Atoi(jobIDStr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid jobid",
		})
		return
	}
	res, err := model.GetJobLog(accessToken, projectID, jobID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(400, gin.H{
		"log": res,
	})

}

func GetAllRegex(c *gin.Context) {
	//session := sessions.Default(c)
	//accessToken, _ := session.Get("access_token").(string)

	//todo fix:check permission:
	projectIDStr := c.Query("projectid")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid projectid",
		})
		return
	}
	data := model.GetAllRegex(projectID)
	c.JSON(200, gin.H{
		"data": data,
	})

}

func CreateRegex(c *gin.Context) {
	//todo fix:check permission
	reqData := make(map[string]interface{})
	c.BindJSON(&reqData)
	projectIdFloat, ok := reqData["project_id"].(float64)
	if !ok {
		c.JSON(400, gin.H{
			"error": "invalid project_id",
		})
		return
	}
	projectId := int(projectIdFloat)
	regex, ok := reqData["regex"].(string)
	if !ok || regex == "" {
		c.JSON(400, gin.H{
			"error": "invalid regex",
		})
		return
	}
	regexType, ok := reqData["regex_type"].(string)
	if !ok || (!(regexType == db.COVERAGE || regexType == db.LINT)) {
		c.JSON(400, gin.H{
			"error": "invalid regex_type",
		})
		return
	}
	comment, ok := reqData["comment"].(string)
	if comment == "" {
		c.JSON(400, gin.H{
			"error": "invalid comment",
		})
		return
	}
	err := model.CreateRegex(projectId, regex, regexType, comment)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid projectid",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func DeleteRegex(c *gin.Context) {
	regexIDStr := c.Query("regexid")
	regexID, err := strconv.Atoi(regexIDStr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid regex_id",
		})
		return
	}
	model.DeleteRegex(regexID)
	c.JSON(200, gin.H{
		"message": "success",
	})
}
