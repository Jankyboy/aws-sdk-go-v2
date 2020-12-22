// Code generated by smithy-go-codegen DO NOT EDIT.

package dax

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dax/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a DAX cluster. All nodes in the cluster run the same DAX caching
// software.
func (c *Client) CreateCluster(ctx context.Context, params *CreateClusterInput, optFns ...func(*Options)) (*CreateClusterOutput, error) {
	if params == nil {
		params = &CreateClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCluster", params, optFns, addOperationCreateClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateClusterInput struct {

	// The cluster identifier. This parameter is stored as a lowercase string.
	// Constraints:
	//
	// * A name must contain from 1 to 20 alphanumeric characters or
	// hyphens.
	//
	// * The first character must be a letter.
	//
	// * A name cannot end with a
	// hyphen or contain two consecutive hyphens.
	//
	// This member is required.
	ClusterName *string

	// A valid Amazon Resource Name (ARN) that identifies an IAM role. At runtime, DAX
	// will assume this role and use the role's permissions to access DynamoDB on your
	// behalf.
	//
	// This member is required.
	IamRoleArn *string

	// The compute and memory capacity of the nodes in the cluster.
	//
	// This member is required.
	NodeType *string

	// The number of nodes in the DAX cluster. A replication factor of 1 will create a
	// single-node cluster, without any read replicas. For additional fault tolerance,
	// you can create a multiple node cluster with one or more read replicas. To do
	// this, set ReplicationFactor to a number between 3 (one primary and two read
	// replicas) and 10 (one primary and nine read replicas). If the AvailabilityZones
	// parameter is provided, its length must equal the ReplicationFactor. AWS
	// recommends that you have at least two read replicas per cluster.
	//
	// This member is required.
	ReplicationFactor int32

	// The Availability Zones (AZs) in which the cluster nodes will reside after the
	// cluster has been created or updated. If provided, the length of this list must
	// equal the ReplicationFactor parameter. If you omit this parameter, DAX will
	// spread the nodes across Availability Zones for the highest availability.
	AvailabilityZones []string

	// A description of the cluster.
	Description *string

	// The Amazon Resource Name (ARN) of the Amazon SNS topic to which notifications
	// will be sent. The Amazon SNS topic owner must be same as the DAX cluster owner.
	NotificationTopicArn *string

	// The parameter group to be associated with the DAX cluster.
	ParameterGroupName *string

	// Specifies the weekly time range during which maintenance on the DAX cluster is
	// performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H
	// Clock UTC). The minimum maintenance window is a 60 minute period. Valid values
	// for ddd are:
	//
	// * sun
	//
	// * mon
	//
	// * tue
	//
	// * wed
	//
	// * thu
	//
	// * fri
	//
	// * sat
	//
	// Example:
	// sun:05:00-sun:09:00 If you don't specify a preferred maintenance window when you
	// create or modify a cache cluster, DAX assigns a 60-minute maintenance window on
	// a randomly selected day of the week.
	PreferredMaintenanceWindow *string

	// Represents the settings used to enable server-side encryption on the cluster.
	SSESpecification *types.SSESpecification

	// A list of security group IDs to be assigned to each node in the DAX cluster.
	// (Each of the security group ID is system-generated.) If this parameter is not
	// specified, DAX assigns the default VPC security group to each node.
	SecurityGroupIds []string

	// The name of the subnet group to be used for the replication group. DAX clusters
	// can only run in an Amazon VPC environment. All of the subnets that you specify
	// in a subnet group must exist in the same VPC.
	SubnetGroupName *string

	// A set of tags to associate with the DAX cluster.
	Tags []types.Tag
}

type CreateClusterOutput struct {

	// A description of the DAX cluster that you have created.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCluster{}, middleware.After)
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
	if err = addOpCreateClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dax",
		OperationName: "CreateCluster",
	}
}
