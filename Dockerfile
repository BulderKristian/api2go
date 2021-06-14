FROM golang:1.16-alpine AS build
WORKDIR /src/
COPY main.go go.* /src/
COPY src /src/src
COPY cmd /src/cmd
COPY templates /src/templates
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/api2go

FROM scratch
COPY --from=build /bin/api2go /bin/api2go
COPY --from=build /src/templates/ /templates
ENTRYPOINT ["/bin/api2go"]