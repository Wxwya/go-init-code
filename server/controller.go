package server

import (
	"xwya/model"
	"xwya/model/repository"
	"xwya/server/dop"
)

func GenerateController(data *model.Controller) error {
	if data.ID == 0 {
		return Db.Create(data).Error
	}
	return Db.Where(&model.Controller{ID: data.ID}).Updates(data).Error
}

func DeleteController(ids *map[string][]int) error {
	return Db.Where("id in (?)", (*ids)["ids"]).Delete(&model.Controller{}).Error
}

func GetControllerInfo(id string) (*model.Controller, error) {
	var info model.Controller
	if err := Db.Where("id=?", id).First(&info).Error; err != nil {
		return &info, err
	}
	return &info, nil
}

func GetControllerList(page *repository.QueryController) (*[]model.Controller, *int64, error) {
	var list []model.Controller
	var total int64
	if err := Db.Model(&model.Controller{}).Where("version like ? ", "%"+page.Version+"%").Count(&total).Error; err != nil {
		return &list, &total, err
	}
	db := Db.Where("version like ? ", "%"+page.Version+"%")
	dop.AddPagination(db, page.PageNum, page.PageSize, "desc")
	if err := db.Find(&list).Error; err != nil {
		return &list, &total, err
	}

	return &list, &total, nil
}
