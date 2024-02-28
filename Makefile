BINARY_NAME=x
 
build:
	go build -o ${BINARY_NAME} cmd/x/main.go
 
clean:
	go clean
	rm ${BINARY_NAME}

fmt:
	go fmt ./...

install: build
	mkdir -p ~/.local/bin
	install x ~/.local/bin/x
