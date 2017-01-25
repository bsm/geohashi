default: vet test

test:
	go test ./... -v

vet:
	go vet ./...

bench:
	go test ./... -run=NONE -bench=.
