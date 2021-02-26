package handler

import (
	"graduation_design/internal/app/model"
	"regexp"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//
func CreateMilestone(c *gin.Context){
	session := sessions.Default(c)
	accessToken,_:=session.Get("access_token").(string)
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
	err:=model.CreateMileStone(accessToken,projectId,title,description,startDateStr,dueDateStr)
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

func GetAllProjectMilestones(c *gin.Context){
	session := sessions.Default(c)
	accessToken,_:=session.Get("access_token").(string)
	projectIDStr:=c.Query("projectid")
	projectID,err:=strconv.Atoi(projectIDStr)
	if err!=nil{
		c.JSON(400,gin.H{
			"error":"invalid project id",
		})
		return
	}
	resp,err:=model.GetAllProjectMilestones(accessToken,projectID)
	if err!=nil{
		c.JSON(400,gin.H{
			"error":err.Error(),
		})
		return
	}
	c.Header("Content-Type", "application/json")
	c.String(200, resp)
}

func GetProjectMilestone(c *gin.Context){
	session := sessions.Default(c)
	accessToken,_:=session.Get("access_token").(string)
	projectIDStr:=c.Query("projectid")
	projectID,err:=strconv.Atoi(projectIDStr)
	if err!=nil{
		c.JSON(400,gin.H{
			"error":"invalid project id",
		})
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
	resp,err:=model.GetProjectMilestone(accessToken,projectID,milestoneID)
	if err!=nil{
		c.JSON(400,gin.H{
			"error":err.Error(),
		})
		return
	}
	c.Header("Content-Type", "application/json")
	c.String(200, resp)
}

func EditMilestone(c *gin.Context){
	session := sessions.Default(c)
	accessToken,_:=session.Get("access_token").(string)
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
	milestoneIDFloat,ok:=reqData["milestone_id"].(float64)
	if !ok{
		c.JSON(400,gin.H{
			"error":"invalid milestone_id",
		})
		return
	}
	milestoneID:=int(milestoneIDFloat)
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
	err:=model.EditMileStone(accessToken,projectId,milestoneID,title,description,startDateStr,dueDateStr)
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

func DeleteProjectMilestone(c *gin.Context){
	session := sessions.Default(c)
	accessToken,_:=session.Get("access_token").(string)
	projectIDStr:=c.Query("projectid")
	projectID,err:=strconv.Atoi(projectIDStr)
	if err!=nil{
		c.JSON(400,gin.H{
			"error":"invalid project id",
		})
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
	err=model.DeleteMilestone(accessToken,projectID,milestoneID)
	if err!=nil{
		c.JSON(400,gin.H{
			"error":err.Error(),
		})
		return
	}
	c.JSON(200,gin.H{
		"message":"success",
	})
	
}

