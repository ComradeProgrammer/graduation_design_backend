package model

import (
	"fmt"
	"graduation_design/internal/app/config"
	"graduation_design/internal/app/db"
	"graduation_design/internal/pkg/logs"
	"regexp"
)

func CheckQuality(token string, projectID int) (map[string]interface{}, error) {
	pipeline, pipelinestatus, err := getLatestFinishedPipelineOnMasterBranch(token, projectID)
	if err != nil {
		return nil, err
	}

	jobs, err := getPipelineJob(token, projectID, pipeline)
	if err != nil {
		return nil, err
	}
	// return map[string]interface{}{
	// 	"status": pipelinestatus,
	// 	"data":   jobs,
	// }, nil
	build := checkBuild(jobs)
	test := checkTest(jobs)
	lints := checkLint(jobs, projectID, token)
	coverages := checkCoverage(jobs, projectID, token)
	return map[string]interface{}{
		"status":   pipelinestatus,
		"build":    build,
		"test":     test,
		"lint":     lints,
		"coverage": coverages,
	}, nil
}

func GetAllRegex(projectID int) map[string]interface{} {
	coverage, err := db.GetCoverageRegex(projectID)
	if err != nil {
		logs.Error("db.GetCoverageRegex:%s", err)
	}
	lints, err := db.GetLintRegex(projectID)
	if err != nil {
		logs.Error("db.GetLintRegex:%s", err)
	}
	return map[string]interface{}{
		"coverage": coverage,
		"lint":     lints,
	}
}

func checkBuild(jobs []Job) bool {
	for _, j := range jobs {
		if j.Name == config.BUILD && j.Status == "success" {
			return true
		}
	}
	return false
}

func checkTest(jobs []Job) bool {
	for _, j := range jobs {
		if j.Name == config.TEST && j.Status == "success" {
			return true
		}
	}
	return false
}

func checkLint(jobs []Job, projectID int, token string) []map[string]interface{} {
	for _, j := range jobs {
		if j.Name == config.LINT {
			coverage, err := db.GetLintRegex(projectID)
			if err != nil {
				logs.Error("checkLint:%s", err)
			}
			log, err := GetJobLog(token, projectID, j.Id)
			if err != nil {
				logs.Error("checkLint %s", err)
			}
			var res = make([]map[string]interface{}, 0)
			for _, r := range coverage {
				reg, err := regexp.Compile(r.Regex)
				if err != nil {
					logs.Error("checkLint:%s", err)
				}
				var tmp [][]byte = reg.FindAll([]byte(log), -1)
				var lints = make([]string, 0)
				for _, arr := range tmp {
					lints = append(lints, string(arr))
				}
				res = append(res, map[string]interface{}{
					"description": r,
					"res":         lints,
				})
			}
			return res

		}
	}
	return nil
}

func checkCoverage(jobs []Job, projectID int, token string) []map[string]interface{} {
	for _, j := range jobs {
		if j.Name == config.COVERAGE {
			coverage, err := db.GetCoverageRegex(projectID)
			if err != nil {
				logs.Error("checkLint:%s", err)
			}
			log, err := GetJobLog(token, projectID, j.Id)
			logs.Info("log:%s",log)
			if err != nil {
				logs.Error("checkLint %s", err)
			}
			var res = make([]map[string]interface{}, 0)
			for _, r := range coverage {
				reg, err := regexp.Compile(r.Regex)
				if err != nil {
					logs.Error("checkLint:%s", err)
				}
				var tmp [][]byte = reg.FindAll([]byte(log), -1)
				var lints = make([]string, 0)
				for _, arr := range tmp {
					lints = append(lints, string(arr))
				}
				res = append(res, map[string]interface{}{
					"description": r,
					"res":         lints,
				})
			}
			return res

		}
	}
	return nil
}

func CreateRegex(projectID int, regex string, regexType string, comment string) error {
	if regexType != db.COVERAGE && regexType != db.LINT {
		return fmt.Errorf("invalid regex_type")
	}
	var tmp = db.Regex{
		ProjectID: projectID,
		Regex:     regex,
		RegexType: regexType,
		Comment:   comment,
	}
	return tmp.SaveRegex()
}

func DeleteRegex(id int) error {
	return db.DeleteRegex(id)
}

func findIndexByRegex(regexStr string, str string) (int, int, error) {
	regex, err := regexp.Compile(regexStr)
	if err != nil {
		return -1, -1, err
	}
	res := regex.FindStringIndex(str)
	if len(res) == 0 {
		return -1, -1, fmt.Errorf("findIndexByRegex:no match %s", regexStr)
	}
	return res[0], res[1], nil
}

func findSubStringByRegex(startRegex, endRegex string, str string) (string, error) {
	_, start, err := findIndexByRegex(startRegex, str)
	if err != nil {
		return "", err
	}
	end, _, err := findIndexByRegex(endRegex, str)
	if err != nil {
		return "", err
	}
	return str[start:end], nil
}
