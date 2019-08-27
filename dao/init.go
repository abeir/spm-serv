package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"spm-serv/core"
)

var LastVersionDaoImpl LastVersionDao
var PackageProfileDaoImpl PackageProfileDao
var UpgradeVersionDaoImpl UpgradeVersionDao

var db *gorm.DB

func CloseDb(){
	if db!=nil {
		err := db.Close()
		if err!=nil {
			core.Log.Panicf("cannot close database: %+v", errors.WithStack(err))
		}
	}
}

//初始化Dao
func InitDao(config *core.Config){
	db, err := gorm.Open(config.Database.Name, config.Database.Url)
	if err!=nil {
		core.Log.Panicf("cannot open database: %s, %s, %+v", config.Database.Name, config.Database.Url, errors.WithStack(err))
	}
	db.LogMode(true)
	db.SetLogger(core.Log)

	LastVersionDaoImpl = LastVersionDao{db}
	PackageProfileDaoImpl = PackageProfileDao{db}
	UpgradeVersionDaoImpl = UpgradeVersionDao{db}
}


func UUID() string{
	return uuid.NewV4().String()
}