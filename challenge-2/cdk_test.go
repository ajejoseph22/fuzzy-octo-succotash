package main

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	assertions "github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

func TestCdkStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := NewCdkStack(app, "MyStack", nil)

	// THEN
	template := assertions.Template_FromStack(stack)

	template.HasResourceProperties(jsii.String("AWS::EC2::VPC"), map[string]interface{}{
		"CidrBlock": jsii.String("10.0.0.0/16"),
	})

	template.ResourceCountIs(jsii.String("AWS::EC2::VPC"), jsii.Number(1))
	// representing the 4 AZs in us-east-1
	template.ResourceCountIs(jsii.String("AWS::EC2::Subnet"), jsii.Number(4))
	template.ResourceCountIs(jsii.String("AWS::EC2::InternetGateway"), jsii.Number(1))
	// representing one NAT GW in each public subnet/AZ
	template.ResourceCountIs(jsii.String("AWS::EC2::NatGateway"), jsii.Number(2))
	// Elastic IPs for the NAT GWs
	template.ResourceCountIs(jsii.String("AWS::EC2::EIP"), jsii.Number(2))
	template.ResourceCountIs(jsii.String("AWS::EC2::VPCGatewayAttachment"), jsii.Number(1))

}
