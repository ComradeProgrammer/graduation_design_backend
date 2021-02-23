package handler

import (
	"graduation_design/internal/app/oauth"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
		"TEST":    "TEST",
	})
}

func Login(c *gin.Context) {
	c.Redirect(302, oauth.GenOauthUrl())
}

func Oauth(c *gin.Context) {
	code := c.Query("code")
	accessToken, refreshToken, err := oauth.OauthGetToken(code)
	if err != nil {
		c.JSON(401, gin.H{})
		return
	}
	c.JSON(200, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}
