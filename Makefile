BINDIR := bin

PKG := $(shell go list .)

BIN := $(BINDIR)/traceroute

GO_FILES := $(shell find . -type f -name '*.go' -print)

.PHONY: build
build: $(BIN)

$(BIN): $(GO_FILES)
	@go build -o $@ .

.PHONY: clean
clean:
	@$(RM) $(BIN)
