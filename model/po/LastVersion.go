package po

import (
	"time"
)

type LastVersion struct {
	//Id
	Id string				`json:"id" gorm:"primary_key,column:id"`
	//包名
	PkgName string			`json:"pkgName" gorm:"column:pkg_name"`
	//版本号
	PkgVersion string		`json:"pkgVersion" gorm:"column:pkg_version"`
	//包概述信息
	PkgProfileId string		`json:"pkgProfileId" gorm:"column:pkg_profile_id"`
	//创建时间
	CreatedAt time.Time	`json:"createdAt" gorm:"column:created_at"`
	//修改时间
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (l *LastVersion) IsEmpty() bool{
	return l.Id=="" && l.PkgName==""
}

func (LastVersion) TableName() string {
	return "last_version"
}
