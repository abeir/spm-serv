package po

import (
	"time"
)

type PackageProfile struct {
	Id string			`json:"id"`
	//包名
	PkgName string		`json:"pkgName"`
	//包描述
	PkgDesc string		`json:"pkgDesc"`
	//仓库地址
	RepoUrl string		`json:"repoUrl"`
	//版本号
	PkgVersion string	`json:"pkgVersion"`
	//作者姓名
	AuthorName string `json:"authorName"`
	//作者email
	AuthorEmail string `json:"authorEmail"`
	//作者备注
	AuthorDesc string `json:"authorDesc"`
	//pri文件名
	PriFilename string `json:"priFilename"`

	CreatedAt time.Time	`json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (p *PackageProfile) IsEmpty() bool{
	return p.Id==""
}