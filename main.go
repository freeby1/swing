package main

import (
	"ginvue/common"
	"ginvue/router"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := common.InitDb()
	defer db.Close()

	r := gin.Default()
	r = router.CollectRouter(r)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务

}
