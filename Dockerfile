# Base image
FROM golang:1.15-alpine3.12 as builder

# All subsequent actions should 
# be taking from this directory
WORKDIR /go/src/app

# Copy app source to image 
# filesystem (i.e. to WORKDIR)
COPY . .

# Install git
RUN apk add --no-cache git

# Install dependencies and
# also install application
RUN go get -d -v ./...
RUN go build .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /go/bin
COPY --from=builder /go/src/app/dockerized_service .
CMD ["/go/bin/dockerized_service"]
