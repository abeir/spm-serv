package po

import (
	"time"
)

type PackageProfile struct {
	Id string			`json:"id" gorm:"primary_key,column:id"`
	//包名
	PkgName string		`json:"pkgName" gorm:"column:pkg_name"`
	//包描述
	PkgDesc string		`json:"pkgDesc" gorm:"column:pkg_desc"`
	//仓库地址
	RepoUrl string		`json:"repoUrl" gorm:"column:repo_url"`
	//版本号
	PkgVersion string	`json:"pkgVersion" gorm:"column:pkg_version"`
	//作者姓名
	AuthorName string `json:"authorName" gorm:"column:author_name"`
	//作者email
	AuthorEmail string `json:"authorEmail" gorm:"column:author_email"`
	//作者备注
	AuthorDesc string `json:"authorDesc" gorm:"column:author_desc"`
	//pri文件名
	PriFilename string `json:"priFilename" gorm:"column:pri_filename"`

	Status	string 		`json:"status" gorm:"column:status"`

	CreatedAt time.Time	`json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (p *PackageProfile) TableName() string{
	return "package_profile"
}

func (p *PackageProfile) IsEmpty() bool{
	return p.Id == "" && p.PkgName == ""
}

//状态是否为可用
func (p *PackageProfile) IsEnabled() bool{
	return p.Status == "1"
}

//状态是否为不可用
func (p *PackageProfile) IsDisabled() bool{
	return p.Status == "0"
}