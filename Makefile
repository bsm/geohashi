default: vet test

test:
	go test ./... -v 1

vet:
	go vet ./...

bench:
	go test ./... -run=NONE -bench=.
