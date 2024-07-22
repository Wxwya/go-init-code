package server

import (
	"xwya/model"
	"xwya/model/repository"
	"xwya/server/dop"
	"xwya/utils/msg"
)

func ValidateUser(Account, Password string) (int, *model.User) {
	var user model.User
	if err := Db.Where("account=?", Account).First(&user).Error; err != nil {
		return msg.Error_UserNotExist, &user
	}
	if user.ID == 0 {
		return msg.Error_UserNotExist, &user
	}
	if dop.EncryptPassword(Password) != user.Password {
		return msg.Error_UserPasswordWrong, &user
	} else {
		user.Password = ""
		return msg.Success, &user
	}
}
func GetUserInfo(id uint) (*model.User, error) {
	var user model.User
	if err := Db.Where("id=?", id).First(&user).Error; err != nil {
		return nil, err
	}
	user.Password = ""
	return &user, nil
}
func CreateAndUpDateUser(user *model.User) error {
	if user.ID > 0 {
		if err := Db.Where("id=?", user.ID).Updates(user).Error; err != nil {
			return err
		}
		return nil
	}
	user.Password = dop.EncryptPassword(user.Password)
	if err := Db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
func ChangePass(id uint, password string) error {
	if err := Db.Model(&model.User{}).Where("id=?", id).Update("password", dop.EncryptPassword(password)).Error; err != nil {
		return err
	}
	return nil
}
func GetUserList(info *repository.QueryUserList) (*[]model.User, int64, error) {
	var userList []model.User
	var total int64
	totalDb := Db.Model(&model.User{}).Where("username like ?", "%"+info.Username+"%")
	listDb := Db.Where("username like ?", "%"+info.Username+"%")
	if err := totalDb.Count(&total).Error; err != nil {
		return &userList, total, err
	}

	dop.AddPagination(listDb, info.Page.PageNum, info.Page.PageSize, "desc")
	if err := listDb.Find(&userList).Error; err != nil {
		return &userList, total, err
	}
	return &userList, total, nil
}
func DeleteUser(id []int) error {
	if err := Db.Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
