# Verify My Age Test - Users Api

This repository keeps the API application to perform a User CRUD operations.

## API Characteristics

* Stack used: Go v1.17, Fiber, Gorm, Go-Jwt, Go-Bcrypt, Docker
* Requests implemented: GET, POST, PUT, DELETE
<!-- * Deployed at AWS --- [http://ec2-54-80-141-128.compute-1.amazonaws.com:8080](http://ec2-54-80-141-128.compute-1.amazonaws.com:8080) -->
* Public Routes: "/api/users/register", "/api/users/auth", "/swagger"
* Private Routes: "/api/users" (GET, PUT, DELETE), "/api/users/{id} (GET)"

## API Documentation

<!-- API documentation can be found at:
* Swagger UI Documentation: [http://ec2-54-80-141-128.compute-1.amazonaws.com:8080/swagger-ui.html](http://ec2-54-80-141-128.compute-1.amazonaws.com:8080/swagger-ui.html).
* Swagger JSON Documentation: [http://ec2-54-80-141-128.compute-1.amazonaws.com:8080/v2/api-docs](http://ec2-54-80-141-128.compute-1.amazonaws.com:8080/v2/api-docs). -->
* Endpoints Description:

|   Route      |  HTTP Verb     |  Description  |  Body Request |  Example Response |  Status Code |
| :---         | :---           | :---          | :---          | :---              | :---         |
| /api/users/register |   POST     | Register a new user based in informed params  |  ````{"name":"John Doe","age":30,"email":"john@doe.com","password":"123456","repeat_password":"123456","address":"Jd Road, 1234"}```` | ````{"id":15,"name":"John Doe","age":30,"email":"john@doe.com","address":"Jd Road, 1234"}``` | 201 |
| /api/users/auth   |     POST     | Authenticates an user based on informed credentials, returning a JWT token  | ````{"email":"john@doe.com","password":"123456"}```` | ````{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzA4Njk2NzIsImlzcyI6IjE1In0.ZrpH4tzt2qdDtTFynj3ez2rIl8KM9cvmkI5AO1JOKps"}```` | 200 |
| /swagger  |   GET   | Get Swagger UI Documentation HTML  | N/A | HTML file | 200 |
| /api/users  |   GET   | Retrieves a list of users paginated. Can filter result by name param | N/A | ````{"data":[{"id":15,"name":"John Doe","age":30,"email":"john@doe.com","address":"Jd Road, 1234"}],"page":{"page":1,"page_total":1,"total_results":1,"last_page":1}}```` | 200 |
| /api/users/{id}  |   GET   | Retrieves a user by Id passed as URL param  | N/A | ````{"id":15,"name":"John Doe","age":30,"email":"john@doe.com","address":"Jd Road, 1234"}```` | 200 |
| /api/users  |   PUT   | Update user info based on the Id informed   | ````{"name":"John Doe Updated","age":32,"address":"Jd Road, 4321"}```` | ````{"id":15,"name":"John Doe Updated","age":32,"email":"john@doe.com","address":"Jd Road, 4321"}```` | 200 |
| /api/users  |   DELETE   | Delete user that made the request based on informed token | N/A | N/A | 204 |

## Running Locally

To run this API locally, please follow the below steps:

* Clone this repository using with **git clone** command in your Terminal
* Start a local MySql database. Just run the following command inside the root path of this project in your Terminal:
```bash
  sudo docker-compose -f "docker-compose.yml" up -d --build
```
* In case you don't have Docker installed in your machine, please follow the installation steps below:
    * [Install Docker](https://docs.docker.com/get-docker/)
    * [Install Docker Compose](https://docs.docker.com/compose/install/)
* Install dependencies with `go mod tidy`
* Run project with `go run .`

## Authors & Version Control

API developed by **Henrique Guazzelli Mendes - [https://github.com/henriquegmendes](https://github.com/henriquegmendes)** - *VMA-GO-Users-API App Version 1.0* - **Published in Sep-04th of 2021**