FROM golang:latest as builder
WORKDIR /build
COPY . /build
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o wow-api cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/wow-api .
EXPOSE 6000
CMD ["/app/wow-api"]
