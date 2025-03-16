#!/bin/bash

# 첫 번째 인자를 변수에 저장
MIGRATION_VERSION=$1

# 마이그레이션 
migrate -database "mysql://root:1234@tcp(localhost:3306)/wdwb" -path ./internal/migrations force $MIGRATION_VERSION