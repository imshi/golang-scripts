// 对excel进行读写
package main

import (
	"fmt"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
)

type User struct {
	ID      int
	Name    string
	Address string
}

// 主函数，用于将 <结构体类型切片数据> 组织成 <map 数组> 格式通过 excel 写入文件（逐行写入）
func main() {
	// 初始化写入用的数据
	users := []User{
		{
			ID:      1,
			Name:    "孔子",
			Address: "山东曲阜",
		},
		{
			ID:      2,
			Name:    "牛顿",
			Address: "英国伦敦",
		},
		{
			ID:      3,
			Name:    "凯撒",
			Address: "罗马",
		},
	}

	// 创建用于写入的 excel文件
	f := excelize.NewFile()
	sheetName := "人物信息"
	// 删除 Sheet1，创建新Sheet
	f.DeleteSheet("Sheet1")
	sheet := f.NewSheet(sheetName)
	f.SetActiveSheet(sheet)

	// 定义表头信息
	header := map[string]interface{}{"A1": "编号", "B1": "姓名", "C1": "地址"}
	// 组织数据并写入文件
	data := GenerateData(users, header)
	err := Write2File("历史人物", sheetName, data, true)
	if err != nil {
		log.Errorln(err)
	}

	// 读取 Excel 文件
	ReadExcel("历史人物_2021-11-24_.xlsx")
}

// 组织数据，将数据组织成 map 数组格式，每个数组元素为 excel 表格中的一行，map 中的每个键表示 excel 中的一个位置，比如 A1，B2，C3 等，其值为需要填充到该单元格中的值
func GenerateData(users []User, header map[string]interface{}) []map[string]interface{} {
	var maps []map[string]interface{}
	var m map[string]interface{}

	// 将表头数据添加到组织数据的 map 数组中，且为第一行，真实数据从第二行开始
	maps = append(maps, header)
	// 从第二行起，按行组织数据
	rowCount := 2
	for _, user := range users {
		index := strconv.Itoa(rowCount)
		m = map[string]interface{}{
			"A" + index: user.ID,
			"B" + index: user.Name,
			"C" + index: user.Address,
		}

		maps = append(maps, m)
		rowCount = rowCount + 1
	}
	// 返回按行组织好的、map切片格式的数据
	return maps
}

// 将组织好的数据写入 excel 中，每个 map[string]interface{} 为一行数据，键为：A2 B2 C2；A3 B3 C3等等
func Write2File(fileName string, sheetName string, data []map[string]interface{}, dateContaining bool) error {
	f := excelize.NewFile()

	sheet := f.NewSheet(sheetName)
	f.SetActiveSheet(sheet)

	for _, item := range data {
		for k, v := range item {
			err := f.SetCellValue(sheetName, k, v)
			if err != nil {
				log.Errorln(err)
				continue
			}
		}
	}

	if dateContaining {
		return f.SaveAs(fileName + "_" + time.Now().Format("2006-01-02") + "_" + ".xlsx")
	} else {
		return f.SaveAs(fileName + ".xlsx")
	}
}

// 读取 excel 文件
func ReadExcel(filePath string) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	for sheetIndex, sheetName := range f.GetSheetMap() {
		fmt.Println(sheetIndex, "--->", sheetName)
		rows, err := f.Rows(sheetName)
		if err != nil {
			log.Errorln(err)
			continue
		}
		for rows.Next() {
			// 获取一行
			cols, err := rows.Columns()
			if err != nil {
				log.Errorln(err)
				continue
			}
			// 获取行数组中的值
			for _, colCell := range cols {
				fmt.Println(colCell)
			}
		}
	}
}
