FROM golang:1.17-alpine AS base

WORKDIR /base

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM base as bundled-stage

COPY . .

RUN go build cmd/server/main.go

EXPOSE 3000

CMD ["./main"]
