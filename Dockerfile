FROM alpine

LABEL author="Richard Wohlbold" version="0.1"

ENV dir "/data"
ENV config "${dir}/server.conf"
ENV port "8080"

WORKDIR "${dir}"
ADD ./gameserver /

# Add java runtime:
# RUN apk add openjdk8

CMD ["/gameserver", "--config", "${config}", "--port", "${port}"]