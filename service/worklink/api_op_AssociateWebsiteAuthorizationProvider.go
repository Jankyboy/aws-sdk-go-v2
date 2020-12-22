// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/worklink/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a website authorization provider with a specified fleet. This is used
// to authorize users against associated websites in the company network.
func (c *Client) AssociateWebsiteAuthorizationProvider(ctx context.Context, params *AssociateWebsiteAuthorizationProviderInput, optFns ...func(*Options)) (*AssociateWebsiteAuthorizationProviderOutput, error) {
	if params == nil {
		params = &AssociateWebsiteAuthorizationProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateWebsiteAuthorizationProvider", params, optFns, addOperationAssociateWebsiteAuthorizationProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateWebsiteAuthorizationProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateWebsiteAuthorizationProviderInput struct {

	// The authorization provider type.
	//
	// This member is required.
	AuthorizationProviderType types.AuthorizationProviderType

	// The ARN of the fleet.
	//
	// This member is required.
	FleetArn *string

	// The domain name of the authorization provider. This applies only to SAML-based
	// authorization providers.
	DomainName *string
}

type AssociateWebsiteAuthorizationProviderOutput struct {

	// A unique identifier for the authorization provider.
	AuthorizationProviderId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssociateWebsiteAuthorizationProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateWebsiteAuthorizationProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateWebsiteAuthorizationProvider{}, middleware.After)
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
	if err = addOpAssociateWebsiteAuthorizationProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateWebsiteAuthorizationProvider(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateWebsiteAuthorizationProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "worklink",
		OperationName: "AssociateWebsiteAuthorizationProvider",
	}
}
