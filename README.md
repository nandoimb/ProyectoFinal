# E-commerce Web Application

This project is a simple e-commerce web application built using Go for the backend and HTML, CSS, and JavaScript for the frontend. The application allows users to manage products and orders, including creating, updating, and deleting them.

## Table of Contents
- [E-commerce Web Application](#e-commerce-web-application)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Technologies Used](#technologies-used)
  - [Setup Instructions](#setup-instructions)
    - [Prerequisites](#prerequisites)
    - [Backend Setup](#backend-setup)
  - [Running the Go Server](#running-the-go-server)
  - [Frontend Setup](#frontend-setup)
  - [API Endpoints](#api-endpoints)
    - [Products](#products)
    - [Orders](#orders)
  - [Frontend Interface](#frontend-interface)
    - [Forms](#forms)

## Features
- Create, read, update, and delete products
- Create, read, update, and delete orders
- Simple web interface to interact with the API

## Technologies Used
- Backend: Go, Gin, GORM, SQLite
- Frontend: HTML, CSS, JavaScript (Fetch API)

## Setup Instructions

### Prerequisites
- Go 1.16 or later
- Node.js (for running a local server if needed)

### Backend Setup
1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/ecommerce-web-app.git
   cd ecommerce-web-app

## Installation Instructions

To install the project dependencies, follow these steps:

1. Open a terminal and navigate to the project directory
2. Run the following command to install the required dependencies:

```sh
go mod tidy
```

## Running the Go Server

To run the Go server, execute the following steps:

1. Open a terminal and navigate to the project directory
2. Run the following command to start the server:

```sh
go run main.go
```

## Frontend Setup

To serve the `index.html` file using a local server, please follow these steps:

1. Use a local server to serve the `index.html` file. You can use the VS Code Live Server extension or any other static file server.
2. Run the following command to serve the file:

```sh
npx serve .
```

3. Open your browser and navigate to `http://localhost:8080` for the backend and `http://localhost:5000` (or the port your server is running on) for the frontend.

## API Endpoints

### Products

- **GET /products**

  Retrieves a list of all products.

  - Response: 200 OK with JSON array of products.

- **GET /products/{id}**

  Retrieves a specific product by ID.

  - Response: 200 OK with JSON of the product or 404 Not Found if not found.

- **POST /products**

  Creates a new product.

  - Request Body: JSON object with name and price.
  - Response: 201 Created with JSON of the created product or 400 Bad Request if validation fails.

- **PUT /products/{id}**

  Updates an existing product by ID.

  - Request Body: JSON object with name and price.
  - Response: 204 No Content or 400 Bad Request if validation fails.

- **DELETE /products/{id}**

  Deletes a specific product by ID.

  - Response: 204 No Content or 400 Bad Request if validation fails.

### Orders

- **GET /orders**

  Retrieves a list of all orders.

  - Response: 200 OK with JSON array of orders.

- **GET /orders/{id}**

  Retrieves a specific order by ID.

  - Response: 200 OK with JSON of the order or 404 Not Found if not found.

- **POST /orders**

  Creates a new order.

  - Request Body: JSON object with userID.
  - Response: 201 Created with JSON of the created order or 400 Bad Request if validation fails.

- **PUT /orders/{id}**

  Updates an existing order by ID.

  - Request Body: JSON object with userID.
  - Response: 204 No Content or 400 Bad Request if validation fails.

- **DELETE /orders/{id}**

  Deletes a specific order by ID.

  - Response: 204 No Content or 400 Bad Request if validation fails.

- **DELETE /orders/{orderId}/products/{productId}**

  Removes a specific product from an order.

  - Response: 204 No Content or 404 Not Found if not found.

## Frontend Interface

The frontend interface provides a simple way to interact with the API endpoints. It allows users to:

- View a list of all products and orders.
- Create new products and orders.
- Update existing products and orders.
- Delete products and orders.

### Forms

- **Create Product**: A form to input the product's name and price.
- **Update Product**: A form to input the product ID, new name, and new price.
- **Delete Product**: A form to input the product ID to delete.
- **Create Order**: A form to input the user ID.
- **Update Order**: A form to input