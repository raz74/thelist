# Microservice Development with Go and Gin

This project demonstrates the implementation of two microservices using the Go and Gin framework. The microservices handle user management, process events, and store notifications in a database.

## Project Overview

- **User Management Microservice**: Creates a user and publishes an event upon successful creation.
- **Notification Microservice**: Listens for the event, processes it, and stores a notification in a PostgreSQL database.

## Features

- **Hexagonal Architecture**: Applied hexagonal architecture principles for modularity, testability, and separation of concerns.
- **Robust Error Handling**: Descriptive error messages and appropriate HTTP status codes for better debugging.
- **Concurrent Processing**: Utilized Goroutines and channels for improved performance.
- **Dockerized**: Easy deployment and scalability using Docker.

## Technologies Used

- **Programming Language**: Go
- **Framework**: Gin
- **Database**: PostgreSQL
- **Concurrency**: Goroutines and Channels
- **Containerization**: Docker
