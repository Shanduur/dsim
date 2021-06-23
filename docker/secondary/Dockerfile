FROM shanduur/dsip-lib:latest

EXPOSE 4004

ENV CONFIG="/etc/dsip/config.json"
ENV PG_DATABASE=

COPY ./files/ /

LABEL maintainer="mateusz.urbanek.98@gmail.com" \
    type="secondary-node" \
    version="1.0.4"

CMD [ "/dsip/server.run" ]