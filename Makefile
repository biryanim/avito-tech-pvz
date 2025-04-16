include example.env

LOCAL_BIN = $(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.20.0

migration-up:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} up -v

migration-down:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} down -v

migration-status:
	$(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} status -v