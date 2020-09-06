package main

import (
	"fmt"
	"path/filepath"

	"github.com/jameshwc/Million-Singer/conf"
	"github.com/jameshwc/Million-Singer/model"
)

func init() {
	conf.Setup(filepath.Join("conf", "conf.toml"))
	model.Setup()
}

func main() {
	fmt.Println(conf.Conf.DB.Host)
}
