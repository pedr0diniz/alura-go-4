package controllers

import (
	"encoding/json"
	"fmt"
	"go-rest-api/database"
	"go-rest-api/models"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersonas(w http.ResponseWriter, r *http.Request) {

	var p []models.Persona
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetPersonaById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var persona models.Persona

	database.DB.First(&persona, id)
	json.NewEncoder(w).Encode(persona)
}

func CreatePersona(w http.ResponseWriter, r *http.Request) {
	var newPersona models.Persona
	json.NewDecoder(r.Body).Decode(&newPersona)
	database.DB.Create(&newPersona)
	json.NewEncoder(w).Encode(newPersona)
}

func DeletePersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var persona models.Persona

	database.DB.Delete(&persona, id)
	json.NewEncoder(w).Encode(persona)
}

func EditPersona(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var persona models.Persona

	database.DB.First(&persona, id)
	json.NewDecoder(r.Body).Decode(&persona)
	database.DB.Save(&persona)
	json.NewEncoder(w).Encode(persona)

}
