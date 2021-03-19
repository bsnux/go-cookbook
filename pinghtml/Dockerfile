FROM golang:1.14.1 as builder
ENV NAME "pinghtml"
WORKDIR /opt
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -o bin/$NAME main.go

FROM alpine:3.11.5
RUN apk --no-cache add ca-certificates
WORKDIR /opt
COPY --from=builder /opt/bin/$NAME .
CMD ["/opt/pinghtml"]
