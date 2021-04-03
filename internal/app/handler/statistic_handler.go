package handler

import (
	"graduation_design/internal/app/model"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetProjectOverviewStatistic(c *gin.Context) {
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

	ok, err = model.CheckProjectAuthorization(accessToken, int(id))
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

	result, err := model.AnalysisProjectOverview(accessToken, int(id))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": result,
	})
	return
}

func GetProjectUserStatistic(c *gin.Context) {
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

	ok, err = model.CheckProjectAuthorization(accessToken, int(id))
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

	result, err := model.AnalyzeUserView(accessToken, int(id))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": result,
	})
	return
}

func GetProjectCommitStatistic(c *gin.Context) {
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

	ok, err = model.CheckProjectAuthorization(accessToken, int(id))
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

	result, err := model.AnalyzeCommitView(accessToken, int(id))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": result,
	})
	return
}
func GetProjectLanguageStatistic(c *gin.Context) {
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

	ok, err = model.CheckProjectAuthorization(accessToken, int(id))
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

	result, err := model.AnalyzeLanguageView(accessToken, int(id))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": result,
	})
	return
}

func GetProjectActivityStatistic(c *gin.Context) {
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

	ok, err = model.CheckProjectAuthorization(accessToken, int(id))
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

	result, err := model.AnalyzeActivityView(accessToken, int(id))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": result,
	})
	return
}