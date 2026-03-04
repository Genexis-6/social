include .env

create_migration:
	migrate create -ext .sql -dir "${MIGRATION_PATH}" "$(name)"


push_migration:
	migrate -path "${MIGRATION_PATH}" -database "${DATABASE_URL}" up


reverse_migration:
	migrate -path "${MIGRATION_PATH}" -database "${DATABASE_URL}" down 

check_migration_version:
	migrate -path "${MIGRATION_PATH}" -database "${DATABASE_URL}" version


force_migration:
	migrate -path "${MIGRATION_PATH}" -database "${DATABASE_URL}" force "$(version)"