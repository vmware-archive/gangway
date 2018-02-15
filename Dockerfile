FROM golang:1.9
WORKDIR /go/src/github.com/heptio/gangway

RUN go get github.com/golang/dep/cmd/dep github.com/jteeuwen/go-bindata/...
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -v -vendor-only

COPY cmd cmd
COPY templates templates
RUN go-bindata -o cmd/gangway/bindata.go templates/ && \
    CGO_ENABLED=0 GOOS=linux go install -ldflags="-w -s" -v github.com/heptio/gangway/...

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/bin/gangway /bin/gangway