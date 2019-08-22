package dao

import (
	"github.com/abeir/GoMybatis"
	"spm-serv/model/po"
)

var LastVersionDaoImpl = LastVersionDao{}

type LastVersionDao struct {
	GoMybatis.SessionSupport

	SelectByPkgNameLike	func(pkgName string) ([]po.LastVersion, error)	`mapperParams:"pkgName"`

	SelectByPkgName func(pkgName string) (po.LastVersion, error)	`mapperParams:"pkgName"`

	Insert func(session *GoMybatis.Session, po po.LastVersion) (int64, error)

	UpdateByPkgName func(session *GoMybatis.Session, po po.LastVersion) (int64, error)
}
