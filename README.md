# Go Fiber Rest API


Welcome to the **Go Fiber CRUD** GitHub repository! This project is a simple example of a CRUD (Create, Read, Update, Delete) application built using the Go programming language and the Fiber web framework. It demonstrates how to perform basic CRUD operations on a database using the Fiber framework.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
- [Usage](#usage)
- [License](#license)

## Features

- Create, Read, Update, and Delete operations on a database.
- Restful API design using the Fiber web framework.
- Demonstrates best practices for building web applications with Go and Fiber.
- Utilizes Postgres as the database for storage.

## Getting Started

Follow these instructions to get the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (1.20 or higher) should be installed. You can download it from the official [Go website](https://golang.org/dl/).
- Git should be installed for cloning the repository.
- Docker should be installed to run the application in a container.

### Installation

1. Clone the repository using Git:
    ```
    git clone https://github.com/ayoam/go-fiber-crud.git
    ```
2. Navigate to the project directory:
    ```
    cd go-fiber-crud
    ```
## Usage

1. Run the application:
    ```
    docker compose up
    ```

2. The server will start locally at `http://localhost:8000`. You can access the API using tools like [Postman](https://www.postman.com/).

3. The available endpoints are as follows:

- `GET /api/v1/notes`: Retrieve a list of all notes.
- `GET /api/v1/notes/:id`: Retrieve a specific note by ID.
- `POST /api/v1/notes`: Create a new note (provide JSON payload).
- `PUT /api/v1/notes/:id`: Update a note by ID (provide JSON payload).
- `DELETE /api/v1/notes/:id`: Delete a note by ID.


## License

This project is licensed under the [MIT License](LICENSE), which means it is open-source and you can use it for personal and commercial purposes.

---

Happy coding!

Project maintained by [ayoam](https://github.com/yourusername)




