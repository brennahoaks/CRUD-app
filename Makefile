BINARY_NAME=CRUD-app

build:
	go build -o ${BINARY_NAME}-main main.go

run:
	go run main.go

build_and_run: build run

clean:
	go clean
	rm ${BINARY_NAME}-main