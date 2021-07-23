package db

//引入模块
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"goAPI/conf"
)

//初始化数据库对象
var SqlDB *gorm.DB

//默认启动
func init() {
	//初始化错误对象
	var err error
	//初始化配置
	projectConfig := conf.Config{}
	configData := projectConfig.ConfigGetValue()
	//读取配置
	mysqlAddr := configData.MysqlAddr
	mysqlPort := configData.MysqlPort
	mysqlRoot := configData.MysqlRoot
	mysqlPassword := configData.MysqlPassword
	mysqlDataBase := configData.MysqlDataBase
	//数据库链接
	SqlDB, err = gorm.Open("mysql", mysqlRoot+":"+mysqlPassword+"@tcp("+mysqlAddr+":"+mysqlPort+")/"+mysqlDataBase+"?parseTime=true")
	//判断链接是否错误
	if err != nil {
		log.Fatal(err.Error())
	}
	//判断链接是否成功
	if SqlDB.Error != nil {
		log.Fatal(SqlDB.Error)
	}
}