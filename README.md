# Album API

My first API project using Go and Gin.
This project is a simple RESTful API for managing a collection of music albums, built with Go and the Gin web framework.

## Features

- List all albums
- Get a specific album by ID
- Add a new album
- Dockerized for easy deployment

## Getting Started

### Prerequisites

- Go 1.18 or later
- Docker (optional)

### Running the application

1. Clone the repository
2. Navigate to the project directory
3. Run the application:
```go run main.go```
4. The server will start on `localhost:8080`

### API Endpoints

- GET `/albums`: Retrieve all albums
- GET `/albums/:id`: Retrieve a specific album
- POST `/albums`: Add a new album

### Running with Docker

1. Build the Docker image:
```docker build -t golang-album-api .```
2. Run the container:
```docker run -p 8080:8080 golang-album-api```


## Testing

Run the tests using:
``` got test```

