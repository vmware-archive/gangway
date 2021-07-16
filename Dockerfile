FROM golang:1.15 as build

WORKDIR /app

COPY go.mod go.sum templates ./

RUN go mod download

RUN go get -u github.com/mjibson/esc/...

COPY . .

RUN  esc -o cmd/gangway/bindata.go templates/

RUN go build -o gangway cmd/gangway/main.go cmd/gangway/handlers.go cmd/gangway/bindata.go

FROM debian:9.12-slim
RUN apt-get update && apt-get install -y ca-certificates
USER 1001:1001
COPY --from=build /app/gangway /bin/gangway
