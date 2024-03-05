# Golang API

## Dependencies

- [Gin](https://github.com/gin-gonic/gin) - Framework, v1.9.1
- [GORM](gorm.io/gorm) - Object Relational Mapping, v1.25.6
- [Testify](github.com/stretchr/testify) - testing toolkit, v1.8.4
- [swag](github.com/swaggo/swag) - swagger, v1.6.0
- [gin-swagger](github.com/swaggo/swag) - gin middleware to swagger, v1.16.3
  
## Requirements
- docker-engine or docker desktop at least version 25.0.2
- docker compose 2.24.5 or above
- golang 1.21.6 or above for develop

## Tools
- Visual Studio Code - IDE
- Bash - terminal
- Postman - testing

## Building
### local dev
1. go to repo location in your computer and use the following command to download and install all dependencies
```bash
$ go mod tidy
```
### docker
1. Open terminal and go to the repositories's directory
2. Build docker compose
```bash
$ docker compose build
```
2. Start docker compose
```bash
$ docker compose up
```
## Testing
### unit testing
- executes all unit testing
```bash
$ go test ./...
```
### code coverage
1. Open terminal and go to the repositories's directory
2. Execute command to create coverage.out via unit testing
```bash
go test -coverprofile=coverage.out ./...
```
3. Execute command to generate html file
```bash
go tool cover -html=coverage.out -o coverage.html
```
4. Open html file on browser to check coverage
## Documentation
### swagger
[Docs](http://localhost:8080/docs/index.html)
