
migration-up:
	migrate -path ./migrations/postgres -database 'postgres://shahzod:1@0.0.0.0:5432/lms?sslmode=disable' up

migration-down:
	migrate -path ./migrations/postgres -database 'postgres://shahzod:1@0.0.0.0:5432/lms?sslmode=disable' down