package server

import (
	"xwya/model"
	"xwya/model/repository"
	"xwya/server/dop"
)

func GenerateMsgCode(p *model.MsgCode) error {
	if p.ID == 0 {
		return Db.Create(p).Error
	}
	return Db.Model(&model.MsgCode{}).Where(&model.MsgCode{ID: p.ID}).Updates(p).Error
}

func DeleteMsgCode(p *map[string][]int) error {
	return Db.Where("id in (?)", (*p)["ids"]).Delete(&model.MsgCode{}).Error
}

func GetMsgCodeList(queryInfo *repository.QueryMsgCode) (*[]model.MsgCode, *int64, error) {
	// 查询条件
	var msgCodes []model.MsgCode
	var total int64
	if err := Db.Model(&model.MsgCode{}).
		Where("key like ? and sql_host = ?  ", "%"+queryInfo.Key+"%", queryInfo.Code).
		Count(&total).Error; err != nil {
		return &msgCodes, &total, err
	}

	db := Db.Where("key like ? and sql_host = ?  ", "%"+queryInfo.Key+"%", queryInfo.Code)
	dop.AddPagination(db, queryInfo.PageNum, queryInfo.PageSize, "desc")
	if err := db.Find(&msgCodes).Error; err != nil {
		return &msgCodes, &total, err
	}
	return &msgCodes, &total, nil
}

func GetMsgCodeInfo(id string) (*model.MsgCode, error) {
	var msgCode model.MsgCode
	if err := Db.Where("id = ?", id).First(&msgCode).Error; err != nil {
		return nil, err
	}
	return &msgCode, nil
}
