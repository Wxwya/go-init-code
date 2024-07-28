package server

import (
	"xwya/entity"
	"xwya/model"
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

func GetMsgCodeList(queryInfo *entity.QueryMsgCode) (*[]model.MsgCode, *int64, error) {
	// 查询条件
	var msgCodes []model.MsgCode
	var total int64
	totalDb := Db.Model(&model.MsgCode{}).Where("`key` like ?  ", "%"+queryInfo.MsgKey+"%")
	db := Db.Where("msg_key like ? ", "%"+queryInfo.MsgKey+"%")

	dop.IsInt(totalDb, "code", queryInfo.Code)
	dop.IsInt(db, "code", queryInfo.Code)
	if err := totalDb.Count(&total).Error; err != nil {
		return &msgCodes, &total, err
	}

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
