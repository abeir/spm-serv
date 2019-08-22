package dao

import (
	"github.com/abeir/GoMybatis"
	"spm-serv/model/po"
)

var PackageProfileDaoImpl = PackageProfileDao{}

type PackageProfileDao struct {
	GoMybatis.SessionSupport

	SelectByPkgNameAndPkgVersion	func(pkgName, pkgVersion string)(po.PackageProfile, error) `mapperParams:"pkgName,pkgVersion"`

	SelectLastVersionByPkgNameLike func(pkgName string) ([]po.PackageProfile, error)	`mapperParams:"pkgName"`

	Insert func(session *GoMybatis.Session, po po.PackageProfile) (int64, error)

	CountByPkgNameAndPkgVersion func(pkgName, pkgVersion string) (int64, error) `mapperParams:"pkgName,pkgVersion"`
}
