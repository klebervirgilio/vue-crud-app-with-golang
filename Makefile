GINKGO := $$GOPATH/bin/ginkgo

setup: run_services
	@go run ./cmd/db/setup.go

run_services:
	@docker-compose up --build -d

run_tests: run_services
	@${GINKGO} pkg/**

run_server:
	@MONGO_URL=mongodb://mongo_user:mongo_secret@0.0.0.0:27017/kudos PORT=:4444 go run cmd/main.go
