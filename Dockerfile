FROM alpine AS base
ENV GIN_MODE=release
WORKDIR /app
EXPOSE 8080

FROM golang:1.19 AS builder
WORKDIR /src
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .

RUN go build -o ./publish/main main.go


FROM base AS final
COPY --from=builder /src/publish/main .

ENTRYPOINT ["/app/main"]



