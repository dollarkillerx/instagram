package main

//go:generate go mod tidy
//go:generate rm -f ./internal/generated/resolver.go
//go:generate go run github.com/99designs/gqlgen generate --config ./configs/gqlgen.yml --verbose
