// main_test.go

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllCars(t *testing.T) {
	req, err := http.NewRequest("GET", "/cars", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleAllCars)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Check the content type
	expectedContentType := "application/json"
	actualContentType := rr.Header().Get("Content-Type")
	if actualContentType != expectedContentType {
		t.Errorf("Expected Content-Type %s, but got %s", expectedContentType, actualContentType)
	}

	// Check the response body
	var carsResponse []Car
	err = json.NewDecoder(rr.Body).Decode(&carsResponse)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}
}

func TestGetCarByID(t *testing.T) {
	// Initialize the test data
	testCarID := "JHk290Xj"
	expectedCar := Car{
		Make:     "Ford",
		Model:    "F10",
		Package:  "Base",
		Color:    "Silver",
		Year:     2010,
		Category: "Truck",
		Mileage:  120123,
		Price:    1999900,
		ID:       "JHk290Xj",
	}

	req, err := http.NewRequest("GET", "/cars/"+testCarID, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handleCarByID(w, r)
	})

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Check the content type
	expectedContentType := "application/json"
	actualContentType := rr.Header().Get("Content-Type")
	if actualContentType != expectedContentType {
		t.Errorf("Expected Content-Type %s, but got %s", expectedContentType, actualContentType)
	}

	// Check the response body
	var carResponse Car
	err = json.NewDecoder(rr.Body).Decode(&carResponse)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	// Check if the returned car matches the expected car
	if !equalCars(expectedCar, carResponse) {
		t.Errorf("Expected car %+v, but got %+v", expectedCar, carResponse)
	}
}

func TestCreateCar(t *testing.T) {
	// Initialize the test data
	newCar := Car{
		Make:     "Tesla",
		Model:    "Model S",
		Package:  "Performance",
		Color:    "Black",
		Year:     2023,
		Category: "Electric",
		Mileage:  100,
		Price:    8500000,
		// ID will be generated during the creation process
	}

	// Convert the newCar to JSON
	requestBody, err := json.Marshal(newCar)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/cars", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		createCar(w, r)
	})

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, rr.Code)
	}

	// Check the content type
	expectedContentType := "application/json"
	actualContentType := rr.Header().Get("Content-Type")
	if actualContentType != expectedContentType {
		t.Errorf("Expected Content-Type %s, but got %s", expectedContentType, actualContentType)
	}

	// Check the response body
	var createdCarResponse Car
	err = json.NewDecoder(rr.Body).Decode(&createdCarResponse)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	// Check if the created car matches the expected car data
	if !equalCars(newCar, createdCarResponse) {
		t.Errorf("Expected created car %+v, but got %+v", newCar, createdCarResponse)
	}
}

func TestUpdateCar(t *testing.T) {
	// Initialize the test data
	testCarID := "JHk290Xj"
	updatedCarData := Car{
		Make:     "Ford",
		Model:    "F10",
		Package:  "Special Edition",
		Color:    "Black",
		Year:     2010,
		Category: "Truck",
		Mileage:  120123,
		Price:    1999900,
		// ID will remain the same
	}

	// Convert the updatedCarData to JSON
	requestBody, err := json.Marshal(updatedCarData)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/cars/"+testCarID, bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		updateCar(w, r, testCarID)
	})

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Check the content type
	expectedContentType := "application/json"
	actualContentType := rr.Header().Get("Content-Type")
	if actualContentType != expectedContentType {
		t.Errorf("Expected Content-Type %s, but got %s", expectedContentType, actualContentType)
	}

	// Check the response body
	var updatedCarResponse Car
	err = json.NewDecoder(rr.Body).Decode(&updatedCarResponse)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	// Check if the updated car matches the expected car data
	if !equalCars(updatedCarData, updatedCarResponse) {
		t.Errorf("Expected updated car %+v, but got %+v", updatedCarData, updatedCarResponse)
	}
}

func TestDeleteCar(t *testing.T) {
	// Initialize the test data
	existingCarID := "JHk290Xj"
	unregisteredCarID := "NonExistentCarID"

	// Test the case when the car is found and deleted
	req, err := http.NewRequest("DELETE", "/cars/"+existingCarID, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		deleteCar(w, r, existingCarID)
	})

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Check the content type
	expectedContentType := "application/json"
	actualContentType := rr.Header().Get("Content-Type")
	if actualContentType != expectedContentType {
		t.Errorf("Expected Content-Type %s, but got %s", expectedContentType, actualContentType)
	}

	// Check the response body
	var responseMessage map[string]string
	err = json.NewDecoder(rr.Body).Decode(&responseMessage)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	expectedMessage := "Successfully deleted car with ID: " + existingCarID
	actualMessage, ok := responseMessage["message"]
	if !ok {
		t.Error("Expected response message key \"message\" not found")
	} else if actualMessage != expectedMessage {
		t.Errorf("Expected response message \"%s\", but got \"%s\"", expectedMessage, actualMessage)
	}

	// Test the case when attempting to delete an unregistered car
	reqUnregistered, err := http.NewRequest("DELETE", "/cars/"+unregisteredCarID, nil)
	if err != nil {
		t.Fatal(err)
	}

	rrUnregistered := httptest.NewRecorder()
	handlerUnregistered := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		deleteCar(w, r, unregisteredCarID)
	})

	handlerUnregistered.ServeHTTP(rrUnregistered, reqUnregistered)

	if rrUnregistered.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, rrUnregistered.Code)
	}

	// Check the content type
	actualContentTypeUnregistered := rrUnregistered.Header().Get("Content-Type")
	if actualContentTypeUnregistered != expectedContentType {
		t.Errorf("Expected Content-Type %s, but got %s", expectedContentType, actualContentTypeUnregistered)
	}

	// Check the response body for "Car not found" message
	var responseMessageUnregistered map[string]string
	err = json.NewDecoder(rrUnregistered.Body).Decode(&responseMessageUnregistered)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	expectedMessageUnregistered := "Car not found"
	actualMessageUnregistered, ok := responseMessageUnregistered["message"]
	if !ok {
		t.Error("Expected response message key \"message\" not found")
	} else if actualMessageUnregistered != expectedMessageUnregistered {
		t.Errorf("Expected response message \"%s\", but got \"%s\"", expectedMessageUnregistered, actualMessageUnregistered)
	}
}

// Helper function to compare two cars
func equalCars(a, b Car) bool {
	return a.Make == b.Make &&
		a.Model == b.Model &&
		a.Package == b.Package &&
		a.Color == b.Color &&
		a.Year == b.Year &&
		a.Category == b.Category &&
		a.Mileage == b.Mileage &&
		a.Price == b.Price
}
