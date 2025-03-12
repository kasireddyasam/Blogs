# part1 go build application 
FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# building binary source code and build the binary
COPY . . 
RUN go build -o blog-app ./cmd/api/main.go

#part 2 run the application
# alpile to run binary code
FROM  alpine:latest   
WORKDIR /root/
COPY --from=builder /app/blog-app .

# copy build binary from thebuilder stage
EXPOSE 8080
CMD [ "./blog-app" ]

