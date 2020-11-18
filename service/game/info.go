package game

func (srv *Service) GetSupportedLanguages() []string {
	return []string{
		"en",
		"zh-tw",
		"zh-cn",
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
