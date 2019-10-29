package main

import (
	"github.com/blocktop/mp-common/server"
	"github.com/blocktop/mp-common/server/middleware"
	"github.com/blocktop/mp-web/config"
	"github.com/go-chi/chi"
	chim "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	cfg := config.GetConfig()

	r := chi.NewMux()
	r.Use(middleware.HealthMiddleware)
	r.Use(chim.Timeout(time.Duration(cfg.HTTPServerRequestTimeout) * time.Second))

	cors := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
	})
	r.Use(cors.Handler)

	setRoutes(r)

	server.RunHTTPServer(r)

	os.Exit(0)
}

var stellarTOML string

func setRoutes(r *chi.Mux) {
	r.Get("/health", middleware.HealthHandler)

	r.Get("/.well-known/stellar.toml", handleGetStellarTOML)
	r.Handle("/", http.FileServer(Dir(false, "/assets")))
}

func handleGetStellarTOML(w http.ResponseWriter, r *http.Request) {
	if len(stellarTOML) == 0 {
		cfg := config.GetConfig()
		toml, err := ioutil.ReadFile(cfg.StellarTOMLPath)
		if err != nil {
			server.ResponseError(w, http.StatusInternalServerError, server.NOFILE, err)
			return
		}

		server.ResponseText(w, toml)
	}
}
