# Prerequisite

- [Go version `1.22+`](https://go.dev/doc/install)
- [Golang-migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#unversioned)
- [Swaggo CLI](https://github.com/swaggo/swag?tab=readme-ov-file#getting-started)
- [Air CLI (optional)](https://github.com/cosmtrek/air?tab=readme-ov-file#installation)

# Getting started

First thing first, in order to create an API, you will need to follow those steps:

1. Create migration files (up/down) with this migrate CLI command:

```sh
migrate create -ext sql -dir migrations/ -seq my_table_name
```

2. Write entities and interfaces.
3. Implement them and make sure to test the API before commit it and there you go. It is that simple :P.

And remember to create an .env file with those variables:

```
PORT=5000
CLIENT_ORIGIN_URL=http://localhost:5173
AUTH0_AUDIENCE=your-auth0-audience
AUTH0_DOMAIN=your-auth0-domain
DB_URL=libsql://[your-db-name].turso.io?authToken=[your-auth-token]

```

To run the server, you can use air or go CLI:

```sh
air
# or
go run . serve
```
