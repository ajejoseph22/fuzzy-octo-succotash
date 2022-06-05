package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkStackProps struct {
	awscdk.StackProps
}

func NewCdkStack(scope constructs.Construct, id string, props *CdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// create VPC
	vpc := awsec2.NewVpc(stack, jsii.String("StackVPC"), &awsec2.VpcProps{
		Cidr:               jsii.String("10.0.0.0/16"),
		VpcName:            jsii.String("Stack VPC"),
		EnableDnsSupport:   jsii.Bool(true),
		EnableDnsHostnames: jsii.Bool(true)})

	// create IGW
	internetGateway := awsec2.NewCfnInternetGateway(stack, jsii.String("StackIGW"), &awsec2.CfnInternetGatewayProps{})
	internetGateway.Tags().SetTag(jsii.String("Environment"), jsii.String("Production"), jsii.Number(1), jsii.Bool(true))

	// attach IGW to VPC
	awsec2.NewCfnVPCGatewayAttachment(stack, jsii.String("StackVPCGatewayAttachment"),
		&awsec2.CfnVPCGatewayAttachmentProps{
			VpcId:             vpc.VpcId(),
			InternetGatewayId: internetGateway.AttrInternetGatewayId()})

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewCdkStack(app, "CdkStack", &CdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
