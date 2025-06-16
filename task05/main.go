package main

import (
	"fmt"
	"os"
	"task05/dao"
	"task05/model"
	"task05/router"
	"task05/setting"
)

const defaultConfFile = "./config/config.ini"

func main() {
	confFile := defaultConfFile
	if len(os.Args) > 2 {
		fmt.Println("use specified conf file: ", os.Args[1])
		confFile = os.Args[1]
	} else {
		fmt.Println("no configuration file was specified, use ./conf/config.ini")
	}
	// 加载配置文件
	if err := setting.Init(confFile); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}
	// 连接数据库
	err := dao.InitDataBase(setting.Conf.DatabaseConfig)
	if err != nil {
		fmt.Printf("init mysql failed,err:%v\n", err)
		return
	}
	defer dao.Close() // 程序退出关闭数据库连接
	// 根据模型创建数据库表项
	err = dao.DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	if err != nil {
		return
	}

	// 启动gin服务
	r := router.SetRouter()

	// 在指定端口上启动web服务
	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
