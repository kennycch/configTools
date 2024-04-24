package json

import (
	"config_tools/tools/excel"
	"fmt"
	"os"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// StructureJson 构建Json
//
//	@param excelFullPath excel文件路径
//	@param jsonPath json目录路径
//	@return error
func StructureJson(excelFullPath, jsonPath string) error {
	// 读取Excel
	file, err := excelize.OpenFile(excelFullPath)
	if err != nil {
		return err
	}
	// 获取sheet
	sheetMap := file.GetSheetMap()
	// 根据类型构建Json
	tableType := file.GetCellValue(sheetMap[1], "D1")
	mainLocal := &excel.Local{}
	mainLocal.GetLocal(0, 5)
	content := ""
	mainSheet := sheetMap[1]
	if tableType == "object" {
		// object仅读取首行数据
		content = getJsonByRow(mainLocal, file, mainSheet, true)
	} else {
		content += "["
		// 遍历行
		for {
			id := file.GetCellValue(mainSheet, mainLocal.GetLocal(0, 0))
			if id != "" {
				content += fmt.Sprintf("%s,", getJsonByRow(mainLocal, file, mainSheet, true))
				mainLocal.GetLocal(-1, 1)
			} else {
				break
			}
		}
		if len(content) > 1 {
			content = content[:len(content)-1]
		}
		content += "]"
	}
	// 生成json文件
	fullName := fmt.Sprintf("%s/%s.json", jsonPath, mainSheet)
	jsonFile, _ := os.OpenFile(fullName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	defer jsonFile.Close()
	jsonFile.WriteString(content)
	return nil
}

// getJsonByRow 获取单行Json
//
//	@param local sheet单元格定位
//	@param file excel对象
//	@param sheetMap sheet集
//	@return string json内容
func getJsonByRow(local *excel.Local, file *excelize.File, sheetName string, useId bool) string {
	// 初始化类型、键单元格定位
	typeLocal, keyLocal := &excel.Local{}, &excel.Local{}
	typeLocal.GetLocal(0, 3)
	keyLocal.GetLocal(0, 4)
	// 主sheet
	content := "{"
	id := ""
	// 遍历列
	for {
		typeStr := file.GetCellValue(sheetName, typeLocal.GetLocal(0, 0))
		keyStr := file.GetCellValue(sheetName, keyLocal.GetLocal(0, 0))
		valueStr := file.GetCellValue(sheetName, local.GetLocal(0, 0))
		value := ""
		// 如果列为空，结束遍历
		if typeStr == "" || keyStr == "" {
			break
		}
		if typeStr[:2] == "##" { // 数组对象额外处理
			value = getValueByArrayObject(id, typeStr, file)
		} else {
			value = getValueByType(typeStr, valueStr)
		}
		// id为主键
		if keyStr == "id" {
			id = value
		}
		if keyStr != "id" || useId {
			content += fmt.Sprintf("\"%s\":%s,", keyStr, value)
		}
		// 列偏移
		typeLocal.GetLocal(1, 0)
		keyLocal.GetLocal(1, 0)
		local.GetLocal(1, 0)
	}
	// 删除最后一个逗号
	if len(content) > 1 {
		content = content[:len(content)-1]
	}
	content += "}"
	return content
}

// getValueByType 根据类型转换值（数组对象除外）
//
//	@param typeStr 类型
//	@param valueStr 值
//	@return string 转换后的值
func getValueByType(typeStr, valueStr string) string {
	value := ""
	switch typeStr {
	case "bool": // 布尔型
		if valueStr == "0" {
			value = "false"
		} else {
			value = "true"
		}
	case "int32", "float64": // 整型、浮点型
		value = valueStr
	case "string": // 字符串
		value = fmt.Sprintf("\"%s\"", valueStr)
	case "[]bool": // 数组布尔型
		value = "["
		arr := strings.Split(valueStr, ",")
		for _, v := range arr {
			if v == "0" {
				value += "false,"
			} else {
				value += "true,"
			}
		}
		if len(arr) > 0 {
			value = value[:len(value)-1]
		}
		value += "]"
	case "[]int32", "[]float64": // 数组整型、数组浮点型
		value = "["
		arr := strings.Split(valueStr, ",")
		for _, v := range arr {
			value += fmt.Sprintf("%s,", v)
		}
		if len(arr) > 0 {
			value = value[:len(value)-1]
		}
		value += "]"
	case "[]string": // 数组字符串
		value = "["
		arr := strings.Split(valueStr, ",")
		for _, v := range arr {
			value += fmt.Sprintf("\"%s\",", v)
		}
		if len(arr) > 0 {
			value = value[:len(value)-1]
		}
		value += "]"
	}
	return value
}

// getValueByArrayObject 获取数组对象
//
//	@param id 主Sheet行Id
//	@param typeStr 数据类型（副Sheet名）
//	@param file excel对象
func getValueByArrayObject(id, typeStr string, file *excelize.File) string {
	value := "["
	local := &excel.Local{}
	local.GetLocal(0, 5)
	for {
		parentId := file.GetCellValue(typeStr, local.GetLocal(0, 0))
		if parentId == id {
			value += fmt.Sprintf("%s,", getJsonByRow(local, file, typeStr, false))
		} else if parentId == "" {
			break
		}
		local.GetLocal(-1, 1)
	}
	// 删除最后一个逗号
	if len(value) > 1 {
		value = value[:len(value)-1]
	}
	value += "]"
	return value
}
