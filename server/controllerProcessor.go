package server

import (
	"xwya/model"
	"xwya/model/repository"
	"xwya/server/dop"
)

func GenerateControllerProcessor(data *model.ControllerProcessor) error {
	if data.ID == 0 {
		return Db.Create(data).Error
	}
	return Db.Where(&model.ControllerProcessor{ID: data.ID}).Updates(data).Error
}

func DeleteControllerProcessor(ids *map[string][]int) error {
	return Db.Where("id in (?)", (*ids)["ids"]).Delete(&model.ControllerProcessor{}).Error
}

func GetControllerProcessorInfo(id string) (*model.ControllerProcessor, error) {
	var info model.ControllerProcessor
	if err := Db.Where("id=?", id).First(&info).Error; err != nil {
		return &info, err
	}
	return &info, nil
}

func GetControllerProcessorList(page *repository.QueryControllerProcessor) (*[]model.ControllerProcessor, *int64, error) {
	var list []model.ControllerProcessor
	var total int64
	totalDb := Db.Model(&model.ControllerProcessor{})
	db := Db.Where("params like ? and func_name like ?", "%"+page.Params+"%", "%"+page.FuncName+"%")
	dop.IsBool(totalDb, "is_hooks", page.IsHooks)
	dop.IsBool(db, "is_hooks", page.IsHooks)
	if err := totalDb.Where("params like ? and func_name like ?", "%"+page.Params+"%", "%"+page.FuncName+"%").Count(&total).Error; err != nil {
		return &list, &total, err
	}
	dop.AddPagination(db, page.PageNum, page.PageSize, "desc")
	if err := db.Find(&list).Error; err != nil {
		return &list, &total, err
	}

	return &list, &total, nil
}
