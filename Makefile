COVER_PROFILE_FILE := "./coverage.out"

## Standard Targets
all: test check

test: test-race test-unit

build: build/staledesk

check: check-golint

clean:
	rm -rf build/*
	rm -f $(COVER_PROFILE_FILE)

## Custom Targets
build/staledesk:
	go build -o build/staledesk ./cmd/staledesk

check-golint:
	golint -set_exit_status ./...

test-race:
	go test -race ./...

test-unit:
	go test -cover ./...

show-func-coverage: test-coverprofile test-show-func-coverage

show-coverage-html: test-coverprofile test-show-coverage-html

test-show-coverage-html:
	go tool cover -html=$(COVER_PROFILE_FILE)

test-show-func-coverage:
	go tool cover -func $(COVER_PROFILE_FILE)

test-coverprofile:
	go test -coverprofile $(COVER_PROFILE_FILE) -covermode=count ./...

.PHONY: build \
check check-golint \
clean \
test test-race test-unit
