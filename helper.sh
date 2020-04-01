#!/bin/bash

# docker container for postgres
docker run --name pg_db -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres
