// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns descriptions of all resources of the specified stack. For deleted
// stacks, ListStackResources returns resource information for up to 90 days after
// the stack has been deleted.
func (c *Client) ListStackResources(ctx context.Context, params *ListStackResourcesInput, optFns ...func(*Options)) (*ListStackResourcesOutput, error) {
	if params == nil {
		params = &ListStackResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStackResources", params, optFns, addOperationListStackResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStackResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ListStackResource action.
type ListStackResourcesInput struct {

	// The name or the unique stack ID that is associated with the stack, which are not
	// always interchangeable:
	//
	// * Running stacks: You can specify either the stack's
	// name or its unique stack ID.
	//
	// * Deleted stacks: You must specify the unique
	// stack ID.
	//
	// Default: There is no default value.
	//
	// This member is required.
	StackName *string

	// A string that identifies the next page of stack resources that you want to
	// retrieve.
	NextToken *string
}

// The output for a ListStackResources action.
type ListStackResourcesOutput struct {

	// If the output exceeds 1 MB, a string that identifies the next page of stack
	// resources. If no additional page exists, this value is null.
	NextToken *string

	// A list of StackResourceSummary structures.
	StackResourceSummaries []types.StackResourceSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListStackResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListStackResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListStackResources{}, middleware.After)
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
	if err = addOpListStackResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStackResources(options.Region), middleware.Before); err != nil {
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

// ListStackResourcesAPIClient is a client that implements the ListStackResources
// operation.
type ListStackResourcesAPIClient interface {
	ListStackResources(context.Context, *ListStackResourcesInput, ...func(*Options)) (*ListStackResourcesOutput, error)
}

var _ ListStackResourcesAPIClient = (*Client)(nil)

// ListStackResourcesPaginatorOptions is the paginator options for
// ListStackResources
type ListStackResourcesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStackResourcesPaginator is a paginator for ListStackResources
type ListStackResourcesPaginator struct {
	options   ListStackResourcesPaginatorOptions
	client    ListStackResourcesAPIClient
	params    *ListStackResourcesInput
	nextToken *string
	firstPage bool
}

// NewListStackResourcesPaginator returns a new ListStackResourcesPaginator
func NewListStackResourcesPaginator(client ListStackResourcesAPIClient, params *ListStackResourcesInput, optFns ...func(*ListStackResourcesPaginatorOptions)) *ListStackResourcesPaginator {
	options := ListStackResourcesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListStackResourcesInput{}
	}

	return &ListStackResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStackResourcesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListStackResources page.
func (p *ListStackResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStackResourcesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListStackResources(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListStackResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "ListStackResources",
	}
}
