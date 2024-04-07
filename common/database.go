package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// DB 是全局的数据库连接变量
var DB *gorm.DB

func Init(datasourceName string) {
	// 创建数据库连接
	var err error
	DB, err = gorm.Open(mysql.Open(datasourceName), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	fmt.Println("Database connection established successfully")
}
