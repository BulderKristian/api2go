# API 2 GO

## Disclaimer
>This is a personal project, and I work on it when I feel like it.
I would advice that it is not used for Production.

## Motivation
Create a faster, leaner more focused golang api model generator for `swagger`/`openapi`/`asyncapi` for golang

## Goal 
Generating golang models from any valid openapi3 specification

## Docker Example
### Generating
```shell
docker run \
  -v "${PWD}":/local \
  -u ${UID} codedevstem/api2go:latest generate \
    -i /local/path/to/spec.yaml \
    -o /local/output/folder \
    -s openapi
```