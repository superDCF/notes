FROM golang:1.15.8-alpine3.13 as build

# Set the Current Working Directory inside the container
WORKDIR /root

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY server.go .

RUN go build -o /server

FROM alpine:3.13
COPY --from=build /server /server

# Run the executable
CMD ["./server"]