package api

import (
	"github.com/go-chi/chi"
	"job-status-tracker/pkg/database"
)

// Routes defines the API routes
func Routes(db *database.DB) chi.Router {
	router := chi.NewRouter()

	// Services routes
	router.Route("/services", func(r chi.Router) {
		r.Post("/", CreateServiceHandler(db))
		r.Get("/", GetServicesHandler(db))
		r.Route("/{service_id}", func(r chi.Router) {
			r.Get("/", GetServiceHandler(db))
			r.Put("/", UpdateServiceHandler(db))
			r.Delete("/", DeleteServiceHandler(db))
		})
	})

	// Jobs routes
	router.Route("/jobs", func(r chi.Router) {
		r.Post("/", CreateJobHandler(db))
		r.Get("/", GetJobsHandler(db))
		r.Route("/{job_id}", func(r chi.Router) {
			r.Get("/", GetJobHandler(db))
			r.Put("/", UpdateJobHandler(db))
			r.Delete("/", DeleteJobHandler(db))
		})
	})

	// Users routes
	router.Route("/users", func(r chi.Router) {
		r.Post("/", CreateUserHandler(db))
		r.Get("/", GetUsersHandler(db))
		r.Route("/{user_id}", func(r chi.Router) {
			r.Get("/", GetUserHandler(db))
			r.Put("/", UpdateUserHandler(db))
			r.Delete("/", DeleteUserHandler(db))
		})
	})

	return router
}
