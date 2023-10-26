## GO backbone

## Project Overview

A go service using the Gin&GorillarWS with Clean Architecture

## Style guide 

- This project using the uber style guide, please check [uber-go-style-guid](https://github.com/uber-go/guide)
- Project layout from standard project layout, [Project-layout](https://github.com/golang-standards/project-layout)
- Using Gin framework, for more futher feature please check [gin](https://gin-gonic.com/docs/)
- Using gorm for ORM. for more futher feature please check [gorm](https://gorm.io/)

## Prerequisite
GO 1.21 

## Installaion & Setup
(Optional) Setup the Database, This source code provide the docker image for start the database 
```bash
# Start the database [Optional]
$ docker compose up -d 
```
Please feel free to using the Client or cli to create a database, username, password. 
------
```bash
# Start the database [Optional]
$ docker compose up 
# Setup .env
$ cp .env.exmaple .env 
# Download go package 
$ go mod download
# Start the servie in dev mode
$ go run cmd/server/main.go
```

## Database Migration
inn this project, we using golang-migrate to handle the database migration file. [Golang-migrate](https://github.com/golang-migrate/)

```bash
# Install golang-migrate
$ brew install golang-migrate
# Create a new migrate file 
$ migrate create -ext sql -dir db/migrations -seq file_name
# Run migration file, please replace database connection string
$ migrate -database ‘mysql://user:password@tcp(host:port)/dbname?query’ -path db/migrations up
```


## Testing 
To run test on this project. make sure to generate mock data before run test
We using mockgen to generate mock repository and service.
for intall mockgen please find on this page (uber-go/mock)[https://github.com/uber-go/mock]
```bash
# genereate mock repo, service
make mockgen-all
# run all test
make test-all
```

## Structure 
```
├── cmd
|   ├── server
│         └──  main.go
├── internal
|   |-- api
|   |-- app
|   |-- config
|   |-- model
|   |-- repository
│   └── service
├── test
└── main.go
```
