/*
@Time: 2022/12/24 15:26
@Author: wxw
@File: server
*/
package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"src/com.wxw/project_actual/module/gin-example/common/infra/conf"
	"src/com.wxw/project_actual/module/gin-example/router"
)

// Setup 初始化服务
func Setup() {
	gin.SetMode(conf.Viper.GetString("server.run_mode"))
	routersInit := router.Setup()
	readTimeout := conf.Viper.GetDuration("server.read_timeout")
	writeTimeout := conf.Viper.GetDuration("server.write_timeout")
	endPoint := fmt.Sprintf(":%d", conf.Viper.GetInt("server.port"))
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
