package apps

import (
	"errors"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"wails_hashcat/global"
	"wails_hashcat/model"
	"wails_hashcat/utils"
)

func (a *App) ChooseFile() string {
	var err error
	selectFilePath, err = runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		return ""
	}
	global.SugarLogger.Info(selectFilePath)
	fileName := filepath.Base(selectFilePath)
	return fileName
}

func (a *App) ConfirmAddFile(fileType string) error {
	if selectFilePath == "" {
		return errors.New("未选择文件")
	}
	fileName := filepath.Base(selectFilePath)
	fileMD5, err := utils.CalculateMD5(selectFilePath)
	if err != nil {
		return err
	}
	db := global.DB
	var sourceFile = model.SourceHashFile{
		FileName: fileName,
		FilePath: selectFilePath,
		FileType: fileType,
		FileMD5:  fileMD5,
	}
	if err = db.Save(&sourceFile).Error; err != nil {
		return errors.New("文件已添加")
	}
	selectFilePath = ""
	return nil
}

func (a *App) GetFileList() []model.SourceHashFile {
	var db = global.DB
	var sourceFileList []model.SourceHashFile
	db.Preload("Tasks").Find(&sourceFileList)
	//spew.Dump(sourceFileList)
	return sourceFileList
}

func (a *App) RemoveFile(id int) error {
	var db = global.DB
	if err := db.Delete(&model.SourceHashFile{}, id).Error; err != nil {
		return errors.New("删除失败")
	}
	return nil
}
