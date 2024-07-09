# Sales Report Microservice

## Overview

Welcome to the Sales Report Microservice repository! This microservice is designed to generate detailed sales reports based on sales history. Built with Golang, it ensures high performance and reliability, making it suitable for integration with larger sales applications or systems.

## Features

- **Sales Reports**: Generate comprehensive sales reports.
- **High Performance**: Built with Golang for efficient processing.
- **Configuration via Makefile**: Simplifies setup and deployment processes.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following:

- [Golang](https://golang.org/dl/) installed on your machine.
- A configured sales history data source (database, API, etc.).
- Installed Make, btw, here's how to install **Make** for [Windows](https://stackoverflow.com/questions/32127524/how-to-install-and-use-make-in-windows), for [MacOS](https://formulae.brew.sh/formula/make).

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/zmskv/sales-app.git
    ```

2. Navigate to the project directory:

    ```sh
    cd sales-app
    ```

3. Install all dependencies:

    ```sh
    go mod tidy
    ```

4. Set up environment variables by creating a `.env` file in the root directory, also create migrations.
    ```sh
    make migrate-up
    ```

### Usage

1. Run the application:

    ```sh
    make app-start
    ```

2. The application will start and listen for requests to generate sales reports.


3. To do this, you can use the [swagger](http://localhost:8000/swagger/index.html)
### Makefile Commands

The included Makefile provides several commands to simplify common tasks:

- `make build`: Compiles the application.
- `make app-start`: Runs sales-app.
- `make migrate-up`: Create migrations.
- `make migrate-down`: Drop migrations.

## Contact

For any questions or issues, please open an issue in this repository or contact the maintainer.

---