tidy:
	go mod tidy
	go mod vendor
	
server:
	go run cmd/api/*.go