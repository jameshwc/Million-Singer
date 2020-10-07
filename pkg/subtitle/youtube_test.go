package subtitle

import (
	"reflect"
	"testing"
)

var (
	sampleLanguages = map[string]string{
		"ar":    "العربية",
		"bn":    "বাংলা",
		"de":    "Deutsch",
		"en":    "English",
		"es":    "Español",
		"es-ES": "Español (España)",
		"fa":    "فارسی",
		"fr":    "Français",
		"hr":    "Hrvatski",
		"hu":    "Magyar",
		"it":    "Italiano",
		"iw":    "עברית",
		"ja":    "日本語",
		"my":    "မြန်မာ",
		"pl":    "Polski",
		"pt-BR": "Português (Brasil)",
		"ru":    "Русский",
		"sk":    "Slovenčina",
		"sr":    "Српски",
		"th":    "ไทย",
		"tr":    "Türkçe",
		"uk":    "Українська",
		"vi":    "Tiếng Việt",
		"zh-CN": "中文（简体）",
		"zh-TW": "中文（繁體）",
	}
)

func Test_parseVideoID(t *testing.T) {
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
			got, err := parseVideoID(tt.args.URL)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseVideoID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseVideoID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewYoutubeDownloader(t *testing.T) {
	type args struct {
		URL string
	}
	tests := []struct {
		name    string
		args    args
		want    *youtubeDownloader
		wantErr bool
	}{
		{"http short url", args{URL: "http://youtu.be/5MgBikgcWnY"}, &youtubeDownloader{"http://youtu.be/5MgBikgcWnY", "5MgBikgcWnY", sampleLanguages}, false},
		{"https short url", args{URL: "https://youtu.be/5MgBikgcWnY"}, &youtubeDownloader{"https://youtu.be/5MgBikgcWnY", "5MgBikgcWnY", sampleLanguages}, false},
		{"http long url", args{URL: "http://www.youtube.com/watch?v=5MgBikgcWnY&feature=feed"}, &youtubeDownloader{"http://www.youtube.com/watch?v=5MgBikgcWnY&feature=feed", "5MgBikgcWnY", sampleLanguages}, false},
		{"https long url", args{URL: "https://www.youtube.com/watch?v=5MgBikgcWnY&feature=feed"}, &youtubeDownloader{"https://www.youtube.com/watch?v=5MgBikgcWnY&feature=feed", "5MgBikgcWnY", sampleLanguages}, false},
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

func TestyoutubeDownloader_getAvailableLanguages(t *testing.T) {
	type fields struct {
		URL       string
		VideoID   string
		Languages map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"http short url", fields{"http://youtu.be/5MgBikgcWnY", "5MgBikgcWnY", sampleLanguages}, false},
		{"https short url", fields{"https://youtu.be/5MgBikgcWnY", "5MgBikgcWnY", sampleLanguages}, false},
		{"http long url", fields{"http://www.youtube.com/watch?v=5MgBikgcWnY&feature=feed", "5MgBikgcWnY", sampleLanguages}, false},
		{"https long url", fields{"https://www.youtube.com/watch?v=5MgBikgcWnY&feature=feed", "5MgBikgcWnY", sampleLanguages}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &youtubeDownloader{
				URL:       tt.fields.URL,
				VideoID:   tt.fields.VideoID,
				Languages: tt.fields.Languages,
			}
			if err := y.getAvailableLanguages(); (err != nil) != tt.wantErr {
				t.Errorf("youtubeDownloader.getAvailableLanguages() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
