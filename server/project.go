package server

import (
	"xwya/model"
	"xwya/model/repository"
	"xwya/server/dop"
)

func GenerateProject(p *model.Project) error {
	if p.ID == 0 {
		return Db.Debug().Create(p).Error
	}
	return Db.Model(&model.Project{}).Where(&model.Project{ID: p.ID}).Updates(p).Error
}

func DeleteProject(p *map[string][]int) error {
	return Db.Where("id in (?)", (*p)["ids"]).Delete(&model.Project{}).Error
}

func GetProjectList(queryInfo *repository.QueryProject) (*[]model.Project, *int64, error) {
	// 查询条件
	var projects []model.Project
	var total int64
	if err := Db.Model(&model.Project{}).
		Where("project_name like ? and sql_host like ?  ", "%"+queryInfo.ProjectName+"%", "%"+queryInfo.SqlHost+"%").
		Count(&total).Error; err != nil {
		return &projects, &total, err
	}

	db := Db.Where("project_name like ? and sql_host like ?  ", "%"+queryInfo.ProjectName+"%", "%"+queryInfo.SqlHost+"%")
	dop.AddPagination(db, queryInfo.PageNum, queryInfo.PageSize, "desc")
	if err := db.Find(&projects).Error; err != nil {
		return &projects, &total, err
	}
	return &projects, &total, nil
}

func GetProjectInfo(id string) (*model.Project, error) {
	var project model.Project
	if err := Db.Where("id = ?", id).First(&project).Error; err != nil {
		return nil, err
	}
	return &project, nil
}
