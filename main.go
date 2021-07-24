package main

//引入模块
import (
	"goAPI/db"
	"goAPI/conf"
	"github.com/fvbock/endless"
	"log"
)

//定义主入口
func main() {
	//初始化配置
	projectConfig := conf.Config{}
	configData := projectConfig.ConfigGetValue()
	//获取端口号
	projectPort := configData.ProjectPort
	//定义防错机制，防止出现错误，消耗数据库链接
	defer db.SqlDB.Close()
	//定义路由
	router := initRouter()
	//监听端口
	err := endless.ListenAndServe(":"+projectPort, router)
	if err != nil {
		log.Println("err:", err)
	}
	//router.Run(":" + projectPort)
}
