package excel

import (
	"config_tools/app/dao"
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// GenerateDemo 创建配置表示例
func GenerateDemo(table *dao.Table, fields []*dao.Field, excelPath string) {
	excel := &excel{
		table:  table,
		fields: fields,
		file:   excelize.NewFile(),
	}
	// 重命名sheet
	excel.file.SetSheetName("Sheet1", table.Name)
	// 设置列宽
	excel.file.SetColWidth(table.Name, "A", "AZ", 24)
	// 头单元格样式
	head, _ := excel.file.NewStyle(headStyle)
	// 标题单元格样式
	title, _ := excel.file.NewStyle(titleStyle)
	// 值单元格样式
	cell, _ := excel.file.NewStyle(cellStyle)
	excel.titleStyle = title
	excel.cellStyle = cell
	// 设置单元格样式
	excel.file.SetCellStyle(table.Name, "A1", "AZ1", head)
	excel.file.SetCellStyle(table.Name, "A2", "AZ5", title)
	excel.file.SetCellStyle(table.Name, "A6", "AZ1000", cell)
	// 头数据
	mainLocal := &Local{}
	datas := map[string]string{
		mainLocal.GetLocal(0, 0): "表名称：",
		mainLocal.GetLocal(1, 0): table.Comment,
		mainLocal.GetLocal(1, 0): "配置类型：",
		mainLocal.GetLocal(1, 0): tableTypeMap[table.TableType],
		mainLocal.GetLocal(1, 0): "哈希标记（请勿改动）：",
		mainLocal.GetLocal(1, 0): table.Hash,
		mainLocal.GetLocal(1, 0): "当前版本：",
		mainLocal.GetLocal(1, 0): fmt.Sprint(table.Version),
	}
	// 固定Id字段
	datas[mainLocal.GetLocal(-1, 1)] = "编号"
	datas[mainLocal.GetLocal(0, 1)] = "固定默认字段"
	datas[mainLocal.GetLocal(0, 1)] = typeMap[2]
	datas[mainLocal.GetLocal(0, 1)] = "id"
	datas[mainLocal.GetLocal(0, 1)] = "1"
	mainLocal.GetLocal(1, -1)
	// 额外sheet
	extraSheets := map[string]uint32{}
	// 主表字段数据
	for _, field := range fields {
		// 过滤非主表字段
		if field.ParentId != 0 {
			continue
		}
		datas[mainLocal.GetLocal(0, 1)] = field.Chinese
		datas[mainLocal.GetLocal(0, 1)] = field.Comment
		// 数组对象需要特殊处理
		if field.FieldType == 9 {
			fieldType := fmt.Sprintf("## %s", field.Name)
			datas[mainLocal.GetLocal(0, 1)] = fieldType
			extraSheets[fieldType] = field.Id
		} else {
			datas[mainLocal.GetLocal(0, 1)] = typeMap[field.FieldType]
		}
		datas[mainLocal.GetLocal(0, 1)] = field.Name
		datas[mainLocal.GetLocal(0, 1)] = field.Example
		mainLocal.GetLocal(1, -1)
	}
	// 填入数据
	for clounm, value := range datas {
		excel.file.SetCellValue(table.Name, clounm, value)
	}
	// 额外sheet
	for sheetName, id := range extraSheets {
		excel.handleExtraSheet(sheetName, id)
	}
	// 保存文件
	excel.file.SaveAs(fmt.Sprintf("%s/%s.xlsx", excelPath, table.Comment))
}

// handleExtraSheet 处理额外Sheet
//
//	@param sheetName sheet名
//	@param parentId 上级Id
//	@param fields 字段集
//	@param f excel对象
//	@param style 单元格样式
func (e *excel) handleExtraSheet(sheetName string, parentId uint32) {
	extra := e.file.NewSheet(sheetName)
	extraLocal := &Local{}
	datas := map[string]string{}
	extraSheets := map[string]uint32{}
	// 设置列宽
	e.file.SetColWidth(sheetName, "A", "AZ", 24)
	// 设置单元格样式
	e.file.SetCellStyle(sheetName, "A2", "AZ5", e.titleStyle)
	e.file.SetCellStyle(sheetName, "A6", "AZ1000", e.cellStyle)
	// 固定Id字段
	datas[extraLocal.GetLocal(0, 1)] = "父级行编号"
	datas[extraLocal.GetLocal(0, 1)] = "固定默认字段"
	datas[extraLocal.GetLocal(0, 1)] = typeMap[2]
	datas[extraLocal.GetLocal(0, 1)] = "id"
	datas[extraLocal.GetLocal(0, 1)] = "1"
	extraLocal.GetLocal(1, -1)
	for _, field := range e.fields {
		if field.ParentId == parentId {
			datas[extraLocal.GetLocal(0, 1)] = field.Chinese
			datas[extraLocal.GetLocal(0, 1)] = field.Comment
			// 数组对象需要特殊处理
			if field.FieldType == 9 {
				fieldType := fmt.Sprintf("## %s", field.Name)
				datas[extraLocal.GetLocal(0, 1)] = fieldType
				extraSheets[fieldType] = field.Id
			} else {
				datas[extraLocal.GetLocal(0, 1)] = typeMap[field.FieldType]
			}
			datas[extraLocal.GetLocal(0, 1)] = field.Name
			datas[extraLocal.GetLocal(0, 1)] = field.Example
			extraLocal.GetLocal(1, -1)
		}
	}
	// 填入数据
	for clounm, value := range datas {
		e.file.SetCellValue(sheetName, clounm, value)
	}
	// 将工作表设为活动工作表
	e.file.SetActiveSheet(extra - 1)
	// 额外sheet
	for sheetName, id := range extraSheets {
		e.handleExtraSheet(sheetName, id)
	}
}

// GetLocal 获取单元格定位
//
//	@param cap 列
//	@param row 行
//	@return string Excel单元格定位
func (l *Local) GetLocal(cap, row int) string {
	if cap < 0 {
		l.cap = 0
	} else {
		l.cap += cap
	}
	if row < 0 {
		l.row = 0
	} else {
		l.row += row
	}
	// 定位前缀
	capArr := strings.Split(caps, "")
	prefix := capArr[l.cap%len(capArr)]
	if l.cap >= len(capArr) {
		prefix = capArr[l.cap/len(capArr)] + prefix
	}
	// 拼接单元格定位
	return fmt.Sprintf("%s%d", prefix, l.row+1)
}
