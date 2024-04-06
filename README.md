# Basic HelloWorld! From NATS JetStream in Golang

## Run Instructions
1. cd into the project directory and run `go mod tidy`
2. my docker configuration for NATS JetStream,
- Run following,
- create nats-network named as `nats-net`
  ```
  docker network create nats-net
  ```
- then run the image,
  ```
   docker run --name nats-server --network nats --rm -p 4222:4222 -p 8222:8222 nats:latest -js
  ```
3. finally after all the setup run,
   ```
   go run main.go
   ```
