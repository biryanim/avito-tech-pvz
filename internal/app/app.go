package app

import (
	"context"
	"github.com/biryanim/avito-tech-pvz/internal/api/middleware"
	"github.com/biryanim/avito-tech-pvz/internal/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {

	err := a.runHTTPServer()
	if err != nil {
		log.Fatalf("failed to run http server: %v", err)
	}

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load("example.env")
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	router := gin.Default()

	authMiddleware := middleware.AuthMiddleware(a.serviceProvider.AuthService(ctx))
	api := router.Group("/api")
	{
		api.POST("/dummyLogin", a.serviceProvider.AuthImpl(ctx).DummyLogin)
		api.POST("/register", a.serviceProvider.AuthImpl(ctx).Register)
		api.POST("/login", a.serviceProvider.AuthImpl(ctx).Login)

		protected := api.Group("/")
		protected.Use(authMiddleware)
	}

	a.httpServer = &http.Server{
		Addr:    a.serviceProvider.HTTPConfig().Address(),
		Handler: router,
	}

	return nil
}

func (a *App) runHTTPServer() error {
	log.Printf("HTTP server is running on %s", a.serviceProvider.HTTPConfig().Address())

	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}
