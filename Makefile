BINARY_NAME=fetch

build:
	CGO_ENABLED=0 go build -o ${BINARY_NAME} main.go

run: build 
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
