#!/bin/bash

aws cloudformation create-stack --stack-name ComputeStack --template-body file://$(pwd)/template.yaml --parameters file://$(pwd)/parameters.json