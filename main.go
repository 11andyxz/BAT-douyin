package main

import (
	"BAT-douyin/dal"
	"github.com/gin-gonic/gin"
)

func Init() {
	dal.MysqlInit()
}

func main() {
	Init()
	r := gin.Default()
	Router(r)

	r.Run()
}
