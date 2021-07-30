package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/xion6/golang-echo/pkg/config"
	"github.com/xion6/golang-echo/pkg/data"
	"github.com/xion6/golang-echo/pkg/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	server  *echo.Echo
	userSvc services.IUserService
	cfg     *config.Settings
}

func New(cfg *config.Settings, client *mongo.Client) *App {
	server := echo.New()

	// middleware
	server.Use(middleware.Recover())
	server.Use(middleware.RequestID())

	// providers
	userProvider := data.NewUserProvider(cfg, client)

	// services
	userSvc := services.NewUserService(cfg, userProvider)

	return &App{
		server:  server,
		userSvc: userSvc,
		cfg:     cfg,
	}
}

func (a App) ConfigureRoutes() {
	a.server.GET("/v1/public/healthy", a.HealthCheck)
	a.server.POST("/v1/public/account/register", a.Register)
	a.server.POST("/v1/public/account/login", a.Login)
}

func (a App) Start() {
	a.ConfigureRoutes()
	a.server.Start(":5000")
}
