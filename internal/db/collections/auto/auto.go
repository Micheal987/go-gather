package auto

import (
	"fmt"
	"gintest/internal/db/collections"
	general "gintest/internal/error"

	"gorm.io/gorm"
)

// 自動遷移
// 自動創建數據表
// 自动迁移 (这是GORM自动创建表的一种方式--译者注)
func CreateTableDB[T collections.CollectionSqlImpl](db *gorm.DB, info *T) {
	// 自动迁移 (这是GORM自动创建表的一种方式--译者注)
	err := db.AutoMigrate(&info)
	errM := general.NewWrapError("errMSql", err)
	fmt.Println(errM)
	errAuto := db.Set("gorm:table_options", "ENGINE=Distributed(cluster, default, hits)").AutoMigrate(&info)
	errAutoM := general.NewWrapError("errMSql", errAuto)
	fmt.Println(errAutoM)
}
