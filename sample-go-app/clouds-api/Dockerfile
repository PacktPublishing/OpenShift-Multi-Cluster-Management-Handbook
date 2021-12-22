FROM golang:alpine AS builder
WORKDIR /app
COPY ./clouds.go ./clouds_test.go ./go.mod ./
RUN go get clouds
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM registry.access.redhat.com/ubi8/ubi-micro
WORKDIR /
COPY --from=builder /app/main /
EXPOSE 8080
CMD ["./main"]
