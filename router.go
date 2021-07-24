package main

/**
路由模块
 */

//引入模块
import (
	"github.com/gin-gonic/gin"
	"strings"
	"net/http"
	. "goAPI/controllers"
	"strconv"
)

//路由
func initRouter() *gin.Engine {
	//初始化路由
	router := gin.Default()
	//跨域中间件
	router.Use(CORSMiddleware())
	//路由中间件
	router.Use(ControlMiddleware)
	//定义入口
	router.Any("/", HomeIndex) //http://localhost:9999
	//情话版-demo
	loveTalkGetData := router.Group("admin/loveTalk")
	{
		loveTalkGetData.Any("/getData", LoveTalkGetData)           //http://localhost:9999/admin/loveTalk/getData
		loveTalkGetData.Any("/getDataList", LoveTalkGetDataList)   //http://localhost:9999/admin/loveTalk/getDataList
		loveTalkGetData.Any("/insert", LoveTalkInsert)             //http://localhost:9999/admin/loveTalk/insert
		loveTalkGetData.Any("/update", LoveTalkUpdate)             //http://localhost:9999/admin/loveTalk/update
		loveTalkGetData.Any("/delete", LoveTalkDelete)             //http://localhost:9999/admin/loveTalk/delete
		loveTalkGetData.Any("/updateStatus", LoveTalkUpdateStatus) //http://localhost:9999/admin/loveTalk/updateStatus
	}
	//图片上传工具类型
	UploadFile := router.Group("admin/UploadFile")
	{
		UploadFile.Any("/uploadImage", UploadImage) //http://localhost:9999/admin/UploadFile/uploadImage
	}
	//返回路由
	return router
}

//设置总路由中间件【权鉴校验】
func ControlMiddleware(context *gin.Context) {
	//获取当前访问路由
	routerPath := context.FullPath()
	//判断是否命中管理端路由【触发令牌检验】【采用中间件】
	if find := strings.Contains(routerPath, "admin/"); find {
		AdminControlMiddleware(context)
	}
	//判断是否命中API端路由【触发令牌检验】【采用中间件】
	if find := strings.Contains(routerPath, "api/"); find {
		APIControlMiddleware(context)
	}
}

//小程序接口-设置路由中间件【权鉴校验】
func APIControlMiddleware(context *gin.Context) {
}

//后台接口-设置路由中间件【权鉴校验】
func AdminControlMiddleware(context *gin.Context) {
	//白名单状态 1：否 2：是
	whileRouteStatus := 1
	//获取当前路由
	path := context.FullPath()
	//设置路由白名单
	var whiteListRouter = []string{"/", "/admin/UploadFile/uploadImage"}
	//循环判断是否命中白名单
	for i := 0; i < len(whiteListRouter); i++ {
		if path == whiteListRouter[i] {
			whileRouteStatus = 2
		}
	}
	//判断是否需要校验
	if whileRouteStatus == 1 {
		//获取头部参数
		token := context.GetHeader("token")
		tokenTime := context.GetHeader("tokenTime")
		if token == "" || tokenTime == "" {
			//令牌缺失返回
			context.Abort()
			ReturnErrorOther(502, "令牌缺失", context)
		} else {
			//转化数据
			tokenTimeInt, _ := strconv.ParseInt(tokenTime, 10, 64)
			//检验密钥是否正确
			tokenResult := IdentificationData(token, tokenTimeInt, path)
			if tokenResult != 1 {
				context.Abort()
				ReturnErrorOther(502, "令牌有误", context)
			} else {
				//控制权交换给路由
				context.Next()
			}
		}
	}
}

//路由中间件-防跨域
func CORSMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,TokenId,TokenPower")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
