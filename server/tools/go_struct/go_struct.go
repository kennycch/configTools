package go_struct

import (
	"config_tools/tools/excel"
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/kennycch/gotools/general"
)

// StructureGo  构建Go结构体
//
//	@param excelFullPath  excel文件路径
//	@param goPath Go目录路径
//	@return error
func StructureGo(excelFullPath, goPath string) error {
	// 读取Excel
	file, err := excelize.OpenFile(excelFullPath)
	if err != nil {
		return err
	}
	// 获取sheet
	sheetMap := file.GetSheetMap()
	mainSheet := sheetMap[1]
	// 开始解析excel
	goStruct := analysisExcel(file, mainSheet)
	// 根据配置类型生成go文件内容
	tableType := file.GetCellValue(mainSheet, "D1")
	content := ""
	if tableType == "object" {
		content = getContentByOjbect(goStruct)
	} else {
		content = getContentByArray(goStruct)
	}
	content = content[:len(content)-2]
	fullName := fmt.Sprintf("%s/%s.go", goPath, mainSheet)
	goFile, _ := os.OpenFile(fullName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	defer goFile.Close()
	goFile.WriteString(content)
	// 生成管理Go文件
	createManagerStruct(goPath)
	return nil
}

// analysisExcel 解析Excel
//
//	@param file excel文件对象
//	@param mainSheet 主sheet名
//	@return *GoStruct
func analysisExcel(file *excelize.File, mainSheet string) *GoStruct {
	goStruct := &GoStruct{
		mainStruct: &Struct{
			lowName: general.HumpFormat(mainSheet, false),
			upName:  general.HumpFormat(mainSheet, true),
			keys:    make([]*key, 0),
		},
		extraStructs: make([]*Struct, 0),
		fileName:     mainSheet,
	}
	// 解析主sheet
	annotationLocal := &excel.Local{}
	annotationLocal.GetLocal(0, 1)
	typeLocal := &excel.Local{}
	typeLocal.GetLocal(0, 3)
	keyLocal := &excel.Local{}
	keyLocal.GetLocal(0, 4)
	// 遍历列
	for {
		annotationStr := file.GetCellValue(mainSheet, annotationLocal.GetLocal(0, 0))
		typeStr := file.GetCellValue(mainSheet, typeLocal.GetLocal(0, 0))
		keyStr := file.GetCellValue(mainSheet, keyLocal.GetLocal(0, 0))
		// 遍历到空列直接退出
		if len(typeStr) == 0 {
			break
		}
		key := &key{
			annotation: annotationStr,
			lowName:    general.HumpFormat(keyStr, false),
			upName:     general.HumpFormat(keyStr, true),
			json:       general.HumpFormat(keyStr, false),
		}
		// 数组对象额外处理
		if typeStr[:2] == "##" {
			key.typeStr = handleExtra(file, mainSheet, typeStr, goStruct)
		} else {
			key.typeStr = typeStr
		}
		goStruct.mainStruct.keys = append(goStruct.mainStruct.keys, key)
		annotationLocal.GetLocal(1, 0)
		typeLocal.GetLocal(1, 0)
		keyLocal.GetLocal(1, 0)
	}
	return goStruct
}

// handleExtra
//
//	@param file
//	@param mainSheet
//	@param typeStr
//	@return string
func handleExtra(file *excelize.File, mainSheet, extraSheet string, goStruct *GoStruct) string {
	low, up := getExtraStruct(mainSheet, extraSheet)
	extraStruct := &Struct{
		lowName: low,
		upName:  up,
		keys:    make([]*key, 0),
	}
	annotationLocal := &excel.Local{}
	annotationLocal.GetLocal(0, 1)
	typeLocal := &excel.Local{}
	typeLocal.GetLocal(0, 3)
	keyLocal := &excel.Local{}
	keyLocal.GetLocal(0, 4)
	// 遍历列
	for {
		annotationS := file.GetCellValue(extraSheet, annotationLocal.GetLocal(0, 0))
		typeS := file.GetCellValue(extraSheet, typeLocal.GetLocal(0, 0))
		keyS := file.GetCellValue(extraSheet, keyLocal.GetLocal(0, 0))
		annotationLocal.GetLocal(1, 0)
		typeLocal.GetLocal(1, 0)
		keyLocal.GetLocal(1, 0)
		// 遍历到空列直接退出
		if len(typeS) == 0 {
			break
		}
		// 父级Id直接略过
		if keyS == "id" {
			continue
		}
		key := &key{
			annotation: annotationS,
			lowName:    general.HumpFormat(keyS, false),
			upName:     general.HumpFormat(keyS, true),
			json:       general.HumpFormat(keyS, false),
		}
		// 数组对象额外处理
		if typeS[:2] == "##" {
			key.typeStr = handleExtra(file, mainSheet, typeS, goStruct)
		} else {
			key.typeStr = typeS
		}
		extraStruct.keys = append(extraStruct.keys, key)
	}
	goStruct.extraStructs = append(goStruct.extraStructs, extraStruct)
	return "[]" + low
}

// getExtraStruct
//
//	@param mainSheet
//	@param typeStr
//	@return low
//	@return up
func getExtraStruct(mainSheet, typeStr string) (low, up string) {
	str := general.HumpFormat(typeStr[3:], true)
	fullStr := mainSheet + str
	return general.HumpFormat(fullStr, false), general.HumpFormat(fullStr, true)
}

func getContentByOjbect(goStruct *GoStruct) string {
	// 生成结构体、Json结构体和调用方法
	structStr := getStructStr(goStruct.mainStruct, true)
	for _, extraStruct := range goStruct.extraStructs {
		structStr += getStructStr(extraStruct, false)
	}
	// 填充占位符
	content := fmt.Sprintf(objectTemplate,
		goStruct.mainStruct.lowName, goStruct.mainStruct.lowName, goStruct.mainStruct.upName, // 变量
		goStruct.mainStruct.lowName, goStruct.mainStruct.upName, // 结构体
		goStruct.mainStruct.upName,                             // 注册cl
		goStruct.mainStruct.upName, goStruct.mainStruct.upName, // 结构体名称
		goStruct.mainStruct.upName, goStruct.fileName, // 文件名称
		goStruct.mainStruct.upName, goStruct.mainStruct.lowName, goStruct.mainStruct.lowName, goStruct.mainStruct.lowName, // 获取配置
		goStruct.mainStruct.upName,                                                                                                                    // 全部配置迭代器
		goStruct.mainStruct.upName, goStruct.mainStruct.lowName, goStruct.mainStruct.lowName, goStruct.mainStruct.upName, goStruct.mainStruct.lowName, // 解析Json
		structStr, // 结构体部分
	)
	return content
}

// getContentByArray 获取数组配置内容
//
//	@param goStruct 结构体集
//	@return string
func getContentByArray(goStruct *GoStruct) string {
	// 生成结构体、Json结构体和调用方法
	structStr := getStructStr(goStruct.mainStruct, true)
	for _, extraStruct := range goStruct.extraStructs {
		structStr += getStructStr(extraStruct, false)
	}
	// 填充占位符
	content := fmt.Sprintf(arrayTemplate,
		goStruct.mainStruct.lowName, goStruct.mainStruct.lowName, goStruct.mainStruct.upName, // 变量
		goStruct.mainStruct.lowName, goStruct.mainStruct.upName, // 结构体
		goStruct.mainStruct.upName,                             // 注册cl
		goStruct.mainStruct.upName, goStruct.mainStruct.upName, // 结构体名称
		goStruct.mainStruct.upName, goStruct.fileName, // 文件名称
		goStruct.mainStruct.upName, goStruct.mainStruct.lowName, goStruct.mainStruct.lowName, goStruct.mainStruct.lowName, // 获取配置
		goStruct.mainStruct.upName, goStruct.mainStruct.lowName, goStruct.mainStruct.lowName, goStruct.mainStruct.lowName, // 全部配置迭代器
		goStruct.mainStruct.upName, goStruct.mainStruct.lowName, goStruct.mainStruct.lowName, goStruct.mainStruct.lowName,
		goStruct.mainStruct.upName, goStruct.mainStruct.upName, goStruct.mainStruct.lowName, // 解析Json
		structStr, // 结构体部分
	)
	return content
}

func getStructStr(s *Struct, isBase bool) string {
	// 定义结构体名
	structName := s.lowName
	if isBase {
		structName = s.upName
	}
	// 初始化基础结构体、Json结构体和复制方法
	baseStruct := fmt.Sprintf("type %s struct {\n", structName)
	jsonStruct := fmt.Sprintf("type %sJson struct {\n", structName)
	copyFunc := fmt.Sprintf("func (cj %sJson) copy() %s {\n	c := %s{\n", structName, structName, structName)
	// 计算空格数量
	baseLen, jsonLen := 0, 0
	for _, key := range s.keys {
		if len(key.lowName) > baseLen {
			baseLen = len(key.lowName)
		}
		typeLen := len(key.typeStr)
		if len(key.typeStr) > 2 && key.typeStr[:2] == "[]" && !general.InArray(arrayOrder, key.typeStr) {
			typeLen += 4
		}
		if typeLen > jsonLen {
			jsonLen = typeLen
		}
	}
	baseLen += 1
	jsonLen += 1
	arrayKeys := []*key{}
	funcs := ""
	// 遍历键
	for _, key := range s.keys {
		baseStruct += fmt.Sprintf("	%s", key.lowName)
		jsonStruct += fmt.Sprintf("	%s", key.upName)
		// 数组与基本类型分开处理
		isArray := len(key.typeStr) > 2 && key.typeStr[:2] == "[]"
		isArrayObject := isArray && !general.InArray(arrayOrder, key.typeStr)
		if !isArrayObject {
			copyFunc += fmt.Sprintf("		%s:", key.lowName)
		} else {
			arrayKeys = append(arrayKeys, key)
		}
		// 补全空格
		baseSpace := baseLen - len(key.lowName)
		for i := 0; i < baseSpace; i++ {
			baseStruct += " "
			jsonStruct += " "
			if !isArrayObject {
				copyFunc += " "
			}
		}
		baseStruct += fmt.Sprintf("%s\n", key.typeStr)
		if !isArrayObject {
			if !isArray {
				copyFunc += fmt.Sprintf("cj.%s,\n", key.upName)
			} else {
				copyFunc += fmt.Sprintf("arrayCopy(cj.%s),\n", key.upName)
			}
		}
		jsonStruct += key.typeStr
		jsonSpace := jsonLen - len(key.typeStr)
		if isArrayObject {
			jsonStruct += "Json"
			jsonSpace -= 4
		}
		// 补全Json空格
		for i := 0; i < jsonSpace; i++ {
			jsonStruct += " "
		}
		jsonStruct += fmt.Sprintf("`json:\"%s\"`\n", key.json)
		// 键对应方法
		if !isArray {
			funcs += fmt.Sprintf("\nfunc (c %s) %s() %s {\n	return c.%s\n}\n", structName, key.upName, key.typeStr, key.lowName)
		} else {
			funcs += fmt.Sprintf("\nfunc (c %s) %s() %s {\n	return arrayCopy(c.%s)\n}\n", structName, key.upName, key.typeStr, key.lowName)
		}
	}
	// 补全基础结构体、Json结构体和复制方法
	baseStruct += "}\n\n"
	jsonStruct += "}\n\n"
	copyFunc += "	}\n"
	for _, key := range arrayKeys {
		copyFunc += fmt.Sprintf("	%s := make(%s, 0)\n	for _, ex := range cj.%s {\n		%s = append(%s, ex.copy())\n	}\n	c.%s = %s\n",
			key.lowName, key.typeStr, key.upName, key.lowName, key.lowName, key.lowName, key.lowName)
	}
	copyFunc += "	return c\n}\n"
	return fmt.Sprintf("%s%s%s%s\n", baseStruct, jsonStruct, copyFunc, funcs)
}

// createManagerStruct 生成管理Go文件
//
//	@param goPath Go目录路径
func createManagerStruct(goPath string) {
	fullName := fmt.Sprintf("%s/manager.go", goPath)
	goFile, _ := os.OpenFile(fullName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	defer goFile.Close()
	goFile.WriteString(managerTemplate)
}
