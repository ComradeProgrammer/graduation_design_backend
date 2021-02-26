package handler

import (
	"graduation_design/internal/app/model"

	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//return :json string of all users' projects
//route:/projects
//sample response
/*
	{
		tracked:[
			{
				id: 3,
				name: "grouptest",
				name_with_namespace: "172317 / 团队项目 / grouptest",
				web_url: "http://127.0.0.1/172317/team-projects/grouptest",
				ssh_url_to_repo: "git@127.0.0.1:172317/team-projects/grouptest.git"
			},...
		],
		untracked:[  ...list of projects...   ]
	}
*/
func GetProjects(c *gin.Context) {
	session := sessions.Default(c)
	accessToken, ok := session.Get("access_token").(string)
	if !ok || accessToken == "" {
		c.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}
	tracked, untracked, err := model.GetUserProjects(accessToken)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"tracked":   tracked,
		"untracked": untracked,
	})
}

//route: /project/track?id=<project id>
//function:start tracking a project
//return code 200 if succeeded,else failede
func TrackProject(c *gin.Context) {
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
	err = model.TrackProject(accessToken, int(id))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
	return
}

//route: /project/untrack?id=<project id>
//function:stop tracking a project
//return code 200 if succeeded,else failed
func UntrackProject(c *gin.Context) {
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
	err = model.UntrackProject(accessToken, int(id))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
	return
}
