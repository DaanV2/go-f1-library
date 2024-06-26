assembly:
	go tool compile -S x.go

test:
	go test ./... --cover -coverprofile reports/coverage

coverage-report: test
	go tool cover -html reports/coverage

local-upd-server:
	go run cmd/upd/main.go