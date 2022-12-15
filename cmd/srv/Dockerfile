# Use the offical golang image to create a binary.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.19.3-buster as builder

# Create and change to the app directory.
WORKDIR /app

# Copy local code to the container image.
COPY . ./
# Add vendor dependencies (committed or not)
# I typically commit the vendor dependencies as it
# makes the final build more reproducible and less
# dependant on dependency managers.
ADD ./vendor ./vendor
COPY go.mod ./
# COPY go.sum ./

# Build the binary.
# RUN go build -mod=readonly -v -o supermarket
RUN go build -o srv

# RUN go test -v -cover ./...

FROM gcr.io/distroless/base
########################

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/srv /app/srv

# Run the web service on container startup.
CMD ["/app/srv"]

# FROM gcr.io/distroless/base-debian10
# WORKDIR /usr/src/app
# COPY server .
# CMD [ "/usr/src/app/server" ]