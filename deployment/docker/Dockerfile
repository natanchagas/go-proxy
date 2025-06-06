FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /app/proxy-server ./cmd/main.go

FROM alpine:3.18
RUN apk add --no-cache ca-certificates openssl

RUN openssl req -x509 -out /localhost.crt -keyout /localhost.key \
    -newkey rsa:2048 -nodes -sha256 \
    -subj '/CN=localhost' -extensions EXT -config <( \
    printf "[dn]\nCN=localhost\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS:localhost\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")

# Trust the self-signed certificate
RUN cp /localhost.crt /usr/local/share/ca-certificates/localhost.crt && update-ca-certificates

COPY --from=builder /app/proxy-server /proxy-server
ENTRYPOINT ["/proxy-server"]