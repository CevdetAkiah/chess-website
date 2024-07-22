FROM golang:1.22.1 AS builder-go

WORKDIR /workdir
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app cmd/main.go


EXPOSE 8080



# TODO: Learn about multi-stage builds :)
COPY --from=builder-go /workdir/database/postgres/migrations /database/postgres/migrations
CMD ["/app"]
