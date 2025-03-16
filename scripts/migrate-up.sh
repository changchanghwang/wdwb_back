#!/bin/bash

# Run the migrate command in local environment
migrate -database "mysql://root:1234@tcp(localhost:3306)/wdwb" -path ./internal/migrations up