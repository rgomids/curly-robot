build:
	@go get
	@rm -rf ./bin; mkdir bin
	@go build -v -o bin/upocwin

br: build
	@./bin/upocwin

install: build
	@sh ./scripts/install.sh