package initialize

import (
	"wails_hashcat/global"
	DB "wails_hashcat/utils/databases/orm"
)

func initOrm() {
	global.DB = DB.InitDB()
	DB.RegisterTables()
}
