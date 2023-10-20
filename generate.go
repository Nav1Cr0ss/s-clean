package main

////go:generate mockery --all --keeptree --case=underscore --dir ./internal --output ./mocks
////go:generate go run gen/gen.go
//go:generate oapi-codegen --config design/models.cfg.yaml docs/openapi.yaml
//go:generate oapi-codegen --config design/server.cfg.yaml docs/openapi.yaml
//go:generate go run github.com/ogen-go/ogen/cmd/ogen --target ./api -package api --clean ./docs/openapi.yaml
