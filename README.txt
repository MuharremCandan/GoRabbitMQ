
RabbitMQ and Golang Example

This repository demonstrates a simple interaction with RabbitMQ using Golang. 
It covers the process of setting up RabbitMQ within a Docker container, creating 
a queue and an exchange, binding the queue to the exchange, and sending/consuming a message.

Prerequisites
Docker installed on your machine

Usage
Run RabbitMQ Docker Container:
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.12-management

Clone this repository:

git clone https://github.com/MuharremCandan/GoRabbitMQ.git
cd rabbitmqgo

Install Dependencies:
go mod tidy

https://www.rabbitmq.com/tutorials/tutorial-one-go.html
