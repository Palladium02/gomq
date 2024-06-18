FROM golang:1.22 as builder

WORKDIR /broker

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o broker ./cmd/broker

FROM debian:bookworm-slim

COPY --from=builder /broker /broker

CMD ["/broker/broker"]
EXPOSE 1883