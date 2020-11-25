package game

func (srv *Service) GetSupportedLanguages() []string {
	return []string{
		"en",
		"zh-TW", // TODO: include zh-Hant & zh-TW
		"zh-CN", // TODO: include zh-Hans & zh-CN
		"fr",
		"ja",
		"ko",
		"es",
	}
}

func (srv *Service) GetGenres() []string {
	return []string{
		"Hip-hop",
		"Pop",
		"Blues",
		"Cover",
		"Jazz",
		"Country",
		"Edm",
		"Alternative",
		"Anime",
		"Dance",
		"Electronic",
		"Indie-pop",
		"Inspirational",
		"J-pop",
		"K-pop",
		"Latin",
		"Metal",
		"New age",
		"Opera",
		"R&B",
		"Reggae",
		"Rock",
		"Soundtrack",
		"World",
	}
}
