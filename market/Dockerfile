FROM alpine:latest

COPY market /app/
COPY templates /app/templates
COPY sql/* /app/sql/
COPY files/* /app/files/
COPY config.json /app/
COPY certificate.* /app/

EXPOSE 5500
WORKDIR /app
ENTRYPOINT ["./market"]
