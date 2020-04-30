FROM golang:latest as build

WORKDIR /go/src/github.com/tomocy/go-echoargs

COPY . .

WORKDIR /go/src/github.com/tomocy/go-echoargs/cmd/echoargs

ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0

RUN go build -o app .

FROM alpine:latest

COPY --from=build /go/src/github.com/tomocy/go-echoargs/cmd/echoargs/app /

ENTRYPOINT ["/app"]