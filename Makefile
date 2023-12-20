postgres:
	docker run --name my-project -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123 -d postgres:latest

createdb:
	docker exec -it my-project createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it my-project dropdb simple_bank

migrate-up:
	migrate -path pkg/db/migration -database "postgresql://root:123@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrate-dowm:
	migrate -path pkg/db/migration -database "postgresql://root:123@localhost:5432/simple_bank?sslmode=disable" -verbose down
