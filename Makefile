.PHONY: init
init:
	go mod download
	docker compose build

.PHONY: build-debug
build-debug:
	mkdir -p tmp
	go build -gcflags "all=-N -l" -o ./tmp/hands-on-202401 ./cmd/hands-on-202401