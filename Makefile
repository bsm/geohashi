default: test

test:
	go test ./... -v 1

bench:
	go test ./... -run=NONE -bench=.
