// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Amazon Redshift subnet group. You must provide a list of one or
// more subnets in your existing Amazon Virtual Private Cloud (Amazon VPC) when
// creating Amazon Redshift subnet group. For information about subnet groups, go
// to Amazon Redshift Cluster Subnet Groups
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-cluster-subnet-groups.html)
// in the Amazon Redshift Cluster Management Guide.
func (c *Client) CreateClusterSubnetGroup(ctx context.Context, params *CreateClusterSubnetGroupInput, optFns ...func(*Options)) (*CreateClusterSubnetGroupOutput, error) {
	if params == nil {
		params = &CreateClusterSubnetGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateClusterSubnetGroup", params, optFns, addOperationCreateClusterSubnetGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateClusterSubnetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateClusterSubnetGroupInput struct {

	// The name for the subnet group. Amazon Redshift stores the value as a lowercase
	// string. Constraints:
	//
	// * Must contain no more than 255 alphanumeric characters or
	// hyphens.
	//
	// * Must not be "Default".
	//
	// * Must be unique for all subnet groups that
	// are created by your AWS account.
	//
	// Example: examplesubnetgroup
	//
	// This member is required.
	ClusterSubnetGroupName *string

	// A description for the subnet group.
	//
	// This member is required.
	Description *string

	// An array of VPC subnet IDs. A maximum of 20 subnets can be modified in a single
	// request.
	//
	// This member is required.
	SubnetIds []string

	// A list of tag instances.
	Tags []types.Tag
}

type CreateClusterSubnetGroupOutput struct {

	// Describes a subnet group.
	ClusterSubnetGroup *types.ClusterSubnetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateClusterSubnetGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateClusterSubnetGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateClusterSubnetGroup{}, middleware.After)
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
	if err = addOpCreateClusterSubnetGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateClusterSubnetGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateClusterSubnetGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "CreateClusterSubnetGroup",
	}
}
