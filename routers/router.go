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
	gamesAPI.GET("/games/:id", game.GetGame)
	gamesAPI.POST("/levels/new", game.AddLevel)
	gamesAPI.GET("/levels/:id", game.GetLevel)
	gamesAPI.POST("/songs/new", game.AddSong)
	gamesAPI.GET("/songs/:id", game.GetSongInstance)
	gamesAPI.GET("/lyrics", game.GetLyricsWithSongID)
	return r
}
