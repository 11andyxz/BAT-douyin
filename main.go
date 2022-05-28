package main

import (
	"BAT-douyin/dal"
	"github.com/gin-gonic/gin"
)

func Init() {
	dal.MysqlInit()
}

func main() {

	r := gin.Default()
	Router(r)
	Init()
	r.Run()
}
