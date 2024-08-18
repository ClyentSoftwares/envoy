FROM golang:1.23.0-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o envoy .

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /app/envoy .
CMD ["./envoy"]
