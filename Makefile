tidy:
	go mod tidy
	go mod vendor
	
build: tidy
	go build -o build/go-movies main.go

server:
	go run main.go