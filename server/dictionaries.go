package server

import (
	"fmt"
	"xwya/model"
	"xwya/model/repository"
	"xwya/server/dop"
)

func GenerateDictionary(data *model.Dictionaries) error {
	if data.ID == 0 {
		return Db.Create(data).Error
	}
	return Db.Where(&model.Dictionaries{ID: data.ID}).Updates(data).Error
}

func DeleteDictionary(p map[string][]int) error {
	return Db.Where("id in (?)", p["ids"]).Delete(&model.Dictionaries{}).Error
}

func GetDictionaryList(page *repository.QueryDictionary) (*[]model.Dictionaries, *int64, error) {
	var data []model.Dictionaries
	var total int64
	totalDb := Db.Model(&model.Dictionaries{})
	listDb := Db
	fmt.Println(page.Code, "code")
	if page.Code != 0 {
		totalDb = totalDb.Where(&model.Dictionaries{Code: page.Code})
		listDb = listDb.Where(&model.Dictionaries{Code: page.Code})
	}
	if err := totalDb.Count(&total).Error; err != nil {
		return &data, &total, err
	}
	dop.AddPagination(listDb, page.PageNum, page.PageSize, "desc")
	if err := listDb.Find(&data).Error; err != nil {
		return &data, &total, err
	}
	return &data, &total, nil
}

func GetDictionaryInfo(id string) (*model.Dictionaries, error) {
	var data model.Dictionaries
	if err := Db.Where("id = ?", id).First(&data).Error; err != nil {
		return &data, err
	}
	return &data, nil
}
