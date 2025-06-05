FROM golang:1.23.4
WORKDIR /app
COPY . .
RUN go install github.com/air-verse/air@v1.52.3
RUN go install github.com/go-delve/delve/cmd/dlv@latest
CMD ["air", "-c", ".air.toml"]