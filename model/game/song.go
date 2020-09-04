package game

type Song struct {
	SongID    int
	Lyrics    []Lyric
	URL       string
	StartTime string
	EndTime   string
	Language  string
	Name      string
	Singer    string
	Genre     []string
}

type GameSong struct {
	Song
	ID          int
	MissLyricID int
}
