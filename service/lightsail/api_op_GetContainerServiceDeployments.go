// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the deployments for your Amazon Lightsail container service A deployment
// specifies the settings, such as the ports and launch command, of containers that
// are deployed to your container service. The deployments are ordered by version
// in ascending order. The newest version is listed at the top of the response. A
// set number of deployments are kept before the oldest one is replaced with the
// newest one. For more information, see Amazon Lightsail endpoints and quotas
// (https://docs.aws.amazon.com/general/latest/gr/lightsail.html) in the AWS
// General Reference.
func (c *Client) GetContainerServiceDeployments(ctx context.Context, params *GetContainerServiceDeploymentsInput, optFns ...func(*Options)) (*GetContainerServiceDeploymentsOutput, error) {
	if params == nil {
		params = &GetContainerServiceDeploymentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetContainerServiceDeployments", params, optFns, addOperationGetContainerServiceDeploymentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetContainerServiceDeploymentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetContainerServiceDeploymentsInput struct {

	// The name of the container service for which to return deployments.
	//
	// This member is required.
	ServiceName *string
}

type GetContainerServiceDeploymentsOutput struct {

	// An array of objects that describe deployments for a container service.
	Deployments []types.ContainerServiceDeployment

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetContainerServiceDeploymentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetContainerServiceDeployments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetContainerServiceDeployments{}, middleware.After)
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
	if err = addOpGetContainerServiceDeploymentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetContainerServiceDeployments(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetContainerServiceDeployments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetContainerServiceDeployments",
	}
}
