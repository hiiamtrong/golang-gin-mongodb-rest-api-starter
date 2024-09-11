# Project golang-gin-mongodb-rest-api-starter

## Description
This is a starter project for building RESTful APIs using Golang, Gin, and MongoDB. It is a simple CRUD API that allows you to create, read, update, and delete todo.

## Prerequisites
- Go 1.21 or higher
- Docker 
- Make

## Installation
### Docker
1. Copy env example and change value for docker
```sh
cp .env.example .env
```

2. Run `make compose-up` to start

### Local
1. Copy env example and change value for local
```sh
cp .env.example .env
```

2. Start other services use docker
```sh
make compose-local-up 
```

3. Run `make server` to start the server

## Swagger
Swagger UI local available at localhost:3000/swagger/index.html

### Generate swagger docs
```
make docs
```


