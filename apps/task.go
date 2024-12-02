package apps

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"wails_hashcat/global"
	"wails_hashcat/model"
	"wails_hashcat/model/request"
	service "wails_hashcat/server"
	"wails_hashcat/utils/hashcat"
)

func (a *App) CreatTask(req *request.CreateTask) error {
	var sf model.SourceHashFile
	db := global.DB
	workHome, _ := os.Getwd()
	cmd := " "
	db.Model(&model.SourceHashFile{}).Where("id = ?", req.FileID).First(&sf)
	switch sf.FileType {
	case "original":
		cmd += "-m 1000"
	case "vc":
		cmd += "-m 29421"
	case "hc22000":
		cmd += "-m 22000"
	}
	req.OutFilePath = filepath.Join(workHome, "out/"+sf.FileName+".txt")
	ht := model.HashCatTask{
		FileID:       req.FileID,
		OutFilePath:  req.OutFilePath,
		AttackMode:   req.AttackMode,
		IncrementMin: req.IncrementMin,
		IncrementMax: req.IncrementMax,
		MaskType:     req.MaskType,
		Status:       model.TaskStatusInit,
	}

	cmd += fmt.Sprintf(" %s -a %d ",
		sf.FilePath, req.AttackMode)
	switch req.AttackMode {
	case 0:
		{
			DictFilePath := filepath.Join(workHome, "static/dict/"+"cain.txt")
			//cmd += fmt.Sprintf(" %s ", req.DictFilePath)
			cmd += fmt.Sprintf(" %s ", DictFilePath)
		}
	case 3:
		{
			mask := service.GenerateMask(req.MaskType, req.IncrementMax)
			cmd += mask
			cmd += fmt.Sprintf(" -i --increment-min %d --increment-max %d ",
				req.IncrementMin, req.IncrementMax)
		}
	default:

		return errors.New("参数错误")
	}
	cmd += fmt.Sprintf(" -o %s ", req.OutFilePath)
	ht.CMD = cmd
	err := service.CreateOrUpdateHashCatTask(ht)
	//go hashcat.HT.StartSession(cmd)
	return err
}

func (a *App) DeleteTask(id int) error {
	db := global.DB
	var ht model.HashCatTask
	if err := db.First(&ht, id).Error; err != nil {
		return err
	}
	runningStatus := []model.TaskStatus{model.TaskStatusRunning}
	if contains(runningStatus, ht.Status) {
		return errors.New("此状态下不可删除")
	}
	if err := db.Delete(&ht).Error; err != nil {
		return err
	}
	return nil
}

func contains(slice []model.TaskStatus, value model.TaskStatus) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func (a *App) StartTask(id int) error {
	db := global.DB
	var task model.HashCatTask
	if err := db.Preload("File").
		Where("id = ?", id).
		First(&task).Error; err != nil {
		return err
	}
	go hashcat.InitTask(int(task.ID))
	go hashcat.HT.StartSession(task.CMD)
	db.Model(&task).Update("status", model.TaskStatusRunning)
	return nil
}
func (a *App) StopTask(id int) error {
	if hashcat.HT != nil {
		go hashcat.HT.StopTask()
	}
	db := global.DB
	var task model.HashCatTask
	if err := db.Where("id = ?", id).
		First(&task).Error; err != nil {
		return err
	}
	db.Model(&task).Update("status", model.TaskStatusStop)
	return nil
}

func (a *App) RestartTask(id int) error {
	hashcat.InitTask(id)
	go hashcat.HT.ReStartSession()
	db := global.DB
	var task model.HashCatTask
	if err := db.Where("id = ?", id).
		First(&task).Error; err != nil {
		return err
	}
	db.Model(&task).Update("status", model.TaskStatusRunning)
	return nil
}
