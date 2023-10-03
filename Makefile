BINARY_NAME=fetch

test: 
	go test ./... -v

build: test
	CGO_ENABLED=0 go build -o ${BINARY_NAME} main.go

run: build 
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
