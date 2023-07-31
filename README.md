**Car Service API**

The Car Service API is a RESTful web service that allows users to manage cars. It provides endpoints to retrieve a list of cars, retrieve a car by its ID, create a new car, update an existing car, and delete a car by its ID.

**Table of Contents**

1. [Endpoints](#endpoints)
2. [Data Models](#data-models)
3. [Test Cases](#test-cases)
4. [How to Run](#how-to-run)

## Endpoints

### Get All Cars

- **Endpoint**: `/cars`
- **HTTP Method**: GET
- **Description**: Retrieves the list of all cars.
- **Responses**:
  - 200: Returns an array of cars.
  - 500: Server error.

### Get Car by ID

- **Endpoint**: `/cars/{id}`
- **HTTP Method**: GET
- **Description**: Retrieves a car by its ID.
- **Parameters**:
  - `id` (path) [Required]: The ID of the car to retrieve.
- **Responses**:
  - 200: Returns the car details.
  - 404: Car not found.
  - 500: Server error.

### Create Car

- **Endpoint**: `/cars`
- **HTTP Method**: POST
- **Description**: Creates a new car.
- **Request Body**: Car object in JSON format.
- **Responses**:
  - 201: Returns the created car details.
  - 400: Bad request (e.g., invalid data).
  - 500: Server error.

### Update Car

- **Endpoint**: `/cars/{id}`
- **HTTP Method**: PUT
- **Description**: Updates an existing car by its ID.
- **Parameters**:
  - `id` (path) [Required]: The ID of the car to update.
- **Request Body**: Car object in JSON format containing the updated data.
- **Responses**:
  - 200: Returns the updated car details.
  - 400: Bad request (e.g., invalid data).
  - 404: Car not found.
  - 500: Server error.

### Delete Car

- **Endpoint**: `/cars/{id}`
- **HTTP Method**: DELETE
- **Description**: Deletes a car by its ID.
- **Parameters**:
  - `id` (path) [Required]: The ID of the car to delete.
- **Responses**:
  - 200: Car successfully deleted.
  - 404: Car not found.
  - 500: Server error.

## Data Models

### Car

Represents the structure of a car.

- `category` (string): The category of the car.
- `color` (string): The color of the car.
- `id` (string): The unique identifier of the car.
- `make` (string): The make of the car.
- `mileage` (integer): The mileage of the car.
- `model` (string): The model of the car.
- `package` (string): The package of the car.
- `price` (integer): The price of the car.
- `year` (integer): The manufacturing year of the car.

## Test Cases

The API includes test cases to ensure the correctness of its functionalities. The following test cases are implemented:

- `TestGetAllCars`: Tests the "Get All Cars" endpoint to retrieve a list of all cars.
- `TestGetCarByID`: Tests the "Get Car by ID" endpoint to retrieve a car by its ID.
- `TestCreateCar`: Tests the "Create Car" endpoint to create a new car.
- `TestUpdateCar`: Tests the "Update Car" endpoint to update an existing car.
- `TestDeleteCar`: Tests the "Delete Car" endpoint to delete a car by its ID.

## How to Run

To run the Car Service API, follow these steps:

1. Clone the repository to your local machine.
2. Open the terminal and navigate to the root directory of the project.
3. Execute the following command to start the server:

```bash
go run main.go
```

4. The server will start running at `http://localhost:8080`.

5. To run the test cases, use the following command:

```bash
go test
```

Ensure that all test cases pass successfully.

Please note that this API is for demonstration purposes and does not include features such as database integration or authentication. For a production-ready application, additional components like a database and authentication mechanisms should be implemented.