## Car Inventory API

The Car Inventory API allows you to manage a collection of cars, including retrieving all cars, getting a car by its ID, adding a new car, and updating an existing car.

### Base URL

`http://localhost:8080`

### Endpoints

#### 1. Get All Cars

- URL: `/cars`
- Method: `GET`
- Description: Retrieves a list of all available cars in the inventory.
- Response: Returns a JSON array containing car objects with the following properties:
  - `make` (string): The make of the car.
  - `model` (string): The model of the car.
  - `package` (string): The package of the car.
  - `color` (string): The color of the car.
  - `year` (int): The manufacturing year of the car.
  - `category` (string): The category of the car (e.g., Sedan, SUV, Truck).
  - `mileage` (int): The mileage of the car.
  - `price` (int): The price of the car.
  - `id` (string): The unique identifier of the car.

#### 2. Get Car by ID

- URL: `/cars/{id}`
- Method: `GET`
- Description: Retrieves a specific car by its unique ID.
- URL Parameters:
  - `{id}` (string): The unique identifier of the car to be retrieved.
- Response: Returns a JSON object representing the car with the provided ID or a 404 error if the car is not found.

#### 3. Add New Car

- URL: `/cars/new`
- Method: `POST`
- Description: Adds a new car to the inventory.
- Request Body: The payload should be a JSON object containing the car details with the following properties:
  - `make` (string): The make of the car.
  - `model` (string): The model of the car.
  - `package` (string): The package of the car.
  - `color` (string): The color of the car.
  - `year` (int): The manufacturing year of the car.
  - `category` (string): The category of the car (e.g., Sedan, SUV, Truck).
  - `mileage` (int): The mileage of the car.
  - `price` (int): The price of the car.
- Response: Returns a JSON object representing the newly added car with a unique `id` field.

#### 4. Update Existing Car

- URL: `/cars/update`
- Method: `PUT`
- Description: Updates the details of an existing car in the inventory.
- Request Body: The payload should be a JSON object containing the updated car details with the following properties:
  - `id` (string): The unique identifier of the car to be updated (this field is required and cannot be changed).
  - `make` (string): The updated make of the car.
  - `model` (string): The updated model of the car.
  - `package` (string): The updated package of the car.
  - `color` (string): The updated color of the car.
  - `year` (int): The updated manufacturing year of the car.
  - `category` (string): The updated category of the car (e.g., Sedan, SUV, Truck).
  - `mileage` (int): The updated mileage of the car.
  - `price` (int): The updated price of the car.
- Response: Returns a JSON object representing the updated car with the provided ID or a 404 error if the car is not found.

---