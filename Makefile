migrateup:
	migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down