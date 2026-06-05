.PHONY: install test cov format doc generate-website

install:
	@if [ "$(CLEAN)" = "true" ]; then \
		go mod download; \
	else \
		go mod tidy; \
	fi

test:
	go test ./library/... ./library-extension/...

cov:
	go test -coverprofile=coverage.out ./library/... ./library-extension/...

format:
	go fmt ./...

doc:
	rm -rf build/doc2go/
	doc2go -out build/doc2go/ ./library/... ./library-extension/...

generate-website: doc
	rm -rf website/static/api/
	mkdir -p website/static/api/
	mv build/doc2go/ website/static/api/godoc/
	cd website && hugo --minify
