package model

import (
	"graduation_design/internal/app/config"
	"graduation_design/internal/pkg/git"
)
//这次修改之后，我究竟需要怎样的需求呢？
/*
	1.项目概况 应当提供：
	项目代码总行数 项目总commit数 项目总issue数 项目总Disscution数 项目总MR数
	项目所有参与者 项目语言结构
	2. 开发者视角 对于每一个开发者展示 
		此人提交代码总行数
		此人提交的语言分布
		此人commit列表或概况图
		此人issue列表或概况图
		此人comment列表或概况图
	3. 语言视角 提供每一种语言的各位作者贡献分布
	4. gitlog commit分析 提供commit列表并进行问题分析
	5. 活跃度视角 提供近14天内的commmit issue discussions MR报表及分析

*/

func AnalysisProjectOverview(token string, projectId int) (map[string]interface{}, error) {
	prj, err := getProjectFromGitlab(token, projectId)
	if err != nil {
		return nil, err
	}
	g, err := git.NewGit(prj.SSHURL, config.CACHEDIR, true)
	if err != nil {
		return nil, err
	}
	defer g.Clear()
	files, err := g.GetAllTrackedFiles()

	var totalLines int = 0
	var linesByLanguage = make(map[string]int)
	var linesByAuthor = make(map[string]int)
	var linesByAuthorLanguage = make(map[string]map[string]int)
	var linesByLanguageByAuthor = make(map[string]map[string]int)

	g.BlameAllFile(files, func(filename, blame string) {
		info, err := git.ResolveBlame(blame)
		if err != nil {
			return
		}
		totalLines++

		if _, ok := linesByAuthor[info.Author]; !ok {
			linesByAuthor[info.Author] = 0
		}
		linesByAuthor[info.Author]++

		if _, ok := linesByLanguage[info.Suffix]; !ok {
			linesByLanguage[info.Suffix] = 0
		}
		linesByLanguage[info.Suffix]++

		if _, ok := linesByAuthorLanguage[info.Author]; !ok {
			linesByAuthorLanguage[info.Author] = make(map[string]int)
		}

		if _, ok := linesByAuthorLanguage[info.Author][info.Suffix]; !ok {
			linesByAuthorLanguage[info.Author][info.Suffix] = 0
		}
		linesByAuthorLanguage[info.Author][info.Suffix]++

		if _, ok := linesByLanguageByAuthor[info.Suffix]; !ok {
			linesByLanguageByAuthor[info.Suffix] = make(map[string]int)
		}
		if _, ok := linesByLanguageByAuthor[info.Suffix][info.Author]; !ok {
			linesByLanguageByAuthor[info.Suffix][info.Author] = 0
		}
		linesByLanguageByAuthor[info.Suffix][info.Author]++
	})
	return map[string]interface{}{
		"totalLines":              totalLines,
		"linesByLanguage":         linesByLanguage,
		"linesByAuthor":           linesByAuthor,
		"linesByAuthorByLanguage": linesByAuthorLanguage,
		"linesByLanguageByAuthor": linesByLanguageByAuthor,
	}, nil

}
