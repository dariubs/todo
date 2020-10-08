all: build

run:
	go run .

install:
	sudo mv todo /usr/bin/
	sudo cp todo.toml /usr/bin/

build:
	go build -o todo
	
