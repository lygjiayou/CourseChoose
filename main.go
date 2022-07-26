package main

import (
	"CourseChoose/model"
	"CourseChoose/server"
)

func main() {
	//g := gin.Default()
	//g.Run(":8080")
	r := server.NewRoute()
	r.Run(":8080")
}

func init() {
	model.InitMysql()
	model.CreateIndexList()
}