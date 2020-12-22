// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new cross-cluster search connection from a source domain to a
// destination domain.
func (c *Client) CreateOutboundCrossClusterSearchConnection(ctx context.Context, params *CreateOutboundCrossClusterSearchConnectionInput, optFns ...func(*Options)) (*CreateOutboundCrossClusterSearchConnectionOutput, error) {
	if params == nil {
		params = &CreateOutboundCrossClusterSearchConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateOutboundCrossClusterSearchConnection", params, optFns, addOperationCreateOutboundCrossClusterSearchConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateOutboundCrossClusterSearchConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the CreateOutboundCrossClusterSearchConnection
// operation.
type CreateOutboundCrossClusterSearchConnectionInput struct {

	// Specifies the connection alias that will be used by the customer for this
	// connection.
	//
	// This member is required.
	ConnectionAlias *string

	// Specifies the DomainInformation for the destination Elasticsearch domain.
	//
	// This member is required.
	DestinationDomainInfo *types.DomainInformation

	// Specifies the DomainInformation for the source Elasticsearch domain.
	//
	// This member is required.
	SourceDomainInfo *types.DomainInformation
}

// The result of a CreateOutboundCrossClusterSearchConnection request. Contains the
// details of the newly created cross-cluster search connection.
type CreateOutboundCrossClusterSearchConnectionOutput struct {

	// Specifies the connection alias provided during the create connection request.
	ConnectionAlias *string

	// Specifies the OutboundCrossClusterSearchConnectionStatus for the newly created
	// connection.
	ConnectionStatus *types.OutboundCrossClusterSearchConnectionStatus

	// Unique id for the created outbound connection, which is used for subsequent
	// operations on connection.
	CrossClusterSearchConnectionId *string

	// Specifies the DomainInformation for the destination Elasticsearch domain.
	DestinationDomainInfo *types.DomainInformation

	// Specifies the DomainInformation for the source Elasticsearch domain.
	SourceDomainInfo *types.DomainInformation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateOutboundCrossClusterSearchConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateOutboundCrossClusterSearchConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateOutboundCrossClusterSearchConnection{}, middleware.After)
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
	if err = addOpCreateOutboundCrossClusterSearchConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateOutboundCrossClusterSearchConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateOutboundCrossClusterSearchConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "CreateOutboundCrossClusterSearchConnection",
	}
}
