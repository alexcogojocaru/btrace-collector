FROM golang:1.17-alpine

WORKDIR /app

COPY . .
RUN go build -o /golang-collector-btrace

CMD [ "/golang-collector-btrace" ]
