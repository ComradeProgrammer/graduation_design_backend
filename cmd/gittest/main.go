package main

import (
	"fmt"
	"graduation_design/internal/pkg/git"
)

func main() {
	res, err := git.NewGit("git@127.0.0.1:172317/team-projects/grouptest.git", "", true)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer res.Clear()

	// files, err := res.GetAllTrackedFiles()

	// res.BlameAllFile(files, func(filename, blame string) {
	// 	if blame==""{
	// 		return
	// 	}
	// 	fmt.Println(git.ResolveBlame(blame))
	// })

	err=res.ReadAllCommit(func(gitlog string){
		fmt.Println(gitlog)
		res,_:=git.ResolveGitlog(gitlog)
		fmt.Println(res)

	})
	if err!=nil{
		fmt.Println(err.Error())
	}

	return
}
