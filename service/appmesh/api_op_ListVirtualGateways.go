// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of existing virtual gateways in a service mesh.
func (c *Client) ListVirtualGateways(ctx context.Context, params *ListVirtualGatewaysInput, optFns ...func(*Options)) (*ListVirtualGatewaysOutput, error) {
	if params == nil {
		params = &ListVirtualGatewaysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVirtualGateways", params, optFns, addOperationListVirtualGatewaysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVirtualGatewaysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVirtualGatewaysInput struct {

	// The name of the service mesh to list virtual gateways in.
	//
	// This member is required.
	MeshName *string

	// The maximum number of results returned by ListVirtualGateways in paginated
	// output. When you use this parameter, ListVirtualGateways returns only limit
	// results in a single page along with a nextToken response element. You can see
	// the remaining results of the initial request by sending another
	// ListVirtualGateways request with the returned nextToken value. This value can be
	// between 1 and 100. If you don't use this parameter, ListVirtualGateways returns
	// up to 100 results and a nextToken value if applicable.
	Limit *int32

	// The AWS IAM account ID of the service mesh owner. If the account ID is not your
	// own, then it's the ID of the account that shared the mesh with your account. For
	// more information about mesh sharing, see Working with shared meshes
	// (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).
	MeshOwner *string

	// The nextToken value returned from a previous paginated ListVirtualGateways
	// request where limit was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value.
	NextToken *string
}

type ListVirtualGatewaysOutput struct {

	// The list of existing virtual gateways for the specified service mesh.
	//
	// This member is required.
	VirtualGateways []types.VirtualGatewayRef

	// The nextToken value to include in a future ListVirtualGateways request. When the
	// results of a ListVirtualGateways request exceed limit, you can use this value to
	// retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListVirtualGatewaysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListVirtualGateways{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListVirtualGateways{}, middleware.After)
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
	if err = addOpListVirtualGatewaysValidationMiddleware(stack); err != nil {
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

// ListVirtualGatewaysAPIClient is a client that implements the ListVirtualGateways
// operation.
type ListVirtualGatewaysAPIClient interface {
	ListVirtualGateways(context.Context, *ListVirtualGatewaysInput, ...func(*Options)) (*ListVirtualGatewaysOutput, error)
}

var _ ListVirtualGatewaysAPIClient = (*Client)(nil)

// ListVirtualGatewaysPaginatorOptions is the paginator options for
// ListVirtualGateways
type ListVirtualGatewaysPaginatorOptions struct {
	// The maximum number of results returned by ListVirtualGateways in paginated
	// output. When you use this parameter, ListVirtualGateways returns only limit
	// results in a single page along with a nextToken response element. You can see
	// the remaining results of the initial request by sending another
	// ListVirtualGateways request with the returned nextToken value. This value can be
	// between 1 and 100. If you don't use this parameter, ListVirtualGateways returns
	// up to 100 results and a nextToken value if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListVirtualGatewaysPaginator is a paginator for ListVirtualGateways
type ListVirtualGatewaysPaginator struct {
	options   ListVirtualGatewaysPaginatorOptions
	client    ListVirtualGatewaysAPIClient
	params    *ListVirtualGatewaysInput
	nextToken *string
	firstPage bool
}

// NewListVirtualGatewaysPaginator returns a new ListVirtualGatewaysPaginator
func NewListVirtualGatewaysPaginator(client ListVirtualGatewaysAPIClient, params *ListVirtualGatewaysInput, optFns ...func(*ListVirtualGatewaysPaginatorOptions)) *ListVirtualGatewaysPaginator {
	options := ListVirtualGatewaysPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListVirtualGatewaysInput{}
	}

	return &ListVirtualGatewaysPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListVirtualGatewaysPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListVirtualGateways page.
func (p *ListVirtualGatewaysPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListVirtualGatewaysOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListVirtualGateways(ctx, &params, optFns...)
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
