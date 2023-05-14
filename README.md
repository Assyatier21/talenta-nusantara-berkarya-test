# Ralali Cake Store

Welcome to the Cake Store RESTful API Service, a robust and scalable platform built on the reliable and efficient Echo framework. Our API offers a range of features that enable you to easily manage your cakes, from adding new items and updating existing ones to retrieving detailed information about each product.

We take pride in implementing the best software development practices, including the Clean Architecture principles coined by the renowned software engineer, Robert C. Martin (also known as Uncle Bob). Our architecture ensures that our codebase is organized, testable, and maintainable, allowing us to continuously provide a high-quality service to our clients.

## Getting Started

### Prerequisites

- [Go 1.19.3](https://go.dev/dl/)
- [MySQL](https://www.mysql.com/downloads/)

### Installation

- Clone the git repository:

```
$ git clone https://github.com/Assyatier21/ralali-technical-test.git
$ cd ralali-technical-test
```

- Install Dependencies

```
$ go mod tidy
```

- Create `config` folder in root path, then create a file `connection.go` in that folder containing this following code:

```
package config

const (
	Host     = "127.0.0.1"
	Port     = "3306"
	User     = "YOUR_USERNAME_HERE"
	Password = "YOUR_PASSWORD_HERE"
	Database = "YOUR_DATABASE_HERE"
)
```

alternatively, we can just run this following command using makefile:

```
$ make config
```

### Running

```
$ go run cmd/main.go
```

or simply running this command

```
$ make run
```

### Running on Docker

```
$ docker build -t ralali-cakes .
$ docker run -p 8800:8800 ralali-cakes

```

### Migration Database

```
$ make migrate

```

### Features

This service has the following API endpoints:

- `` `GET /cakes` ``: get a list of cakes
- `` `POST /cakes` ``: insert a new cake
- `` `GET /cakes/{id}` ``: get details of a cake with the specified ID
- `` `PUT /cakes/{id}` ``: update a cake with the specified ID
- `` `DELETE /cakes/{id}` ``: delete a cake with the specified ID

We can test the endpoint using the postman collection in `tnb-technical-test/tools`.

### Testing

```
$ go test -v -coverprofile coverage.out ./...
```

## Install Local Sonarqube

please follow this [tutorial](https://techblost.com/how-to-setup-sonarqube-locally-on-mac/) as well.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/Assyatier21/simple-tnb-technical-test/blob/master/LICENSE) file for details.
