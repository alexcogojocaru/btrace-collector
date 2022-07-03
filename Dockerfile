FROM golang:1.17-alpine

WORKDIR /data

COPY . .
RUN go build -o /data/bin-collector

CMD [ "/data/bin-collector" ]
