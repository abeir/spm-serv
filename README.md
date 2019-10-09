# spm-serv

基于go语言实现的Qt包管理工具（[spm](https://github.com/syberos-team/spm)）的服务端

服务启动时，首先尝试从环境变量（SPM_CONFIG）指定的路径中读取配置文件，若未配置环境变量，则在服务启动程序所在的位置下寻找config目录，并从中依次查找spm.yml、spm.yaml、spm.json

---

Web框架
[Gin](https://github.com/gin-gonic/gin)

ORM框架
[Gorm](https://github.com/jinzhu/gorm)

Log框架
[Logrus](https://github.com/sirupsen/logrus)
