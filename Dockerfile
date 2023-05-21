# Use golang:1.17 as the base image
FROM registry.durp.info/golang:1.20

ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin

# Set the working directory inside the container
WORKDIR /app

# Copy the Go project files into the container
COPY ./output/main .

# Expose the port the application listens on
EXPOSE 8080

# Run the application
CMD ["./main"]
