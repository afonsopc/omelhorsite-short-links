FROM golang:latest as builder

WORKDIR /app

COPY short-links/go.mod short-links/go.sum ./

RUN go mod download

COPY short-links/ .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM alpine:latest as runner

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]