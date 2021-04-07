package core

import (
	"fmt"
	"gin-learn/global"
	"gin-learn/initialize"
	"gin-learn/model"
	"time"
)

func RunApplicationServer() {
	global.G_DB = initialize.GormMysql()
	if global.G_DB != nil {
		global.G_DB.AutoMigrate(
			model.SysUser{},
			model.SysAuthority{},
			model.SysBaseMenu{},
			model.SysBaseMenuParameter{},
			// model.SysMenu{},
			model.LuckyK{},
			model.SysOperationRecord{},
			model.CasbinModel{},
		)
		db, _ := global.G_DB.DB()
		defer db.Close()
	}

	Router := initialize.Routes()

	Router.Run(":8999")
	time.Sleep(10 * time.Microsecond)
	fmt.Printf(`
		后端服务启动成功
		server on port: %d
		`, 8999)
}
