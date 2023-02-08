FROM golang:latest as builder
LABEL author="ductnn"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /your_bin


FROM alpine:latest
LABEL author="ductnn"

RUN apk --no-cache add ca-certificates \
    && rm -rf /var/cache/apk/*

WORKDIR /root/
COPY --from=builder /app/your_bin .

# Expose port
EXPOSE 8090

CMD ["./your_bin"]
