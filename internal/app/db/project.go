package db

import (
	"graduation_design/internal/pkg/logs"

	"github.com/jinzhu/gorm"
)

type ProjectDB struct {
	ID                int    `json:"id" gorm:"primaryKey;column:id"`
	Name              string `json:"name" gorm:"column:name"`
	NameWithNameSpace string `json:"name_with_namespace" gorm:"column:name_with_namespace"`
	WebUrl            string `json:"web_url" gorm:"column:web_url"`
	SSHUrl            string `json:"ssh_url_to_repo" gorm:"column:ssh_url_to_repo"`
}

func (p *ProjectDB) TableName() string {
	return "project_tracked"
}

//untested
func (p *ProjectDB) DeleteProject() error {
	logs.Info("Delete Project From DB,id %d,name %s", p.ID, p.NameWithNameSpace)
	return db.Delete(p).Error
}

func (p *ProjectDB) SaveProject() error {
	logs.Info("Save Project to DB,id %d,name %s", p.ID, p.NameWithNameSpace)
	return db.Create(p).Error
}

func FindProjectByID(id int) (ProjectDB, error) {
	var res = ProjectDB{}
	err := db.Where("id=?", id).First(&res).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			logs.Error("FindProjectByID:%s", err)
		}
		return res, err
	}
	return res, nil

}

func GetAllProjects() []ProjectDB {
	var res = make([]ProjectDB, 0)
	db.Find(&res)
	return res
}
