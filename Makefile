BUILD_DIR=./bin
EXEC=z

build:
	date >> build_log.txt
	go build -v -o $(BUILD_DIR)/$(EXEC) main.go

run:
	date >> build_log.txt
	go run main.go

test:
	date >> build_log.txt
	go test ./...