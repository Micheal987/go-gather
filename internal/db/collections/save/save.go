package save

import (
	"fmt"
	"gintest/internal/db/collections"

	"gorm.io/gorm"
)

// SaveDbById 更新
func SaveDbById[T collections.CollectionSqlImpl](db *gorm.DB, info []*T, id int) {
	db.Save(&info).Where("id=?", id)
	fmt.Println("插入数据", db.RowsAffected, "条数")
}
