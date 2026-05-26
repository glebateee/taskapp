include .env
export

env-up:
	docker compose up -d taskapp-pg

env-down:
	docker compose down taskapp-pg

env-cleanup:
	@read -p "Do you really wanna delete all files? [y/N]: " ans; \
	case "$$ans" in \
		y|Y|yes|YES) \
			docker compose down -v && \
			rm -rv out/pgdata && \
			echo "All done!" ;; \
		*) \
			echo "Cleanup cancelled" ;; \
	esac

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Missing parameter seq, usage: make migrate-create seq=init"; \
		exit 1; \
	fi; \
	docker compose run --rm pg-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)" 

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Missing action parameter, usage: make migrate-action action=up"; \
		exit 1; \
	fi; \
	docker compose run --rm pg-migrate \
		-path /migrations \
		-database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@taskapp-pg:5432/$(POSTGRES_DB)?sslmode=disable" \
		$(action)

env-port-forward:
	@ docker compose up -d port-forwarder

env-port-close:
	@ docker compose down port-forwarder


env-logs:
	@ docker compose logs -f pg

env-shell:
	@ docker exec -it taskapp-pg-1 psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)

taskapp-run:
	@ export LOGGER_FOLDER=./out/logs &&  export POSTGRES_HOST=localhost && go mod tidy && 	go run cmd/taskapp/main.go