package subtitle

import (
	"reflect"
	"testing"
)

var (
	sampleLanguages = map[string]string{
		"Deutsch":            "de",
		"English":            "en",
		"Español":            "es",
		"Español (España)":   "es-ES",
		"Français":           "fr",
		"Hrvatski":           "hr",
		"Italiano":           "it",
		"Magyar":             "hu",
		"Polski":             "pl",
		"Português (Brasil)": "pt-BR",
		"Slovenčina":         "sk",
		"Tiếng Việt":         "vi",
		"Türkçe":             "tr",
		"default":            "en",
		"Русский":            "ru",
		"Српски":             "sr",
		"Українська":         "uk",
		"עברית":              "iw",
		"العربية":            "ar",
		"فارسی":              "fa",
		"বাংলা":              "bn",
		"ไทย":                "th",
		"မြန်မာ":             "my",
		"中文（简体）":             "zh-CN",
		"中文（繁體）":             "zh-TW",
		"日本語":                "ja",
	}
	sampleAndDefault = sampleLanguages
)

func Test_ParseVideoID(t *testing.T) {
	type args struct {
		URL string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"http short url", args{URL: "http://youtu.be/5MgBikgcWnY"}, "5MgBikgcWnY", false},
		{"https short url", args{URL: "https://youtu.be/5MgBikgcWnY"}, "5MgBikgcWnY", false},
		{"http long url", args{URL: "http://www.youtube.com/watch?v=5MgBikgcWnY&feature=feed"}, "5MgBikgcWnY", false},
		{"https long url", args{URL: "http://www.youtube.com/watch?v=5MgBikgcWnY&feature=feed"}, "5MgBikgcWnY", false},
		{"no protocol url", args{URL: "www.youtube.com/watch?v=5MgBikgcWnY"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseVideoID(tt.args.URL)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseVideoID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseVideoID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewYoutubeDownloader(t *testing.T) {
	sampleAndDefault["default"] = "en"
	type args struct {
		URL string
	}
	tests := []struct {
		name    string
		args    args
		want    *youtubeDownloader
		wantErr bool
	}{
		{"http short url", args{URL: "http://youtu.be/5MgBikgcWnY"}, &youtubeDownloader{"http://youtu.be/5MgBikgcWnY", "5MgBikgcWnY", sampleAndDefault}, false},
		{"https short url", args{URL: "https://youtu.be/5MgBikgcWnY"}, &youtubeDownloader{"https://youtu.be/5MgBikgcWnY", "5MgBikgcWnY", sampleAndDefault}, false},
		{"http long url", args{URL: "http://www.youtube.com/watch?v=5MgBikgcWnY&feature=feed"}, &youtubeDownloader{"http://www.youtube.com/watch?v=5MgBikgcWnY&feature=feed", "5MgBikgcWnY", sampleAndDefault}, false},
		{"https long url", args{URL: "https://www.youtube.com/watch?v=5MgBikgcWnY&feature=feed"}, &youtubeDownloader{"https://www.youtube.com/watch?v=5MgBikgcWnY&feature=feed", "5MgBikgcWnY", sampleAndDefault}, false},
		{"no protocol url", args{URL: "www.youtube.com/watch?v=5MgBikgcWnY&feature=feed"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newYoutubeDownloader(tt.args.URL)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewyoutubeDownloader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewyoutubeDownloader() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
