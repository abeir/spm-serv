package po

import (
	"time"
)

type LastVersion struct {
	Id string				`json:"id" gorm:"primary_key,column:id"`
	PkgName string			`json:"pkgName" gorm:"column:pkg_name"`
	PkgVersion string		`json:"pkgVersion" gorm:"column:pkg_version"`
	PkgProfileId string		`json:"pkgProfileId" gorm:"column:pkg_profile_id"`
	CreatedAt time.Time	`json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (l *LastVersion) IsEmpty() bool{
	return l.Id=="" && l.PkgName==""
}

func (LastVersion) TableName() string {
	return "last_version"
}
