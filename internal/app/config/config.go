package config

// config for this app
const (

	//APPHOST ip of this app
	APPHOST string = "http://127.0.0.1"
	//APPPORT listening port of this app
	APPPORT string = ":3333"
)

//config for gitlab api
const (
	//GITLABURL url of gitlab
	GITLABURL string = "http://127.0.0.1"
	//APPID appid for calling gitlab api
	APPID string = "ef7e8b5810623eb01bb81af58a303b12ee0afaf07485ff09cde2a89e38dcf475"
	//SECRET secret for calling gitlab api
	APPSECRET string = "28a15f0b219038d4714de031bf307a134db06fa362bfd94e230765a04d987206"
	//REDIRECTROUTE route for gitlab oauth redirect
	REDIRECTROUTE string = "/oauth/callback"
	//REDIRECTURI actual url for gitlab oauth redirect
	REDIRECTURI string = APPHOST + APPPORT + REDIRECTROUTE
)
