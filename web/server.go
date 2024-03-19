package web

import (
	"log"
	"otp-webapp/configs"
	"otp-webapp/handlers/api"
	"otp-webapp/handlers/page"
	"otp-webapp/models/mailer"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html/v2"
	"github.com/rs/zerolog"
)

type WebServer struct {
	AppName string
	Cfg     configs.WebConf
	Log     *zerolog.Logger
}

func NewWebServer(cfg configs.WebConf, lg *zerolog.Logger) *WebServer {
	return &WebServer{
		AppName: "OTP WebApp",
		Cfg:     cfg,
		Log:     lg,
	}
}

func (w *WebServer) Start() {
	mlr := mailer.NewMailer(w.Log)
	rd := configs.NewRedisClient()

	_, err := rd.Ping().Result()
	if err != nil {
		w.Log.Error().Str("error", err.Error()).Msg("failed to connect redis")
	}

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		AppName: w.AppName,
		Views: engine,
		Prefork: false,
	})

	app.Use(requestid.New())
	app.Use(recover.New())

	app.Static("/", "./views/assets")
	app.Static("/", "./views/static")

	apiHandler := &api.Handler{
		Log: w.Log,
		Mailer: mlr,
		Redis: rd,
	}
	pageHandler := &page.Handler{
		Log: w.Log,
	}

	ApiRoutes(app, apiHandler)
	WebViews(app, pageHandler)

	log.Fatal(app.Listen(w.Cfg.ListenAddr()))
}