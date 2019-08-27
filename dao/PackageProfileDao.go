package dao

import (
	"github.com/jinzhu/gorm"
	"spm-serv/model/pkg"
	"spm-serv/model/po"
)

type PackageProfileDao struct {
	db *gorm.DB
}

func (p *PackageProfileDao) SetDb(db *gorm.DB){
	p.db = db
}

func (p *PackageProfileDao) SelectByPkgNameAndPkgVersion(pkgName, pkgVersion string) po.PackageProfile{
	pkgProfile := po.PackageProfile{}
	//pkgVersion!=null
	p.db.Where(&po.PackageProfile{PkgName:pkgName, PkgVersion:pkgVersion}).First(&pkgProfile)
	return pkgProfile
}

func (p *PackageProfileDao) SelectLastVersionByPkgNameLike(pkgName string) []po.PackageProfile {
	var list []po.PackageProfile
	p.db.Raw(`
SELECT p.id, p.pkg_name, p.pkg_desc, p.repo_url, p.pkg_version, p.author_name, p.author_email,
p.author_desc, p.pri_filename, p.created_at, p.updated_at
FROM last_version v
JOIN package_profile p ON v.pkg_profile_id = p.id
WHERE v.pkg_name LIKE ?
ORDER BY v.created_at DESC
`, pkgName).Scan(&list)
	return list
}

func (p *PackageProfileDao) Insert(po po.PackageProfile) int64{
	return p.db.Create(&po).RowsAffected
}

func (p *PackageProfileDao) CountByPkgNameAndPkgVersion(pkgName, pkgVersion string) int64{
	var count int64
	p.db.Model(&po.PackageProfile{}).Where("pkg_name = ? and pkg_version = ?", pkgName, pkgVersion).Count(&count)
	return count
}
func (p *PackageProfileDao) SelectPageList(req pkg.PackageListReq) []po.PackageProfile {
	var list []po.PackageProfile
	p.db.Offset(req.GetPage()).Limit(req.GetRow()).Where(&po.PackageProfile{PkgName:req.PkgName, PkgVersion:req.PkgVersion}).Find(&list)
	return list
}

func (p *PackageProfileDao) Tx(f func()error) error{
	tx := p.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := f(); err!=nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
