package main

import (
	"fmt"

	"github.com/jameshwc/Million-Singer/conf"
)

func init() {
	conf.Setup("conf.toml")
}

func main() {
	fmt.Println(conf.Conf.DB.Host)
}
