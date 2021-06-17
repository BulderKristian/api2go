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
    -s asyncapi
```

## Supported Features
### AsyncApi
#### Components
##### Schema
|Feature|Flag
|:---:|:---:|
|title|-|
|type|-|
|required|-|
|multipleOf|-|
|maximum|-|
|exclusiveMaximum|-|
|minimum|-|
|exclusiveMinimum|-|
|maxLength|-|
|minLength|-|
|pattern|-|
|maxItems|-|
|minItems|-|
|uniqueItems|-|
|maxProperties|-|
|minProperties|-|
|enum|Done|
|const|-|
|examples|-|
|if/then/else|-|
|readOnly|-|
|writeOnly|-|
|properties|-|
|patternProperties|-|
|additionalProperties|-|
|additionalItems|-|
|items|-|
|propertyNames|-|
|contains|-|
|allOf|-|
|oneOf|-|
|anyOf|-|
|not|-|
                                