package script

import "xwya/model"

type Func func() *[]model.Field

var HandleField = map[string]Func{
	"User":       GetUserFormFieldData,
	"Role":       GetRoleFormFieldData,
	"Permission": GetPermissionFormFieldData,
	"Log":        GetLogFormFieldData,
	"Dict":       GetDictFormFieldData,
}

func GetDefaultFormData(id uint) *[]model.Form {
	var DefaultFormData = []model.Form{
		{ProjectID: id, FormName: "User", FormComment: "用户表"},
		{ProjectID: id, FormName: "Role", FormComment: "角色表"},
		{ProjectID: id, FormName: "Permission", FormComment: "权限表"},
		{ProjectID: id, FormName: "Log", FormComment: "日志表"},
		{ProjectID: id, FormName: "Dict", FormComment: "字典表"},
	}
	return &DefaultFormData
}
func GetFormFieldData(name string) *[]model.Field {
	return HandleField[name]()
}
func GetUserFormFieldData() *[]model.Field {
	return nil
}
func GetRoleFormFieldData() *[]model.Field       { return nil }
func GetPermissionFormFieldData() *[]model.Field { return nil }
func GetLogFormFieldData() *[]model.Field        { return nil }
func GetDictFormFieldData() *[]model.Field       { return nil }
