# Build stage
FROM docker.io/library/golang:1.19.0-alpine3.16 as  builder
LABEL stage=builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM docker.io/library/alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .

CMD ["/app/main"]
