FROM golang:latest AS builder
WORKDIR /app
COPY ./articles.go ./go.mod ./
RUN go get articles
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

#FROM registry.access.redhat.com/ubi8/ubi-micro
FROM registry.access.redhat.com/ubi8/ubi-minimal
WORKDIR /
COPY --from=builder /app/main /
RUN microdnf install curl -y
EXPOSE 8080
CMD ["./main"]
