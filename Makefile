build:
	go build -o ./bin/lb ./cmd/lb/*
	go build -o ./bin/test-client ./cmd/test-client/*
	go build -o ./bin/test-server ./cmd/test-server/*

clean:
	rm -rf ./bin

run-lb: build
	./bin/lb

run-test-client: build
	./bin/test-client

run-test-server-1: build
	./bin/test-server 9090

run-test-server-2: build
	./bin/test-server 9091

run-test-server-3: build
	./bin/test-server 9092

