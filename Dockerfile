FROM golang:1.12.4-alpine as builder

# Build project
WORKDIR /go/src/github.com/batazor/go-logger
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch:latest

RUN addgroup -S 997 && adduser -S -g 997 997
USER 997

WORKDIR /app/
COPY --from=builder /go/src/github.com/batazor/go-logger/app .
CMD ["./app"]
