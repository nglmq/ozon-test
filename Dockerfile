FROM golang:1.22

WORKDIR ./app

COPY . .

RUN go mod tidy && \
    go build -o /build ./cmd/shortener && \
    go clean -cache -modcache

EXPOSE 8080
EXPOSE 5050

CMD ["/build", "-f", "true", "-p", ":8080", "-b", "http://localhost:8080", "-g", "false"]
