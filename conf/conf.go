package conf

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Database struct {
	User     string
	Password string
	Host     string
}

type Server struct {
	RunMode  string
	HttpPort int
}

type Config struct {
	DB     Database `toml:"database"`
	Server Server
}

var Conf = &Config{}

func Setup(configPath string) {
	if _, err := toml.DecodeFile(configPath, Conf); err != nil {
		log.Fatalf("Error: read config %s", configPath)
	}
}
