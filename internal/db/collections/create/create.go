package create

import (
	"gintest/internal/db/collections"

	"gorm.io/gorm"
)

// 插入數據多條
func CreateDb[T collections.CollectionSqlImpl](db *gorm.DB, err error, info []*T) []*T {
	db.Create(info)
	return info
}

// 插入數據單條
func CreateSingle[T collections.CollectionSqlImpl](db *gorm.DB, err error, info *T) *T {
	db.Create(info)
	return info
}
