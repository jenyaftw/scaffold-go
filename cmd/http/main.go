package main

import (
	"fmt"

	"github.com/jenyaftw/scaffold-go/internal/adapters/config"
	"github.com/jenyaftw/scaffold-go/internal/adapters/delivery/http"
	"github.com/jenyaftw/scaffold-go/internal/adapters/delivery/http/handlers"
	"github.com/jenyaftw/scaffold-go/internal/adapters/storage/postgres"
	"github.com/jenyaftw/scaffold-go/internal/adapters/storage/postgres/repos"
	"github.com/jenyaftw/scaffold-go/internal/core/services"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	db, err := postgres.InitDb(cfg.Db)
	if err != nil {
		panic(err)
	}

	userRepo := repos.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	authService := services.NewAuthService(cfg.Jwt, userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	protectedHandler := handlers.NewProtectedHandler()

	r := http.NewRouter(userHandler, authHandler, protectedHandler)

	fmt.Printf("Listening on http://%s:%d\n", cfg.Http.Host, cfg.Http.Port)
	if err := r.ListenAndServe(cfg.Http.Host, cfg.Http.Port); err != nil {
		panic(err)
	}
}
