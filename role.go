package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CasbinRule struct {
	gorm.Model
	ID    uint `gorm:"primaryKey;autoIncrement"`
	Ptype string
	V0    string
	V1    string
	V2    string
	V3    string
	V4    string
	V5    string
}

func InitCasbin() {
	//
	// e, err := casbin.NewEnforcer("./rules/model.conf", "./rules/policy.csv")
	//mysql driver
	// adapter, errs := gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/casbin1.0", true) // Your driver and data source.
	// if errs != nil {
	// 	fmt.Printf("%s", errs)
	// }
	//數據庫連接
	dsn := "root:123456@tcp(127.0.0.1:3306)/casbin1.0?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		fmt.Printf("%s", err)
	}
	//
	a, _ := gormadapter.NewAdapterByDBWithCustomTable(db, &CasbinRule{})
	//
	enforcer, err := casbin.NewEnforcer("internal/validate/rules/model.conf", a)
	if err != nil {
		fmt.Printf("%s", err)
	}
	//Newvalidate
	Newvalidate(enforcer)
}
func Newvalidate(e *casbin.Enforcer) {
	e.AddFunction("my_func", KeyMatchFunc)
	sub := "user6" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "write" // the operation that the user performs on the resource.
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		fmt.Printf("%s", err)
	}

	if ok {
		fmt.Println("ok")
		// permit alice to read data1
	} else {
		fmt.Println("not")
		// deny the request, show an error
	}
	rules := [][]string{
		{"user6", "data1", "write"},
	}
	ok1, err1 := e.AddGroupingPoliciesEx(rules)
	if err1 != nil {
		fmt.Printf("%s", err1)
	} else {
		fmt.Println("ok1", ok1)
	}
}
func KeyMatch(key1 string, key2 string) bool {
	return key1 == key2
}
func KeyMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return (bool)(KeyMatch(name1, name2)), nil
}
