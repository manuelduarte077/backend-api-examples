# Shop Phoenix API REST

This is a REST API project built with Elixir and Phoenix web framework. It allows users to perform CRUD operations on products in a shop database. The project uses Postgres as the database and Docker Compose to run the app and the database in containers.

## Features

- Create, read, update, and delete products
- Validate product attributes and handle errors
- Use Ecto for database interactions and migrations
- Use Phoenix generators and scaffolding
- Use Docker Compose to run the app and the database

## Getting Started

To run this project, you need to have Docker and Docker Compose installed on your machine. Then, follow these steps:

- Clone the repository from GitHub
- Navigate to the project directory and run `docker-compose up`
- Visit `localhost:4000` from your browser to see the app running
- Use a tool like Postman or curl to test the API endpoints

## API Endpoints

The project exposes the following endpoints:

- `GET /api/products` - List all products
- `GET /api/products/:id` - Show a single product by id
- `POST /api/products` - Create a new product
- `PUT /api/products/:id` - Update an existing product by id
- `DELETE /api/products/:id` - Delete an existing product by id

The product attributes are:

- `name` - The name of the product (required, string)
- `description` - The description of the product (optional, string)
- `price` - The price of the product in cents (required, integer)
- `stock` - The number of units available in stock (required, integer)
