PACKAGE := github.com/31z4/harvest
BUILD_DIR := build

test:
	go test -v -race -coverprofile=coverage.out $(PACKAGE)

build:
	go build -o $(BUILD_DIR)/harvest $(PACKAGE)

clean:
	$(RM) *.out
	$(RM) -r $(BUILD_DIR)

.PHONY: test build clean
