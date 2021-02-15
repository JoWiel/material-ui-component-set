FROM alpine/git as intermediate

WORKDIR /public/repositories
RUN git clone https://github.com/bettyblocks/material-ui-component-set.git

FROM alpine:3.9 AS node_builder

ENV METEOR_VERSION=1.8.1
ENV METEOR_ALLOW_SUPERUSER true
ENV NODE_VERSION 8.15
RUN apk add --no-cache --repository=http://dl-cdn.alpinelinux.org/alpine/v3.8/main/ nodejs=8.14.0-r0 npm 

WORKDIR /

RUN npm install -g @betty-blocks/cli@latest
COPY generator/scripts* /generator/scripts/

WORKDIR /public
RUN  mkdir generated && mkdir build && mkdir uploaded


FROM golang:alpine AS builder

RUN apk add --no-cache go
# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    ENV=production

ADD . /build
# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go build -o main .

FROM alpine:latest
# Copy the code into the container
COPY --from=intermediate / ./
COPY --from=node_builder / ./
COPY --from=builder ./build/main /main

# Move to /dist directory as the place for resulting binary folder
WORKDIR /

# Copy binary from build to main folder
# RUN cp /build/main .
RUN chmod +x ./main

# Export necessary port
EXPOSE 5000

# Command to run when starting the container
CMD ["/main"]