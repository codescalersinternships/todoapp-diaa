# Todo App
This is a Todo application built using Vue.js with TypeScript. The client interacts with the server-side implemented in Golang. The application allows users to manage a list of Todo items, including adding, updating, and deleting items.


![ui](https://github.com/codescalersinternships/todoapp-diaa/assets/77173710/a43ffea8-7b3a-4d19-abad-ea23be96bf97)

## Prerequisites

Node.js - The client is built with Vue.js and requires Node.js to run.
Package Manager (npm or yarn) - To install dependencies and run the application.
Golang - The server is built using Golang.


## Start the App
1. Clone the repository:
```bash
git clone <repository-url>
```

2. To run the client
- 1. Navigate to client
```bash
cd client
```
- 2. Install Dependencies
```bash
npm install
# or
yarn install
```
- 3. Add .env file
check .env.example

- 4. Run the client
```bash
npm run dev
```
- 5. Now you can access the [client](http://localhost:5173)

3. To run the server
- 1. Navigate to server
```bash
cd ..
cd server
```
- 2. Install Dependencies
```bash
go mod download

```
- 3. Build
```bash
go build -o bin/app cmd/server.go
```
- 4. Run the server
    ```bash
    ./bin/app -p <port> -db<db path>
    ```
    Note: **port is 5000 by default**

## Run using docker-compose
```bash
docker-compose up -d
```


## Features

1. Add Todo: Allows users to add new Todo items to the list.

2. Delete Todo: Allows users to delete a Todo item from the list. 

3. Edit Todo: Allows users to edit the name of a Todo item.

4. Mark Todo as Finished/Unfinished: Allows users to mark a Todo item as finished or unfinished.

## API Endpoints
The server exposes the following API endpoints:

- POST /todo: Adds a new Todo item to the database. Expects a JSON payload representing the Todo name.

- GET /todo: Retrieves the list of all Todo items from the database.

- PUT /todo: Updates an existing Todo item in the database. Expects a JSON payload representing the updated Todo item.

- DELETE /todo/{id}: Deletes a Todo item from the database based on its ID.

- GET /todo/{id}: Retrieves a single Todo item from the database based on its ID.

## API Documentation with Swagger
The server API is documented using Swagger, making it easy to understand and interact with the endpoints. You can access the Swagger UI at:
```bash
http://localhost:<port>/docs/index.html
```

