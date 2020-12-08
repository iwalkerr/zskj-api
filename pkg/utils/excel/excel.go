package excel

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type Excelize struct {
	File   *excelize.File
	Stream *excelize.StreamWriter
}

// 新建文件
func NewFile() (*Excelize, error) {
	file := excelize.NewFile()
	w, err := file.NewStreamWriter("Sheet1")
	if err != nil {
		return nil, err
	}
	return &Excelize{
		File:   file,
		Stream: w,
	}, nil
}

// 打开已经存在的文件
func OpenFile(path, sheetName string) (*Excelize, error) {
	file, err := excelize.OpenFile(path)
	if err != nil {
		return nil, err
	}
	w, err := file.NewStreamWriter(sheetName)
	if err != nil {
		return nil, err
	}
	return &Excelize{
		File:   file,
		Stream: w,
	}, nil
}

// 按照行写入
func (e *Excelize) SetRow(i int, list []interface{}) {
	cell, _ := excelize.CoordinatesToCellName(1, i)
	_ = e.Stream.SetRow(cell, list)
}

// 保存到文件
func (e *Excelize) Save() {
	_ = e.Stream.Flush()
	_ = e.File.Save()
}

// 文件另存为
func (e *Excelize) SaveAs(filename string) (string, error) {
	_ = e.Stream.Flush()

	curDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	if filename == "" {
		curdate := time.Now().UnixNano()
		filename = strconv.FormatInt(curdate, 10) + ".xlsx"
	}
	filePath := curDir + "/public/upload/" + filename + ".xlsx"

	if err = CreateFilePath(filePath); err != nil {
		log.Printf("%s", err.Error())
		return "", err
	}

	return filename + ".xlsx", e.File.SaveAs(filePath)
}

//  创建路径
func CreateFilePath(filePath string) error {
	// 路径不存在创建路径
	path, _ := filepath.Split(filePath) // 获取路径
	_, err := os.Stat(path)             // 检查路径状态，不存在创建
	if err != nil || os.IsExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
	}
	return err
}
