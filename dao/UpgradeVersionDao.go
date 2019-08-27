package dao

import (
	"github.com/jinzhu/gorm"
	"spm-serv/model/po"
)

type UpgradeVersionDao struct {
	db *gorm.DB
}

//查询最新的版本
func (u *UpgradeVersionDao) SelectLatestVersion() po.UpgradeVersion{
	result := po.UpgradeVersion{}
	u.db.Where("version_sort = ?", u.db.Table("upgrade_version").Select("max(version_sort)").QueryExpr()).First(&result)
	return result
}

func (u *UpgradeVersionDao) SelectByVersion(version string) po.UpgradeVersion{
	result := po.UpgradeVersion{}
	u.db.Where("version = ?", version).First(&result)
	return result
}
//指定版本号的数量
func (u *UpgradeVersionDao) CountByVersion(version string) int64 {
	var count int64
	u.db.Where("version = ?", version).Count(&count)
	return count
}


func (u *UpgradeVersionDao) Insert(po *po.UpgradeVersion) int64{
	return u.db.Create(po).RowsAffected
}