package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/conf"
	"github.com/jameshwc/Million-Singer/model"
	"github.com/jameshwc/Million-Singer/routers"
)

func init() {
	conf.Setup(filepath.Join("conf", "conf.toml"))
	model.Setup()
}

func main() {
	gin.SetMode(conf.Conf.Server.RunMode)
	routersInit := routers.InitRouters()
	endPoint := fmt.Sprintf(":%d", conf.Conf.Server.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
