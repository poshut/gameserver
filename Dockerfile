FROM golang:latest

LABEL author="Richard Wohlbold" version="0.1"

ENV config "/config/server.conf"
ENV port "8080"

WORKDIR "/app"
ADD ./cmd/gameserver/ /app/
RUN go build -o gameserver .

# Add java runtime:
# RUN apk add openjdk8

CMD ["/app/gameserver", "--config", "/config/server.conf", "--port", "8080"]
