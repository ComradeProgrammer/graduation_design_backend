package handler

import (
	"graduation_design/internal/app/oauth"
	"graduation_design/internal/pkg/logs"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	session := sessions.Default(c)
	//TODO：REMOVE DEBUG INFO 
	accessToken,_:=session.Get("access_token").(string)
	refreshToken,_:=session.Get("refresh_token").(string)
	c.JSON(200, gin.H{
		"message": "pong",
		"access_token":  accessToken,
		"refresh_token": refreshToken,

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
		c.JSON(401, gin.H{})
		return
	}
	session := sessions.Default(c)
	session.Set("access_token",accessToken)
	session.Set("refresh_token",refreshToken)
	session.Save()
	logs.Info("oauth:session saved,token:%s,refresh_token:%s",accessToken,refreshToken)
	//have already logined，jump to main page
	//TODO:Implement
	//TODO: remove debug output 

	c.JSON(200, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}
