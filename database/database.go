package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	// const MYSQL = "root:@tcp(127.0.0.1:3306)/learn-go?charset=utf8mb4&parseTime=True&loc=Local"
	const POSTGRESQL = "host=ep-shiny-paper-a1zyqayd-pooler.ap-southeast-1.aws.neon.tech user=default password=5LkQIZpBY7Tf dbname=verceldb sslmode=require TimeZone=Asia/Shanghai"

	dsn := POSTGRESQL
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("cannot connect to database")
	}
	fmt.Println("Connected to database")
}
