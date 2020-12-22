// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the roots, organizational units (OUs), and accounts that the specified
// policy is attached to. Always check the NextToken response parameter for a null
// value when calling a List* operation. These operations can occasionally return
// an empty set of results even when there are more results available. The
// NextToken response parameter value is null only when there are no more results
// to display. This operation can be called only from the organization's management
// account or by a member account that is a delegated administrator for an AWS
// service.
func (c *Client) ListTargetsForPolicy(ctx context.Context, params *ListTargetsForPolicyInput, optFns ...func(*Options)) (*ListTargetsForPolicyOutput, error) {
	if params == nil {
		params = &ListTargetsForPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTargetsForPolicy", params, optFns, addOperationListTargetsForPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTargetsForPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTargetsForPolicyInput struct {

	// The unique identifier (ID) of the policy whose attachments you want to know. The
	// regex pattern (http://wikipedia.org/wiki/regex) for a policy ID string requires
	// "p-" followed by from 8 to 128 lowercase or uppercase letters, digits, or the
	// underscore character (_).
	//
	// This member is required.
	PolicyId *string

	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific to
	// the operation. If additional items exist beyond the maximum you specify, the
	// NextToken response element is present and has a value (is not null). Include
	// that value as the NextToken request parameter in the next call to the operation
	// to get the next part of the results. Note that Organizations might return fewer
	// results than the maximum even when there are more results available. You should
	// check NextToken after every operation to ensure that you receive all of the
	// results.
	MaxResults *int32

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string
}

type ListTargetsForPolicyOutput struct {

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null.
	NextToken *string

	// A list of structures, each of which contains details about one of the entities
	// to which the specified policy is attached.
	Targets []types.PolicyTargetSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTargetsForPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTargetsForPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTargetsForPolicy{}, middleware.After)
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
	if err = addOpListTargetsForPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTargetsForPolicy(options.Region), middleware.Before); err != nil {
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

// ListTargetsForPolicyAPIClient is a client that implements the
// ListTargetsForPolicy operation.
type ListTargetsForPolicyAPIClient interface {
	ListTargetsForPolicy(context.Context, *ListTargetsForPolicyInput, ...func(*Options)) (*ListTargetsForPolicyOutput, error)
}

var _ ListTargetsForPolicyAPIClient = (*Client)(nil)

// ListTargetsForPolicyPaginatorOptions is the paginator options for
// ListTargetsForPolicy
type ListTargetsForPolicyPaginatorOptions struct {
	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific to
	// the operation. If additional items exist beyond the maximum you specify, the
	// NextToken response element is present and has a value (is not null). Include
	// that value as the NextToken request parameter in the next call to the operation
	// to get the next part of the results. Note that Organizations might return fewer
	// results than the maximum even when there are more results available. You should
	// check NextToken after every operation to ensure that you receive all of the
	// results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTargetsForPolicyPaginator is a paginator for ListTargetsForPolicy
type ListTargetsForPolicyPaginator struct {
	options   ListTargetsForPolicyPaginatorOptions
	client    ListTargetsForPolicyAPIClient
	params    *ListTargetsForPolicyInput
	nextToken *string
	firstPage bool
}

// NewListTargetsForPolicyPaginator returns a new ListTargetsForPolicyPaginator
func NewListTargetsForPolicyPaginator(client ListTargetsForPolicyAPIClient, params *ListTargetsForPolicyInput, optFns ...func(*ListTargetsForPolicyPaginatorOptions)) *ListTargetsForPolicyPaginator {
	options := ListTargetsForPolicyPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	if params == nil {
		params = &ListTargetsForPolicyInput{}
	}

	return &ListTargetsForPolicyPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTargetsForPolicyPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListTargetsForPolicy page.
func (p *ListTargetsForPolicyPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTargetsForPolicyOutput, error) {
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

	result, err := p.client.ListTargetsForPolicy(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTargetsForPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "ListTargetsForPolicy",
	}
}
