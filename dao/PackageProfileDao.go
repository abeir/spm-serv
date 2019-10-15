package dao

import (
	"github.com/jinzhu/gorm"
	"spm-serv/model/pkg"
	"spm-serv/model/po"
)

func NewPackageProfileDao(db *gorm.DB) *PackageProfileDao{
	return &PackageProfileDao{
		CommonDao: CommonDao{db},
		db:          db,
	}
}

type PackageProfileDao struct {
	CommonDao
	db *gorm.DB
}

// 通过包名、版本号查询状态是可用的包信息
func (p *PackageProfileDao) SelectByPkgNameAndPkgVersion(pkgName, pkgVersion string) po.PackageProfile{
	pkgProfile := po.PackageProfile{}
	//pkgVersion!=null
	p.db.Where(&po.PackageProfile{PkgName:pkgName, PkgVersion:pkgVersion, Status:"1"}).First(&pkgProfile)
	return pkgProfile
}

// 使用包名模糊查询最新的包
func (p *PackageProfileDao) SelectLastVersionByPkgNameLike(pkgName string) []po.PackageProfile {
	var list []po.PackageProfile
	p.db.Raw(`
SELECT p.id, p.pkg_name, p.pkg_desc, p.repo_url, p.pkg_version, p.author_name, p.author_email,
p.author_desc, p.pri_filename, p.created_at, p.updated_at
FROM last_version v
JOIN package_profile p ON v.pkg_profile_id = p.id
WHERE v.pkg_name LIKE ? AND p.status = '1'
ORDER BY v.created_at DESC
`, "%" + pkgName + "%").Scan(&list)
	return list
}

func (p *PackageProfileDao) Insert(po po.PackageProfile) int64{
	return p.db.Create(&po).RowsAffected
}

func (p *PackageProfileDao) CountByPkgNameAndPkgVersion(pkgName, pkgVersion string) int64{
	var count int64
	p.db.Model(&po.PackageProfile{}).Where("pkg_name = ? and pkg_version = ? and status = '1'", pkgName, pkgVersion).Count(&count)
	return count
}
func (p *PackageProfileDao) SelectPageList(req pkg.PackageListReq) []po.PackageProfile {
	var list []po.PackageProfile
	p.db.Offset(req.GetPage()).Limit(req.GetRow()).Where(&po.PackageProfile{PkgName:req.PkgName, PkgVersion:req.PkgVersion, Status:"1"}).Find(&list)
	return list
}