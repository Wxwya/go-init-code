package server

import (
	"xwya/entity"
	"xwya/model"
	"xwya/server/dop"
)

func GenerateDictionary(data *model.Dict) error {
	if data.ID == 0 {
		return Db.Create(&data).Error
	}
	return Db.Where(&model.Dict{ID: data.ID}).Updates(data).Error
}

func DeleteDictionary(p map[string][]int) error {
	return Db.Where("id in (?)", p["ids"]).Delete(&model.Dict{}).Error
}

func GetDictionaryList(page *entity.QueryDictionary) (*[]model.Dict, *int64, error) {
	var data []model.Dict
	var total int64
	totalDb := Db.Model(&model.Dict{})
	listDb := Db
	dop.IsInt(totalDb, "code", page.Code)
	dop.IsInt(listDb, "code", page.Code)
	if err := totalDb.Count(&total).Error; err != nil {
		return &data, &total, err
	}
	dop.AddPagination(listDb, page.PageNum, page.PageSize, "desc")
	if err := listDb.Find(&data).Error; err != nil {
		return &data, &total, err
	}
	return &data, &total, nil
}

func GetDictionaryInfo(id string) (*model.Dict, error) {
	var data model.Dict
	if err := Db.Where("id = ?", id).First(&data).Error; err != nil {
		return &data, err
	}
	return &data, nil
}

func Cheshi() []string {
	var data []string
	Db.Model(&model.Dict{}).Pluck("value", &data)
	return data
}
