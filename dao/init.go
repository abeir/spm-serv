package dao

import (
	"github.com/gookit/color"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"spm-serv/core"
)

var LastVersionDaoImpl *LastVersionDao
var PackageProfileDaoImpl *PackageProfileDao
var UpgradeVersionDaoImpl *UpgradeVersionDao

var db *gorm.DB

//关闭db，应在关闭服务器时调用
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
	color.Println("<light_green>open database:</>", config.Database.Name)
	db.LogMode(true)
	db.SetLogger(GormLogger{})

	LastVersionDaoImpl = NewLastVersionDao(db)
	PackageProfileDaoImpl = NewPackageProfileDao(db)
	UpgradeVersionDaoImpl = NewUpgradeVersionDao(db)
}

//GormLogger Gorm日志
type GormLogger struct {
}

func (GormLogger) Print(v ...interface{}){
	core.Log.Printf("| %s | %s | %s | %s | %s | %d ", v...)
}


//dao公共组件
type CommonDao struct {
	db *gorm.DB
}

func (CommonDao) UUID() string{
	return uuid.NewV4().String()
}

func (a *CommonDao) Tx(f func()error) error{
	tx := a.db.Begin()
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

