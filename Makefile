PACKAGE := github.com/31z4/harvest
BUILD_DIR := build
RELEASE_DIR := release

build:
	go build -o $(BUILD_DIR)/harvest $(PACKAGE)

release:
	CGO_ENABLED=0 go build -ldflags '-s' -o $(RELEASE_DIR)/harvest $(PACKAGE)

docker:
	docker build --force-rm -t harvest .

test:
	go test -tags test -v -race -coverprofile=coverage.out $(PACKAGE)

clean:
	$(RM) *.out
	$(RM) -r $(BUILD_DIR)
	$(RM) -r $(RELEASE_DIR)

.PHONY: test build release docker clean
