package handler

import (
	"graduation_design/internal/app/model"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetProjectStatistic(c *gin.Context) {
	session := sessions.Default(c)
	accessToken, ok := session.Get("access_token").(string)
	if !ok || accessToken == "" {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}

	id, err := strconv.ParseInt(c.Query("id"), 0, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	result,err:=model.AnalysisProject(accessToken,int(id))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200,gin.H{
		"data":result,
	})
	return 
}