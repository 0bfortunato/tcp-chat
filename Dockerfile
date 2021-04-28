FROM golang:1.16.3-alpine3.13

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/tcp-chat

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the package
RUN make build

# Run the executable
CMD ["tcp-chat"]

# Expose ports to external world
EXPOSE 8080/tcp
