/*
@Time: 2022/12/24 15:26
@Author: wxw
@File: server
*/
package server

import (
	"net/http"
)

// Setup 初始化服务
func Setup() {
	//gin.SetMode(setting.ServerSetting.RunMode)
	//routersInit := routers.InitRouter()
	//readTimeout := setting.ServerSetting.ReadTimeout
	//writeTimeout := setting.ServerSetting.WriteTimeout
	//endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		//Addr:           endPoint,
		//Handler:        routersInit,
		//ReadTimeout:    readTimeout,
		//WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	//log.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
}
