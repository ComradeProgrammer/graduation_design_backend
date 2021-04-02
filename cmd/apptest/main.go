package main

import (
	"fmt"
	"graduation_design/internal/app/config"
	"graduation_design/internal/pkg/request"
)


func main(){
	token:="2852b707ac3a644f68688d19ce84cc406a20be0d144a0e0cb05e509b55987e78"
	res,err:=request.RequestForGitlabPagination(config.GITLABAPIURL+"/projects/3/issues?per_page=1","GET",
	map[string]string{
		"Authorization": "Bearer " + token,
	},
	"",
	5,)
	if err!=nil{
		fmt.Print(err.Error())
	}
	fmt.Println(len(res))
	for _,tmp:=range res{
		fmt.Println(tmp)
	}
	
}