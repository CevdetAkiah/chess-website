FROM golang:1.21 AS builder-go

WORKDIR /workdir
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN ls -altr
RUN CGO_ENABLED=0 GOOS=linux go build -o /app cmd/main.go

EXPOSE 8080

# TODO: Learn about multi-stage builds :)
CMD ["/app"]
