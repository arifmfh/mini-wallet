# Mini Wallet 
This Project is fully written in GO and using Redis as database. This code structure implement Clean Architecture Pattern, separated in 3-layer (delivery/presentation, usecase, repository)

And still need some Improvement below:
- Implement FIFO Queue (Apache Kafka/rabbitMQ/Google Pubsub) for procesing transaction
- Add Unit Testing 
- Migrate to SQL database if necessary

## Prerequisites
- Golang
- Redis

## How to Run
1. If you are not using redis locally with default port please update `REDIS_ADDR` in `.env` 
2. `go run main.go`


