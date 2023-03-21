# AWS Lambda with Go

Serverless application using AWS Lambda and Go

## Setup

### Set AWS configuration

See [Configuration basics](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html)

### Creating the Role of the Lambda Function

```shell
aws iam create-role --role-name lambda-ex --assume-role-policy-document '{"Version": "2012-10-17","Statement": [{ "Effect": "Allow", "Principal": {"Service": "lambda.amazonaws.com"}, "Action": "sts:AssumeRole"}]}'
```

In case the above command doesn't work, you need to create the trust policy for the role. In this case the trust policy allows the lambda service to assume the role. Save the contents of the trust policy in a file named `trust-policy.json`

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}

```

Then run:

```shell
aws iam create-role --role-name lambda-example-role --assume-role-policy-document file://trust-policy.json
```

The next step is to attach a policy to the role. This policy grants permissions to the lambda function to log to CloudWatch.

```bash
aws iam attach-role-policy --role-name lambda-ex --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
```

### Build the application

**Note**: If you are using Windows, you have to set the environment variables GOARCH and GOOS to "AMD64" and "linux" respectively as the Lambda function run on Linux. This must **be done as administrator** (eg right click the "command prompt" or "Powershell" shortcut and click "Run as Administrator"). For command prompt, use:

```bash
set GOARCH=amd64
set GOOS=linux
```

If you use Powershell, use:

```powershell
$Env:GOOS = "linux"; $Env:GOARCH = "amd64"
```

```shell
go build main.go

zip function.zip main
```

### Creating the Lambda Function with AWS CLI

```shell
aws lambda create-function --function-name go-lambda-ex --runtime go1.x --zip-file fileb://function.zip --handler main --role "arn:aws:iam::694942875169:role/lambda-ex"
```

### Invoking a Lambda Function with AWS CLI

Create a file named `event.json` with the following `json` in it:

```json
{
  "name": "John Doe",
  "ager": 42
}
```

Run the `invoke` command:

```shell
aws lambda invoke --function-name go-lambda-ex --cli-binary-format raw-in-base64-out --payload file://event.json response.json
```

### Updating a Lambda Function with AWS CLI

To update a lambda function's code using the AWS CLI, run the `update-function-code` command:

```shell
aws lambda update-function-code --function-name go-lambda-ex --zip-file fileb://function.zip
```

### Delete a Lambda Function with AWS CLI

```shell
aws lambda delete-function --function-name go-lambda-ex
```
