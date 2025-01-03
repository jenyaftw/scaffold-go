package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/jenyaftw/scaffold-go/internal/adapters/delivery/http/handlers"
	"github.com/jenyaftw/scaffold-go/internal/adapters/delivery/http/middleware"
)

type Router struct {
	router *chi.Mux
}

func NewRouter(
	userHandler handlers.UserHandler,
	authHandler handlers.AuthHandler,
	protectedHandler handlers.ProtectedHandler,
) Router {
	r := chi.NewRouter()

	r.Use(chiMiddleware.Logger)

	userRouter := chi.NewRouter()
	userRouter.Post("/", userHandler.Register)
	userRouter.Post("/{id}/verify", userHandler.ResendVerify)
	userRouter.Post("/{id}/verify/{code}", userHandler.Verify)

	protectedUsers := chi.NewRouter()
	protectedUsers.Use(middleware.AuthMiddleware)
	protectedUsers.Get(("/me"), userHandler.Me)

	userRouter.Mount("/", protectedUsers)

	authRouter := chi.NewRouter()
	authRouter.Post("/login", authHandler.Login)

	protectedRouter := chi.NewRouter()
	protectedRouter.Use(middleware.AuthMiddleware)
	protectedRouter.Get("/protected", protectedHandler.TestRoute)

	r.Mount("/users", userRouter)
	r.Mount("/auth", authRouter)
	r.Mount("/", protectedRouter)

	return Router{
		router: r,
	}
}

func (r Router) ListenAndServe(host string, port int) error {
	return http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), r.router)
}
