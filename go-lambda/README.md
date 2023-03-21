# AWS Lambda with Go

Serverless application using AWS Lambda and Go

## Setup

### Set AWS configuration

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
