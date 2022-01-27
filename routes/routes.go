package routes

import (
	"go-rest-api/controllers"
	"go-rest-api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personas", controllers.AllPersonas).Methods("Get")
	r.HandleFunc("/api/personas/{id}", controllers.GetPersonaById).Methods("Get")
	r.HandleFunc("/api/personas", controllers.CreatePersona).Methods("Post")
	r.HandleFunc("/api/personas/{id}", controllers.DeletePersona).Methods("Delete")
	r.HandleFunc("/api/personas/{id}", controllers.EditPersona).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
