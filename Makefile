MODULE_PATH=github.com/petabite/gomachine

.PHONY: all
all: gomachine

.PHONY: gomachine
gomachine:
	go build -v ./cmd/gomachine

.PHONY: run
run:
	go run $(MODULE_PATH)/cmd/gomachine $(SOURCE)

.PHONY: format
format:
	gofmt -l -s -w .
