FROM golang:1.14.2-stretch
WORKDIR /go/src/github.com/heptiolabs/gangway

RUN go get -u github.com/mjibson/esc/...
COPY . .
ADD https://raw.githubusercontent.com/PrismJS/prism/v1.16.0/components/prism-bash.js assets/
ADD https://raw.githubusercontent.com/PrismJS/prism/v1.16.0/prism.js assets/
ADD https://raw.githubusercontent.com/PrismJS/prism/v1.16.0/themes/prism.css assets/
ADD https://raw.githubusercontent.com/PrismJS/prism/v1.16.0/components/prism-powershell.js assets/
ADD https://raw.githubusercontent.com/Dogfalo/materialize/v1-dev/dist/css/materialize.min.css assets/
ADD https://raw.githubusercontent.com/Dogfalo/materialize/v1-dev/dist/js/materialize.min.js assets/
RUN esc -o cmd/gangway/bindata.go templates/ assets/

ENV GO111MODULE on
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux go install -ldflags="-w -s" -v github.com/heptiolabs/gangway/...

FROM debian:9.12-slim
RUN apt-get update && apt-get install -y ca-certificates
USER 1001:1001
COPY --from=0 /go/bin/gangway /bin/gangway
