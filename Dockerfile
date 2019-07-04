FROM golang:1.12-alpine as builder

LABEL author="Richard Wohlbold" version="0.1"

WORKDIR "/go/src/app"
ENV CGO_ENABLED=0 GOOS=linux
COPY ./cmd .
RUN go get -d -v ./...
RUN go build -o /go/bin/gameserver ./gameserver/

FROM scratch
COPY --from=builder /go/bin/gameserver /
COPY ./config /config

# CMD ["/gameserver", "--config", "/config/server.conf", "--port", "8080"]
EXPOSE 8080

CMD ["/gameserver", "--config", "/config/server.conf", "--port", "8080", "--output", "/dev/null"]


