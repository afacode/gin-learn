package initialize

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GormMysql() *gorm.DB {
	dns := "root:zxcv1234@tcp(127.0.0.1:3306)/gva?charset=utf8mb4&parseTime=True&loc=Local"
	if db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 数据表单数 user => user
			// SingularTable: true,
			// 表前缀
			TablePrefix: "",
		},
	}); err != nil {
		fmt.Println("数据库连接出错")
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(10)
		fmt.Println("数据库初始化成功")
		return db
	}
}
