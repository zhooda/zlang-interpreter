BUILD_DIR=./bin
EXEC=z

build:
	go build -v -o $(BUILD_DIR)/$(EXEC) main.go

run:
	go run main.go

test:
	go test ./...