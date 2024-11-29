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
	}

	//hashcat.InitTask(int(sf.ID))
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
