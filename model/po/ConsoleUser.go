package po

import "time"
//后台用户表
type ConsoleUser struct {
	//Id
	Id string `json:"id" gorm:"primary_key,column:id"`
	//用户名称
	Username string `json:"username" gorm:"column:username"`
	//密码
	Password string `json:"password" gorm:"column:password"`
	//状态 0为不可用，1为可用
	Status	string 		`json:"status" gorm:"column:status"`
	//创建时间
	CreatedAt time.Time	`json:"createdAt" gorm:"column:created_at"`
	//修改时间
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (c *ConsoleUser) TableName() string{
	return "console_user"
}

func (c *ConsoleUser) IsEmpty() bool{
	return c.Id == "" && c.Username == ""
}

func (c *ConsoleUser) IsEnable() bool{
	return c.Status=="1"
}

func (c *ConsoleUser) IsDisable() bool{
	return c.Status=="0"
}