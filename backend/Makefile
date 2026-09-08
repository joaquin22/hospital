.PHONY: run build test tidy seed clean db-up db-down db-logs docker-up docker-down

run:
	go run ./cmd/api

build:
	go build -o bin/hotel-api ./cmd/api

test:
	go test ./... -v

tidy:
	go mod tidy

seed:
	bash scripts/seed.sh

clean:
	rm -f bin/hotel-api

# --- Postgres local (solo la base de datos, para correr la API con `make run`) ---
db-up:
	docker compose -f build/docker-compose.yml up -d postgres

db-down:
	docker compose -f build/docker-compose.yml down

db-logs:
	docker compose -f build/docker-compose.yml logs -f postgres

# --- Stack completo dockerizado (Postgres + API) ---
docker-up:
	docker compose -f build/docker-compose.yml up -d --build

docker-down:
	docker compose -f build/docker-compose.yml down
