package main

import (
	"log"
	"net/http"
	"os"
	"teas/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dataPath := os.Getenv("TEA_DATA_PATH")
	if dataPath == "" {
		dataPath = "data/tes.json"
	}

	if err := handlers.LoadTeaData(dataPath); err != nil {
		log.Fatalf("Error al cargar los datos de té: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", handlers.TeaAllData)
	r.Get("/category", handlers.ListCategories)
	r.Get("/category/{id}", handlers.GetCategoryByID)

	log.Printf("Servidor funcionando en el puerto %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
