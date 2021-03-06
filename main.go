package main

import (
	"github.com/gobuffalo/packr"
	config "github.com/cqlsy/yeelib/conf"
	"github.com/cqlsy/yeelib/webGin"
	"github.com/cqlsy/yeelib/yeeDb"
	"github.com/cqlsy/yeelib/yeeGin"
	"github.com/cqlsy/yeelib/yeelog"
)

const configPath = "./config.json"

func main() {
	// 初始化配置文件
	config.ParseConf(configPath)
	yeelog.MustInitLog(config.Conf.Web.LogPath, config.Conf.Web.RunMode)
	database := yeeDb.InitDataBase()
	gin := yeeGin.New()
	initStatic(gin)
	// 保存相关的实例以便我们在其他地方的直接调用
	webGin.Init(gin, database)
	// 开开启web服务
	gin.StartListen(config.Conf.Web.Ip, config.Conf.Web.Port)
}

func initStatic(gin *yeeGin.WebGin) {
	gin.AddStaticPath("/", packr.NewBox("./static"))
}
