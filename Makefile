postgres:
	docker run --name go-movies-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

tidy:
	go mod tidy
	go mod vendor
	
build: tidy
	go build -o build/go-movies main.go

server:
	go run main.go