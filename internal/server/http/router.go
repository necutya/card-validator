package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	swagMiddleware "github.com/go-openapi/runtime/middleware"
	"github.com/necutya/card_validator/internal/server/http/controllers"
)

type Router struct {
	chi.Router
	controller controllers.Controller
}

func NewRouter(controller controllers.Controller) Router {
	r := Router{
		Router:     chi.NewRouter(),
		controller: controller,
	}

	r.initMiddlewares()
	r.initRoutes()

	return r
}

func (r *Router) initMiddlewares() {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
}

func (r *Router) initRoutes() {
	sw := swagMiddleware.SwaggerUI(swagMiddleware.SwaggerUIOpts{SpecURL: "/swagger.yml"}, nil)
	r.Get("/docs", sw.ServeHTTP)
	r.Get("/swagger.yml", r.controller.Swagger)

	r.Get("/healthcheck", r.controller.HealthCheck)
	r.Put("/card/validate", r.controller.ValidateCard)
}
