package po

import (
	"spm-serv/core"
	"strconv"
	"time"
)

type UpgradeVersion struct {
	//ID
	Id string			`json:"id" gorm:"primary_key,column:id"`
	//版本号
	Version string 		`json:"version" gorm:"column:version"`
	//版本排序字段
	VersionSort string	`json:"versionSort" gorm:"column:version_sort"`
	//描述信息
	Description string 	`json:"description" gorm:"column:description"`
	//文件路径
	Path string 		`json:"path" gorm:"column:path"`
	//发布状态 0为未发布， 1为已发布， 2为已下架
	Status	string 		`json:"status" gorm:"column:status"`
	//创建时间
	CreatedAt time.Time	`json:"createdAt" gorm:"column:created_at"`
	//修改时间
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

func (u *UpgradeVersion) TableName() string {
	return "upgrade_version"
}

func (u *UpgradeVersion) IsEmpty() bool{
	return u.Id=="" && u.Version==""
}

//状态是否为未发布
func (u *UpgradeVersion) IsUnreleased() bool{
	return u.Status == "0"
}

//状态是否为已发布
func (u *UpgradeVersion) IsReleased() bool{
	return u.Status == "1"
}

//状态是否为已下架
func (u *UpgradeVersion) IsDetain() bool{
	return u.Status == "2"
}

//将Version属性值转换成VersionSort值
func (u *UpgradeVersion) ToVersionSort() string{
	if u.Version=="" {
		return ""
	}
	ver, _ := core.NewVersion(u.Version)

	major := strconv.Itoa(ver.Major)
	minor := strconv.Itoa(ver.Minor)
	revision := strconv.Itoa(ver.Revision)

	major, _ = core.PadLeft(major, "0", 6)
	minor, _ = core.PadLeft(minor, "0", 6)
	revision, _ = core.PadLeft(revision, "0", 6)
	return major + "-" + minor + "-" + revision
}
