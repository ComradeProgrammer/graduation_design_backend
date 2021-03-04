package db

import (
	"graduation_design/internal/pkg/logs"

	"github.com/jinzhu/gorm"
)

type Regex struct {
	ID        int    `json:"id" gorm:"primary_key;auto_increment"`
	ProjectID int    `json:"prohect_id" gorm:"column:project_id"`
	Regex     string `json:"regex" gorm:"column:regex"`
	RegexType string `json:"regex_type" gorm:"column:regex_type"`
	Comment   string `json:"comment" gorm:"column:comment"`
}

const COVERAGE = "coverage"
const LINT = "lint"

func (p *Regex) TableName() string {
	return "project_regex"
}

func GetCoverageRegex(projectID int) ([]Regex, error) {
	var res = make([]Regex, 0)
	err := db.Where("project_id=? AND regex_type=?", projectID, COVERAGE).Find(&res).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			logs.Error("GetCoverageRegex:%s", err)
		}
		return res, err
	}
	return res, nil
}
func GetLintRegex(projectID int) ([]Regex, error) {
	var res = make([]Regex, 0)
	err := db.Where("project_id=? AND regex_type=?", projectID, LINT).Find(&res).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			logs.Error("GetCoverageRegex:%s", err)
		}
		return res, err
	}
	return res, nil
}

func (p *Regex) SaveRegex() error {
	return db.Create(p).Error
}

func (p *Regex) UpdateRegex() error {
	return db.Save(p).Error
}
func DeleteRegex(id int) error {
	return db.Delete(&Regex{ID: id}).Error
}
