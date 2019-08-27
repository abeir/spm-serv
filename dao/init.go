package dao

import (
	"github.com/abeir/GoMybatis"
	"github.com/abeir/GoMybatis/tx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"spm-serv/core"
)

//初始化Dao
func InitDao(config *core.Config){
	logImpl := goMybatisLogger{}
	engine := GoMybatis.GoMybatisEngine{}.New()
	engine.SetLog(&logImpl)
	_, err := engine.Open(config.Database.Name, config.Database.Url)
	if err!=nil {
		panic(err)
	}
	mapper(&engine, &LastVersionDaoImpl, "config/mapper/LastVersionMapper.xml")
	mapper(&engine, &PackageProfileDaoImpl, "config/mapper/PackageProfileMapper.xml")
	mapper(&engine, &UpgradeVersionDaoImpl, "config/mapper/UpgradeVersionMapper.xml")
}

func mapper(engine *GoMybatis.GoMybatisEngine, ptr interface{}, xmlPath string){
	bytes, err := ioutil.ReadFile(xmlPath)
	if err!=nil {
		panic(err)
	}
	engine.WriteMapperPtr(ptr, bytes)
}


//GoMybatis 日志实现
type goMybatisLogger struct {
}

//日志消息队列长度
func (it *goMybatisLogger) QueueLen() int {
	//默认50万个日志消息缓存队列
	return 500000
}

//日志输出方法实现
func (it *goMybatisLogger) Println(v []byte) {
	core.Log.Debugln(string(v))
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