package routers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jameshwc/Million-Singer/middleware/iplog"
	"github.com/jameshwc/Million-Singer/middleware/prom"
	"github.com/jameshwc/Million-Singer/middleware/version"
	"github.com/jameshwc/Million-Singer/routers/api/game"
	"github.com/jameshwc/Million-Singer/routers/api/user"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	r.Use(iplog.TraceIP())
	r.Use(static.Serve("/", static.LocalFile("dist", true)))

	r.GET("/metrics", prom.PromHandler(promhttp.Handler()))

	r.HandleMethodNotAllowed = true
	r.NoMethod(methodNotAllowed)
	userAPI := r.Group("/api/users")

	userAPI.GET("/check", user.ValidateUser)

	userAPI.POST("/register", user.Register)

	userAPI.POST("/login", user.Login)

	gamesAPI := r.Group("/api/game")

	v2gamesAPI := r.Group("/api/v2/game")
	gamesAPI.GET("/tours/:id", game.GetTour)

	gamesAPI.DELETE("/tours/:id", game.DelTour)

	gamesAPI.GET("/tours", game.GetTotalTours)

	v2gamesAPI.GET("/tours", game.GetTours)

	gamesAPI.DELETE("/collects/:id", game.DelCollect)

	gamesAPI.GET("/collects/:id", game.GetCollect)

	gamesAPI.GET("/collects", game.GetCollects)

	gamesAPI.GET("/songs/:id", game.GetSongInstance)

	gamesAPI.GET("/songs", game.GetSongs)

	gamesAPI.GET("/lyrics/:id", game.GetLyricsWithSongID)

	gamesAPI.GET("/languages", game.GetSupportedLanguages)

	gamesAPI.GET("/genres", game.GetGenres)

	gamesAPI.GET("/captions/youtube", game.ListYoutubeCaptionLanguages)

	gamesAPI.GET("/subtitles/youtube", game.DownloadYoutubeSubtitle)
	gamesAPI.GET("/youtube/title", game.GetYoutubeTitle) // TODO: make api endpoints more tidy
	gamesAPI.POST("/subtitles/convert", game.ConvertFileToSubtitle)

	// gamesAPI.Use(jwt.JWT())
	// {
	gamesAPI.POST("/tours/new", game.AddTour)

	gamesAPI.POST("/collects/new", game.AddCollect)

	gamesAPI.POST("/songs/new", game.AddSong)

	gamesAPI.DELETE("/songs/:id", game.DelSong)

	// }

	r.NoRoute(static.Serve("/", static.LocalFile("dist", true)))

	return r
}
