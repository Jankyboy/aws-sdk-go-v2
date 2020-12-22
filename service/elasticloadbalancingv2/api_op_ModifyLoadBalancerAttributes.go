// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the specified attributes of the specified Application Load Balancer,
// Network Load Balancer, or Gateway Load Balancer. If any of the specified
// attributes can't be modified as requested, the call fails. Any existing
// attributes that you do not modify retain their current values.
func (c *Client) ModifyLoadBalancerAttributes(ctx context.Context, params *ModifyLoadBalancerAttributesInput, optFns ...func(*Options)) (*ModifyLoadBalancerAttributesOutput, error) {
	if params == nil {
		params = &ModifyLoadBalancerAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyLoadBalancerAttributes", params, optFns, addOperationModifyLoadBalancerAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyLoadBalancerAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyLoadBalancerAttributesInput struct {

	// The load balancer attributes.
	//
	// This member is required.
	Attributes []types.LoadBalancerAttribute

	// The Amazon Resource Name (ARN) of the load balancer.
	//
	// This member is required.
	LoadBalancerArn *string
}

type ModifyLoadBalancerAttributesOutput struct {

	// Information about the load balancer attributes.
	Attributes []types.LoadBalancerAttribute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyLoadBalancerAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyLoadBalancerAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyLoadBalancerAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpModifyLoadBalancerAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyLoadBalancerAttributes(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opModifyLoadBalancerAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "ModifyLoadBalancerAttributes",
	}
}
