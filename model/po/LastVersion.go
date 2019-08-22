package po

import (
	"time"
)

type LastVersion struct {
	Id string				`json:"id"`
	PkgName string			`json:"pkgName"`
	PkgVersion string		`json:"pkgVersion"`
	PkgProfileId string		`json:"pkgProfileId"`
	CreatedAt time.Time	`json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (l *LastVersion) IsEmpty() bool{
	return l.Id=="" && l.PkgName==""
}
