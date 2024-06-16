# oapi-codegen-share

This is a simple introduction of how to generate code from OpenAPI(Swagger) specifications using the `oapi-codegen`
tool.

## What?

`oapi-codegen` is a command-line tool and library to convert OpenAPI specifications to Go code

* server-side boilerplate
* client API boilerplate
* HTTP models.

## Why?

### Pros

1. less hand-coding => less errors
2. template-based generation => consistent code
3. save time => save money

### Cons

1. coding style relies on templates
2. not all features are supported
3. might increase complexity, you can't get tailored templates by default

## How?

```shell
# binary install
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

oapi-codegen --help
Usage of oapi-codegen:
  -alias-types
    	Alias type declarations of possible.
  -config string
    	A YAML config file that controls oapi-codegen behavior.
  -exclude-operation-ids string
    	Exclude operations with the given operation-ids. Comma-separated list of operation-ids.
  -exclude-schemas string
    	A comma separated list of schemas which must be excluded from generation.
  -exclude-tags string
    	Exclude operations that are tagged with the given tags. Comma-separated list of tags.
  -generate string
    	Comma-separated list of code to generate; valid options: "types", "client", "chi-server", "server", "gin", "gorilla", "spec", "skip-fmt", "skip-prune", "fiber", "iris", "std-http". (default "types,client,server,spec")
  -h	Same as -help.
  -help
    	Show this help and exit.
  -import-mapping string
    	A dict from the external reference to golang package path.
  -include-operation-ids string
    	Only include operations with the given operation-ids. Comma-separated list of operation-ids.
  -include-tags string
    	Only include operations with the given tags. Comma-separated list of tags.
  -initialism-overrides
    	Use initialism overrides.
  -o string
    	Where to output generated code, stdout is default.
  -old-config-style
    	Whether to use the older style config file format.
  -output-config
    	When true, outputs a configuration file for oapi-codegen using current settings.
  -package string
    	The package name for generated code.
  -response-type-suffix string
    	The suffix used for responses types.
  -templates string
    	Path to directory containing user templates.
  -version
    	When specified, print version and exit.
```

### Server

```shell
oapi-codegen -generate std-http -package gen -o gen/std-server.gen.go api.yaml
```

### Client

```shell
oapi-codegen -generate client -package gen -o gen/client.gen.go api.yaml
```

### Models

```shell
oapi-codegen -generate types -package gen -o gen/models.gen.go api.yaml
```

## Customization

