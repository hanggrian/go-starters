.PHONY: install test cov format generate-website

install:
	@if [ "$(CLEAN)" = "true" ]; then \
		go mod download; \
	else \
		go mod tidy; \
	fi

test:
	go test ./main/...

cov:
	go test -coverprofile=coverage.out ./main/...

format:
	go fmt ./...

generate-website:
	cd website && hugo --minify
