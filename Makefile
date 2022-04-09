run:
	@cd src/cmd && go run ./*.go

build:
	@cd $(PWD)/src/cmd && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o skm .

migration:
	@cd $(PWD)/src/migrations && read -p "Enter migration name: " migration_name; \
	goose create $${migration_name} sql

migrate:
	@bash $(PWD)/scripts/migration.sh

gen_docs:
	@cp src/cmd/main.go src/
	@cd src && swag init
	@rm -rf src/main.go
	@rm -rf src/cmd/docs
	@mv src/docs src/cmd
gen_mock_repo:
	cd $(PWD)/src && mockery --case=underscore --dir=$(PWD)/src/repositories --output $(PWD)/src/mocks/repositories --all ; \

test:
	@cd src && go test --cover -p 1 -v -failfast -coverprofile=src.cov `go list ./...`
	@cd src && cat src.cov | grep -v "fake" > fine.cov
	@cd src && go tool cover -func=fine.cov
