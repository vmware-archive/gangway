FROM golang:1.12
WORKDIR /go/src/github.com/heptiolabs/gangway

RUN go get github.com/golang/dep/cmd/dep
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -v -vendor-only

RUN go get -u github.com/mjibson/esc/...
COPY cmd cmd
COPY templates templates
COPY internal internal
RUN esc -o cmd/gangway/bindata.go templates/

RUN CGO_ENABLED=0 GOOS=linux go install -ldflags="-w -s" -v github.com/heptiolabs/gangway/...

FROM alpine:latest
RUN apk --no-cache add ca-certificates
USER 1001:1001
COPY --from=0 /go/bin/gangway /bin/gangway
