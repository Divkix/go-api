FROM gcr.io/distroless/static
COPY go-api /
CMD ["/go-api"]
