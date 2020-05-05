serve:
	go run cmd/server/main.go

migrate_up:
	go run cmd/migrate/main.go up

migrate_down:
	go run cmd/migrate/main.go down

migrate_refresh: migrate_down migrate_up
