FROM golang:1.20.1-alpine3.17
RUN mkdir /app
WORKDIR /app

# Set the mongoURI environment variable
ENV CONFIG_MONGO_URI=mongodb://root:example@mongo:27017/

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o main .
EXPOSE 8080
CMD ["/app/main"]