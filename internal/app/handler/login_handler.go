package handler

import (
	"graduation_design/internal/app/model"
	"graduation_design/internal/app/oauth"
	"graduation_design/internal/pkg/logs"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//todo:fix refresh session

func Ping(c *gin.Context) {
	session := sessions.Default(c)
	//TODO：REMOVE DEBUG INFO 
	accessToken,_:=session.Get("access_token").(string)
	refreshToken,_:=session.Get("refresh_token").(string)
	userName,_:=session.Get("user_name").(string)
	userId,_:=session.Get("id").(int)
	//avator:=session.Get("avator").(string)
	isAdmin,_:=session.Get("is_admin").(bool)
	
	c.JSON(200, gin.H{
		"message": "pong",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"username":userName,
		"userId":userId,
		"isAdmin":isAdmin,

	})
}

func Login(c *gin.Context) {
	session := sessions.Default(c)
	accessToken,_:=session.Get("access_token").(string)
	refreshToken,_:=session.Get("refresh_token").(string)
	if accessToken==""||refreshToken==""{
		//haven't logged in,jump to login
		c.Redirect(302, oauth.GenOauthUrl())
	}
	//have already logined，jump to main page
	//TODO:Implement
	
}

func Oauth(c *gin.Context) {
	code := c.Query("code")
	accessToken, refreshToken, err := oauth.OauthGetToken(code)
	if err != nil {
		c.JSON(401, gin.H{
			"message":"unauthorized",
		})
		return
	}
	
	//load user data to session
	currentUser,err:=model.GetCurrentUser(accessToken)
	if err!=nil{
		c.JSON(401, gin.H{
			"message":"unauthorized",
		})
		return
	}
	//save to session
	session := sessions.Default(c)
	session.Set("access_token",accessToken)
	session.Set("refresh_token",refreshToken)
	session.Set("user_name",currentUser.UserName)
	session.Set("id",currentUser.Id)
	session.Set("avator",currentUser.Avatar)
	session.Set("is_admin",currentUser.IsAdmin)
	session.Save()
	logs.Info("oauth:session saved,token:%s,refresh_token:%s",accessToken,refreshToken)
	//have already logined,
	//TODO:Implement jump to main page
	//TODO: remove debug output 
	c.JSON(200, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"id":currentUser.Id,
		"user_name":currentUser.UserName,
		"is_admin":currentUser.IsAdmin,
	})

}

func Test(c *gin.Context){
	session := sessions.Default(c)
	accessToken,_:=session.Get("access_token").(string)
	projectIDStr:=c.Query("projectid")
	projectID,_:=strconv.Atoi(projectIDStr)
	err:=model.CheckAndCreateLabels(accessToken,projectID)
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