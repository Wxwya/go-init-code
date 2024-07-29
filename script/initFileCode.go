package script

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"xwya/model"
	"xwya/utils"

	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/duke-git/lancet/v2/strutil"
)

type Folders struct {
	ProjectName string
	Children    []Folders
}

var f = Folders{
	ProjectName: "template",
	Children: []Folders{
		{ProjectName: "api", Children: []Folders{
			{ProjectName: "v1"},
		}},
		{ProjectName: "model"},
		{ProjectName: "config"},
		{ProjectName: "cmd"},
		{ProjectName: "enetity"},
		{ProjectName: "router"},
		{ProjectName: "server"},
		{ProjectName: "utils"},
	},
}

func InitProject(data *model.Project) (err error) {
	rootDir, _ := os.Getwd()

	url := filepath.Join(rootDir, f.ProjectName, data.ProjectName)
	fmt.Println(url)
	if err = initProjectDir(url); err != nil {
		utils.WriteLog(err)
		return
	}
	if err = initModelFile(url, data.Forms, data.ProjectName); err != nil {
		utils.WriteLog(err)
		return
	}
	return
}

func initModelFile(url string, models []model.Form, pn string) (err error) {
	// filepath.Join(url, "model", model.FormName+".go")
	for _, model := range models {
		name := strutil.CamelCase(model.FormName)
		rd := filepath.Join(url, "model", name+".go")
		if model.IsChange != nil && *model.IsChange {
			continue
		}
		if fileutil.IsExist(rd) {
			if err = fileutil.RemoveFile(rd); err != nil {
				return
			}
		}

		file, err := os.Create(rd)
		if err != nil {
			return err
		}
		defer file.Close()
		str, isSqlType := createStruct(model.FormName, model.Fields)
		if isSqlType {
			str = "import (\n\"" + pn + "/enetity\"\n\"" + pn + "/enetity/sqltype\"\n)\n\n" + str
		}
		if _, err := file.WriteString("package model \n\n" + str); err != nil {
			return err
		}
		if err := formatGoFile(rd); err != nil {
			return err
		}
	}
	return
}
func createStruct(name string, fields []model.Field) (str string, isSqlType bool) {
	structName := strutil.UpperFirst(strutil.CamelCase(name))
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("type %s struct {\n", structName))

	for _, field := range fields {
		isSqlType = isSqlType || strings.Contains(field.FieldType, "sqltype")
		sb.WriteString(formatField(field) + "\n")
	}

	sb.WriteString("enetity.Global\n}\n")
	return sb.String(), isSqlType
}

func formatField(field model.Field) string {
	var sb strings.Builder

	fieldName := strutil.UpperFirst(strutil.CamelCase(field.FieldName))
	sb.WriteString("\t" + fieldName + "\t")

	if field.IsPointer != nil && *field.IsPointer {
		sb.WriteString("*")
	}
	sb.WriteString(field.FieldType + "\t" + "`json:\"")

	if field.IsValNull != "" {
		if field.IsValNull == "-" {
			sb.WriteString("-")
		} else {
			sb.WriteString(strutil.SnakeCase(strutil.CamelCase(field.FieldName)))
			if field.IsValNull == "omitempty" {
				sb.WriteString(",omitempty")
			}
		}
	} else {
		sb.WriteString(strutil.SnakeCase(strutil.CamelCase(field.FieldName)))
	}
	sb.WriteString("\" \tgorm:\"")

	sb.WriteString(formatGormTags(field))
	sb.WriteString("\" `")

	return sb.String()
}

func formatGormTags(field model.Field) string {
	var sb strings.Builder

	if field.IsPrimaryKey != nil && *field.IsPrimaryKey {
		sb.WriteString("primaryKey;")
	}
	if field.AutoIncrement != nil && *field.AutoIncrement {
		sb.WriteString("autoIncrement;")
	}
	if field.IsUnique != nil && *field.IsUnique {
		sb.WriteString("unique;")
	}
	if field.IsNull != nil && *field.IsNull {
		sb.WriteString("not null;")
	}
	if field.Embedded != nil && *field.Embedded {
		sb.WriteString("embedded;")
	}
	if field.IsQuery != nil && *field.IsQuery {
		if field.IndexName != "" {
			sb.WriteString(fmt.Sprintf("index:%s;", field.IndexName))
		} else {
			sb.WriteString("index;")
		}
	}
	if field.FieldSize != nil && *field.FieldSize != 0 {
		sb.WriteString(fmt.Sprintf("size:%d;", *field.FieldSize))
	}
	if field.DefaultValue != "" {
		sb.WriteString(fmt.Sprintf("default:%s;", field.DefaultValue))
	}
	if field.ForeignKey != "" {
		sb.WriteString(fmt.Sprintf("foreignKey:%s;", strutil.UpperFirst(strutil.CamelCase(field.ForeignKey))))
	}
	if field.References != "" {
		sb.WriteString(fmt.Sprintf("references:%s;", strutil.UpperFirst(strutil.CamelCase(field.References))))
	}
	if field.Many != "" {
		sb.WriteString(fmt.Sprintf("many2many:%s;", strutil.SnakeCase(field.Many)))
	}
	if field.JoinReferences != "" {
		sb.WriteString(fmt.Sprintf("joinReferences:%s;", strutil.UpperFirst(strutil.CamelCase(field.JoinReferences))))
	}
	if field.JoinForeignKey != "" {
		sb.WriteString(fmt.Sprintf("joinForeignKey:%s;", strutil.UpperFirst(strutil.CamelCase(field.JoinForeignKey))))
	}
	if field.Constraint != nil && *field.Constraint {
		sb.WriteString("constraint:")
		if field.OnUpdate != "" {
			sb.WriteString(fmt.Sprintf("onUpdate:%s,", field.OnUpdate))
		}
		if field.OnDelete != "" {
			sb.WriteString(fmt.Sprintf("onDelete:%s;", field.OnDelete))
		}
	}
	if field.EmbeddedPrefix != "" {
		sb.WriteString(fmt.Sprintf("embeddedPrefix:%s;", field.EmbeddedPrefix))
	}
	if field.Precision != "" {
		sb.WriteString(fmt.Sprintf("precision:%s;", field.Precision))
	}
	if field.Scale != "" {
		sb.WriteString(fmt.Sprintf("scale:%s;", field.Scale))
	}
	if field.Polymorphic != "" {
		sb.WriteString(fmt.Sprintf("polymorphic:%s;", field.Polymorphic))
	}
	if field.PolymorphicValue != "" {
		sb.WriteString(fmt.Sprintf("polymorphicValue:%s;", field.PolymorphicValue))
	}
	if field.Check != "" {
		sb.WriteString(fmt.Sprintf("check:%s;", field.Check))
	}

	return sb.String()
}

func initProjectDir(url string) (err error) {
	if err = createFolder(url); err != nil {
		return
	}
	return forEachCreateFolder(url, f.Children)

}
func forEachCreateFolder(url string, folders []Folders) error {
	for _, folder := range folders {
		if err := createFolder(filepath.Join(url, folder.ProjectName)); err != nil {
			return err
		}
		if len(folder.Children) > 0 {
			if err := forEachCreateFolder(filepath.Join(url, folder.ProjectName), folder.Children); err != nil {
				return err
			}
		}
	}
	return nil
}
func createFolder(s string) (err error) {
	if fileutil.IsExist(s) {
		return
	}
	return fileutil.CreateDir(s)
}
func formatGoFile(filename string) error {
	cmd := exec.Command("gofmt", "-w", filename)
	return cmd.Run()
}
