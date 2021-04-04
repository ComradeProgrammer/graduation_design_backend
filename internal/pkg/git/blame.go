package git

import (
	"fmt"
	"path"
	"regexp"
)

type GitBlame struct {
	FileName string
	Suffix   string
	Author   string
	Date     string
	Time     string
}

func ResolveBlame(blame string) (GitBlame, error) {
	var res = GitBlame{}
	//group2:path group1:hash group3:author
	var regex = `(\S*)\s*(\S*)\s*\((\S*)\s*(\S*)\s*(\S*)`
	reg, err := regexp.Compile(regex)
	if err != nil {
		return res, err
	}
	var tmp = reg.FindStringSubmatch(blame)
	if tmp == nil {
		return res, fmt.Errorf("unmatched line:%s", blame)
	}
	res.FileName = tmp[2]
	res.Author = tmp[3]
	res.Date = tmp[4]
	res.Time = tmp[5]
	res.Suffix = path.Ext(tmp[2])
	return res, nil
}
