FROM golang:latest as builder
WORKDIR /build
COPY . /build
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cmd cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/cmd .
EXPOSE 6000
CMD ["/app/cmd"]
