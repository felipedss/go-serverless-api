
# Go Serverless Api - AWS

## The application structure will have two components:
- API Gateway: manages all tasks related to receiving and handling requests for the API.
- Lambda: code to be executed in order to verify the request made, obtain the necessary data, process it and return to the API.

## Top Level System Diagram

![component](/docs/arch.png)


## Compile the code in the main file
```console
GOOS=linux go build -o main
```

## Put your executable into a zip file
```console
zip main.zip main
```

Thereafter, upload your code in your lambda console

<b>Further improvements:</b> CI/CD pipeline