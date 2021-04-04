package git

import (
	"fmt"
	"regexp"
	"strconv"
)

type GitCommit struct {
	Hash        string `json:"hash"`
	Author      string `json:"author"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Title       string `json:"title"`
	FileChanged int    `json:"file_changed"`
	Insertions  int    `json:"insertions"`
	Deletions   int    `json:"deletions"`
	Total       int    `json:"total"`
}

var testString = `
c87656a 2962928213@qq.com 2021-03-27 23:24:52 +0800 fix: authorization
 12 files changed, 454 insertions(+), 91 deletions(-)

a8625b2 2962928213@qq.com 2021-03-27 14:49:18 +0800 feat: add contribute analysis
 8 files changed, 207 insertions(+), 13 deletions(-)

fc9132b 2962928213@qq.com 2021-03-26 21:00:14 +0800 add exec commandline
 5 files changed, 81 insertions(+)
 `

func ResolveGitlog(gitlog string) (GitCommit, error) {
	var res = GitCommit{}
	//(1:hash)(2:author)(3:date)(4:time)+0800(5:title)\n
	//(6:num) files changed,
	var regex = `(\S+)\s*(\S+)\s*(\S+)\s*(\S+)\s*\S+\s*(.*)\n` +
		`(?:\s*(\d*)\s* files? changed)?,?(?:\s*(\d*)\s* insertions?\(\+\))?,?(?:\s*(\d*)\s* deletions?\(\-\))?`
	reg, err := regexp.Compile(regex)
	if err != nil {
		return res, err
	}
	var tmp = reg.FindStringSubmatch(gitlog)
	if tmp == nil {
		return res, fmt.Errorf("unmatched line:%s", gitlog)
	}
	res.Hash = tmp[1]
	res.Author = tmp[2]
	res.Date = tmp[3]
	res.Time = tmp[4]
	res.Title = tmp[5]
	res.FileChanged, err = strconv.Atoi(tmp[6])
	if err != nil {
		res.FileChanged = 0
	}
	res.Insertions, err = strconv.Atoi(tmp[7])
	if err != nil {
		res.Insertions = 0
	}
	res.Deletions, err = strconv.Atoi(tmp[8])
	if err != nil {
		res.Deletions = 0
	}
	res.Total = res.Insertions - res.Deletions
	//fmt.Printf("debug:'%s %s %s'\n",tmp[6],tmp[7],tmp[8])
	return res, nil
}
