// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates a CIDR block from a subnet. Currently, you can disassociate an
// IPv6 CIDR block only. You must detach or delete all gateways and resources that
// are associated with the CIDR block before you can disassociate it.
func (c *Client) DisassociateSubnetCidrBlock(ctx context.Context, params *DisassociateSubnetCidrBlockInput, optFns ...func(*Options)) (*DisassociateSubnetCidrBlockOutput, error) {
	if params == nil {
		params = &DisassociateSubnetCidrBlockInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateSubnetCidrBlock", params, optFns, addOperationDisassociateSubnetCidrBlockMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateSubnetCidrBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateSubnetCidrBlockInput struct {

	// The association ID for the CIDR block.
	//
	// This member is required.
	AssociationId *string
}

type DisassociateSubnetCidrBlockOutput struct {

	// Information about the IPv6 CIDR block association.
	Ipv6CidrBlockAssociation *types.SubnetIpv6CidrBlockAssociation

	// The ID of the subnet.
	SubnetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateSubnetCidrBlockMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDisassociateSubnetCidrBlock{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDisassociateSubnetCidrBlock{}, middleware.After)
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
	if err = addOpDisassociateSubnetCidrBlockValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateSubnetCidrBlock(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateSubnetCidrBlock(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DisassociateSubnetCidrBlock",
	}
}
