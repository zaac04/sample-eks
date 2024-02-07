package main

import (
	models "backend/Models"
	"backend/db"
	"backend/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("config.env")
	if err != nil {
		fmt.Println(err)
	}

}

func main() {

	db.Pg.Connect()
	db.Pg.Client.AutoMigrate(models.Todo{})

	router := chi.NewRouter()
	router.Use(CORS)
	router.Group(func(r chi.Router) {
		r.Post("/add", handlers.AddTodo)
		r.Get("/get", handlers.GetTodos)
	})
	http.ListenAndServe(":5000", router)
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})

}
