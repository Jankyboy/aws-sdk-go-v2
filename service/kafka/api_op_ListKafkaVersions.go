// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of Kafka versions.
func (c *Client) ListKafkaVersions(ctx context.Context, params *ListKafkaVersionsInput, optFns ...func(*Options)) (*ListKafkaVersionsOutput, error) {
	if params == nil {
		params = &ListKafkaVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListKafkaVersions", params, optFns, addOperationListKafkaVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListKafkaVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListKafkaVersionsInput struct {

	// The maximum number of results to return in the response. If there are more
	// results, the response includes a NextToken parameter.
	MaxResults int32

	// The paginated results marker. When the result of the operation is truncated, the
	// call returns NextToken in the response. To get the next batch, provide this
	// token in your next request.
	NextToken *string
}

type ListKafkaVersionsOutput struct {
	KafkaVersions []types.KafkaVersion

	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListKafkaVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListKafkaVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListKafkaVersions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListKafkaVersions(options.Region), middleware.Before); err != nil {
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

// ListKafkaVersionsAPIClient is a client that implements the ListKafkaVersions
// operation.
type ListKafkaVersionsAPIClient interface {
	ListKafkaVersions(context.Context, *ListKafkaVersionsInput, ...func(*Options)) (*ListKafkaVersionsOutput, error)
}

var _ ListKafkaVersionsAPIClient = (*Client)(nil)

// ListKafkaVersionsPaginatorOptions is the paginator options for ListKafkaVersions
type ListKafkaVersionsPaginatorOptions struct {
	// The maximum number of results to return in the response. If there are more
	// results, the response includes a NextToken parameter.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListKafkaVersionsPaginator is a paginator for ListKafkaVersions
type ListKafkaVersionsPaginator struct {
	options   ListKafkaVersionsPaginatorOptions
	client    ListKafkaVersionsAPIClient
	params    *ListKafkaVersionsInput
	nextToken *string
	firstPage bool
}

// NewListKafkaVersionsPaginator returns a new ListKafkaVersionsPaginator
func NewListKafkaVersionsPaginator(client ListKafkaVersionsAPIClient, params *ListKafkaVersionsInput, optFns ...func(*ListKafkaVersionsPaginatorOptions)) *ListKafkaVersionsPaginator {
	options := ListKafkaVersionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListKafkaVersionsInput{}
	}

	return &ListKafkaVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListKafkaVersionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListKafkaVersions page.
func (p *ListKafkaVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListKafkaVersionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListKafkaVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListKafkaVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafka",
		OperationName: "ListKafkaVersions",
	}
}
