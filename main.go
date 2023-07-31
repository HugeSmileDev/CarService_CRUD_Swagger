package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

// Car represents the structure of a car.
type Car struct {
	Make     string `json:"make"`
	Model    string `json:"model"`
	Package  string `json:"package"`
	Color    string `json:"color"`
	Year     int    `json:"year"`
	Category string `json:"category"`
	Mileage  int    `json:"mileage"`
	Price    int    `json:"price"`
	ID       string `json:"id"`
}

var cars = []Car{
	{
		Make:     "Ford",
		Model:    "F10",
		Package:  "Base",
		Color:    "Silver",
		Year:     2010,
		Category: "Truck",
		Mileage:  120123,
		Price:    1999900,
		ID:       "JHk290Xj",
	},
	{
		Make:     "Toyota",
		Model:    "Camry",
		Package:  "SE",
		Color:    "White",
		Year:     2019,
		Category: "Sedan",
		Mileage:  3999,
		Price:    2899000,
		ID:       "fWl37la",
	},
	{
		Make:     "Toyota",
		Model:    "Rav4",
		Package:  "XSE",
		Color:    "Red",
		Year:     2018,
		Category: "SUV",
		Mileage:  24001,
		Price:    2275000,
		ID:       "1j3xjRllc",
	},
	{
		Make:     "Ford",
		Model:    "Bronco",
		Package:  "Badlands",
		Color:    "Burnt Orange",
		Year:     2022,
		Category: "SUV",
		Mileage:  1,
		Price:    4499000,
		ID:       "dku43920s",
	},
}

var ErrCarNotFound = errors.New("car not found")

func findCarByID(id string) (*Car, error) {
	for i, car := range cars {
		if car.ID == id {
			return &cars[i], nil
		}
	}
	return nil, ErrCarNotFound
}

// swagger:parameters getCarByID
type CarIDParam struct {
	// The ID of the car to retrieve
	// in: path
	ID string `json:"id"`
}

// swagger:response carResponse
type CarResponse struct {
	// in: body
	Body Car
}

// swagger:route GET /cars/{id} getCars getCarByID
// Retrieves a car by ID.
// Responses:
//   200: carResponse
//   404: errorResponse
//   500: errorResponse

// swagger:response carsResponse
type CarsResponse struct {
	// in: body
	Body []Car
}

// swagger:route GET /cars getCars getAllCars
// Retrieves the list of cars.
// Responses:
//   200: carsResponse

// swagger:parameters createCar
type CreateCarParams struct {
	// The car to create.
	// in: body
	Body Car
}

// swagger:response carCreatedResponse
type CarCreatedResponse struct {
	// in: body
	Body Car
}

// swagger:route POST /cars createCar createCar
// Creates a new car.
// Responses:
//   201: carCreatedResponse
//   400: errorResponse
//   500: errorResponse

// swagger:parameters updateCar
type UpdateCarParams struct {
	// The ID of the car to update.
	// in: path
	ID string `json:"id"`
	// The updated car data.
	// in: body
	Body Car
}

// swagger:route PUT /cars/{id} updateCar updateCar
// Updates an existing car.
// Responses:
//   200: carResponse
//   400: errorResponse
//   404: errorResponse
//   500: errorResponse

func handleAllCars(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Println("Get all cars")
		getAllCars(w, r)
	case http.MethodPost:
		log.Println("Create car")
		createCar(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func handleCarByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/cars/"):]
	switch r.Method {
	case http.MethodGet:
		log.Println("Get car by ID")
		getCarByID(w, r, id)
	case http.MethodPut:
		log.Println("Update car by ID")
		updateCar(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cars", handleAllCars)
	mux.HandleFunc("/cars/", handleCarByID)

	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {

		log.Println("Here is Swagger.json URL")
		filePath := "./swagger.json"
		file, err := os.Open(filePath)
		if err != nil {
			http.Error(w, "Failed to open swagger.json", http.StatusInternalServerError)
			log.Println("Error opening swagger.json:", err)
			return
		}
		defer file.Close()

		// Set the appropriate content type for the response
		w.Header().Set("Content-Type", "application/json")

		_, err = io.Copy(w, file)
		if err != nil {
			http.Error(w, "Failed to serve swagger.json", http.StatusInternalServerError)
			log.Println("Error serving swagger.json:", err)
			return
		}
	})

	// Serve the Swagger UI
	swaggerUIHandler := http.FileServer(http.Dir("./swagger-ui/"))
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", swaggerUIHandler))

	log.Println("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func getAllCars(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Add observability: Log the request
	log.Printf("GET /cars")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

func getCarByID(w http.ResponseWriter, r *http.Request, id string) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Find the car by ID
	car, err := findCarByID(id)
	if err != nil {
		if errors.Is(err, ErrCarNotFound) {
			http.Error(w, "Car not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	// Add observability: Log the request with car ID
	log.Printf("GET /cars/%s", id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)
}

func createCar(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Add observability: Log the request
	log.Printf("POST /cars")

	// Parse the request body to create a new car
	var newCar Car
	err := json.NewDecoder(r.Body).Decode(&newCar)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Assign a new ID for the car
	newID := uuid.New().String()[:8]
	newCar.ID = newID

	// Add the new car to the slice
	cars = append(cars, newCar)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Respond with the details of the newly added car
	json.NewEncoder(w).Encode(newCar)
}

func updateCar(w http.ResponseWriter, r *http.Request, id string) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Add observability: Log the request
	log.Printf("PUT /cars")

	log.Println(r.URL.Path)

	// Find the car by ID
	car, err := findCarByID(id)
	if err != nil {
		if errors.Is(err, ErrCarNotFound) {
			http.Error(w, "Car not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	// Parse the request body to update an existing car
	var updatedCar Car
	err = json.NewDecoder(r.Body).Decode(&updatedCar)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Update the car fields
	car.Make = updatedCar.Make
	car.Model = updatedCar.Model
	car.Package = updatedCar.Package
	car.Color = updatedCar.Color
	car.Year = updatedCar.Year
	car.Category = updatedCar.Category
	car.Mileage = updatedCar.Mileage
	car.Price = updatedCar.Price

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Respond with the details of the updated car
	json.NewEncoder(w).Encode(car)
}
