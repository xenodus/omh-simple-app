package server

import (
	"omh-simple-app/controllers"

	"github.com/gorilla/mux"
)

func initRoutes(r *mux.Router) {

	// Middleware
	r.Use(checkAPIKey)

	// Country
	r.HandleFunc("/countries", controllers.GetCountries).Methods("GET").Name("GetCountry")
	r.HandleFunc("/countries/{id:[0-9]+}", controllers.GetCountryByID).Methods("GET").Name("GetCountryByID")
	r.HandleFunc("/countries", controllers.CreateCountry).Methods("POST").Name("CreateCountry")
	r.HandleFunc("/countries/{id:[0-9]+}", controllers.UpdateCountry).Methods("PUT").Name("UpdateCountryByID")
	r.HandleFunc("/countries/{id:[0-9]+}", controllers.DeleteCountry).Methods("DELETE").Name("DeleteCountry")

	// Property
	r.HandleFunc("/properties", controllers.GetProperties).Methods("GET").Name("GetProperty")
	r.HandleFunc("/properties/{id:[0-9]+}", controllers.GetPropertyByID).Methods("GET").Name("GetCountryByID")
	r.HandleFunc("/properties", controllers.CreateProperty).Methods("POST").Name("CreateProperty")
	r.HandleFunc("/properties/{id:[0-9]+}", controllers.UpdateProperty).Methods("PUT").Name("UpdatePropertyByID")
	r.HandleFunc("/properties/{id:[0-9]+}", controllers.DeleteProperty).Methods("DELETE").Name("DeleteProperty")
}
