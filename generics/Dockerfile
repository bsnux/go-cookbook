FROM golang:1.19 as builder
ENV NAME "generics"
WORKDIR /opt
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -o bin/$NAME main.go

FROM scratch
WORKDIR /opt
COPY --from=builder /opt/bin/$NAME .
ENTRYPOINT ["/opt/generics"]
