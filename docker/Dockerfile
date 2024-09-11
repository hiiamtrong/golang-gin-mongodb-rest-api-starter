FROM golang:1.21.9-alpine AS builder
WORKDIR /src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /bin/app/server ./cmd/server

FROM alpine:latest AS runner
WORKDIR /bin/app
COPY --from=builder /bin/app .
COPY --from=builder /src/app/config.yml ./config.yml

ARG SERVER_PORT=3000
ENV SERVER_PORT=${SERVER_PORT}

EXPOSE ${SERVER_PORT}

CMD ["./server"]
