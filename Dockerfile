FROM golang:1.12
WORKDIR /go/src/github.com/heptiolabs/gangway

RUN go get -u github.com/mjibson/esc/...
COPY . .
RUN esc -o cmd/gangway/bindata.go templates/

ENV GO111MODULE on
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux go install -ldflags="-w -s" -v github.com/heptiolabs/gangway/...

FROM alpine:latest
RUN apk --no-cache add ca-certificates
USER 1001:1001
COPY --from=0 /go/bin/gangway /bin/gangway
