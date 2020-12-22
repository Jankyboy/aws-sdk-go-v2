// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation replaces the current set of name servers for the domain with the
// specified set of name servers. If you use Amazon Route 53 as your DNS service,
// specify the four name servers in the delegation set for the hosted zone for the
// domain. If successful, this operation returns an operation ID that you can use
// to track the progress and completion of the action. If the request is not
// completed successfully, the domain registrant will be notified by email.
func (c *Client) UpdateDomainNameservers(ctx context.Context, params *UpdateDomainNameserversInput, optFns ...func(*Options)) (*UpdateDomainNameserversOutput, error) {
	if params == nil {
		params = &UpdateDomainNameserversInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDomainNameservers", params, optFns, addOperationUpdateDomainNameserversMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDomainNameserversOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Replaces the current set of name servers for the domain with the specified set
// of name servers. If you use Amazon Route 53 as your DNS service, specify the
// four name servers in the delegation set for the hosted zone for the domain. If
// successful, this operation returns an operation ID that you can use to track the
// progress and completion of the action. If the request is not completed
// successfully, the domain registrant will be notified by email.
type UpdateDomainNameserversInput struct {

	// The name of the domain that you want to change name servers for.
	//
	// This member is required.
	DomainName *string

	// A list of new name servers for the domain.
	//
	// This member is required.
	Nameservers []types.Nameserver

	// The authorization key for .fi domains
	//
	// Deprecated: This member has been deprecated.
	FIAuthKey *string
}

// The UpdateDomainNameservers response includes the following element.
type UpdateDomainNameserversOutput struct {

	// Identifier for tracking the progress of the request. To query the operation
	// status, use GetOperationDetail
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html).
	//
	// This member is required.
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDomainNameserversMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDomainNameservers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDomainNameservers{}, middleware.After)
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
	if err = addOpUpdateDomainNameserversValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDomainNameservers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDomainNameservers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "UpdateDomainNameservers",
	}
}
