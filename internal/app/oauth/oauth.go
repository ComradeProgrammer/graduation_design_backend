package oauth

import (
	"fmt"
	"graduation_design/internal/app/config"
	"graduation_design/internal/pkg/logs"
	"graduation_design/internal/pkg/request"
)

//Oauth: get token after user have logged in through api and redirected to our website with code
func OauthGetToken(code string) (string, string, error) {
	// map[string]interface{}{
	// 	"client_id":config.APPID,
	// 	"client_secret":config.APPSECRET,
	// 	"code":code,
	// 	"grant_type":"authorization_code",
	// 	"redirect_uri":config.REDIRECTURI,
	// },

	status, data, err := request.StringForJson(
		genOauthTokenUrl(),
		"POST",
		nil,
		fmt.Sprintf("client_id=%s&client_secret=%s&code=%s&grant_type=authorization_code&redirect_uri=%s",
			config.APPID, config.APPSECRET, code, config.REDIRECTURI),
		5,
	)
	if err != nil {
		logs.Error("OauthGetToken Failed,%s", err)
		return "", "", err
	}
	if status != 200 {
		logs.Error("OauthGetToken Failed,Code %d", status)
		return "", "", fmt.Errorf("OauthGetToken Failed,Code %d", status)
	}
	logs.Info("code %d,body %v", status, data)
	accessToken, ok := data["access_token"].(string)
	if !ok {
		logs.Error("no access_token found,body %v", data)
		return "", "", fmt.Errorf("no access_tken found,body %v", data)
	}
	refreshToken, ok := data["refresh_token"].(string)
	if !ok {
		logs.Error("no refresh_token found,body %v", data)
		return "", "", fmt.Errorf("no refresh_token found,body %v", data)
	}
	return accessToken, refreshToken, nil
}

//GenOauthUrl generate url for redirecting users to gitlab login oage
func GenOauthUrl() string {
	return fmt.Sprintf(
		"%s/oauth/authorize?client_id=%s&redirect_uri=%s&response_type=code",
		config.GITLABURL,
		config.APPID,
		config.REDIRECTURI)
}

func genOauthTokenUrl() string {
	return fmt.Sprintf("%s/oauth/token", config.GITLABURL)
}
