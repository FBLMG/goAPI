package models

/**
数据模型-情话板
 */

//引入模块
import (
	orm "personApi/db"
)

//表结构构造
type LoveTalk struct {
	Id        int    `json:"id" form:"id"`
	TypeId    int    `json:"type_id" form:"type_id"`
	Motto     string `json:"motto" form:"motto"`
	MottoCard string `json:"motto_card" form:"motto_card"`
	Card      string `json:"card" form:"card"`
	Status    int    `json:"status" form:"status"`
	Sort      int    `json:"sort" form:"sort"`
}

/**
根据条件获取单条记录
 */
func (loveTalk *LoveTalk) LoveTalkWithWhere() (LoveTalk) {
	//初始化容器
	var data LoveTalk
	//获取单条数据
	orm.SqlDB.Where(&loveTalk).First(&data)
	//返回数据
	return data
}

/**
获取数据列表【带分页】
 */
func (loveTalk *LoveTalk) LoveTalkGetDataList(offset, pageNumber int) ([]LoveTalk) {
	//初始化数据容器
	var dataList []LoveTalk
	//获取数据
	orm.SqlDB.Where(&loveTalk).Limit(pageNumber).Offset(offset).Order("sort asc").Find(&dataList)
	//返回
	return dataList
}

/**
获取数据总数
 */
func (loveTalk *LoveTalk) LoveTalkGetDataCount() int {
	//初始化数据容器
	var dataList []LoveTalk
	var count int
	//获取数据
	orm.SqlDB.Where(&loveTalk).Find(&dataList).Count(&count)
	//返回
	return count
}

/**
添加记录
 */
func (loveTalk LoveTalk) LoveTalkInsert() int {
	//添加数据
	result := orm.SqlDB.Create(&loveTalk)
	id := loveTalk.Id
	if result.Error != nil {
		return 0
	}
	return id
}

/**
更新数据
 */
func (loveTalk LoveTalk) LoveTalkUpdate(updateData LoveTalk) int {
	//更新数据
	result := orm.SqlDB.Model(&updateData).Updates(&loveTalk)
	if result.Error != nil {
		return 0
	}
	//返回更新Id
	return 1
}

/**
条件更新
 */
func (loveTalk LoveTalk) LoveTalkUpdateWithWhere(updateWhere, updateData LoveTalk) int {
	//更新数据
	result := orm.SqlDB.Model(&updateData).Where(&updateWhere).Updates(&loveTalk)
	if result.Error != nil {
		return 0
	}
	//返回更新Id
	return 1
}

/**
删除数据
 */
func (loveTalk *LoveTalk) LoveTalkDelete() int64 {
	//删除数据
	result := orm.SqlDB.Delete(&loveTalk)
	if result.Error != nil {
		return 0
	}
	//返回删除行数
	return 1
}

/**
获取数据列表【不带分页】
 */
func (loveTalk *LoveTalk) LoveTalkGetDataListWithHome() ([]LoveTalk) {
	//初始化数据容器
	var dataList []LoveTalk
	//获取数据
	orm.SqlDB.Where(&loveTalk).Order("sort asc").Find(&dataList)
	//返回
	return dataList
}
