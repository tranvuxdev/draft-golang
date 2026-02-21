run:
# 	go run cmd/main.go
	air
build:
	go build ./...
docker up:
	docker compose up -d
docker down:
	docker compose down -v