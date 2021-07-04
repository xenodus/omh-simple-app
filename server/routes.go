package server

import (
	"omh-simple-app/controllers"

	"github.com/gorilla/mux"
)

func initRoutes(r *mux.Router) {

	// Middleware
	r.Use(checkAPIKey)

	// v1.0
	s := r.PathPrefix("/api/v1").Subrouter()

	// Country
	s.HandleFunc("/countries", controllers.GetCountries).Methods("GET").Name("GetCountry")
	s.HandleFunc("/countries/{id:[0-9]+}", controllers.GetCountryByID).Methods("GET").Name("GetCountryByID")
	s.HandleFunc("/countries", controllers.CreateCountry).Methods("POST").Name("CreateCountry")
	s.HandleFunc("/countries/{id:[0-9]+}", controllers.UpdateCountry).Methods("PUT").Name("UpdateCountryByID")
	s.HandleFunc("/countries/{id:[0-9]+}", controllers.DeleteCountry).Methods("DELETE").Name("DeleteCountry")

	// Property
	s.HandleFunc("/properties", controllers.GetProperties).Methods("GET").Name("GetProperty")
	s.HandleFunc("/properties/{id:[0-9]+}", controllers.GetPropertyByID).Methods("GET").Name("GetCountryByID")
	s.HandleFunc("/properties", controllers.CreateProperty).Methods("POST").Name("CreateProperty")
	s.HandleFunc("/properties/{id:[0-9]+}", controllers.UpdateProperty).Methods("PUT").Name("UpdatePropertyByID")
	s.HandleFunc("/properties/{id:[0-9]+}", controllers.DeleteProperty).Methods("DELETE").Name("DeleteProperty")
}
