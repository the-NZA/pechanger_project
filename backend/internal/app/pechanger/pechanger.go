package pechanger

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-pkgz/lgr"
	"github.com/rs/cors"
)

// Author: Kozlov Roman
type Server struct {
	config *Config
	logger *lgr.Logger
	router *chi.Mux
}

// Author: Kozlov Roman
// newLogger configure logger in DEBUG or PRODUCTION mode
// Possible log levels TRACE, DEBUG, INFO, WARN, ERROR, PANIC and FATAL
func newLogger(dbg bool) *lgr.Logger {
	if dbg {
		return lgr.New(lgr.Msec, lgr.Debug, lgr.CallerFile, lgr.CallerFunc, lgr.LevelBraces)
	}

	return lgr.New(lgr.Msec, lgr.LevelBraces)
}

// Author: Kozlov Roman
// New creates fresh instance of Server struct
func New(config *Config) *Server {
	s := &Server{
		config: config,
		router: chi.NewRouter(),
		logger: newLogger(config.LogDebug),
	}

	return s
}

// Author: Kozlov Roman
// Return configured CORS middleware
func newCorsMiddleware() func(http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "OPTIONS", "HEAD"},
		AllowedHeaders: []string{"Authorization", "X-Auth-Key", "X-Auth-Secret", "Content-Type", "X-Requested-With"},
		// Debug:          true,
	}).Handler
}

// Author: Kozlov Roman
// configureRouter setup all routes for application's router
func (s *Server) configureRouter() {
	s.router.Use(newCorsMiddleware())

	s.router.Route("/api", func(r chi.Router) {
		r.Get("/", s.handleGetApi())
		r.Get("/exchange", s.handleGetExchange())
		r.Get("/exchange_at", s.handleGetExchangeAt())
		r.Get("/exchange_between", s.handleGetExchangeBetween())
		r.Post("/email_share", s.handlePostEmailShare())
	})
}

// Author: Kozlov Roman
// Start configure's router and other internal things
// and run HTTP server
func (s *Server) Start() error {
	s.configureRouter()

	// Another stuff

	s.logger.Logf("INFO Server is starting at port %v...\n", s.config.BindAddr)

	return http.ListenAndServe(s.config.BindAddr, s.router)
}
