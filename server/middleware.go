package server

import (
	"net/http"
	"omh-simple-app/models"
	"os"
	"strings"
)

func checkAPIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

		if req.Header.Get("Content-type") != "application/json" {
			var jsonResponse models.JSONResponse
			jsonResponse.PrintJSONResponse(res, http.StatusUnsupportedMediaType, "wrong content-type", nil)
			return
		}

		apikey := strings.TrimSpace(req.FormValue("apikey"))

		if apikey != os.Getenv("API_KEY") {
			var jsonResponse models.JSONResponse
			jsonResponse.PrintUnauthorizeResponse(res)
			return
		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(res, req)
	})
}
