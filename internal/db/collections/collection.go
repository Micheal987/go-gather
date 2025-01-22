package collections

import (
	articleModel "gintest/internal/db/model/article"
	userModel "gintest/internal/db/model/user"
)

// 自動遷移
// 自動創建數據表
// 自动迁移 (这是GORM自动创建表的一种方式--译者注)
type CollectionSql interface {
	userModel.User | articleModel.Article
}
type CollectionSqlImpl interface {
	CollectionSql
}
