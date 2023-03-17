FROM golang:1.20-alpine3.17 as builder

RUN mkdir /build

ADD . /build/

WORKDIR /build

RUN go build -o main .

# Stage 2

FROM alpine

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY . /app

COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 3000

CMD ["./main"]