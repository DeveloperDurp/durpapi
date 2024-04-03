FROM registry.internal.durp.info/golang:1.22-alpine

RUN chmod -R o=,g=rwX /go
RUN mkdir /app

RUN adduser \
--disabled-password \
--gecos "" \
--home "/nonexistent" \
--shell "/sbin/nologin" \
--no-create-home \
--uid "10001" \
"durp"

WORKDIR /app
COPY ./output/* .
RUN chown -R durp /app

USER durp
EXPOSE 8080
CMD ["./main"]
