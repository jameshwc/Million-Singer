package subtitle

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/beevik/etree"
	"github.com/jameshwc/Million-Singer/model"
)

type youtubeDownloader struct {
	URL       string
	VideoID   string
	Languages map[string]string
}

func newYoutubeDownloader(URL string) (*youtubeDownloader, error) {
	y := new(youtubeDownloader)
	y.URL = URL
	var err error
	y.VideoID, err = ParseVideoID(URL)
	if err != nil {
		return nil, err
	}

	y.Languages = make(map[string]string)
	if err := y.getAvailableLanguages(); err != nil {
		return nil, err
	}

	return y, nil
}

// Examples:
// - http://youtu.be/5MgBikgcWnY
// - http://www.youtube.com/watch?v=5MgBikgcWnY&feature=feed
// - http://www.youtube.com/embed/5MgBikgcWnY
// - http://www.youtube.com/v/5MgBikgcWnY?version=3&amp;hl=en_US
func ParseVideoID(URL string) (string, error) {
	u, err := url.Parse(URL)
	if err != nil {
		return "", err
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return "", fmt.Errorf("not supported scheme (http/https)")
	}

	switch u.Host {

	case "youtu.be":
		return u.Path[1:], nil

	case "youtube.com", "www.youtube.com":

		switch {

		case strings.HasPrefix(u.Path, "/watch"):
			return u.Query().Get("v"), nil

		case strings.HasPrefix(u.Path, "/v/"), strings.HasPrefix(u.Path, "/embed/"):
			return strings.Split(u.Path, "/")[2], nil

		default:
			return "", fmt.Errorf("path not correct")

		}

	default:
		return "", fmt.Errorf("host not correct")
	}
}

func (y *youtubeDownloader) getAvailableLanguages() error {
	URL := fmt.Sprintf("http://www.youtube.com/api/timedtext?v=%s&type=list", y.VideoID)

	resp, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc := etree.NewDocument()
	if _, err := doc.ReadFrom(resp.Body); err != nil {
		return err
	}

	root := doc.SelectElement("transcript_list")
	for _, s := range root.SelectElements("track") {
		y.Languages[s.SelectAttr("lang_original").Value] = s.SelectAttr("lang_code").Value
		if s.SelectAttrValue("lang_default", "false") == "true" {
			y.Languages["default"] = s.SelectAttr("lang_code").Value
		}
	}
	if _, ok := y.Languages["default"]; !ok {
		y.Languages["default"] = root.SelectElements("track")[0].SelectAttr("lang_code").Value // TODO: is zero index always correct lang?
	}
	return nil
}

func (y *youtubeDownloader) getLyrics() ([]model.Lyric, error) {
	doc, err := y.download()
	if err != nil {
		return nil, err
	}
	var lyrics []model.Lyric
	idx := 1
	for _, s := range doc.SelectElements("text") {
		var l model.Lyric
		start, err := strconv.ParseFloat(s.SelectAttr("start").Value, 64)
		if err != nil {
			return nil, err
		}
		dur, err := strconv.ParseFloat(s.SelectAttr("dur").Value, 64)
		if err != nil {
			return nil, err
		}
		l.StartAt = time.Duration(start * float64(time.Second))
		l.EndAt = time.Duration((start + dur) * float64(time.Second))
		l.Line = strings.ReplaceAll(s.Text(), "&#39;", "'")
		l.Index = idx
		idx++
		lyrics = append(lyrics, l)
	}
	return lyrics, nil
}

func (y *youtubeDownloader) download() (*etree.Element, error) {
	downloadUrl := fmt.Sprintf("http://www.youtube.com/api/timedtext?v=%s&lang=%s", y.VideoID, y.Languages["default"])

	resp, err := http.Get(downloadUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc := etree.NewDocument()
	if _, err := doc.ReadFrom(resp.Body); err != nil {
		return nil, err
	}
	return doc.SelectElement("transcript"), nil

}

func GetLyricsFromYoutubeSubtitle(url string) ([]model.Lyric, error) {
	y, err := newYoutubeDownloader(url)
	if err != nil {
		return nil, err
	}

	return y.getLyrics()
}
