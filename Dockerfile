FROM golang:1.22.5-alpine as builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go test -v ./...

RUN go build -o /bin/app cmd/datacollector/main.go

FROM alpine:3.9

COPY --from=builder /bin/app /bin/app

CMD ["/bin/app"]
