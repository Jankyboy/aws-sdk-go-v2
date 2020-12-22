// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the list of event trackers associated with the account. The response
// provides the properties for each event tracker, including the Amazon Resource
// Name (ARN) and tracking ID. For more information on event trackers, see
// CreateEventTracker.
func (c *Client) ListEventTrackers(ctx context.Context, params *ListEventTrackersInput, optFns ...func(*Options)) (*ListEventTrackersOutput, error) {
	if params == nil {
		params = &ListEventTrackersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEventTrackers", params, optFns, addOperationListEventTrackersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEventTrackersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEventTrackersInput struct {

	// The ARN of a dataset group used to filter the response.
	DatasetGroupArn *string

	// The maximum number of event trackers to return.
	MaxResults *int32

	// A token returned from the previous call to ListEventTrackers for getting the
	// next set of event trackers (if they exist).
	NextToken *string
}

type ListEventTrackersOutput struct {

	// A list of event trackers.
	EventTrackers []types.EventTrackerSummary

	// A token for getting the next set of event trackers (if they exist).
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListEventTrackersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListEventTrackers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEventTrackers{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEventTrackers(options.Region), middleware.Before); err != nil {
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

// ListEventTrackersAPIClient is a client that implements the ListEventTrackers
// operation.
type ListEventTrackersAPIClient interface {
	ListEventTrackers(context.Context, *ListEventTrackersInput, ...func(*Options)) (*ListEventTrackersOutput, error)
}

var _ ListEventTrackersAPIClient = (*Client)(nil)

// ListEventTrackersPaginatorOptions is the paginator options for ListEventTrackers
type ListEventTrackersPaginatorOptions struct {
	// The maximum number of event trackers to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEventTrackersPaginator is a paginator for ListEventTrackers
type ListEventTrackersPaginator struct {
	options   ListEventTrackersPaginatorOptions
	client    ListEventTrackersAPIClient
	params    *ListEventTrackersInput
	nextToken *string
	firstPage bool
}

// NewListEventTrackersPaginator returns a new ListEventTrackersPaginator
func NewListEventTrackersPaginator(client ListEventTrackersAPIClient, params *ListEventTrackersInput, optFns ...func(*ListEventTrackersPaginatorOptions)) *ListEventTrackersPaginator {
	options := ListEventTrackersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListEventTrackersInput{}
	}

	return &ListEventTrackersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEventTrackersPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListEventTrackers page.
func (p *ListEventTrackersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEventTrackersOutput, error) {
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

	result, err := p.client.ListEventTrackers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEventTrackers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "ListEventTrackers",
	}
}
