package server

import (
	"xwya/entity"
	"xwya/model"
	"xwya/server/dop"
)

func GenerateForm(data *model.Form) error {
	if data.ID == 0 {
		data.Dict = []model.Dict{
			{DictLabel: data.FormName, Description: data.FormName + "表类型", DictValue: data.FormName, ProjectID: data.ProjectID},
			{DictLabel: "[]" + data.FormName, Description: data.FormName + "表类型", DictValue: "[]" + data.FormName, ProjectID: data.ProjectID},
		}
		return Db.Create(data).Error
	}
	return Db.Where(&model.Form{ID: data.ID}).Updates(data).Error
}

func DeleteForm(ids *map[string][]int) error {
	return Db.Where("id in (?)", (*ids)["ids"]).Delete(&model.Form{}).Error
}

func GetFormInfo(id string) (*model.Form, error) {
	var info model.Form
	if err := Db.Where("id=?", id).First(&info).Error; err != nil {
		return &info, err
	}
	return &info, nil
}

func GetFormList(page *entity.QueryForm) (*[]model.Form, *int64, error) {
	var list []model.Form
	var total int64
	if err := Db.Model(&model.Form{}).Where("project_id = ? and form_name like ?", page.ProjectID, "%"+page.FormName+"%").Count(&total).Error; err != nil {
		return &list, &total, err
	}
	db := Db.Where("project_id = ? and form_name like ?", page.ProjectID, "%"+page.FormName+"%")
	dop.AddPagination(db, page.PageNum, page.PageSize, "desc")
	if err := db.Find(&list).Error; err != nil {
		return &list, &total, err
	}

	return &list, &total, nil
}
