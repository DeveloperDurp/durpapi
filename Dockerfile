FROM registry.durp.info/golang:1.20-alpine

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
COPY . .
RUN chown -R durp /app

USER durp

ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin

EXPOSE 8080

# Run the application
CMD ["./main"]
