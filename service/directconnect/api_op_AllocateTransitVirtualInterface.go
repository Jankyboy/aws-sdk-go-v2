// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provisions a transit virtual interface to be owned by the specified AWS account.
// Use this type of interface to connect a transit gateway to your Direct Connect
// gateway. The owner of a connection provisions a transit virtual interface to be
// owned by the specified AWS account. After you create a transit virtual
// interface, it must be confirmed by the owner using
// ConfirmTransitVirtualInterface. Until this step has been completed, the transit
// virtual interface is in the requested state and is not available to handle
// traffic.
func (c *Client) AllocateTransitVirtualInterface(ctx context.Context, params *AllocateTransitVirtualInterfaceInput, optFns ...func(*Options)) (*AllocateTransitVirtualInterfaceOutput, error) {
	if params == nil {
		params = &AllocateTransitVirtualInterfaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AllocateTransitVirtualInterface", params, optFns, addOperationAllocateTransitVirtualInterfaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AllocateTransitVirtualInterfaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AllocateTransitVirtualInterfaceInput struct {

	// The ID of the connection on which the transit virtual interface is provisioned.
	//
	// This member is required.
	ConnectionId *string

	// Information about the transit virtual interface.
	//
	// This member is required.
	NewTransitVirtualInterfaceAllocation *types.NewTransitVirtualInterfaceAllocation

	// The ID of the AWS account that owns the transit virtual interface.
	//
	// This member is required.
	OwnerAccount *string
}

type AllocateTransitVirtualInterfaceOutput struct {

	// Information about a virtual interface.
	VirtualInterface *types.VirtualInterface

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAllocateTransitVirtualInterfaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAllocateTransitVirtualInterface{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAllocateTransitVirtualInterface{}, middleware.After)
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
	if err = addOpAllocateTransitVirtualInterfaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAllocateTransitVirtualInterface(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAllocateTransitVirtualInterface(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "AllocateTransitVirtualInterface",
	}
}
