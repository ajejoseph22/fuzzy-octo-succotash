#!/bin/bash

#TMPL=$(cfn-include template.yml -y)
#sed s/T00:00:00.000Z//g <<< "$TMPL" > combined.yml
#aws cloudformation create-stack --stack-name "$1" --template-body file://combined.yml --parameters file://"$2" --capabilities CAPABILITY_NAMED_IAM

aws cloudformation create-stack --stack-name "$1" --template-body file://"$2" --parameters file://"$3" --capabilities CAPABILITY_NAMED_IAM
