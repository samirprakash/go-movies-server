# Go Watch a Movie

[![Build and Test](https://github.com/samirprakash/go-movies-server/actions/workflows/build.yaml/badge.svg?branch=main)](https://github.com/samirprakash/go-movies-server/actions/workflows/build.yaml)

Backend server for go-movies application

#### Database setup

- Start docker engine
- Run `make postgres` to download and start a postgres DB instance locally
- Execute `make createdb` to cretae a database
- Execute `make populatedb` to create tables and populate data in the database
- If there is an error in any of the steps, run `make cleanpostgres` to clean up your postgres environment and try again

#### Start application

- Run `make server`
