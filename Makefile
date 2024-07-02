codegen:
	oapi-codegen -generate std-http -package api -o api/std-server.gen.go api.yaml
	oapi-codegen -generate client -package api -o api/client.gen.go api.yaml
	oapi-codegen -generate types -package api -o api/models.gen.go api.yaml

generate:
	go generate ./...

run:
	go run main.go

client:
	go run client.go

tests:
	./tests.sh
