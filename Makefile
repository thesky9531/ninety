# Makefile
.PHONY: test prod

test:
	go build -ldflags="-X main.port=8889" -o ninety-test-rd

prod:
	go build -ldflags="-X main.port=8888" -o ninety-prod-rd

