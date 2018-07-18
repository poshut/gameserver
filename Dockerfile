FROM archlinux/base

LABEL author="Richard Wohlbold" version="0.1"

RUN pacman -Syy --noconfirm go

WORKDIR "/app"
ADD ./cmd/gameserver/ /app/
RUN go build -o gameserver .

# Add java runtime:
# RUN pacman -S jdk10-openjdk

CMD ["/app/gameserver", "--config", "/config/server.conf", "--port", "8080"]
