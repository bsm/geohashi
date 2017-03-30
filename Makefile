default: vet test

test:
	go test ./... -v

vet:
	go vet ./...

bench:
	go test ./... -run=NONE -bench=.

README.md: README.md.tpl $(wildcard *.go)
	becca -package $(subst $(GOPATH)/src/,,$(PWD))
