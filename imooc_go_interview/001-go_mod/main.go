package main

import "github.com/beego/beego/v2/core/logs"

func main() {
	log := logs.NewLogger(1000)
	log.SetLogger("console", "")
	log.Warn("go_mod")
}
