package user

type User struct {
	Id         int32 `grom:"primaryKey;autoIncrement;unique"`
	Name       string
	Password   *string
	Email      *string
	CreateTime int64 `gorm:"autoCreateTime:milli"`
	UpdateTime int64 `gorm:"autoUpdateTime:milli"`
}
