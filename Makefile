install:
	go install ./cmd/game
run:
	GIN_MODE=release go run ./cmd/game/main.go
dev:
	GIN_MODE=debug go run ./cmd/game/main.go

docker-build:
	docker build -f build/docker/Dockerfile -t game-rest-api-mysql:latest .
docker-run:
	docker run -it --name game-rest-api -p 3306:3306 game-rest-api-mysql:latest
docker-stop:
	docker container stop game-rest-api && docker rm game-rest-api

test:
	bash scripts/test.sh
static:
	bash scripts/static.sh
