package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/routers/api/game"
	"github.com/jameshwc/Million-Singer/routers/api/user"
)

func InitRouters() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gamesAPI := r.Group("/api/game")
	gamesAPI.GET("/tours/:id", game.GetTour)
	gamesAPI.POST("/levels/new", game.AddLevel)
	gamesAPI.GET("/levels/:id", game.GetLevel)
	gamesAPI.POST("/songs/new", game.AddSong)
	gamesAPI.GET("/songs/:id", game.GetSongInstance)
	gamesAPI.GET("/lyrics", game.GetLyricsWithSongID)

	userAPI := r.Group("/api/users")
	userAPI.GET("/check", user.ValidateUser)
	return r
}
