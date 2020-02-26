start:
	go run app/api/main.go

mod:
	go mod download

test:
	go test -cover -coverprofile=coverage.out ./...

coverage:
	go tool cover -html=coverage.out

vendor:
	export GO111MODULE=on
	go mod vendor
