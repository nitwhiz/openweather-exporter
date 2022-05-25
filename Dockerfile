FROM golang:1.18.1-buster as builder

WORKDIR /app

COPY ./ /app

RUN CGO_ENABLED=0 go build -o ./build/exporter ./cmd/exporter

FROM alpine:3.16.0

COPY --from=builder /app/build/exporter /exporter

EXPOSE 9924

CMD [ "/exporter" ]
