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

	userAPI := r.Group("/api/users")
	userAPI.GET("/check", user.ValidateUser)
	userAPI.POST("/register", user.Register)
	userAPI.POST("/login", user.Login)

	gamesAPI := r.Group("/api/game")
	gamesAPI.GET("/tours/:id", game.GetTour)
	gamesAPI.GET("/tours", game.GetTotalTours)
	gamesAPI.GET("/levels/:id", game.GetLevel)
	gamesAPI.GET("/songs/:id", game.GetSongInstance)
	gamesAPI.GET("/lyrics/:id", game.GetLyricsWithSongID)

	// gamesAPI.Use(jwt.JWT())
	{
		gamesAPI.POST("/levels/new", game.AddLevel)
		gamesAPI.POST("/songs/new", game.AddSong)
	}
	return r
}
