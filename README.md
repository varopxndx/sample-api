# Sample API

Sample API uses dependency injection, unit tests using mocks, and Clean Architecture.

## Architecture

[Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) consists on separate the logic into different packages. The main proposal is the _Usecase_ package.
The Usecase package is used to handle the _bussines logic_, and call other packages according to the different flows.

Clean Architecture proposes a lot of packages, this Sample API contains the principal ones:

* Model
* Service
* Usecase
* Controller
* Router

### Model

Model package contains all the structures to handle the JSON data and responses.
Is the only package that could be used in all the layers.

### Service

Service package (also called Repository) contains all the logic to handle the operations. This logic is separated into several functions, which allows to test in an easier way, have decoupled logic, and maintain the code in a better way.
If some logic is added in the future, a new function will be created and the Usecase will call it.

### Usecase

Usecase package contains the business logic, according to the flows, it will call the specific function in the Service package to get or store data.

### Controller

Controller package contains the handlers, it calls the usecase flows, returns the responses and handles the http errors.

### Router

Router package contains the endpoints, each enpoint is linked to a handler.

## Requirements

* [Go](https://golang.org/doc/install) - Version _go1.15.2_ or above.

## Setup

1. Clone the [repository](https://github.com/varopxndx/sample-api.git)

1. Open the terminal and go to the project path

1. In the root level, run the following commands:

    ```sh
        make build

        ./sample-api
    ```

1. The Samplle API will be running on:

    ```sh
        http://localhost:8080/v1
    ```

## Running the tests

The _Service_ layer does not contain tests due to is mocked data.

* Run all the tests with the following command:

    ```sh
        make unit-test
    ```

## Endpoints

Sample API contains two GET endpoints:

```sh
    http://localhost:8080/v1/ping
    http://localhost:8080/v1/sample
```
