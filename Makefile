COMPOSE := docker-compose -f build/docker-compose.local.yaml -p assignment

pg:
	${COMPOSE} up -d pg

setup: pg build-image

build-image:
	docker build -f build/api.Dockerfile -t api-image .

run:
	${COMPOSE} run --rm --service-ports -w /api api sh -c "go run cmd/main.go"

pg-redo:
	${COMPOSE} run --rm pg-migrate sh -c 'migrate -path /migration -database $$PG_URL up'

api-gen-models:
	${COMPOSE} run --rm --service-ports -w /api api sh -c 'sqlboiler  --wipe psql && GOFLAGS="-mod=vendor"'

api-vendors:
	${COMPOSE} run --rm --service-ports -w /api api sh -c 'go mod vendor & go mod tidy'

api-gen-mocks:
	${COMPOSE} run  --name mockery --rm -w /api --entrypoint '' mockery /bin/sh -c "\
		mockery --dir internal/controller --all --recursive --inpackage && \
		mockery --dir internal/repository --all --recursive --inpackage"
