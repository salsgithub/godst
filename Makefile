.PHONY: audit tidy test bench

audit:
	go mod tidy -diff
	go mod verify
	go vet ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

tidy:
	go mod tidy -v
	go fmt ./...

test:
	go test -v -race ./...

bench:
	go test -run=NO_TEST -bench . -benchmem -benchtime 1s ./...