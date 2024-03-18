package main

import (
	"otp-webapp/configs"
	"otp-webapp/web"
)

func main() {
	configs.LoadEnv()
	zlog := configs.InitLogger()

	ws := web.NewWebServer(configs.EnvWebConf(), zlog)
	ws.Start()
}