PACKAGE := github.com/31z4/harvest
BUILD_DIR := build

build:
	go build -o $(BUILD_DIR)/harvest $(PACKAGE)

test:
	go test -tags test -v -race -coverprofile=coverage.out $(PACKAGE)

clean:
	$(RM) *.out
	$(RM) -r $(BUILD_DIR)

.PHONY: test build clean
