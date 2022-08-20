#!/bin/sh
export ENV=development
export PORT=8081
# postgres://postgres:123456@127.0.0.1:5432/dummy
export POSTGRES_DATABASE_URL=postgres://postgres:postgres@localhost:5432/timer?sslmode=disable
