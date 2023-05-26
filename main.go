package main

import (
	"github.com/hospital-management/controller"
	"github.com/hospital-management/utils"
)

func main() {
	utils.InitializeDatabase()
	controller.StartServer()
}
