#!/bin/bash

NUM_DOWN=${1:-1}

# Run the migrate command in local environment
migrate -database "mysql://root:1234@tcp(localhost:3306)/wdwb" -path ./internal/migrations down $NUM_DOWN