# Build Stage: Build bot using the alpine image, also install doppler in it
FROM golang:1.19.4-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=`go env GOHOSTOS` GOARCH=`go env GOHOSTARCH` go build -o out/go-api

# Run Stage: Run bot using the bot and doppler binary copied from build stage
FROM gcr.io/distroless/static:latest
COPY --from=builder /app/out/go-api /
CMD ["/go-api"]
