DB_URL=postgresql://nav1cr0ss:0608@localhost:5436/s-auth?sslmode=disable

postgres:
	docker run --name postgres-NearBy -p 5436:5432 -e POSTGRES_USER=nav1cr0ss -e POSTGRES_PASSWORD=0608 -d postgres:14-alpine

createdb:
	docker exec -it postgres-NearBy createdb --username=nav1cr0ss s-auth

dropdb:
	docker exec -it postgres-NearBy dropdb --username=nav1cr0ss s-auth

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc --file db/sqlcConfig.yaml generate

openapi:
	swagger-cli bundle ./design/_docs.yaml --outfile api/openapi.yaml --type yaml
api:
	oapi-codegen --config design/models.cfg.yaml docs/openapi.yaml && oapi-codegen --config design/server.cfg.yaml docs/openapi.yaml
run:



.PHONY: api