package dao

import (
	"github.com/jinzhu/gorm"
	"spm-serv/model/po"
)

type LastVersionDao struct {
	db *gorm.DB
}

func (l *LastVersionDao) SetDb(db *gorm.DB){
	l.db = db
}

func (l *LastVersionDao) SelectByPkgNameLike(pkgName string) []po.LastVersion {
	var versions = make([]po.LastVersion, 10)
	l.db.Where("pkg_name LIKE ?", "%" + pkgName + "%").Find(&versions)
	return versions
}

func (l *LastVersionDao) SelectByPkgName(pkgName string) po.LastVersion {
	version := po.LastVersion{}
	l.db.Where("pkg_name = ?", pkgName).Find(&version)
	return version
}

func (l *LastVersionDao) Insert(po po.LastVersion) {
	l.db.Create(po)
}

func (l *LastVersionDao) UpdateByPkgName(p po.LastVersion) {
	l.db.Model(po.LastVersion{}).Where("pkg_name = ?", p.PkgName).Updates(p)
}



