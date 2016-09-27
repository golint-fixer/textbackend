# textbackend Makefile

VERSION=0.1.0

test:
	@echo "\nrun textbackend v${VERSION} tests & benchmark\n"
	@golint
	@go test -bench=. -benchmem -v