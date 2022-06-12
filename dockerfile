FROM golang:1.18 as builder

##
## Build
##
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -v -o tiktok_demo ./cmd/

##
## Build
##
FROM alpine:3.14

WORKDIR /root/
RUN apk add --no-cache ffmpeg
COPY --from=builder /app/tiktok_demo ./tiktok_demo

EXPOSE 8080

ENTRYPOINT ["./tiktok_demo"]
