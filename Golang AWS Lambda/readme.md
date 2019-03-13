## AWS Lambda Base Golang

#### Check Go

```bash
go 
#use brew
brew install go
brew upgrade go
```
`go to` [Download](https://golang.org/dl/)

#### Create Path
```bash
mkdir ~/go
echo "export GOPATH=\$HOME/go" >> ~/.bash_profile
source ~/.bash_profile
```

#### Create Project
```bash
mkdir ~/go/helloworld
cd ~/go/helloworld
```

#### Touch File
```bash
touch helloworld.go
```

#### Install Lambda

```bash
go get github.com/aws/aws-lambda-go/lambda
```

#### HelloWorld
```go
package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	)

func main() {
	lambda.Start(Hello)
	
}

func Hello()  {
	fmt.Println("Hello lambda")
}
```

#### Go Build

```bash
go build helloworld.go
```

#### Zip 

```bash
zip helloworld.zip helloworld
```

#### Write lambda-trust-policy.json

```json
{
	"Version": "2012-10-17",
	"Statement": {
  		"Effect": "Allow",
  		"Principal": {
    		"Service": "lambda.amazonaws.com"
  		},
  		"Action": "sts:AssumeRole"
	}
}
```

#### AWS Prepare
```bash
aws configure
aws ecr get-login --no-include-email --region "region_name"

aws iam create-role --role-name lambda-basic-execution--assume-role-policy-document file://lambda-trust-policy.json

aws iam attach-role-policy --role-name lambda-basic-execution --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole

aws iam get-role --role-name lambda-basic-execution

#or

aws lambda create-function \
--function-name helloworld_go \
--zip-file fileb://helloworld.zip \
--handler helloworld \
--runtime go1.x \
--role "arn:aws:iam::<YOUR_ACCOUNT_ID>:role/lambda-basic-execution"
```

### Invoke

```bash
aws lambda invoke \
--function-name helloworld_go \
--invocation-type "RequestResponse" \
response.txt

```