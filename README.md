# Language Management

Welcome to the Language Store RESTful API Service, a robust and scalable platform built on the reliable and efficient Echo framework. Our API offers a range of features that enable you to easily manage programming languages, from adding new languages and updating existing ones to retrieving detailed information about each language.

We take pride in implementing the best software development practices, including the Clean Architecture principles coined by the renowned software engineer, Robert C. Martin (also known as Uncle Bob). Our architecture ensures that our codebase is organized, testable, and maintainable, allowing us to continuously provide a high-quality service to our clients.

## Getting Started

### Prerequisites

- [Go 1.19.3](https://go.dev/dl/)

### Installation

- Clone the git repository:

```
$ git clone https://github.com/Assyatier21/tnb-technical-test.git
$ cd tnb-technical-test
```

- Install Dependencies

```
$ go mod tidy
```

### Running

```
$ go run cmd/main.go
```

or simply running this command

```
$ make run
```

### Features

This service has the following API endpoints:

- `` `GET /languages` ``: get a list of language
- `` `POST /language` ``: insert a new language
- `` `GET /language/{id}` ``: get details of a language with the specified ID
- `` `PUT /language/{id}` ``: update a language with the specified ID
- `` `DELETE /language/{id}` ``: delete a language with the specified ID

## Install Local Sonarqube

please follow this [tutorial](https://techblost.com/how-to-setup-sonarqube-locally-on-mac/) as well.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/Assyatier21/simple-tnb-technical-test/blob/master/LICENSE) file for details.
