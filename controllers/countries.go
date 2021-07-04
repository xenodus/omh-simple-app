package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"omh-simple-app/database"
	"omh-simple-app/models"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetCountries(res http.ResponseWriter, req *http.Request) {

	var jsonResponse models.JSONResponse

	var countries = []*models.Country{}
	results := database.DB.Find(&countries)

	if results.Error != nil {
		jsonResponse.PrintUnexpectedErrorResponse(res)
		return
	}

	// success
	jsonResponse.PrintJSONResponse(res, http.StatusOK, "success", countries)
}

func GetCountryByID(res http.ResponseWriter, req *http.Request) {

	var jsonResponse models.JSONResponse

	vars := mux.Vars(req)
	countryID, err := strconv.Atoi(vars["id"])

	if err != nil {
		jsonResponse.PrintUnexpectedErrorResponse(res)
		return
	}

	var country *models.Country
	result := database.DB.First(&country, countryID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			jsonResponse.PrintNotFoundResponse(res)
		} else {
			jsonResponse.PrintUnexpectedErrorResponse(res)
		}
		return
	}

	// success
	jsonResponse.PrintJSONResponse(res, http.StatusOK, "success", country)
}

func CreateCountry(res http.ResponseWriter, req *http.Request) {

	var jsonResponse models.JSONResponse

	// json decode
	var newCountry *models.Country
	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		jsonResponse.PrintUnexpectedErrorResponse(res)
		return
	}

	marErr := json.Unmarshal(reqBody, &newCountry)

	if marErr != nil {
		jsonResponse.PrintJSONResponse(res, http.StatusBadRequest, marErr.Error(), nil)
		return
	}

	// validation
	validateErr := newCountry.Validate()

	if validateErr != nil {
		jsonResponse.PrintJSONResponse(res, http.StatusBadRequest, validateErr.Error(), nil)
		return
	}

	result := database.DB.Create(&newCountry)

	if result.Error != nil {

		if strings.Contains(result.Error.Error(), "Duplicate entry") {
			jsonResponse.StatusCode = http.StatusConflict
		} else {
			jsonResponse.StatusCode = http.StatusInternalServerError
		}

		jsonResponse.PrintJSONResponse(res, jsonResponse.StatusCode, result.Error.Error(), nil)
		return
	}

	// success
	jsonResponse.PrintJSONResponse(res, http.StatusCreated, "success", newCountry)
}

func UpdateCountry(res http.ResponseWriter, req *http.Request) {

	var jsonResponse models.JSONResponse

	vars := mux.Vars(req)
	countryID, err := strconv.Atoi(vars["id"])

	if err != nil {
		jsonResponse.PrintUnexpectedErrorResponse(res)
		return
	}

	var country models.Country
	getResult := database.DB.First(&country, countryID)

	if getResult.Error != nil {
		if errors.Is(getResult.Error, gorm.ErrRecordNotFound) {
			jsonResponse.PrintNotFoundResponse(res)
		} else {
			jsonResponse.PrintUnexpectedErrorResponse(res)
		}
		return
	}

	// json decode
	var newCountry *models.Country
	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		jsonResponse.PrintUnexpectedErrorResponse(res)
		return
	}

	json.Unmarshal(reqBody, &newCountry)

	marErr := json.Unmarshal(reqBody, &newCountry)

	if marErr != nil {
		jsonResponse.PrintJSONResponse(res, http.StatusBadRequest, marErr.Error(), nil)
		return
	}

	// validation
	validateErr := newCountry.Validate()

	if validateErr != nil {
		jsonResponse.PrintJSONResponse(res, http.StatusBadRequest, validateErr.Error(), nil)
		return
	}

	// update values
	country.Name = newCountry.Name
	country.UpdatedAt = time.Now()

	updateResult := database.DB.Save(&country)

	if updateResult.Error != nil {

		if strings.Contains(updateResult.Error.Error(), "Duplicate entry") {
			jsonResponse.StatusCode = http.StatusConflict
		} else {
			jsonResponse.StatusCode = http.StatusInternalServerError
		}

		jsonResponse.PrintJSONResponse(res, jsonResponse.StatusCode, updateResult.Error.Error(), nil)
		return
	}

	// success
	jsonResponse.PrintJSONResponse(res, http.StatusOK, "success", country)
}

func DeleteCountry(res http.ResponseWriter, req *http.Request) {

	var jsonResponse models.JSONResponse

	vars := mux.Vars(req)
	countryID, err := strconv.Atoi(vars["id"])

	if err != nil {
		jsonResponse.PrintUnexpectedErrorResponse(res)
		return
	}

	var country models.Country
	deleteResult := database.DB.Delete(&country, countryID)

	if deleteResult.Error != nil {
		if errors.Is(deleteResult.Error, gorm.ErrRecordNotFound) {
			jsonResponse.PrintNotFoundResponse(res)
		} else {
			jsonResponse.PrintUnexpectedErrorResponse(res)
		}
		return
	}

	// success
	jsonResponse.PrintJSONResponse(res, http.StatusOK, "success", nil)
}
