# ###########
# Migrations
# ###########

# Creates new migration up/down files in the 'migration' folder with the provided name.
.PHONY: migration/create
migration/create:
	@read -p "Enter migration name: " migration; \
    ${GOPATH}/bin/migrate create -ext sql -dir ./gateways/postgres/migrations -seq $$migration

# Execute the migrations up to the most recent one. Needs the following environment variables:
# MIGRATION_DATABASE_HOST: database url
# MIGRATION_DATABASE_USER: database user
# MIGRATION_DATABASE_PASS: database password
# MIGRATION_DATABASE_NAME: database name
.PHONY: migration/up
migration/up:
	dsn="postgres://user:password@localhost/transactionsdb?sslmode=disable"; \
	${GOPATH}/bin/migrate -source file://gateways/postgres/migrations -database $$dsn up

# Rollback the migrations up to the oldest one. Needs the following environment variables:
# MIGRATION_DATABASE_HOST: database url
# MIGRATION_DATABASE_USER: database user
# MIGRATION_DATABASE_PASS: database password
# MIGRATION_DATABASE_NAME: database name
.PHONY: migration/down
migration/down:
	dsn="postgres://user:password@localhost/transactionsdb?sslmode=disable"; \
	${GOPATH}/bin/migrate -source file://migrations -database $$dsn down
