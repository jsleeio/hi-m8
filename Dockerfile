# Start by building the application.
FROM golang:1.10 as build

WORKDIR /go/src/github.com/jsleeio/hi-m8
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build

# Now copy it into our base image.
FROM scratch
USER 1000
COPY --from=build /go/src/github.com/jsleeio/hi-m8/hi-m8 /hi-m8
ENTRYPOINT ["/hi-m8"]
