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
COPY --from=builder ./build/main /main
# COPY --from=super_graph_builder /config ./config
# COPY --from=node_builder /build ./frontend/whisperspot-web/build
# Build the application

# Move to /dist directory as the place for resulting binary folder
WORKDIR /

# Copy binary from build to main folder
# RUN cp /build/main .
RUN chmod +x ./main

# Export necessary port
EXPOSE 5000

# Command to run when starting the container
CMD ["/main"]