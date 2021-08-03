FROM golang:latest as builder

RUN mkdir /build

ADD . /build/

WORKDIR /build

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 8080

CMD ["./main", "server"]