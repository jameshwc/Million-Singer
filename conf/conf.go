package conf

import (
	"log"
	"os"
	"strconv"
)

type Database struct {
	User        string
	Password    string
	Host        string
	Type        string
	Name        string
	TablePrefix string
}

type Server struct {
	RunMode   string
	HttpPort  int
	JwtSecret string
}

var DBconfig = &Database{}
var ServerConfig = &Server{}

func Setup() {
	DBconfig = &Database{
		User:        os.Getenv("db_user"),
		Password:    os.Getenv("db_passwd"),
		Host:        os.Getenv("db_host"),
		Type:        os.Getenv("db_type"),
		Name:        os.Getenv("db_name"),
		TablePrefix: os.Getenv("db_table_prefix"),
	}
	port, err := strconv.Atoi(os.Getenv("server_port"))
	if err != nil {
		log.Fatal("read config: server_port is not a number")
	}
	ServerConfig = &Server{
		RunMode:   os.Getenv("server_runmode"),
		HttpPort:  port,
		JwtSecret: os.Getenv("jwt_secret"),
	}
}
