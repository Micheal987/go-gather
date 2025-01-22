package update

import (
	"gintest/internal/db/collections"

	"gorm.io/gorm"
)

// 更新
// 更新多列
func UpdateSingle[T collections.CollectionSqlImpl](db *gorm.DB, info *T, id int) {
	db.Model(&info).Where("id=?", id).Updates(&info)
}
