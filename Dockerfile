# Start from golang base image
FROM golang:1.21.0-alpine as builder

# Set the current working directory
WORKDIR /thelist

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine

COPY --from=builder /thelist/default.env .env
COPY --from=builder /thelist/main .

# Expose ports to the outside world
EXPOSE 8080

#Command to run the executable
ENTRYPOINT ["./main"]
