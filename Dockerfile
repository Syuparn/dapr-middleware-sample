FROM golang:1.19

WORKDIR /work

COPY . .
RUN go build

CMD ["./dapr-middleware-sample"]
