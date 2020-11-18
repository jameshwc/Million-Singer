package service

import (
	"github.com/jameshwc/Million-Singer/model"
	"github.com/jameshwc/Million-Singer/service/game"
	"github.com/jameshwc/Million-Singer/service/user"
)

type GameService interface {
	GetCollect(param string) (*model.Collect, error)
	AddCollect(songs []int, title string) (int, error)
	GetLyricsWithSongID(param string) ([]*model.Lyric, error)
	GetSupportedLanguages() []string
	GetGenres() []string
	AddSong(s *game.Song) (int, error)
	GetSongInstance(param string, hasLyrics bool) (*game.SongInstance, error)
	GetSongs() ([]*model.Song, error)
	DeleteSong(param string) error
	AddTour(collectsID []int) (int, error)
	GetTotalTours() (int, error)
	GetTour(param string) (*model.Tour, error)
}

type UserService interface {
	Auth(username, password string) (string, error)
	Validate(username string, email string) error
	Register(username, email, password string) error
}

var Game GameService
var User UserService

func init() {
	Game = game.NewGameService()
	User = user.NewUserService()
}
