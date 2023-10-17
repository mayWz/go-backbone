build:
	go build -o bin/main cmd/server/main.go

run:
	go run cmd/server/main.go

mockgen-repo:
	mockgen -source=internal/repository/user.go -destination=test/mocks/repository/user.go -package=mocks

mockgen-service:
	mockgen -source=internal/service/user.go -destination=test/mocks/service/user.go -package=mocks

mockgen-all:
	make mockgen-repo
	make mockgen-service

test-all:
	go test -v ./...

test-service:
	go test -v ./internal/service
