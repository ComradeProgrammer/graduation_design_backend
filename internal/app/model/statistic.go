package model

import (
	"graduation_design/internal/app/config"
	"graduation_design/internal/pkg/git"
)

func AnalysisProject(token string, projectId int) (map[string]interface{}, error) {
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
