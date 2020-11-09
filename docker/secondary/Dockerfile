FROM shanduur/pluggabl-lib:latest

EXPOSE 4004

ENV CONFIG=
ENV PG_DATABASE=

COPY ./files/ /

LABEL maintainer="mateusz.urbanek.98@gmail.com" \
    type="secondary-node" \
    version="0.0.1-init"

ENTRYPOINT [ "/pluggabl/secondary.run" ]