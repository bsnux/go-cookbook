FROM golang:1.18 as build

WORKDIR /app
COPY . .

RUN go mod download \
  && CGO_ENABLED=0 go build -o /app/bin/app

# Using distroless from Google https://github.com/GoogleContainerTools/distroless/
FROM gcr.io/distroless/static-debian11

COPY --from=build /app/bin/app /
CMD ["/app"]
