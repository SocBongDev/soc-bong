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