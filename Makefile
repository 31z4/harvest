BUILD_DIR := build

coverage: coverage.out

coverage.out: test
	cp harvest.out $@
	grep -h -v '^mode:' harvest.trie.out >> $@

test: harvest.out harvest.trie.out

harvest.out: $(wildcard *.go)
	go test -v -race -coverprofile=$@ github.com/31z4/harvest

harvest.trie.out: $(wildcard trie/*.go)
	go test -v -race -coverprofile=$@ github.com/31z4/harvest/trie

build:
	go build -o $(BUILD_DIR)/harvest github.com/31z4/harvest

clean:
	$(RM) *.out
	$(RM) -r $(BUILD_DIR)

.PHONY: build clean
