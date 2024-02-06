package main

import (
	"fmt"

	"github.com/TNXG/ProcessReporterWingo/packages/getconf"
)

var conf = getconf.ReadFile()
var endpoint = conf.Server.Endpoint
var token = conf.Server.Token

func main() {
	conf := getconf.ReadFile()
	url := conf.Server.Endpoint
	fmt.Println("Endpoint: ", url)
}
