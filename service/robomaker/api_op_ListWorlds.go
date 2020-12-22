// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists worlds.
func (c *Client) ListWorlds(ctx context.Context, params *ListWorldsInput, optFns ...func(*Options)) (*ListWorldsOutput, error) {
	if params == nil {
		params = &ListWorldsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorlds", params, optFns, addOperationListWorldsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorldsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorldsInput struct {

	// Optional filters to limit results. You can use status.
	Filters []types.Filter

	// When this parameter is used, ListWorlds only returns maxResults results in a
	// single page along with a nextToken response element. The remaining results of
	// the initial request can be seen by sending another ListWorlds request with the
	// returned nextToken value. This value can be between 1 and 100. If this parameter
	// is not used, then ListWorlds returns up to 100 results and a nextToken value if
	// applicable.
	MaxResults *int32

	// If the previous paginated request did not return all of the remaining results,
	// the response object's nextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListWorlds again and assign that token to the
	// request object's nextToken parameter. If there are no remaining results, the
	// previous response object's NextToken parameter is set to null.
	NextToken *string
}

type ListWorldsOutput struct {

	// If the previous paginated request did not return all of the remaining results,
	// the response object's nextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListWorlds again and assign that token to the
	// request object's nextToken parameter. If there are no remaining results, the
	// previous response object's NextToken parameter is set to null.
	NextToken *string

	// Summary information for worlds.
	WorldSummaries []types.WorldSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListWorldsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWorlds{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWorlds{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorlds(options.Region), middleware.Before); err != nil {
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

// ListWorldsAPIClient is a client that implements the ListWorlds operation.
type ListWorldsAPIClient interface {
	ListWorlds(context.Context, *ListWorldsInput, ...func(*Options)) (*ListWorldsOutput, error)
}

var _ ListWorldsAPIClient = (*Client)(nil)

// ListWorldsPaginatorOptions is the paginator options for ListWorlds
type ListWorldsPaginatorOptions struct {
	// When this parameter is used, ListWorlds only returns maxResults results in a
	// single page along with a nextToken response element. The remaining results of
	// the initial request can be seen by sending another ListWorlds request with the
	// returned nextToken value. This value can be between 1 and 100. If this parameter
	// is not used, then ListWorlds returns up to 100 results and a nextToken value if
	// applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorldsPaginator is a paginator for ListWorlds
type ListWorldsPaginator struct {
	options   ListWorldsPaginatorOptions
	client    ListWorldsAPIClient
	params    *ListWorldsInput
	nextToken *string
	firstPage bool
}

// NewListWorldsPaginator returns a new ListWorldsPaginator
func NewListWorldsPaginator(client ListWorldsAPIClient, params *ListWorldsInput, optFns ...func(*ListWorldsPaginatorOptions)) *ListWorldsPaginator {
	options := ListWorldsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListWorldsInput{}
	}

	return &ListWorldsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorldsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListWorlds page.
func (p *ListWorldsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorldsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListWorlds(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWorlds(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "ListWorlds",
	}
}
