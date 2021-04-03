package model

import (
	"graduation_design/internal/app/config"
	"graduation_design/internal/pkg/git"
	"graduation_design/internal/pkg/logs"
)

//这次修改之后，我究竟需要怎样的需求呢？
/*
	1.项目概况 应当提供：
	项目代码总行数 项目总commit数 项目总issue数  项目总MR数
	项目所有参与者 项目语言结构
	2. 开发者视角 对于每一个开发者展示
		此人提交代码总行数
		此人提交的语言分布
		此人commit列表或概况图
		此人issue列表或概况图
		此人comment列表或概况图
	3. 语言视角 提供每一种语言的各位作者贡献分布
	4. gitlog commit分析 提供commit列表并进行问题分析
	5. 活跃度视角 提供近14天内的commmit issue discussions MR报表及自定义分析

*/

func AnalysisProjectOverview(token string, projectId int) (map[string]interface{}, error) {
	prj, err := getProjectFromGitlab(token, projectId)
	if err != nil {
		return nil, err
	}
	//go routine 1,for analysisByLabguage
	var languageChan=make(chan map[string]interface{})
	go func(){
		res,err:=analysisByLanguageAndCommitNum(prj.SSHUrl)
		if err!=nil{
			logs.Error("error when analysisByLanguage,%s",err.Error())
			languageChan<-map[string]interface{}{}
		}
		languageChan<-res
	}()
	
	//goroutine 2 for issue statistic
	var issueStatisticChan=make(chan int)
	go func(){
		res,err:= GetProjectIssueNum(token,projectId)
		if err!=nil{
			logs.Error("error when GetProjectIssueStatistic,%s",err.Error())
			issueStatisticChan<--1
		}
		issueStatisticChan<-res
	}()

	//goroutine3 for mr statistic
	var mrCountsChan=make(chan int)
	go func(){
		total,err:=GetProjectMrNum(token,projectId)
		if err!=nil{
			logs.Error("error when get project mr nums:%s",err.Error())
			mrCountsChan<--1
		}
		mrCountsChan<-total
	}()

	

	languageResult:=<-languageChan

	issueResult:=<-issueStatisticChan
	languageResult["issueNum"]=issueResult

	mrCount:=<-mrCountsChan
	languageResult["mrnum"]=mrCount


	return languageResult,nil

}

/*2. 开发者视角 对于每一个开发者展示
		此人提交代码总行数
		此人提交的语言分布
		此人commit列表或概况图
		此人issue列表或概况图
		此人comment列表或概况图*/

func AnalyzeUserView(token string,projectId int)(map[string]interface{},error){
	prj, err := getProjectFromGitlab(token, projectId)
	if err != nil {
		return nil, err
	}
	
	var languageChan=make(chan map[string]interface{})
	var commitsChan=make(chan []git.GitCommit)
	var gitlabCommitChan=make(chan map[string]GitlabCommit)
	//go routine 1,for analysisByLabguage
	go func(){
		res,commits,err:=analysisByLanguageAndCommitList(prj.SSHUrl)
		if err!=nil{
			logs.Error("error when analysisByLanguage,%s",err.Error())
			languageChan<-map[string]interface{}{}
			commitsChan<-nil
		}
		languageChan<-res
		commitsChan<-commits
	}()
	go func(){
		res,err:=GetAllCommitsInMap(token,projectId)
		if err!=nil{
			logs.Error("error when analysisByLanguage,%s",err.Error())
			gitlabCommitChan<-nil
		}
		gitlabCommitChan<-res
	}()

	languageResult:=<-languageChan
	commitResult:=<-commitsChan
	gitlabCommitResult:=<-gitlabCommitChan
	fullCommits,gitAuthorToGitlabUser:=mergeCommitsFromGitAndGitlab(commitResult,gitlabCommitResult)
	languageResult["commitList"]=fullCommits
	languageResult["gitAuthorToGitlabUser"]=gitAuthorToGitlabUser
	
	
	//logs.Warning("commits:%d,gitlab:%d",len(commitResult),len(gitlabCommitResult))

	return languageResult,nil

}

func AnalyzeLanguageView(token string, projectId int) (map[string]interface{}, error){
	prj, err := getProjectFromGitlab(token, projectId)
	if err != nil {
		return nil, err
	}
	res,err:=analysisByLanguageAndCommitNum(prj.SSHUrl)
	if err!=nil{
		logs.Error("error when analysisByLanguage,%s",err.Error())
		return nil,err
	}
	return res,nil
}


func AnalyzeCommitView(token string, projectId int)([]FullCommit,error){
	prj, err := getProjectFromGitlab(token, projectId)
	if err != nil {
		return nil, err
	}

	var commitsChan=make(chan []git.GitCommit)
	var gitlabCommitChan=make(chan map[string]GitlabCommit)
	go func(){
		res,err:=analysisByCommit(prj.SSHUrl)
		if err!=nil{
			logs.Error("error when analysisByLanguage,%s",err.Error())
		
		}
		
		commitsChan<-res
	}()
	go func(){
		res,err:=GetAllCommitsInMap(token,projectId)
		if err!=nil{
			logs.Error("error when analysisByLanguage,%s",err.Error())
			gitlabCommitChan<-nil
		}
		gitlabCommitChan<-res
	}()


	commitResult:=<-commitsChan
	gitlabCommitResult:=<-gitlabCommitChan
	fullCommits,_:=mergeCommitsFromGitAndGitlab(commitResult,gitlabCommitResult)
	return fullCommits,nil

}

func AnalyzeActivityView(token string,projectId int)(map[string]interface{},error){
	prj, err := getProjectFromGitlab(token, projectId)
	if err != nil {
		return nil, err
	}
	var issueChan=make(chan []map[string]interface{})
	var mrChan=make(chan []map[string]interface{})
	var languageChan=make(chan map[string]interface{})
	var commitsChan=make(chan []git.GitCommit)
	var gitlabCommitChan=make(chan map[string]GitlabCommit)
	//go routine 1,for analysisByLabguage
	go func(){
		res,commits,err:=analysisByLanguageAndCommitList(prj.SSHUrl)
		if err!=nil{
			logs.Error("error when analysisByLanguage,%s",err.Error())
			languageChan<-map[string]interface{}{}
			commitsChan<-nil
		}
		languageChan<-res
		commitsChan<-commits
	}()
	go func(){
		res,err:=GetAllCommitsInMap(token,projectId)
		if err!=nil{
			logs.Error("error when analysisByLanguage,%s",err.Error())
			gitlabCommitChan<-nil
		}
		gitlabCommitChan<-res
	}()
	go func(){
		res,err:=GetAllProjectIssueInObject(token,projectId)
		if err!=nil{
			logs.Error("get all issue failed when AnalysisByActivity")
			issueChan<-nil
		}
		issueChan<-res
	}()

	go func(){
		res,err:=GetAllProjectMrInObject(token,projectId)
		if err!=nil{
			logs.Error("get all issue failed when AnalysisByActivity")
			mrChan<-nil
		}
		mrChan<-res
	}()

	languageResult:=<-languageChan
	commitResult:=<-commitsChan
	gitlabCommitResult:=<-gitlabCommitChan
	fullCommits,gitAuthorToGitlabUser:=mergeCommitsFromGitAndGitlab(commitResult,gitlabCommitResult)
	
	allIssues:=<-issueChan
	allMrs:=<-mrChan
	return map[string]interface{}{
		"issues":allIssues,
		"mrs":allMrs,
		"language":languageResult,
		"commit":fullCommits,
		"gitAuthorToGitlabUser":gitAuthorToGitlabUser,
	},nil
}
	

func analysisByLanguageAndCommitList(ssh string)(map[string]interface{},[]git.GitCommit,error){
	g, err := git.NewGit(ssh, config.CACHEDIR, true)
	if err != nil {
		return nil,nil, err
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
	var commitList=make([]git.GitCommit,0)

	g.ReadAllCommit(func(gitlog string){
		res,_:=git.ResolveGitlog(gitlog)
		commitList = append(commitList, res)
	})
	return map[string]interface{}{
		"totalLines":              totalLines,
		"linesByLanguage":         linesByLanguage,
		"linesByAuthor":           linesByAuthor,
		"linesByAuthorByLanguage": linesByAuthorLanguage,
		"linesByLanguageByAuthor": linesByLanguageByAuthor,
	},commitList, nil
}

func analysisByLanguageAndCommitNum(ssh string)(map[string]interface{},error){
	g, err := git.NewGit(ssh, config.CACHEDIR, true)
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
	var commitNum=0

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
	g.ReadAllCommit(func(string){
		commitNum++
	})
	return map[string]interface{}{
		"totalLines":              totalLines,
		"linesByLanguage":         linesByLanguage,
		"linesByAuthor":           linesByAuthor,
		"linesByAuthorByLanguage": linesByAuthorLanguage,
		"linesByLanguageByAuthor": linesByLanguageByAuthor,
		"commitNum":commitNum,
	}, nil
}


func analysisByLanguage(ssh string)(map[string]interface{},error){
	g, err := git.NewGit(ssh, config.CACHEDIR, true)
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

func analysisByCommit(ssh string)([]git.GitCommit,error){
	g, err := git.NewGit(ssh, config.CACHEDIR, true)
	if err != nil {
		return nil, err
	}
	defer g.Clear()
	var commitList=make([]git.GitCommit,0)

	g.ReadAllCommit(func(gitlog string){
		res,_:=git.ResolveGitlog(gitlog)
		commitList = append(commitList, res)
	})
	return commitList,nil
}
