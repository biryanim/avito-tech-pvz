package main

import (
	"context"
	"github.com/biryanim/avito-tech-pvz/internal/api/auth"
	"github.com/biryanim/avito-tech-pvz/internal/config"
	"github.com/biryanim/avito-tech-pvz/internal/repository/user"
	authServ "github.com/biryanim/avito-tech-pvz/internal/service/auth"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func main() {
	err := config.Load("./example.env")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatal(err)
	}

	pool, err := pgxpool.New(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatal(err)
	}

	jwtCfg, err := config.NewJWTConfig()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := user.NewRepository(pool)
	userService := authServ.NewService(userRepo, jwtCfg)
	authImpl := auth.NewImplementation(userService)

	router := gin.Default()

	router.POST("/login", authImpl.Login)
	router.POST("/register", authImpl.Register)

	serverAddr := ":8080"
	log.Printf("server listening on %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatal(err)
	}
}
