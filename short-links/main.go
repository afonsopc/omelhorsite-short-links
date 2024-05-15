package main

import (
	"fmt"
	"net/http"

	linksRoutes "github.com/afonsopc/omelhorsite-short-links/routes/links"
	rootRoutes "github.com/afonsopc/omelhorsite-short-links/routes/root"
	"github.com/afonsopc/omelhorsite-short-links/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

/*

ROUTES:

GET / - SERVICE MOTO

GET /:linkID - REDIRECT TO LINK

GET /link - PARAMS: id - GET LINK

GET /links - GET ALL LINKS

DELETE /link - PARAMS: id - DELETE LINK

POST /link - PARAMS: link - CREATE LINK

*/

func main() {
	godotenv.Load()

	utils.CheckAllConfigurationVariables()

	utils.DatabaseInit()

	apiConfigurations := utils.GetApiConfiguration()

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		MaxAge:         300,
	}))

	router.Get("/", rootRoutes.RootHandler)
	router.Get("/{linkID}", rootRoutes.RootLinkHandler)
	router.Get("/link", linksRoutes.GetLinkHandler)
	router.Get("/links", linksRoutes.GetAllLinksHandler)
	router.Delete("/link", linksRoutes.DeleteLinkHandler)
	router.Post("/link", linksRoutes.CreateLinkHandler)

	fmt.Printf("Listening at %s\n", apiConfigurations.Endpoint)
	http.ListenAndServe(apiConfigurations.Endpoint, router)
}
