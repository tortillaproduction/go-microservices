.PHONY: run-command run-command-local run-query run-query-local

# --- Run command-service ---
## Docker
run-command:
	cd apps/command-service && go run main.go

## Local
run-command-local:
	cd apps/command-service && MYSQL_HOST=localhost go run main.go

# --- Run query-service ---
## Docker
run-query:
	cd apps/query-service && go run main.go

## Local
run-query-local:
	cd apps/query-service && MYSQL_HOST=localhost go run main.go
