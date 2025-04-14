FROM golang:1.18-alpine

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go mod tidy
RUN go build -o orderapp

CMD ["/app/orderapp"]
