# Base image
FROM golang:1.13-alpine3.11

# All subsequent actions should 
# be taking from this directory
WORKDIR /go/src/app

# Copy app source to image 
# filesystem (i.e. to WORKDIR)
COPY . .

# Install git
RUN apk add --no-cache git

# Make ssh key available
RUN mkdir -p /root/.ssh/
ARG SSH_PRIVATE_KEY
RUN echo "${SSH_PRIVATE_KEY}" > /root/.ssh/id_rsa
RUN cat /root/.ssh/id_rsa

# Install dependencies and
# also install application
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
