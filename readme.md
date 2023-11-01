# Weather Service Backend

Welcome to the Weather Service Backend repository. This backend server provides weather data retrieval, authentication, and more for the Weather UI project.

## Getting Started

### 1. Clone the Repository
To get started, clone the repository using the following command:

```bash
git clone https://github.com/SouravChoudhary/weather_backend_server.git
```

### 2. Install Dependencies

Navigate to the project folder and install the required dependencies:
```bash
go mod tidy
```

### 3. Running the Server

To run the server, use the following command:

```bash
 1) cd weather_backend_server
 2) go run ./cmd/server/main.go  
```

### 4. Configuration

#### Database Credentials and API Keys

You will need to set your database DSN and API keys for services like OpenWeatherMap in the `config.yaml` file. Ensure the following fields are correctly configured:

```yaml
[Database]
DSN = "your-dsn-here"

[OpenWeatherMap]
APIKey = "your-api-key-here"
```

The server uses GORM, so no need to create database tables manually; they will be auto-migrated.

## 5. Unit Test 

- Unit Test Case is written for api/service/weather.go file to 
showcase general unit test case strategy. Similar test template can be used to write more unit tests.

```go 
go test ./... -coverprofile=coverage.out 
go tool cover -html=coverage 
```

## 6. API Documentation
-  check documentation directory for api documentation.

## 7. Limitations

- The server may have rate limits imposed by external APIs like OpenWeatherMap.
- Unit Test Case is written for api/service/weather.go file to 
showcase general unit test case strategy and to save time 
- proper error and logging package is build but not used extensively across the repo 


## 8. Repo Architecture
The architecture of the Weather Service Backend is organized and modular, making it easy to maintain, scale, and extend. It provides routes for user authentication and weather data retrieval, with well-structured controllers handling the business logic. External services like OpenWeatherMap are accessed through service components, and the project follows best practices. It's designed to serve as a robust backend for the Weather UI project.

Feel free to explore and contribute to the codebase, and don't hesitate to reach out if you have any questions or need further details.