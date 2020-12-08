FROM shanduur/pluggabl-lib:latest

EXPOSE 4004

ENV CONFIG="/etc/pluggabl/config.json"
ENV PG_DATABASE=

COPY ./files/ /

LABEL maintainer="mateusz.urbanek.98@gmail.com" \
    type="secondary-node" \
    version="1.0.3"

CMD [ "/pluggabl/server.run" ]