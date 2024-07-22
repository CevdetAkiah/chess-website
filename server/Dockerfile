FROM golang:1.22.1 AS builder-go

WORKDIR /workdir
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /app cmd/main.go

EXPOSE 8080

# Copy the binary and migrations from the build stage
COPY --from=builder-go /app /app
COPY --from=builder-go /workdir/database/postgres/migrations /database/postgres/migrations

# TODO: Learn about multi-stage builds :)
CMD ["/app"]
