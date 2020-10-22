package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/conf"
	_ "github.com/jameshwc/Million-Singer/docs"
	"github.com/jameshwc/Million-Singer/pkg/gredis"
	"github.com/jameshwc/Million-Singer/pkg/log"
	"github.com/jameshwc/Million-Singer/repo/mysql"
	"github.com/jameshwc/Million-Singer/routers"
	_ "github.com/joho/godotenv/autoload"
	swaggerFile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	conf.Setup()
	log.Setup()
	gredis.Setup()
	mysql.Setup()
}

// @title Million Singer API
// @version 1.0
// @description To add a level into database
// @contact.name jameshwc
// @contact.url https://jameshsu.csie.org
// @contact.email jameshwc@gmail.com
// @license.name GPL v3.0
// @license.url https://www.gnu.org/licenses/gpl-3.0.html
// @host ms.csie.org
// @BasePath /api
// @schemes http
func main() {
	gin.SetMode(conf.ServerConfig.RunMode)
	routers := routers.InitRouters()
	endPoint := fmt.Sprintf(":%d", conf.ServerConfig.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routers,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Infof("start http server listening %s", endPoint)
	if gin.Mode() == gin.DebugMode {
		apiDocURL := ginSwagger.URL(fmt.Sprintf(":%d/swagger/doc.json", conf.ServerConfig.HttpPort))
		routers.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFile.Handler, apiDocURL))
	}

	server.ListenAndServe()
}
