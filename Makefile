.PHONY: build run test docker

build:
	go build -o stress-cli ./cmd/stress-cli/main.go

run:
	./stress-cli --url=http://localhost --requests=100 --concurrency=10

test:
	go test ./...

docker:
	docker build -t stress-cli .

docker-run:
	docker run stress-cli --url=http://localhost --requests=100 --concurrency=10
