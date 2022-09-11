FROM golang:1.17-alpine AS base
WORKDIR /base

RUN apk add --no-cache ca-certificates && update-ca-certificates

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM base as bundled-stage

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/migrate cmd/migrate/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/server cmd/server/main.go

FROM busybox
WORKDIR /app
COPY --from=bundled-stage /base/bin/ /app/
COPY --from=bundled-stage /base/app/config/ /app/config/
COPY --from=bundled-stage /base/internal/migrations/ /app/internal/migrations

ARG RUN_ENV
ENV VIDYARTHA_ENV=${RUN_ENV}

EXPOSE 3000