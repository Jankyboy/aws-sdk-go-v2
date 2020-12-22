// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves (queries) statistical data and other information about one or more S3
// buckets that Amazon Macie monitors and analyzes.
func (c *Client) DescribeBuckets(ctx context.Context, params *DescribeBucketsInput, optFns ...func(*Options)) (*DescribeBucketsOutput, error) {
	if params == nil {
		params = &DescribeBucketsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBuckets", params, optFns, addOperationDescribeBucketsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBucketsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBucketsInput struct {

	// The criteria to use to filter the query results.
	Criteria map[string]types.BucketCriteriaAdditionalProperties

	// The maximum number of items to include in each page of the response. The default
	// value is 50.
	MaxResults int32

	// The nextToken string that specifies which page of results to return in a
	// paginated response.
	NextToken *string

	// The criteria to use to sort the query results.
	SortCriteria *types.BucketSortCriteria
}

type DescribeBucketsOutput struct {

	// An array of objects, one for each bucket that meets the filter criteria
	// specified in the request.
	Buckets []types.BucketMetadata

	// The string to use in a subsequent request to get the next page of results in a
	// paginated response. This value is null if there are no additional pages.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeBucketsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeBuckets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeBuckets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBuckets(options.Region), middleware.Before); err != nil {
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

// DescribeBucketsAPIClient is a client that implements the DescribeBuckets
// operation.
type DescribeBucketsAPIClient interface {
	DescribeBuckets(context.Context, *DescribeBucketsInput, ...func(*Options)) (*DescribeBucketsOutput, error)
}

var _ DescribeBucketsAPIClient = (*Client)(nil)

// DescribeBucketsPaginatorOptions is the paginator options for DescribeBuckets
type DescribeBucketsPaginatorOptions struct {
	// The maximum number of items to include in each page of the response. The default
	// value is 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeBucketsPaginator is a paginator for DescribeBuckets
type DescribeBucketsPaginator struct {
	options   DescribeBucketsPaginatorOptions
	client    DescribeBucketsAPIClient
	params    *DescribeBucketsInput
	nextToken *string
	firstPage bool
}

// NewDescribeBucketsPaginator returns a new DescribeBucketsPaginator
func NewDescribeBucketsPaginator(client DescribeBucketsAPIClient, params *DescribeBucketsInput, optFns ...func(*DescribeBucketsPaginatorOptions)) *DescribeBucketsPaginator {
	options := DescribeBucketsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &DescribeBucketsInput{}
	}

	return &DescribeBucketsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeBucketsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeBuckets page.
func (p *DescribeBucketsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeBucketsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeBuckets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeBuckets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "DescribeBuckets",
	}
}
