# Makefile
.PHONY: test prod

test:
	rm -rf /usr/local/ninety-test/ninetyrd101/ninety-test-rd
	go build -ldflags="-X main.port=8889" -o ninety-test-rd
	mv ninety-test-rd /usr/local/ninety-test/ninetyrd101/
prod:
	go build -ldflags="-X main.port=8888" -o ninety-prod-rd

