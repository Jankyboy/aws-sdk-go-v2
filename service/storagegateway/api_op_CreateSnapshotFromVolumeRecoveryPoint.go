// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Initiates a snapshot of a gateway from a volume recovery point. This operation
// is only supported in the cached volume gateway type. A volume recovery point is
// a point in time at which all data of the volume is consistent and from which you
// can create a snapshot. To get a list of volume recovery point for cached volume
// gateway, use ListVolumeRecoveryPoints. In the
// CreateSnapshotFromVolumeRecoveryPoint request, you identify the volume by
// providing its Amazon Resource Name (ARN). You must also provide a description
// for the snapshot. When the gateway takes a snapshot of the specified volume, the
// snapshot and its description appear in the AWS Storage Gateway console. In
// response, the gateway returns you a snapshot ID. You can use this snapshot ID to
// check the snapshot progress or later use it when you want to create a volume
// from a snapshot. To list or delete a snapshot, you must use the Amazon EC2 API.
// For more information, see DescribeSnapshots
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSnapshots.html)
// or DeleteSnapshot
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DeleteSnapshot.html)
// in the Amazon Elastic Compute Cloud API Reference.
func (c *Client) CreateSnapshotFromVolumeRecoveryPoint(ctx context.Context, params *CreateSnapshotFromVolumeRecoveryPointInput, optFns ...func(*Options)) (*CreateSnapshotFromVolumeRecoveryPointOutput, error) {
	if params == nil {
		params = &CreateSnapshotFromVolumeRecoveryPointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSnapshotFromVolumeRecoveryPoint", params, optFns, addOperationCreateSnapshotFromVolumeRecoveryPointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSnapshotFromVolumeRecoveryPointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSnapshotFromVolumeRecoveryPointInput struct {

	// Textual description of the snapshot that appears in the Amazon EC2 console,
	// Elastic Block Store snapshots panel in the Description field, and in the AWS
	// Storage Gateway snapshot Details pane, Description field.
	//
	// This member is required.
	SnapshotDescription *string

	// The Amazon Resource Name (ARN) of the iSCSI volume target. Use the
	// DescribeStorediSCSIVolumes operation to return to retrieve the TargetARN for
	// specified VolumeARN.
	//
	// This member is required.
	VolumeARN *string

	// A list of up to 50 tags that can be assigned to a snapshot. Each tag is a
	// key-value pair. Valid characters for key and value are letters, spaces, and
	// numbers representable in UTF-8 format, and the following special characters: + -
	// = . _ : / @. The maximum length of a tag's key is 128 characters, and the
	// maximum length for a tag's value is 256.
	Tags []types.Tag
}

type CreateSnapshotFromVolumeRecoveryPointOutput struct {

	// The ID of the snapshot.
	SnapshotId *string

	// The Amazon Resource Name (ARN) of the iSCSI volume target. Use the
	// DescribeStorediSCSIVolumes operation to return to retrieve the TargetARN for
	// specified VolumeARN.
	VolumeARN *string

	// The time the volume was created from the recovery point.
	VolumeRecoveryPointTime *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateSnapshotFromVolumeRecoveryPointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateSnapshotFromVolumeRecoveryPoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateSnapshotFromVolumeRecoveryPoint{}, middleware.After)
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
	if err = addOpCreateSnapshotFromVolumeRecoveryPointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSnapshotFromVolumeRecoveryPoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSnapshotFromVolumeRecoveryPoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "CreateSnapshotFromVolumeRecoveryPoint",
	}
}
