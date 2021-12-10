# API 2 GO

## Motivation
Create a faster, leaner more focused golang api model generator for `swagger`/`openapi`/`asyncapi` for golang

## Docker Example

```shell
docker run \
  -v "${PWD}":/local \
  -u ${UID} codedevstem/api2go:latest generate \
    -i /local/path/to/spec.yaml \
    -o /local/output/folder \
    -s openapi
```