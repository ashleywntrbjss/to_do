# ToDo App

## Running

Please use the main.go file in the `runall` dir to run all aspects of the application.

When running the application you will need to provide the `-r` parameter to select a repo type

### Repo Types
 - In Memory (`memory`)
 - Postgres (`sql`)

If using a postgres repo type, you will also need to provide a connection string with the `-cs` parameter (also include `?sslmode=disable` at the end)

Example

`-r=sql -cs="postgres://postgres:1234@localhost:5432/postgres?sslmode=disable"`