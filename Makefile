GINKGO := $$GOPATH/bin/ginkgo

setup: run_services
	@go run ./cmd/db/setup.go

run_services:
	@docker-compose up -d

run_tests: run_services
	@${GINKGO} pkg/**
