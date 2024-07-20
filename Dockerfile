FROM golang:1.21-alpine3.18 as builder
RUN apk update && apk upgrade --available && sync
WORKDIR /app
COPY . .
# Instalar ufw
RUN apt-get update && apt-get install -y ufw

# Habilitar ufw y permitir puertos
RUN ufw allow 8080/tcp && \
    ufw allow 80/tcp && \
    ufw enable
RUN CGO_ENABLED=0 go build -o /app/fsb -ldflags="-w -s" ./cmd/fsb

FROM scratch
COPY --from=builder /app/fsb /app/fsb
EXPOSE ${PORT}
ENTRYPOINT ["/app/fsb", "run"]
