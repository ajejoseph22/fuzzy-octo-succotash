# Deploy a highly-available web app using Cloudformation
This project deploys a highly available web app spanning four availability zones.

## Architecture

## Solution
The solution consists of two parts:
1. The network stack.
2. The application (server) stack

## Deployment Steps
* Deploy the network stack first.

```
aws cloudformation create-stack --stack-name <network-stack-name> --template-body file://network.yml --parameters file://network-parameters.json --capabilities CAPABILITY_NAMED_IAM
```

* Then deploy the application stack.

```
aws cloudformation create-stack --stack-name <network-stack-name> --template-body file://servers.yml --parameters file://server-parameters.json --capabilities CAPABILITY_NAMED_IAM
```

