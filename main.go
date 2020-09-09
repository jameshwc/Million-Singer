package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/conf"
	"github.com/jameshwc/Million-Singer/model"
	"github.com/jameshwc/Million-Singer/routers"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	conf.Setup()
	model.Setup()
}

func main() {
	gin.SetMode(conf.ServerConfig.RunMode)
	routersInit := routers.InitRouters()
	endPoint := fmt.Sprintf(":%d", conf.ServerConfig.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

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
