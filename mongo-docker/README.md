# Human Resource Management System using Gofiberv2 x mongoDB x Docker

All components are docker-based

## Start development with Docker Compose

### To start the application

Step 1: start mongodb and mongo-express

    docker-compose -f docker-compose.yaml up

_You can access the mongo-express under [localhost:8081](localhost:8081) from your browser_

Step 2: Go to Postman to check application REST API

    http://localhost:8080/employee/

#### Employee struct

    ID     string 
 Name   string
 Salary float64
 Age    float64

#### Get all amployees

    GET http://localhost:8080/employee/

#### Get employee by ID

    GET http://localhost:8080/employee/{id}

#### Create new amployees

    POST http://localhost:8080/employee/

#### Update employee by ID

    PUT http://localhost:8080/employee/{id}

#### Delete employee by ID

    DELETE http://localhost:8080/employee/{id}

### To build a docker image from the application

    docker build -t mygo:1.0 .       

The dot "." at the end of the command denotes location of the Dockerfile.
