package dao

import (
	"github.com/jinzhu/gorm"
	"spm-serv/model/po"
)

func NewUpgradeVersionDao(db *gorm.DB) *UpgradeVersionDao{
	return &UpgradeVersionDao{db}
}


type UpgradeVersionDao struct {
	db *gorm.DB
}

//查询最新的版本
func (u *UpgradeVersionDao) SelectLatestVersion() po.UpgradeVersion{
	result := po.UpgradeVersion{}
	u.db.Where("status = '1'").Order("created_at desc").First(&result)
	//u.db.Where("created_at = (?) and status = '1'", u.db.Table("upgrade_version").Select("max(created_at)").QueryExpr()).First(&result)
	return result
}

func (u *UpgradeVersionDao) SelectByVersionRelease(version string) po.UpgradeVersion{
	result := po.UpgradeVersion{}
	u.db.Where("version = ? and status = '1'", version).First(&result)
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