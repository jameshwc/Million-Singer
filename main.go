package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/conf"
	_ "github.com/jameshwc/Million-Singer/docs"
	"github.com/jameshwc/Million-Singer/model"
	"github.com/jameshwc/Million-Singer/pkg/gredis"
	"github.com/jameshwc/Million-Singer/routers"
	_ "github.com/joho/godotenv/autoload"
	swaggerFile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	conf.Setup()
	model.Setup(nil)
	gredis.Setup()
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

	log.Printf("[info] start http server listening %s", endPoint)
	if gin.Mode() == gin.DebugMode {
		apiDocURL := ginSwagger.URL(fmt.Sprintf(":%d/swagger/doc.json", conf.ServerConfig.HttpPort))
		routers.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFile.Handler, apiDocURL))
	}

	// test function code
	// if f, err := os.Open("/home/james/下載/[Toolbxs]Eminem - Beautiful (Edited) (Explicit)-English.srt"); err != nil {
	// log.Fatalf("open srt file")
	// } else {
	// subtitle.ReadSrtFromFile(f)
	// defer f.Close()
	// }

	// end test function

	server.ListenAndServe()
}
