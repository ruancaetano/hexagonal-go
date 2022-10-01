# Hexagonal Go

Exploration of hexagonal architecture concepts using Go.

The application simulates a simple management of products, where the `application` directory contains all the specifications and rules of the domain, isolated through interfaces, so that it was possible to implement two types of clients/adapters (CLI and an http server) without these implementations to the domain.

## how to play with it?

<br />

### CLI mode

Run the following command to start cli mode. You can use `-h` optios to list all available command options

```go
go run main.go cli -h
```

Example: Get a product by id, `-a` defines the desired action and `-i` is the id of the product.

```go
go run main.go cli -a=get -i=blablabla
```
<br />

### HTTP mode

Run the following command to start http server on port 8080. After that you can send http requests to localhost:8080 


Example: Get product by id

```
curl http://localhost:8080/products/blablabla
```
