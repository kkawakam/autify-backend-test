BINARY_NAME=fetch

build:
	go build -o ${BINARY_NAME} main.go

run: build 
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
