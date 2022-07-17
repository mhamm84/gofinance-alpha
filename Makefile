.PHONY: install
install: audit
	@echo 'Building gofinance-alpha...'
	go install ./...


.PHONY: audit
audit:
	@echo 'Tidy and verify dependencies...'
	go mod tidy
	go mod verify

	@echo 'Running formatting...'
	go fmt ./...

	@echo 'Running vetting'
	go vet ./...

	@echo 'Running staticcheck'
	staticcheck

	@echo 'Running tests...'
	go test -race -vet=off