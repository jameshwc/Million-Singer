package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/routers/api/game"
)

func InitRouters() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gamesAPI := r.Group("/api/game")
	gamesAPI.GET("/games/", game.GetGame)
	return r
}
