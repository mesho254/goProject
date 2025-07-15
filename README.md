### Go Task Management Application:
- This is a simple task management API built with Go, utilizing the Gin framework and GORM for database interactions with SQLite. It supports basic CRUD operations for managing tasks.

  
## Installation
-To set up the project locally, follow these steps:

## Initialize the Go Module:

```code
cd task-manager
```

```go
go mod init task-manager
```


## Install Dependencies:
```go
go get -u gorm.io/gorm

go get -u gorm.io/driver/sqlite

go get -u github.com/gin-gonic/gin

go install github.com/swaggo/swag/cmd/swag@latest

go get github.com/swaggo/gin-swagger

go get github.com/swaggo/files
```



## Running the Application

# Start the Server:
```go
go run main.go
```

- The server will be available at http://localhost:8080.

# Generate Swagger Documentation:
```code
swag init --parseDependency
```

- This generates Swagger files for API documentation, accessible at http://localhost:8080/swagger/index.html.


# Testing the API Endpoints
- You can test the API using tools like curl, Postman, or a browser. Below are examples for each endpoint:

## List all tasks (GET):
```code
curl http://localhost:8080/tasks
```


## Create a task (POST):
```code
curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d '{"title":"Buy milk","description":"Go to the store","due_date":"2023-10-10T00:00:00Z","status":"pending"}'
```

## Get a task by ID (GET):
```code
curl http://localhost:8080/tasks/1
```

## Update a task (PUT):
```code
curl -X PUT http://localhost:8080/tasks/1 -H "Content-Type: application/json" -d '{"title":"Buy milk and bread"}'
```

## Delete a task (DELETE):
```code
curl -X DELETE http://localhost:8080/tasks/1
```


### Additional Information

- Database: Uses SQLite for lightweight, embedded storage.
- Dependencies: Run go mod tidy to ensure all dependencies are resolved.
- Swagger UI: After generating documentation, visit http://localhost:8080/swagger/index.html to explore the API interactively.
