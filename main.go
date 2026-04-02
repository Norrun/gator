package main

import "github.com/Norrun/gator/internal/config"

func main() {
	conf := config.Read()
	conf.SetUser("Trym")
	conf = config.Read()
	println(conf.CurrentUserName)
	println(conf.DbURL)
}
