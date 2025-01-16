gen:
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=api/cfg.yaml api/api.yaml

PG="postgres://wbuser:wbpassword@localhost:5435/postgres?sslmode=disable"

goose-up:
	@goose -dir ./migrations postgres $(PG) up

goose-down:
	@goose -dir ./migrations postgres $(PG) down