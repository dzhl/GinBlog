package model

import (
	"ginblog/utils/errmsg"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

//检查分类是否存在
func CheckCategory(name string) (code int) {
	var cat Category
	db.Where("name=?", name).First(&cat)
	if cat.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCSE
}

//新增分类
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//查询分类列表
func GetCategorys(pageSize int, pageNum int) []Category {
	var cats []Category
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cats).Error
	if err != nil {
		return nil
	}
	return cats
}

//编辑分类
func EditCategory(id int, data *Category) int {

	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := db.Model(&Category{}).Where("id=?", id).Updates(&maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//删除分类
func DeleteCategory(id int) int {
	var cat Category
	err := db.Where("id=?", id).Delete(&cat).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
