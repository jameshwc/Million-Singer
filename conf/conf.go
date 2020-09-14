package conf

import (
	"log"
	"os"
	"strconv"
	"time"
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

type Redis struct {
	Host        string
	Password    string
	MinIdleConn int
	IdleTimeout time.Duration
}

var DBconfig = &Database{}
var ServerConfig = &Server{}
var RedisConfig = &Redis{}

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
		JwtSecret: os.Getenv("server_jwt_secret"),
	}
	minIdleConn, err := strconv.Atoi(os.Getenv("redis_min_idle_conn"))
	if err != nil {
		log.Fatal("read config: redis_max_idle is not a number")
	}
	idleTimeout, err := strconv.Atoi(os.Getenv("redis_idle_timeout"))
	if err != nil {
		log.Fatal("read config: redis_idle_timeout is not a number")
	}
	RedisConfig = &Redis{
		Host:        os.Getenv("redis_host"),
		Password:    os.Getenv("redis_password"),
		MinIdleConn: minIdleConn,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
	}
}
