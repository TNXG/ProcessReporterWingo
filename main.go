package main

import (
	"fmt"

	Core "github.com/TNXG/ProcessReporterWingo/core"
)

var conf = Core.ReadConf()
var endpoint = conf.Server.Endpoint
var token = conf.Server.Token

func main() {
	processName, windowTitle := Core.GetWindowInfo()
	fmt.Println(processName, ":::", windowTitle)
}
