# Microservice API Gateway with Golang, Gin, and gRPC
## Overview

This project is a Microservice API Gateway implemented in Golang using the Gin framework and gRPC. The API Gateway serves as the entry point and orchestrator for various APIs in the system. It provides a unified interface for client applications to interact with multiple backend services, handling authentication and authorization, implementing rate limiting and request validation, and managing load balancing and fault tolerance.
Project Structure

## The project is structured as follows:

    cmd/main.go: The entry point of the API Gateway.
    pkg/auth/pb/auth.proto: Protocol Buffers definition for the authentication service.
    pkg/product/pb/product.proto: Protocol Buffers definition for the product service.
    pkg/order/pb/order.proto: Protocol Buffers definition for the order service.
    pkg/cart/pb/cart.proto: Protocol Buffers definition for the cart service.
    pkg/auth: Package containing authentication-related functionality.
    pkg/product: Package containing product-related functionality.
    pkg/order: Package containing order-related functionality.
    pkg/cart: Package containing cart-related functionality.

## Makefile

### The project includes a Makefile with the following commands:

    make proto: Compiles the Protocol Buffers files.
    make server: Runs the API Gateway server.

### Running the Server

To run the API Gateway server, use the following command:

bash

make server

##Routes

### User Routes

    POST /auth/register: Registers a new user.
    POST /auth/registerValidation: Validates user registration through OTP.
    POST /auth/login: Logs in a user.
    POST /auth/resetPassword: Initiates the password reset process.
    POST /auth/resetPasswordValidation: Validates the password reset through OTP.

### Authentication Routes

    GET /authRoutes/userProfile: Retrieves user profile information.
    POST /authRoutes/addAddress: Adds a new address.

### Cart Routes

    POST /cart/addCart: Adds an item to the cart.
    PATCH /cart/removeCart: Removes an item from the cart.
    GET /cart/displayCart: Displays the contents of the cart.

### Order Routes

    POST /order: Creates a new order.

### Product Routes

    POST /product: Creates a new product.
    GET /product/:id: Retrieves information for a specific product.
    GET /product/products: Lists all products.
    GET /product/search: Searches for products.
    GET /product/sortByPrice: Sorts products by price.

## Middleware

Authentication middleware is applied to routes requiring authentication to ensure proper access control.

## Dependencies

    Golang: v1.16 or later
    Gin: HTTP web framework for Golang
    gRPC: RPC framework for communication between services
    Protocol Buffers: Used for defining service interfaces

## Contributing

Feel free to contribute to the development of this project by opening issues or submitting pull requests.
