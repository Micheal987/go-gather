package article

type Article struct {
	Id         int64 `gorm:"primaryKey;unique;autoIncrement"`
	Title      string
	CreateTime int64 `gorm:"autoCreateTime:milli"`
	UpdateTime int64 `gorm:"autoUpdateTime:milli"`
}
