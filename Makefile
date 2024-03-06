
.PHONY: all
all:
	go build -o bin/main

.PHONY: dev
dev:
	air & pnpm css

.PHONY: fmt
fmt:
	go fmt ./...

