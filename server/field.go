package server

import (
	"xwya/entity"
	"xwya/model"
	"xwya/server/dop"
)

func GenerateField(data *model.Field) (err error) {
	if data.ID == 0 {
		return Db.Create(data).Error
	}
	return Db.Model(&data).Updates(&data).Error
}

func DeleteField(ids []uint) (err error) {
	return Db.Model(&model.Field{}).Delete(ids).Error
}

func GetFieldInfo(id any) (data *model.Field, err error) {
	return data, Db.First(&data, id).Error
}

func GetFieldList(page *entity.Page) (data *model.FieldDto, total int64, err error) {
	listDb := Db.Select("field.*, dict.code as field_type_name").Joins("left join dict on dict.id = field.field_type_id")
	err = Db.Model(&model.Field{}).Count(&total).Error
	if err != nil {
		return
	}
	dop.AddPagination(Db, page.PageNum, page.PageSize, "desc")
	return data, total, listDb.Find(&data).Error
}
