package dao

import (
	"github.com/abeir/GoMybatis"
	"github.com/abeir/GoMybatis/tx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/satori/go.uuid"
	"io/ioutil"
)

const DB = "sqlite3"
const SQLITE_DB = "/home/abeir/workspace/syberos/spm-serv/data.db"

func init(){
	engine := GoMybatis.GoMybatisEngine{}.New()
	_, err := engine.Open(DB, SQLITE_DB)
	if err!=nil {
		panic(err)
	}
	bytes, err := ioutil.ReadFile("dao/mapper/LastVersionMapper.xml")
	if err!=nil {
		panic(err)
	}
	engine.WriteMapperPtr(&LastVersionDaoImpl, bytes)

	bytes, err = ioutil.ReadFile("dao/mapper/PackageProfileMapper.xml")
	if err!=nil {
		panic(err)
	}
	engine.WriteMapperPtr(&PackageProfileDaoImpl, bytes)

	bytes, err = ioutil.ReadFile("dao/mapper/UpgradeVersionMapper.xml")
	if err!=nil {
		panic(err)
	}
	engine.WriteMapperPtr(&UpgradeVersionDaoImp, bytes)
}


func UUID() string{
	return uuid.NewV4().String()
}


func Tx(session *GoMybatis.Session, f func()error) error{
	s := *session
	p := tx.PROPAGATION_REQUIRED
	err := s.Begin(&p)
	if err!=nil {
		return err
	}
	defer s.Close()
	err = f()
	if err!=nil {
		_ = s.Rollback()
		return err
	}
	_ = s.Commit()
	return nil
}