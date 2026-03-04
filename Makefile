include .env

create_migration:
	migrate create -ext .sql -dir ./cmd/migrate/migrations "$(name)"


push_migration:
	migrate -path ./cmd/migrate/migrations -database "${DATABASE_URL}" up


reverse_migration:
	migrate -path ./cmd/migrate/migrations -database "${DATABASE_URL}" down 

check_migration_version:
	migrate -path ./cmd/migrate/migrations -database "${DATABASE_URL}" version


force_migration:
	migrate -path ./cmd/migrate/migrations -database "${DATABASE_URL}" force "$(version)"