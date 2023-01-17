# simple-banking-service

## Project dependencies

- [Postgres Docker Image](https://hub.docker.com/_/postgres) - docker utility container for running postgresql locally
- [Golang migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) - For creating and running migrations in Golang.
- [Sqlc](https://sqlc.dev/) - For compiling sql queries to type-safe Golang code.
- [pq](https://pkg.go.dev/github.com/lib/pq) - Go Postgres driver for database/sql
- [Testify](https://github.com/stretchr/testify) - A toolkit with common assertions and mocks that plays nicely with the standard library.

## Run Commands

A makefile has been included and contains important commands that'll come in handy ex for running migrations, compiling queries using sqlc, running tests

```
$ make test // to run tests
```

## Built with

- [Golang](https://go.dev)
- Postgresql
- Docker
- Kubernetes