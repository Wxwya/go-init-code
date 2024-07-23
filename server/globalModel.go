package server

import (
	"xwya/model"
	"xwya/model/global"
	"xwya/server/dop"
)

func GenerateGlobalModel(p *model.GlobalModel) error {
	if p.ID == 0 {
		return Db.Create(p).Error
	}
	return Db.Model(&model.GlobalModel{}).Where(&model.GlobalModel{ID: p.ID}).Updates(p).Error
}

func DeleteGlobalModel(p *map[string][]int) error {
	return Db.Where("id in (?)", (*p)["ids"]).Delete(&model.GlobalModel{}).Error
}

func GetGlobalModelList(queryInfo *global.Page) (*[]model.GlobalModel, *int64, error) {
	// 查询条件
	var globalModels []model.GlobalModel
	var total int64
	totalDb := Db.Model(&model.GlobalModel{})
	db := Db
	if err := totalDb.Count(&total).Error; err != nil {
		return &globalModels, &total, err
	}
	dop.AddPagination(db, queryInfo.PageNum, queryInfo.PageSize, "desc")
	if err := db.Find(&globalModels).Error; err != nil {
		return &globalModels, &total, err
	}
	return &globalModels, &total, nil
}

func GetGlobalModelInfo(id string) (*model.GlobalModel, error) {
	var globalModel model.GlobalModel
	if err := Db.Where("id = ?", id).First(&globalModel).Error; err != nil {
		return nil, err
	}
	return &globalModel, nil
}
