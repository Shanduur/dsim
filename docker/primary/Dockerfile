FROM ubuntu:latest

EXPOSE 4004

ENV CONFIG="/etc/pluggabl/config.json"
ENV PG_DATABASE=

COPY ./files/ /
RUN chmod -R 777 /etc/pluggabl

LABEL maintainer="mateusz.urbanek.98@gmail.com" \
    type="primary-node" \
    version="1.0.1"

CMD [ "/pluggabl/server.run" ]