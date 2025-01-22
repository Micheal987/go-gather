package delete

import (
	"fmt"
	"gintest/internal/db/collections"

	"gorm.io/gorm"
)

// DeleteDbUniqueById 根據主键的id刪除
func DeleteDbUniqueById[T collections.CollectionSqlImpl](db *gorm.DB, info *T, id int) {
	db.Delete(&info, id)
	fmt.Println("删除数据", info)
}

// DeleteDbNotUniqueById 根據非主键的id刪除
func DeleteDbNotUniqueById[T collections.CollectionSqlImpl](db *gorm.DB, info *T, id int) {
	db.Where("id=?", id)
	fmt.Println("删除数据", info)
}
func DeleteDbBatch[T collections.CollectionSqlImpl](db *gorm.DB, info *T, count ...[]int) {
	db.Delete(&info, count)
	fmt.Println("删除数据", info)
}
