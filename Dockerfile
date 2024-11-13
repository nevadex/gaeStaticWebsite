FROM golang:1.23.2-bookworm AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY ./*.go ./

RUN go build -v -o /usr/local/bin/app ./...

FROM debian:bookworm-slim AS runner

COPY --from=builder /usr/local/bin/app /usr/local/bin/app
COPY ./www ./www

CMD ["app"]