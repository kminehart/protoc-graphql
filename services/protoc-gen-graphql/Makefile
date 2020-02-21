.PHONY: build
build: clean
	go build -o bin/ .

.PHONY: clean
clean:
	@mkdir -p bin
	@mkdir -p out

.PHONY: setup
setup:
	ln -s $$PWD/bin/protoc-gen-graphql $$HOME/.local/bin/protoc-gen-graphql

.PHONY: gen
gen: build
	protoc --graphql_out=./out --proto_path=. example/kitchen.proto
