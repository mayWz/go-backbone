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
