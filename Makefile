BUILD_DIR=./bin
EXEC=z
VERSION=0.2.0

build:
	date >> logs/build_log.txt
	go build -v -o $(BUILD_DIR)/$(EXEC) main.go

run:
	date >> logs/build_log.txt
	go run main.go

noversion:
	go build -v -o $(BUILD_DIR)/$(EXEC) main.go

test:
	date >> logs/build_log.txt
	go test ./...

version:
	python3 logs/build_string.py logs/build_log.txt main.go $(VERSION)

release:
	rm -rf $(BUILD_DIR)
	python3 logs/build_string.py logs/build_log.txt main.go $(VERSION)
	go test ./...
	GOOS=linux GOARCH=amd64 go build -v -o $(BUILD_DIR)/linux_amd64/$(EXEC)
	GOOS=linux GOARCH=arm64 go build -v -o $(BUILD_DIR)/linux_arm64/$(EXEC)
	GOOS=darwin GOARCH=amd64 go build -v -o $(BUILD_DIR)/darwin_amd64/$(EXEC)
	GOOS=darwin GOARCH=arm64 go build -v -o $(BUILD_DIR)/darwin_arm64/$(EXEC)
	GOOS=windows GOARCH=amd64 go build -v -o $(BUILD_DIR)/windows_amd64/$(EXEC).exe
	zip $(BUILD_DIR)/$(EXEC)$(VERSION)_linux_amd64.zip $(BUILD_DIR)/linux_amd64/*
	zip $(BUILD_DIR)/$(EXEC)$(VERSION)_linux_arm64.zip $(BUILD_DIR)/linux_arm64/*
	zip $(BUILD_DIR)/$(EXEC)$(VERSION)_darwin_amd64.zip $(BUILD_DIR)/darwin_amd64/*
	zip $(BUILD_DIR)/$(EXEC)$(VERSION)_darwin_arm64.zip $(BUILD_DIR)/darwin_arm64/*
	zip $(BUILD_DIR)/$(EXEC)$(VERSION)_windows_amd64.zip $(BUILD_DIR)/windows_amd64/*
	rm -r $(BUILD_DIR)/*/