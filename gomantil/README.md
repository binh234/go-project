# Serverless service with AWS Lambda and Go using Mantil

This is a simple implementation of Movie CRUD serverless application with additional authentication using JWT. Code is written in Go and deployed on AWS Lambda using Mantil. Movie data is stored in AWS DynamoDB.

[Mantil](https://github.com/mantil-io/mantil) is a modern open-source framework for writing serverless apps in Go. It allows you to quickly create and deploy applications that use AWS Lambda over a command line interface.

## Mantil documentation

[Documentation](https://docs.mantil.com)

## Deploying the application

Note: If this is the first time you are using Mantil you will need to install Mantil Node on your AWS account. For detailed instructions please follow the [setup guide](https://docs.mantil.com/aws_detailed_setup/aws_credentials)

```bash
mantil aws install
```

Then you can proceed with application deployment.

```bash
mantil deploy
```

This command will create a new stage for your project with the default name `development` and deploy it to your node.

Now you can output the stage endpoint with `mantil env -u`. This is where the website for this project will be availabe. The API endpoints can be invoked by specifying the function and method name in the path, for example `$(mantil env -u)/ping`.

## Interact with the application

### Get JWT token

```bash
mantil invoke auth/setToken
```

or

```bash
curl -X GET $(mantil env --url)/auth/setToken
```

The response will contain a jWT token required for all subsequent requests for authentication purpose. For example:

```text
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImJpbmgiLCJlbWFpbCI6ImJpbmhuZDIzNEBnbWFpbC5jb20iLCJwaG9uZSI6Ijc4ODk0OTk5MyIsImV4cCI6MTY3OTY2OTc3MywiaXNzIjoiQmluaCBMZSJ9.pqyrYupvp0HaXHNAnjUOXuIey6UTyH4FWCGATssqgv0
```

### Create new movie

```bash
mantil invoke movie/create --data '{"token": eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImJpbmgiLCJlbWFpbCI6ImJpbmhuZDIzNEBnbWFpbC5jb20iLCJwaG9uZSI6Ijc4ODk0OTk5MyIsImV4cCI6MTY3OTY2OTc3MywiaXNzIjoiQmluaCBMZSJ9.pqyrYupvp0HaXHNAnjUOXuIey6UTyH4FWCGATssqgv0", "title": "Movie 1", "rating": "4.5"}'
```

### Get all movies

```bash
mantil invoke movie/get --data '{"token": eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImJpbmgiLCJlbWFpbCI6ImJpbmhuZDIzNEBnbWFpbC5jb20iLCJwaG9uZSI6Ijc4ODk0OTk5MyIsImV4cCI6MTY3OTY2OTc3MywiaXNzIjoiQmluaCBMZSJ9.pqyrYupvp0HaXHNAnjUOXuIey6UTyH4FWCGATssqgv0"}'
```

### Update a movie

```bash
mantil invoke movie/create --data '{"token": eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImJpbmgiLCJlbWFpbCI6ImJpbmhuZDIzNEBnbWFpbC5jb20iLCJwaG9uZSI6Ijc4ODk0OTk5MyIsImV4cCI6MTY3OTY2OTc3MywiaXNzIjoiQmluaCBMZSJ9.pqyrYupvp0HaXHNAnjUOXuIey6UTyH4FWCGATssqgv0", "title": "Movie 1", "rating": "4"}'
```

### Delete a movie

```bash
mantil invoke movie/delete --data '{"token": eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImJpbmgiLCJlbWFpbCI6ImJpbmhuZDIzNEBnbWFpbC5jb20iLCJwaG9uZSI6Ijc4ODk0OTk5MyIsImV4cCI6MTY3OTY2OTc3MywiaXNzIjoiQmluaCBMZSJ9.pqyrYupvp0HaXHNAnjUOXuIey6UTyH4FWCGATssqgv0", "id": "4c434389-ddd2-49f6-8c26-2a1692e8a4d8"}'
```

## Cleanup

To remove the created stage from your AWS account destroy it with:

```bash
mantil stage destroy development
```

Uninstall command `mantil aws uninstall` will clean up all created resources and leave the AWS account in the initial state.
