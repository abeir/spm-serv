package dao

import (
	"github.com/jinzhu/gorm"
	"spm-serv/model/po"
)

func NewConsoleUserDao(db *gorm.DB) *ConsoleUserDao{
	return &ConsoleUserDao{db}
}

type ConsoleUserDao struct {
	db *gorm.DB
}

//根据username查询后台用户信息
func (c *ConsoleUserDao) SelectByUserName(username string) po.ConsoleUser{
	result := po.ConsoleUser{}
	c.db.Where("username = ? and status = ?", username, 1).First(&result)
	return result
}
