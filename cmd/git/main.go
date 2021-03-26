package main

import (
	"fmt"
	"graduation_design/internal/pkg/git"

)

func main() {
	res,err:=git.NewGit("git@127.0.0.1:172317/team-projects/grouptest.git","",true)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
}
