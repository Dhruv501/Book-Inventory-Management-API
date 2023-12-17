# Zopsmart-mini-Project
# Gofr Book Inventory API

[Include a logo or banner image if you have one

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

## Database Integration

The Gofr Book Inventory API is seamlessly integrated with a MySQL database for efficient data persistence. Follow the steps below to set up and interact with the database.

### MySQL Database Configuration

1. Ensure you have Docker installed on your machine.

2. Run the following command to start a MySQL Docker container:

   ```bash
   docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=test_db -p 3306:3306 -d mysql:8.0.30
   # To create the necessary table for storing books, execute the following command:  
   docker exec -it gofr-mysql mysql -uroot -proot123 test_db -e "CREATE TABLE books (id INT AUTO_INCREMENT PRIMARY KEY, title VARCHAR(255) NOT NULL, author VARCHAR(255), genre VARCHAR(255), date DATE, price INT,    description TEXT);"

### Database Schema

The database schema includes a table named `books` with the following columns:

- **id:** Integer, Auto-incremented primary key.
- **title:** String, Not null.
- **author:** String.
- **genre:** String.
- **date:** Date.
- **price:** Integer.
- **description:** Text.

## Unit Testing
![Screenshot 2023-12-18 030359](https://github.com/Dhruv501/Zopsmart-mini-Project/assets/75206417/c8b32720-1e77-4dfc-ac24-11d95b454c95)

## Postman Collection
![Screenshot 2023-12-18 013533](https://github.com/Dhruv501/Zopsmart-mini-Project/assets/75206417/011ddfc6-6def-4e3e-abf9-6bd7598be439)
![Screenshot 2023-12-18 013523](https://github.com/Dhruv501/Zopsmart-mini-Project/assets/75206417/86110e2a-5a5c-4f44-889c-24ac600fade7)
![Screenshot 2023-12-18 013508](https://github.com/Dhruv501/Zopsmart-mini-Project/assets/75206417/5cb99399-afa5-4680-8871-fb3f4ceadef5)

