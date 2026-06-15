# Start the database and run the web application.
db:
	docker compose up -d --wait

run:
	go run ./cmd/web

# Stop the database.
stop-db:
	docker compose down -v

.PHONY: db run stop-db
