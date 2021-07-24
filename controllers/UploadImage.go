package controllers

import (
	"github.com/gin-gonic/gin"
	"time"
	"strings"
	"path"
	"fmt"
	"net/url"
	"net/http"
	"context"
	"goAPI/conf"
	"os"
	"github.com/tencentyun/cos-go-sdk-v5"
)

/**
控制器-图片上传-公共基础
 */

/**
上传图片至腾讯云COS
 */
func UploadImage(c *gin.Context) {
	//获取上传控件内容
	file, fileErr := c.FormFile("file")
	//获取上传指定位置
	_, fileAction := GetRequestParameters(c, "fileAction", 1) //上传指定位置
	if fileAction == "" {
		ReturnError("请说明上传指定位置", c)
		return
	}
	//是否存在文件
	if fileErr == nil {
		//初始化配置
		projectConfig := conf.Config{}
		configData := projectConfig.ConfigGetValue()
		//读取腾讯云Cos配置
		uploadFile := configData.UploadFile
		cosSecretId := configData.CosSecretId
		cosSecretKey := configData.CosSecretKey
		cosHost := configData.CosHost
		ImageHost := configData.ImageHost
		//获取上传文件类型
		fileExt := strings.ToLower(path.Ext(file.Filename))
		if fileExt != ".jpg" && fileExt != ".png" && fileExt != ".gif" && fileExt != ".jpeg" {
			ReturnError("上传格式有误，请上传图片", c)
			return
		}
		//获取当前时间戳
		var currentTime = time.Unix(time.Now().Unix(), 0).Format("20060102150405")
		//获取随机字符串
		randString := RandomString(8)
		//生成文件名
		fileName := fileAction + randString + string(currentTime)
		//保存目录
		fileDir := fmt.Sprintf(uploadFile)
		//保存图片至本地
		filePath := fmt.Sprintf("%s%s%s", fileDir, fileName, fileExt)
		saveResult := c.SaveUploadedFile(file, filePath)
		//判断是否失败
		if saveResult != nil {
			ReturnError("上传失败", c)
			return
		}
		//保存本地地址
		localFilePath := uploadFile + fileName + fileExt
		//设置保存cos地址名称【包含目录】
		cosName := "go/" + fileAction + "/" + fileName + fileExt
		//上传至腾讯云cos
		cosStatus := uploadToCos(localFilePath, cosSecretId, cosSecretKey, cosHost, cosName)
		if cosStatus != 200 {
			ReturnError("上传腾讯云失败", c)
			return
		}
		//删除本地图片
		os.Remove(localFilePath)
		//成功返回
		dataList := make(map[string]interface{})
		dataList["fileName"] = cosName
		dataList["fileUrl"] = ImageHost + cosName
		ReturnSuccess("上传成功", c, dataList)
	} else {
		ReturnError("请上传文件", c)
		return
	}
}

/**
上传本地图片至腾讯云Cos
 */
func uploadToCos(fileString, cosSecretId, cosSecretKey, cosHost, cosName string) int {
	//初始化腾讯云Cos配置
	u, _ := url.Parse(cosHost)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  cosSecretId,
			SecretKey: cosSecretKey,
		},
	})
	//上传图片至腾讯云Cos
	_, err := c.Object.PutFromFile(context.Background(), cosName, fileString, nil)
	if err != nil {
		panic(err)
	}
	//成功返回
	return 200
}
