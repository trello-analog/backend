build: go build -v ./cmd/trelloserver
migrationsDown: migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/trello?sslmode=disable" down
migrationsUp: migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/trello?sslmode=disable" up
createMigration migrate create -ext sql -dir migrations <name>