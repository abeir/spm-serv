package dao

import "spm-serv/model/po"

var UpgradeVersionDaoImpl = UpgradeVersionDao{}

type UpgradeVersionDao struct {
	//查询最新的版本
	SelectLatestVersion func() (po.UpgradeVersion, error)

	SelectByVersion func(version string) (po.UpgradeVersion, error)	`mapperParams:"version"`

	//指定版本号的数量
	CountByVersion func(version string) (int64, error)		`mapperParams:"version"`

	Insert	func(*po.UpgradeVersion) (int64, error)

}