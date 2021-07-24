package controllers

/**
引入模块
 */
import (
	"github.com/gin-gonic/gin"
	. "goAPI/models"
)

/**
控制器-情话板
 */

/**
定义主入口
 */
func HomeIndex(c *gin.Context) {
	//获取Ip、头部
	ip := c.ClientIP()
	host := c.GetHeader("Host")
	//组合返回数据
	dataList := make(map[string]interface{})
	dataList["ip"] = ip
	dataList["host"] = host
	//返回
	ReturnSuccess("欢迎来到GoLand世界", c, dataList)
}

/**
获取单条数据
*/
func LoveTalkGetData(c *gin.Context) {
	//获取用户参数
	id, _ := GetRequestParameters(c, "id", 2)
	if id <= 0 {
		ReturnError("请输入Id", c)
		return
	}
	//初始化结构体
	loveTalk := LoveTalk{
		Id: id,
	}
	//根据Id获取参数
	data := loveTalk.LoveTalkWithWhere()
	//处理图片
	showCard := DealImageUrl(data.Card)
	showMottoCard := DealImageUrl(data.MottoCard)
	//判断是否存在数据
	if data.Id <= 0 {
		ReturnError("数据获取失败", c)
		return
	}
	//组合返回数据
	dataList := make(map[string]interface{})
	dataList["data"] = data                   //数据详情
	dataList["showCard"] = showCard           //图片地址
	dataList["showMottoCard"] = showMottoCard //图片地址
	//返回数据
	ReturnSuccess("数据获取成功", c, dataList)
}

/**
获取数据列表
 */
func LoveTalkGetDataList(c *gin.Context) {
	//获取用户参数
	page, _ := GetRequestParameters(c, "page", 2)
	status, _ := GetRequestParameters(c, "status", 2)
	pageNumber, _ := GetRequestParameters(c, "pageNumber", 2)
	//初始化页码
	if page <= 1 {
		page = 1
	}
	//初始化每页展示多少条
	if pageNumber <= 0 {
		pageNumber = 10
	}
	//计算偏移量
	offset := (page - 1) * pageNumber
	//初始化结构体
	loveTalk := LoveTalk{
	}
	//判断条件查询
	if status > 0 {
		loveTalk.Status = status
	}
	//读取列表数据
	resultList := loveTalk.LoveTalkGetDataList(offset, pageNumber)
	//循环获取要的数据
	materialList := make([]map[string]interface{}, len(resultList))
	for k, v := range resultList {
		row := make(map[string]interface{})
		//重新赋值数组
		row["id"] = v.Id
		row["motto"] = v.Motto
		row["mottoCard"] = DealImageUrl(v.MottoCard)
		row["card"] = DealImageUrl(v.Card)
		row["status"] = v.Status
		row["sort"] = v.Sort
		materialList[k] = row
	}
	//获取数据总数
	materialCount := loveTalk.LoveTalkGetDataCount()
	//组合返回数据
	dataList := make(map[string]interface{})
	dataList["materialList"] = materialList   //数据列表
	dataList["materialCount"] = materialCount //数据总数
	dataList["pageNumber"] = pageNumber       //每页展示数据
	dataList["curPage"] = page                //当前页码
	//成功返回
	ReturnSuccess("success", c, dataList)
	return
}

/**
添加数据
 */
func LoveTalkInsert(c *gin.Context) {
	//获取用户参数
	typeId, _ := GetRequestParameters(c, "typeId", 2)       //分类Id
	_, motto := GetRequestParameters(c, "motto", 1)         //文案
	_, mottoCard := GetRequestParameters(c, "mottoCard", 1) //文案封面
	_, card := GetRequestParameters(c, "card", 1)           //卡片封面
	status, _ := GetRequestParameters(c, "status", 2)       //状态 1：显示  2：隐藏
	sort, _ := GetRequestParameters(c, "sort", 2)           //序号
	//判断用户参数
	if typeId <= 0 {
		ReturnError("请选择分类", c)
		return
	}
	if motto == "" {
		ReturnError("文案不能为空", c)
		return
	}
	if mottoCard == "" {
		ReturnError("文案封面不能为空", c)
		return
	}
	if card == "" {
		ReturnError("卡片封面不能为空", c)
		return
	}
	//过滤掉图片域名
	mottoCard = FilterImageUrl(mottoCard)
	card = FilterImageUrl(card)
	//添加数据
	loveTalk := LoveTalk{
		TypeId:    typeId,
		Motto:     motto,
		MottoCard: mottoCard,
		Card:      card,
		Status:    status,
		Sort:      sort,
	}
	insertId := loveTalk.LoveTalkInsert()
	if insertId <= 0 {
		ReturnError("数据添加失败", c)
		return
	}
	//初始化返回数据
	dataList := make(map[string]interface{})
	dataList["insertId"] = insertId
	//成功返回
	ReturnSuccess("数据添加成功", c, dataList)
	return

}

/**
更新数据
 */
func LoveTalkUpdate(c *gin.Context) {
	//获取用户参数
	id, _ := GetRequestParameters(c, "id", 2) //Id
	//获取用户参数
	typeId, _ := GetRequestParameters(c, "typeId", 2)       //分类Id
	_, motto := GetRequestParameters(c, "motto", 1)         //文案
	_, mottoCard := GetRequestParameters(c, "mottoCard", 1) //文案封面
	_, card := GetRequestParameters(c, "card", 1)           //卡片封面
	status, _ := GetRequestParameters(c, "status", 2)       //状态 1：显示  2：隐藏
	sort, _ := GetRequestParameters(c, "sort", 2)           //序号
	//判断用户参数
	if id <= 0 {
		ReturnError("Id获取失败", c)
		return
	}
	if typeId <= 0 {
		ReturnError("请选择分类", c)
		return
	}
	if motto == "" {
		ReturnError("文案不能为空", c)
		return
	}
	if mottoCard == "" {
		ReturnError("文案封面不能为空", c)
		return
	}
	if card == "" {
		ReturnError("卡片封面不能为空", c)
		return
	}
	//过滤掉图片域名
	mottoCard = FilterImageUrl(mottoCard)
	card = FilterImageUrl(card)
	//更新数据
	loveTalkParams := LoveTalk{
		TypeId:    typeId,
		Motto:     motto,
		MottoCard: mottoCard,
		Card:      card,
		Status:    status,
		Sort:      sort,
	}
	loveTalkWhere := LoveTalk{
		Id: id,
	}
	updateRow := loveTalkParams.LoveTalkUpdate(loveTalkWhere)
	if updateRow <= 0 {
		ReturnError("数据编辑失败", c)
		return
	}
	//初始化返回数据
	dataList := make(map[string]interface{})
	dataList["updateId"] = id
	//成功返回
	ReturnSuccess("编辑数据成功", c, dataList)
	return
}

/**
删除数据
 */
func LoveTalkDelete(c *gin.Context) {
	//获取用户参数
	id, _ := GetRequestParameters(c, "id", 2)
	if id <= 0 {
		ReturnError("请输入Id", c)
		return
	}
	//初始化结构对象
	loveTalk := LoveTalk{
		Id: id,
	}
	//删除数据
	result := loveTalk.LoveTalkDelete()
	//判断是否删除成功
	if result <= 0 {
		ReturnError("删除失败", c)
		return
	}
	//组合返回信息
	dataList := make(map[string]interface{})
	dataList["deleteRow"] = result
	dataList["deleteId"] = id
	//成功返回
	ReturnSuccess("删除数据成功", c, dataList)
	return

}

/**
更新状态
 */
func LoveTalkUpdateStatus(c *gin.Context) {
	//获取用户参数
	id, _ := GetRequestParameters(c, "id", 2)         //Id
	status, _ := GetRequestParameters(c, "status", 2) //状态 1：显示  2：隐藏
	//判断Id
	if id <= 0 {
		ReturnError("请输入Id", c)
		return
	}
	if status != 1 && status != 2 {
		ReturnError("状态有误", c)
		return
	}
	//更新数据
	loveTalkParams := LoveTalk{
		Status: status,
	}
	loveTalkWhere := LoveTalk{
		Id: id,
	}
	updateRow := loveTalkParams.LoveTalkUpdate(loveTalkWhere)
	if updateRow <= 0 {
		ReturnError("状态更新失败", c)
		return
	}
	//初始化返回数据
	dataList := make(map[string]interface{})
	dataList["updateId"] = id
	//成功返回
	ReturnSuccess("状态更新成功", c, dataList)
	return
}
