package conf

/**
配置文件
 */

/**
初始化项目配置结构体
 */
type Config struct {
	MysqlAddr         string
	MysqlPort         string
	MysqlRoot         string
	MysqlPassword     string
	MysqlDataBase     string
	ProjectPort       string
	ProjectName       string
	ProjectPrivateKey string
	CosSecretId       string
	CosSecretKey      string
	CosHost           string
	UploadFile        string
	ImageHost         string
}

/**
获取配置信息
 */
func (config *Config) ConfigGetValue() Config {
	//初始化结构体
	conf := Config{
	}
	//赋值
	conf.MysqlAddr = "127.0.0.1"        //数据库地址
	conf.MysqlPort = "3306"             //数据库端口
	conf.MysqlRoot = "root"             //数据库账户
	conf.MysqlPassword = ""             //数据库密码
	conf.MysqlDataBase = "projectApi"   //数据库名
	conf.ProjectPort = "9998"           //项目启动端口
	conf.ProjectName = "Go-API项目"     //项目名称
	conf.ProjectPrivateKey = "goAPI"    //网站验证私钥
	conf.CosSecretId = ""               //腾讯云Cos-SecretID
	conf.CosSecretKey = ""              //腾讯云Cos-SecretKey
	conf.CosHost = ""                   //腾讯云Cos-地址
	conf.UploadFile = "static/uploads/" //图片保存本地目录
	conf.ImageHost = ""                 //图片域名
	//读取配置
	return conf
}
