postgres:
	docker run --name go-movies-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it go-movies-postgres createdb --username=root --owner=root go_movies

populatedb:
	docker cp data/go.sql go-movies-postgres:/
	docker exec -it go-movies-postgres psql -d go_movies -f /go.sql

db: postgres createdb populatedb

cleanpostgres:
	docker rm -f go-movies-postgres

tidy:
	go fmt
	go mod tidy
	go mod vendor
	
build: tidy
	go build -o build/go-movies main.go

server:
	go run main.go

.PHONY: postgres createdb populatedb cleanpostgres db tidy build server