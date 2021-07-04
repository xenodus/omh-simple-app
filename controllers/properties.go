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

func GetProperties(res http.ResponseWriter, req *http.Request) {

	var jsonResponse models.JSONResponse

	var properties = []*models.Property{}
	results := database.DB.Preload("Country").Find(&properties)

	if results.Error != nil {
		jsonResponse.PrintUnexpectedErrorResponse(res)
		return
	}

	// success
	jsonResponse.PrintJSONResponse(res, http.StatusOK, "success", properties)
}

func GetPropertyByID(res http.ResponseWriter, req *http.Request) {

	var jsonResponse models.JSONResponse

	vars := mux.Vars(req)
	propertyID, err := strconv.Atoi(vars["id"])

	if err != nil {
		jsonResponse.PrintUnexpectedErrorResponse(res)
		return
	}

	var property *models.Property
	result := database.DB.First(&property, propertyID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			jsonResponse.PrintNotFoundResponse(res)
		} else {
			jsonResponse.PrintUnexpectedErrorResponse(res)
		}
		return
	}

	// success
	jsonResponse.PrintJSONResponse(res, http.StatusOK, "success", property)
}

func CreateProperty(res http.ResponseWriter, req *http.Request) {

	var jsonResponse models.JSONResponse

	// json decode
	var newProperty *models.Property
	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		jsonResponse.PrintUnexpectedErrorResponse(res)
		return
	}

	json.Unmarshal(reqBody, &newProperty)

	marErr := json.Unmarshal(reqBody, &newProperty)

	if marErr != nil {
		jsonResponse.PrintJSONResponse(res, http.StatusBadRequest, marErr.Error(), nil)
		return
	}

	// validation
	validateErr := newProperty.Validate()

	if validateErr != nil {
		jsonResponse.PrintJSONResponse(res, http.StatusBadRequest, validateErr.Error(), nil)
		return
	}

	result := database.DB.Create(&newProperty)

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
	jsonResponse.PrintJSONResponse(res, http.StatusCreated, "success", newProperty)
}

func UpdateProperty(res http.ResponseWriter, req *http.Request) {

	var jsonResponse models.JSONResponse

	vars := mux.Vars(req)
	propertyID, err := strconv.Atoi(vars["id"])

	if err != nil {
		jsonResponse.PrintUnexpectedErrorResponse(res)
		return
	}

	var property models.Property
	getResult := database.DB.First(&property, propertyID)

	if getResult.Error != nil {
		if errors.Is(getResult.Error, gorm.ErrRecordNotFound) {
			jsonResponse.PrintNotFoundResponse(res)
		} else {
			jsonResponse.PrintUnexpectedErrorResponse(res)
		}
		return
	}

	// json decode
	var newProperty *models.Property
	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		jsonResponse.PrintUnexpectedErrorResponse(res)
		return
	}

	json.Unmarshal(reqBody, &newProperty)

	marErr := json.Unmarshal(reqBody, &newProperty)

	if marErr != nil {
		jsonResponse.PrintJSONResponse(res, http.StatusBadRequest, marErr.Error(), nil)
		return
	}

	// Validation
	validateErr := newProperty.Validate()

	if validateErr != nil {
		jsonResponse.PrintJSONResponse(res, http.StatusBadRequest, validateErr.Error(), nil)
		return
	}

	// Update values
	property.Name = newProperty.Name
	property.UpdatedAt = time.Now()

	updateResult := database.DB.Save(&property)

	if updateResult.Error != nil {
		jsonResponse.PrintJSONResponse(res, http.StatusInternalServerError, updateResult.Error.Error(), nil)
		return
	}

	// success
	jsonResponse.PrintJSONResponse(res, http.StatusOK, "success", property)
}

func DeleteProperty(res http.ResponseWriter, req *http.Request) {

	var jsonResponse models.JSONResponse

	vars := mux.Vars(req)
	propertyID, err := strconv.Atoi(vars["id"])

	if err != nil {
		jsonResponse.PrintUnexpectedErrorResponse(res)
		return
	}

	var property models.Property
	deleteResult := database.DB.Delete(&property, propertyID)

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
