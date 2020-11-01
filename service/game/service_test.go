package game

import (
	"bytes"
	"database/sql"
	"io"
	"strconv"

	"github.com/gomodule/redigo/redis"
	"github.com/jameshwc/Million-Singer/model"
	"github.com/jameshwc/Million-Singer/pkg/subtitle"
)

type repoMockCacheBase struct {
}

type repoMockTourBase struct {
	id []int
}

type repoMockCollectBase struct {
	id    []int
	title []string
}

type repoMockSongBase struct {
	id []int
}

type repoMockUserBase struct {
}

type repoMockCacheServerError struct {
	repoMockCacheBase
}

type repoMockTourServerError struct {
	repoMockTourBase
}

type repoMockCollectServerError struct {
	repoMockCollectBase
}

type repoMockSongServerError struct {
	repoMockSongBase
}

type mockSubtitler struct {
}

type mockWeber struct {
}

func newRepoMockCacheBase() *repoMockCacheBase {
	return &repoMockCacheBase{}
}

func newRepoMockTourBase(id []int) *repoMockTourBase {
	return &repoMockTourBase{id}
}

func newRepoMockCollectBase(id []int) *repoMockCollectBase {
	var title []string
	for _, n := range id {
		title = append(title, "collect-"+strconv.Itoa(n))
	}
	return &repoMockCollectBase{id, title}
}

func newRepoMockSongBase(id []int) *repoMockSongBase {
	return &repoMockSongBase{id}
}

func newRepoMockCacheServerError() *repoMockCacheServerError {
	return &repoMockCacheServerError{}
}

func newRepoMockTourServerError() *repoMockTourServerError {
	return &repoMockTourServerError{}
}

func newRepoMockCollectServerError() *repoMockCollectServerError {
	return &repoMockCollectServerError{}
}

func newRepoMockSongServerError() *repoMockSongServerError {
	return &repoMockSongServerError{}
}

func (r *repoMockCacheBase) Set(key string, data interface{}, timeout int) error {
	return nil
}

func (r *repoMockCacheBase) Get(key string) ([]byte, error) {
	return []byte{'t', 'e', 's', 't'}, nil
}

func (r *repoMockCacheBase) Del(key string) error {
	return nil
}

func (r *repoMockTourBase) Add(collectsID []int) (int, error) {
	return len(r.id) + 1, nil
}

func (r *repoMockTourBase) Get(id int) (*model.Tour, error) {
	for _, n := range r.id {
		if n == id {
			return &model.Tour{n, nil}, nil
		}
	}
	return nil, sql.ErrNoRows
}

func (r *repoMockTourBase) GetTotal() (int, error) {
	return len(r.id), nil
}

func (r *repoMockCollectBase) Add(title string, songsID []int) (int, error) {
	return len(r.id) + 1, nil
}

func (r *repoMockCollectBase) Get(id int) (*model.Collect, error) {
	for _, n := range r.id {
		if n == id {
			return &model.Collect{n, "collect-" + strconv.Itoa(n), nil}, nil
		}
	}
	return nil, sql.ErrNoRows
}

func (r *repoMockCollectBase) CheckManyExist(collectsID []int) (int64, error) {
	cnt := int64(0)
	for i := range collectsID {
		for j := range r.id {
			if collectsID[i] == r.id[j] {
				cnt++
			}
		}
	}
	return cnt, nil
}

func (r *repoMockSongBase) Add(videoID, name, singer, genre, language, missLyrics, startTime, endTime string, lyrics []model.Lyric) (int, error) {
	return len(r.id) + 1, nil
}

func (r *repoMockSongBase) Get(songID int, hasLyrics bool) (*model.Song, error) {
	return &model.Song{1, nil, "", "", "", "", "", "", "", ""}, nil
}

func (r *repoMockSongBase) Delete(id int) error {
	return nil
}

func (r *repoMockSongBase) QueryByVideoID(videoID string) (id int64, err error) {
	return int64(1), nil
}

func (r *repoMockSongBase) CheckManyExist(songsID []int) (int64, error) {
	cnt := int64(0)
	for i := range songsID {
		for j := range r.id {
			if songsID[i] == r.id[j] {
				cnt++
			}
		}
	}
	return cnt, nil
}

func (r *repoMockCacheServerError) Set(key string, data interface{}, timeout int) error {
	return redis.ErrPoolExhausted
}

func (r *repoMockCacheServerError) Get(key string) ([]byte, error) {
	return nil, redis.ErrNil
}

func (r *repoMockCacheServerError) Del(key string) error {
	return redis.ErrPoolExhausted
}

func (r *repoMockTourServerError) Add(collectsID []int) (int, error) {
	return 0, sql.ErrTxDone
}

func (r *repoMockTourServerError) Get(id int) (*model.Tour, error) {
	return nil, sql.ErrConnDone
}

func (r *repoMockTourServerError) GetTotal() (int, error) {
	return 0, sql.ErrConnDone
}

func (r *repoMockCollectServerError) Add(title string, songsID []int) (int, error) {
	return 0, sql.ErrTxDone
}

func (r *repoMockCollectServerError) Get(id int) (*model.Collect, error) {
	return nil, sql.ErrConnDone
}

func (r *repoMockCollectServerError) CheckManyExist(collectsID []int) (int64, error) {
	return int64(0), sql.ErrConnDone
}

func (r *repoMockSongServerError) Add(videoID, name, singer, genre, language, missLyrics, startTime, endTime string, lyrics []model.Lyric) (int, error) {
	return 0, sql.ErrTxDone
}

func (r *repoMockSongServerError) Get(songID int, hasLyrics bool) (*model.Song, error) {
	return nil, sql.ErrConnDone
}

func (r *repoMockSongServerError) Delete(id int) error {
	return sql.ErrConnDone
}

func (r *repoMockSongServerError) QueryByVideoID(videoID string) (id int64, err error) {
	return int64(0), sql.ErrConnDone
}

func (r *repoMockSongServerError) CheckManyExist(songsID []int) (int64, error) {
	return int64(0), sql.ErrConnDone
}

func (s *mockSubtitler) ReadFromBytes(file []byte) ([]subtitle.Line, error) {
	return s.ReadFromFile(bytes.NewReader(file))
}

func (s *mockSubtitler) ReadFromFile(i io.Reader) ([]subtitle.Line, error) {
	return nil, nil
}

func NewMockSubtitler() subtitle.Subtitler {
	return &mockSubtitler{}
}

func (w *mockWeber) GetLines(url string) ([]subtitle.Line, error) {
	return nil, nil
}

func NewMockWeber() subtitle.Weber {
	return &mockWeber{}
}
