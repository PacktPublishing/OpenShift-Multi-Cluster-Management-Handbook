FROM golang:alpine AS builder
WORKDIR /app
COPY ./articles.go ./go.mod ./
RUN go get articles
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM registry.access.redhat.com/ubi8/ubi-micro
WORKDIR /
COPY --from=builder /app/main /
EXPOSE 8080
CMD ["./main"]
