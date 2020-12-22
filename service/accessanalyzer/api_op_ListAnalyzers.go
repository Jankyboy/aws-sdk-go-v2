// Code generated by smithy-go-codegen DO NOT EDIT.

package accessanalyzer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of analyzers.
func (c *Client) ListAnalyzers(ctx context.Context, params *ListAnalyzersInput, optFns ...func(*Options)) (*ListAnalyzersOutput, error) {
	if params == nil {
		params = &ListAnalyzersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAnalyzers", params, optFns, addOperationListAnalyzersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAnalyzersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Retrieves a list of analyzers.
type ListAnalyzersInput struct {

	// The maximum number of results to return in the response.
	MaxResults *int32

	// A token used for pagination of results returned.
	NextToken *string

	// The type of analyzer.
	Type types.Type
}

// The response to the request.
type ListAnalyzersOutput struct {

	// The analyzers retrieved.
	//
	// This member is required.
	Analyzers []types.AnalyzerSummary

	// A token used for pagination of results returned.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAnalyzersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAnalyzers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAnalyzers{}, middleware.After)
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

// ListAnalyzersAPIClient is a client that implements the ListAnalyzers operation.
type ListAnalyzersAPIClient interface {
	ListAnalyzers(context.Context, *ListAnalyzersInput, ...func(*Options)) (*ListAnalyzersOutput, error)
}

var _ ListAnalyzersAPIClient = (*Client)(nil)

// ListAnalyzersPaginatorOptions is the paginator options for ListAnalyzers
type ListAnalyzersPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAnalyzersPaginator is a paginator for ListAnalyzers
type ListAnalyzersPaginator struct {
	options   ListAnalyzersPaginatorOptions
	client    ListAnalyzersAPIClient
	params    *ListAnalyzersInput
	nextToken *string
	firstPage bool
}

// NewListAnalyzersPaginator returns a new ListAnalyzersPaginator
func NewListAnalyzersPaginator(client ListAnalyzersAPIClient, params *ListAnalyzersInput, optFns ...func(*ListAnalyzersPaginatorOptions)) *ListAnalyzersPaginator {
	options := ListAnalyzersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListAnalyzersInput{}
	}

	return &ListAnalyzersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAnalyzersPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListAnalyzers page.
func (p *ListAnalyzersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAnalyzersOutput, error) {
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

	result, err := p.client.ListAnalyzers(ctx, &params, optFns...)
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
