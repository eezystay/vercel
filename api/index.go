package handler

import (
	"net/http"

	cnf "github.com/eezystay/backend/pkg/config"
	"github.com/eezystay/backend/pkg/logger"
	"github.com/eezystay/backend/pkg/server"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// This is needed to set the proper request path in `fiber.Ctx`
	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

// building the fiber application
func handler() http.HandlerFunc {
	config, err := cnf.Load()
	if err != nil {
		logger.FatalErr("Error loading config", err)
	}

	// Setup logger
	logger.Init(config.Environment, config.LogLevel, nil)

	app := server.Create(config)

	return adaptor.FiberApp(app)
}
