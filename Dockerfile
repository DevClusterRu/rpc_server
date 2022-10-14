FROM golang:1.18-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main

FROM alpine:3
COPY --from=builder main /bin/main

EXPOSE 4466

ENTRYPOINT ["/bin/main"]