BUILD_DIR=./bin
EXEC=z
VERSION=0.1.0

build:
	date >> logs/build_log.txt
	go build -v -o $(BUILD_DIR)/$(EXEC) main.go

run:
	date >> logs/build_log.txt
	go run main.go

test:
	date >> logs/build_log.txt
	go test ./...

release:
	go test ./...
	GOOS=linux GOARCH=amd64 go build -v -o $(BUILD_DIR)/linux_amd64/$(EXEC)
	GOOS=linux GOARCH=arm64 go build -v -o $(BUILD_DIR)/linux_arm64/$(EXEC)
	GOOS=darwin GOARCH=amd64 go build -v -o $(BUILD_DIR)/darwin_amd64/$(EXEC)
	GOOS=windows GOARCH=amd64 go build -v -o $(BUILD_DIR)/windows_amd64/$(EXEC).exe
	zip $(BUILD_DIR)/$(EXEC)v$(VERSION)_linux_amd64.zip $(BUILD_DIR)/linux_amd64
	zip $(BUILD_DIR)/$(EXEC)v$(VERSION)_linux_arm64.zip $(BUILD_DIR)/linux_arm64
	zip $(BUILD_DIR)/$(EXEC)v$(VERSION)_darwin_amd64.zip $(BUILD_DIR)/darwin_amd64
	zip $(BUILD_DIR)/$(EXEC)v$(VERSION)_windows_amd64.zip $(BUILD_DIR)/windows_amd64
	rm -r $(BUILD_DIR)/*/