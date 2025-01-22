package instance

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbInstance, ErrorInstance = GormDbPostgresInit()

func GormDbPostgresInit() (*gorm.DB, error) {
	dns := "host=localhost user=postgres password=123456 dbname=Gorm1.0 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dns,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err == nil {
		fmt.Println("数据连接成功了", err)
	}
	return db, err
}
