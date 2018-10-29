all: generate test
	go build

test:
	go test

generate:
	rm global.go
	./gen-global.sh
