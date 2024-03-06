# Regenerate changes made to yml, ORM models and resolvers.
gql-gen:
	go run github.com/99designs/gqlgen generate

# Updates go.mod and tidy's.
mod-update:
	go get -u
	go get -u ./...
	go mod tidy

# Starts docker container(s).
containers:
	docker compose up

clean-containers:
	docker compose down
	