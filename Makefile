postgres:
	docker run --name v001_onelab -e POSTGRES_PASSWORD=password  -p 5436:5432 -v pgdata:/var/lib/postgresql/data --rm -d postgres

migrate-up:
	migrate -database "postgres://postgres:password@localhost:5436/postgres?sslmode=disable" -path ./pkg/database/postgres/schema up

migrate-down:
	migrate -database "postgres://postgres:password@localhost:5436/postgres?sslmode=disable" -path ./pkg/database/postgres/schema down

create-migration:
	migrate create -ext sql -dir ./pkg/database/postgres/schema -seq init

run:
	docker run -dp 8080:8080 --name sad_gould  --rm app

stop:
	docker stop sad_gould