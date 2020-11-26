package subtitle

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/beevik/etree"
)

const watchURL = "https://youtube.com/watch?v=%s"

// TODO: complete this map
var languageCodeMap = map[string]string{
	"ar":      "العربية",
	"az":      "Azərbaycan",
	"zh":      "中文",
	"zh-CN":   "中文（简体）",
	"zh-TW":   "中文（繁體）",
	"cs":      "Čeština",
	"nl":      "Nederlands",
	"en":      "English",
	"fil":     "Filipino",
	"fr":      "Français",
	"de":      "Deutsch",
	"hi":      "हिन्दी",
	"hu":      "Magyar",
	"id":      "Indonesia",
	"ja":      "日本語",
	"jv":      "Jawa",
	"kk":      "Қазақ Тілі",
	"km":      "ខ្មែរ",
	"ko":      "한국어",
	"mk":      "Македонски",
	"nan":     "Min Nan Chinese",
	"pl":      "Polski",
	"pt":      "Português",
	"pt-BR":   "Português (Brasil)",
	"ro":      "Română",
	"ru":      "Русский",
	"sr":      "Српски",
	"es":      "Español",
	"es-ES":   "Español (España)",
	"es-419":  "Español (Latinoamérica)",
	"su":      "Basa Sunda",
	"th":      "ไทย",
	"tr":      "Türkçe",
	"uk":      "Українська",
	"vi":      "Tiếng Việt",
	"zh-Hant": "中文（繁體）",
	"zh-Hans": "中文（简体）",
	"hr":      "Hrvatski",
	"bn":      "বাংলা",
	"sk":      "Slovenčina",
	"it":      "Italiano",
	"my":      "မြန်မာ",
	"iw":      "עברית",
	// "tk":      "土庫曼文",
	// "da":      "丹麥文",
	// "eu":      "巴斯克文",
	// "mi":      "毛利文",
	// "eo":      "世界文",
	// "gl":      "加利西亞文",
	// "ca":      "加泰蘭文",
	// "gu":      "古吉拉特文",
	// "sw":      "史瓦希里文",
	// "ne":      "尼泊爾文",
	// "ny":      "尼揚賈文",
	// "be":      "白俄羅斯文",
	// "lt":      "立陶宛文",
	// "ig":      "伊布文",
	// "is":      "冰島文",
	// "ky":      "吉爾吉斯文",
	// "fy":      "西弗里西亞文",
	// "kn":      "坎那達文",
	// "el":      "希臘文",
	// "hy":      "亞美尼亞文",
	// "ta":      "坦米爾文",
	// "hmn":     "孟文",
	// "la":      "拉丁文",
	// "lv":      "拉脫維亞文",
	// "bs":      "波士尼亞文",
	// "fa":      "波斯文",
	// "fi":      "芬蘭文",
	// "am":      "阿姆哈拉文",
	// "sq":      "阿爾巴尼亞文",
	// "bg":      "保加利亞文",
	// "sd":      "信德文",
	// "af":      "南非荷蘭文",
	// "cy":      "威爾斯文",
	// "co":      "科西嘉文",
	// "xh":      "科薩文",
	// "yo":      "約魯巴文",
	// "haw":     "夏威夷文",
	// "ku":      "庫德文",
	// "no":      "挪威文",
	// "pa":      "旁遮普文",
	// "te":      "泰盧固文",
	// "ht":      "海地文",
	// "uz":      "烏茲別克文",
	// "ur":      "烏都文",
	// "zu":      "祖魯文",
	// "so":      "索馬利文",
	// "ms":      "馬來文",
	// "ml":      "馬來亞拉姆文",
	// "mr":      "馬拉地文",
	// "mg":      "馬達加斯加文",
	// "mt":      "馬爾他文",
	// "ceb":     "宿霧文",
	// "sn":      "紹納文",
	// "ka":      "喬治亞文",
	// "sl":      "斯洛維尼亞文",
	// "ps":      "普什圖文",
	// "tg":      "塔吉克文",
	// "st":      "塞索托文",
	// "yi":      "意第緒文",
	// "et":      "愛沙尼亞文",
	// "ga":      "愛爾蘭文",
	// "sv":      "瑞典文",
	// "si":      "僧伽羅文",
	// "ug":      "維吾爾文",
	// "mn":      "蒙古文",
	// "ha":      "豪撒文",
	// "lo":      "寮文",
	// "or":      "歐迪亞文",
	// "rw":      "盧安達文",
	// "lb":      "盧森堡文",
	// "sm":      "薩摩亞文",
	// "gd":      "蘇格蘭蓋爾文",
	// "tt":      "韃靼文",
}

type youtube struct{}

var (
	ErrTranscriptDisabled          = errors.New("Subtitles are disabled in this video")
	ErrVideoUnavailable            = errors.New("The video is no longer available")
	ErrLanguageDownloadURLNotFound = errors.New("Caption download url not found; perhaps language code is incorrect")
)

type youtubeDownloader struct {
	URL       string
	VideoID   string
	Languages map[string]string
	Caption   captionRenderer
}

type captionRenderer struct {
	PlayerCaptionsRenderer struct {
		BaseURL    string `json:"baseUrl"`
		Visibility string `json:"visibility"`
	} `json:"playerCaptionsRenderer"`
	PlayerCaptionsTracklistRenderer struct {
		CaptionTracks []struct {
			BaseURL string `json:"baseUrl"`
			Name    struct {
				SimpleText string `json:"simpleText"`
			} `json:"name"`
			VssID          string `json:"vssId"`
			LanguageCode   string `json:"languageCode"`
			IsTranslatable bool   `json:"isTranslatable"`
			Kind           string `json:"kind,omitempty"`
			Rtl            bool   `json:"rtl,omitempty"`
		} `json:"captionTracks"`
		AudioTracks []struct {
			CaptionTrackIndices      []int  `json:"captionTrackIndices"`
			DefaultCaptionTrackIndex int    `json:"defaultCaptionTrackIndex"`
			Visibility               string `json:"visibility"`
			HasDefaultTrack          bool   `json:"hasDefaultTrack"`
		} `json:"audioTracks"`
		TranslationLanguages []struct {
			LanguageCode string `json:"languageCode"`
			LanguageName struct {
				SimpleText string `json:"simpleText"`
			} `json:"languageName"`
		} `json:"translationLanguages"`
		DefaultAudioTrackIndex int `json:"defaultAudioTrackIndex"`
	} `json:"playerCaptionsTracklistRenderer"`
}

func newYoutube() *youtube {
	return &youtube{}
}

func newYoutubeDownloader(URL string) (*youtubeDownloader, error) {
	y := new(youtubeDownloader)
	y.URL = URL
	var err error
	y.VideoID, err = ParseVideoID(URL)
	if err != nil {
		return nil, err
	}
	if err := y.Fetch(); err != nil {
		return nil, err
	}
	return y, nil
}

func (y *youtube) ListLanguages(url string) (map[string]string, error) {
	yd, err := newYoutubeDownloader(url)
	if err != nil {
		return nil, err
	}
	return yd.ListLanguages(), nil
}

func (y *youtube) GetLines(url, languageCode string) ([]Line, error) {
	yd, err := newYoutubeDownloader(url)
	if err != nil {
		return nil, err
	}
	return yd.getLyrics(languageCode)
}

func (y *youtubeDownloader) Fetch() error {
	resp, err := http.Get(fmt.Sprintf(watchURL, y.VideoID))
	if err != nil {
		return err
	}
	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	dat = bytes.ReplaceAll(dat, []byte("\\u0026"), []byte{'&'})
	dat = bytes.Trim(dat, "\\")
	sp := bytes.Split(dat, []byte("\"captions\":"))
	if len(sp) <= 1 {
		if bytes.Contains(dat, []byte("\"playabilityStatus\":")) {
			return ErrTranscriptDisabled
		}
		return ErrVideoUnavailable
	}
	var result captionRenderer
	json.Unmarshal(bytes.Trim(bytes.Split(sp[1], []byte(",\"videoDetails"))[0], "\n"), &result)
	y.Caption = result
	return nil
}

func (y *youtubeDownloader) ListLanguages() map[string]string {
	languages := make(map[string]string)
	for _, language := range y.Caption.PlayerCaptionsTracklistRenderer.CaptionTracks {
		languages[language.LanguageCode] = language.BaseURL
	}
	return languages
}

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

func (y *youtubeDownloader) getLyrics(language string) ([]Line, error) {
	doc, err := y.download(language)
	if err != nil {
		return nil, err
	}
	var lyrics []Line
	idx := 1
	for _, s := range doc.SelectElements("text") {
		var l Line
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
		l.Text = strings.ReplaceAll(s.Text(), "&#39;", "'")
		l.Index = idx
		idx++
		lyrics = append(lyrics, l)
	}
	return lyrics, nil
}

func (y *youtubeDownloader) download(languageCode string) (*etree.Element, error) {
	downloadURL := ""
	for _, language := range y.Caption.PlayerCaptionsTracklistRenderer.CaptionTracks {
		if languageCode == language.LanguageCode {
			downloadURL = language.BaseURL
			break
		}
	}
	if downloadURL == "" {
		return nil, ErrLanguageDownloadURLNotFound
	}

	resp, err := http.Get(downloadURL)
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
