# Start the database in the background.
db/up:
	docker compose up -d

# Stop the database but KEEP the data.
db/down:
	docker compose down

# Wipe the database completely and re-run the init script with fresh seed data.
db/reset:
	docker compose down -v
	docker compose up -d

# Open a mysql shell inside the container as the web user.
db/shell:
	docker compose exec mysql mysql -D snippetbox -u web -ppassword

# Run the web application.
run:
	go run ./cmd/web

.PHONY: db/up db/down db/reset db/shell run
