package routers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/jameshwc/Million-Singer/middleware/prom"
	"github.com/jameshwc/Million-Singer/middleware/version"
	"github.com/jameshwc/Million-Singer/routers/api/game"
	"github.com/jameshwc/Million-Singer/routers/api/user"
)

func methodNotAllowed(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{"error": "Method Not Allowed"})
}
func InitRouters() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(cors.Default())

	r.Use(version.RevisionMiddleware())

	r.Use(prom.PromMiddleware())
	r.GET("/metrics", prom.PromHandler(promhttp.Handler()))

	r.HandleMethodNotAllowed = true
	r.NoMethod(methodNotAllowed)
	userAPI := r.Group("/api/users")

	userAPI.GET("/check", user.ValidateUser)

	userAPI.POST("/register", user.Register)

	userAPI.POST("/login", user.Login)

	gamesAPI := r.Group("/api/game")

	gamesAPI.GET("/tours/:id", game.GetTour)

	gamesAPI.GET("/tours", game.GetTotalTours)

	gamesAPI.GET("/collects/:id", game.GetCollect)

	gamesAPI.GET("/songs/:id", game.GetSongInstance)

	gamesAPI.GET("/lyrics/:id", game.GetLyricsWithSongID)

	// gamesAPI.Use(jwt.JWT())
	// {
	gamesAPI.POST("/tours/new", game.AddTour)

	gamesAPI.POST("/collects/new", game.AddCollect)

	gamesAPI.POST("/songs/new", game.AddSong)
	// }
	return r
}
