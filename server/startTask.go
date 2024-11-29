package service

import (
	"errors"
	"strings"

	"wails_hashcat/global"
	"wails_hashcat/model"
)

func CreateOrUpdateHashCatTask(ht model.HashCatTask) error {
	db := global.DB
	var existingTask model.HashCatTask
	result := db.Where(ht).FirstOrCreate(&existingTask, ht)

	if result.Error != nil {
		global.SugarLogger.Error(result.Error)
		return result.Error
	} else {
		if result.RowsAffected > 0 {
			global.SugarLogger.Info("记录已创建")
		} else {
			return errors.New("任务已存在")
		}
	}
	return nil
}

func GenerateMask(maskType int, incrementMax int) string {
	var mask string
	switch maskType {
	//纯数字
	case 0:
		mask = strings.Repeat("?d", incrementMax)
	//小写字母
	case 1:
		mask = strings.Repeat("?l", incrementMax)
	//大写字母
	case 2:
		mask = strings.Repeat("?u", incrementMax)
	//小写字母+数字
	case 3:
		mask = strings.Repeat("?h", incrementMax)
	//大小写字母+数字
	case 4:
		mask = "-1 ?d?u?l "
		mask += strings.Repeat("?1", incrementMax)
	//大小写字母+数字+可见特殊符号
	case 5:
		mask = "-1 ?d?u?l?s "
		mask += strings.Repeat("?1", incrementMax)
	}
	return mask
}
