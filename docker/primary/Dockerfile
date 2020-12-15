FROM ubuntu:latest

RUN apt-get update -y && apt-get install tzdata -y

EXPOSE 4004

ENV CONFIG="/etc/pluggabl/config.json"
ENV PG_DATABASE=

COPY ./files/ /
RUN chmod -R 777 /etc/pluggabl

LABEL maintainer="mateusz.urbanek.98@gmail.com" \
    type="primary-node" \
    version="1.0.2"

CMD [ "/pluggabl/server.run" ]