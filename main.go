package main

import (
	"fmt"
	"path/filepath"

	"github.com/jameshwc/Million-Singer/conf"
)

func init() {
	conf.Setup(filepath.Join("conf", "conf.toml"))
}

func main() {
	fmt.Println(conf.Conf.DB.Host)
}
