package controllers

/**
公共基础
 */

//引入模块
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
	"goAPI/conf"
	"time"
	"strings"
	"crypto/md5"
	"io"
	"fmt"
	"math/rand"
)

/**
错误返回
 */
func ReturnError(msg string, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 404,
		"msg":  msg,
		"data": 0,
	})
}

/**
自定义错误码返回
 */
func ReturnErrorOther(code int, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": 0,
	})
}

/**
成功返回
 */
func ReturnSuccess(msg string, c *gin.Context, dataList interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
		"data": dataList,
	})
}

/**
获取用户参数
paramsType 1：字符串 2：数字
 */
func GetRequestParameters(c *gin.Context, key string, paramsType int) (int, string) {
	//获取Get参数
	value := c.Query(key)
	//获取POST参数
	if value == "" {
		value = c.PostForm(key)
	}
	//判断是否是数字
	if paramsType == 2 {
		valueInt, valueIntErr := strconv.Atoi(value)
		if valueIntErr != nil {
			valueInt = 0
		}
		//返回
		return valueInt, ""
	}
	return 0, value
}

/**
检验秘钥是否正确
 */
func IdentificationData(token string, timeStamp int64, path string) int {
	//获取项目令牌
	projectConfig := conf.Config{}
	configData := projectConfig.ConfigGetValue()
	ProjectPrivateKey := configData.ProjectPrivateKey
	//获取当前时间戳
	timeUnix := time.Now().Unix()
	//设置时间限制
	timeDisparity := timeUnix - timeStamp
	timeLimit := int64(60 * 5) //int转int64
	//判断是否超期
	if timeDisparity > timeLimit {
		return -1
	}
	//判断时间是否提前了
	if timeDisparity < 0 {
		return -1
	}
	//处理路径【去除/】
	pathSign := strings.Replace(path, "/", "", -1)
	//判断秘钥是否正确【int64转字符串】
	localToken := ProjectPrivateKey + pathSign + strconv.FormatInt(timeStamp, 10)
	if localToken != token {
		return -1
	}
	//返回
	return 1
}

/**
获取当前年月日时分秒
 */
func GetDateTime() string {
	timeObj := time.Now()
	var createTime = timeObj.Format("2006-01-02 15:04:05")
	return createTime
}

/**
获取当前年月日
 */
func GetDate() string {
	timeObj := time.Now()
	var createTime = timeObj.Format("2006-01-02")
	return createTime
}

/**
密码加密算法
 */
func Encryption(password, username string) string {
	//拼接加密体【账户+密码】
	conversion := username + password
	//计算md5
	encryption := md5.New()
	io.WriteString(encryption, conversion)                //将str写入到w中
	md5password := fmt.Sprintf("%x", encryption.Sum(nil)) //w.Sum(nil)将w的hash转成[]byte格式
	//返回
	return md5password
}

/**
刷新令牌时间
 */
func RefreshTokenTime() string {
	//获取当前时间戳
	var currentTime = time.Now().Unix()
	//令牌过期时间2个小时
	currentTime = currentTime + (60 * 60 * 2)
	//时间戳转日期
	tokenTime := time.Unix(currentTime, 0).Format("2006-01-02 15:04:05")
	//返回
	return tokenTime
}

/**
获取用户令牌
 */
func RefreshToken(username string) string {
	//获取当前时间戳
	var currentTime = time.Unix(time.Now().Unix(), 0).Format("20060102150405")
	//获取随机字符串
	randString := RandomString(8)
	//拼接加密体【账户+密码】
	conversion := username + randString + string(currentTime)
	//计算md5
	encryption := md5.New()
	io.WriteString(encryption, conversion)                //将str写入到w中
	md5password := fmt.Sprintf("%x", encryption.Sum(nil)) //w.Sum(nil)将w的hash转成[]byte格式
	//返回
	return md5password
}

/**
获取随机字符串
 */
func RandomString(n int) string {
	//设置随机字符串列表
	var allowedChars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	//初始化容器
	var letters []rune
	letters = allowedChars
	//初始化数组
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	//返回
	return string(b)
}

/**
处理图片地址
 */
func DealImageUrl(image string) string {
	//初始化配置
	projectConfig := conf.Config{}
	configData := projectConfig.ConfigGetValue()
	//判断地址
	if image == "" {
		return image
	} else {
		if find := strings.Contains(image, "http"); find {
			return image
		} else {
			return configData.ImageHost + image
		}

	}
}

/**
日期转时间戳
 */
func DealTimeToTimeStamp(datetime string) int64 {
	//转化所需模板
	timeLayout := "2006-01-02 15:04:05"
	//获取时区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	//日期转化为时间戳
	tmp, _ := time.ParseInLocation(timeLayout, datetime, loc)
	//转化为时间戳 类型是int64
	timestamp := tmp.Unix()
	//返回
	return timestamp
}

/**
过滤掉图片请求域名
 */
func FilterImageUrl(imageParams string) string {
	//获取系统配置
	projectConfig := conf.Config{}
	configData := projectConfig.ConfigGetValue()
	//过滤头部域名
	imageData := strings.Replace(imageParams, configData.ImageHost, "", -1)
	//返回
	return imageData
}

/**
处理日期格式
 */
func GetDateValue(change string) string {
	to, _ := time.Parse("2006-01-02T15:04:05Z", change)
	stamp := to.Format("2006-01-02")
	return stamp
}

/**
处理日期时间格式
 */
func GetDateTimeValue(change string) string {
	to, _ := time.Parse("2006-01-02T15:04:05Z", change)
	stamp := to.Format("2006-01-02 15:04:05")
	return stamp
}
