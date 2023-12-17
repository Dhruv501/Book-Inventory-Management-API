# Zopsmart-mini-Project
# Gofr Book Inventory API

[Include a logo or banner image if you have one]

A powerful API for managing books with CRUD operations and MySQL database integration.

---

## Table of Contents

1. [Introduction](#introduction)
2. [Features](#features)
3. [Installation](#installation)
4. [Usage](#usage)
5. [Database Integration](#database-integration)
6. [Unit Testing](#unit-testing)
7. [Postman Collection](#postman-collection)

---

## Introduction

Gofr Book Inventory API is a feature-rich solution for managing book data with CRUD operations. It provides a seamless integration with a MySQL database, ensuring data persistence and reliability. This API is designed to be easy to use and flexible, making it suitable for a variety of applications.

---

## Features

- **CRUD Operations:** Build APIs for create, read, update, and delete operations for all entities.
- **Database Integration:** The API is seamlessly integrated with a MySQL database for data persistence.
- **Unit Testing:** Thorough unit testing ensures the reliability and correctness of the API.
- **Easy Installation:** Get started quickly with a simple installation process.
- **Docker Support:** Easily set up a MySQL database through a freely available Docker image.

---

## Installation

To install the dependencies, run the following command:

```bash
go get ./...
# Run MySQL Docker container
docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 3306:3306 -d mysql:8.0.30
# to run the server
go run main.go 
```

## Usage

To interact with the Book Inventory API, you can make HTTP requests to the specified endpoints. Below are some examples of how to perform basic operations using Postman.

### Create a new book

Send a `POST` request to the `/books` endpoint with the book details in the request body.

1. Open Postman.
2. Set the request type to `POST`.
3. Enter the URL: `http://localhost:8000/books`.
4. In the request body, provide the book details in JSON format:

```json
{
  "title": "Example Book",
  "author": "John Doe",
  "genre": "Fiction",
  "date": "2023-01-01",
  "price": 20,
  "description": "A fascinating book."
}
```
### Retrieve all books

Send a `GET` request to the `/books` endpoint to retrieve a list of all books.

1. Open Postman.
2. Set the request type to `GET`.
3. Enter the URL: `http://localhost:8000/books`.
4. In the request body, provide the book details in JSON format:

### Update a book

Send a `PUT` request to the `/books/{id}` endpoint with the book details in the request body.

1. Open Postman.
2. Set the request type to `PUT`.
3. Enter the URL: `http://localhost:8000/books/1`(replace `1`with the actual book ID).
4. In the request body, provide the book details in JSON format:

```json
{
  "title": "Updated part",
  "author": "Updated part",
  "genre": "Updated part",
  "date": "Updated part",
  "price": "Updated part",
  "description": "Updated part"
}
```
### Delete a book

Send a `DELETE` request to the `/books/{id}` endpoint to delete a specific book

1. Open Postman.
2. Set the request type to `POST`.
3. Enter the URL: `http://localhost:8000/books/1`(replace `1`with the actual book ID).
4. Click "Send" to delete the book.



