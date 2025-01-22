package main

import (
	dbAuto "gintest/internal/db/collections/auto"
	"gintest/internal/db/instance"
	userModel "gintest/internal/db/model/user"
	"gintest/internal/router"
)

func main() {
	InitCasbin()
	dbAuto.CreateTableDB[userModel.User](instance.DbInstance, &userModel.User{})
	router.MainRouter()
}
