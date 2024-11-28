package DB

import (
	"os"

	"wails_hashcat/global"
	"wails_hashcat/model"
)

func RegisterTables() {
	db := global.DB
	err := db.AutoMigrate(
		model.HashCatTask{},
		model.SourceHashFile{},
	)
	if err != nil {
		global.SugarLogger.Error("register table failed")
		os.Exit(0)
	}
	global.SugarLogger.Info("register table success")
}
