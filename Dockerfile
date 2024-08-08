FROM golang:1.22.5-alpine3.9 as builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go test -v ./...

RUN go build -o /out/app

FROM alpine:3.9

COPY --from=builder /out/app /app

CMD ["/app"]
