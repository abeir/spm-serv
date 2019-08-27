package po

import "time"

type ConsoleUser struct {
	Id string `json:"id" gorm:"primary_key,column:id"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`

	Status	string 		`json:"status" gorm:"column:status"`

	CreatedAt time.Time	`json:"createdAt" gorm:"column:created_at"`
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