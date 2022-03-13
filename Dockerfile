# Builder
FROM golang:1.17-alpine AS builder

WORKDIR /usr/src

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build

# Runner
FROM alpine AS runner

WORKDIR /usr/app

COPY --from=builder /usr/src/carving .

EXPOSE 3000

CMD ["./carving"]