// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// Sets the supplied tag-set on an S3 Batch Operations job. A tag is a key-value
// pair. You can associate S3 Batch Operations tags with any job by sending a PUT
// request against the tagging subresource that is associated with the job. To
// modify the existing tag set, you can either replace the existing tag set
// entirely, or make changes within the existing tag set by retrieving the existing
// tag set using GetJobTagging
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetJobTagging.html),
// modify that tag set, and use this action to replace the tag set with the one you
// modified. For more information, see Controlling access and labeling jobs using
// tags
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-managing-jobs.html#batch-ops-job-tags)
// in the Amazon Simple Storage Service Developer Guide.
//
// * If you send this
// request with an empty tag set, Amazon S3 deletes the existing tag set on the
// Batch Operations job. If you use this method, you are charged for a Tier 1
// Request (PUT). For more information, see Amazon S3 pricing
// (http://aws.amazon.com/s3/pricing/).
//
// * For deleting existing tags for your
// Batch Operations job, a DeleteJobTagging
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteJobTagging.html)
// request is preferred because it achieves the same result without incurring
// charges.
//
// * A few things to consider about using tags:
//
// * Amazon S3 limits the
// maximum number of tags to 50 tags per job.
//
// * You can associate up to 50 tags
// with a job as long as they have unique tag keys.
//
// * A tag key can be up to 128
// Unicode characters in length, and tag values can be up to 256 Unicode characters
// in length.
//
// * The key and values are case sensitive.
//
// * For tagging-related
// restrictions related to characters and encodings, see User-Defined Tag
// Restrictions
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/allocation-tag-restrictions.html)
// in the AWS Billing and Cost Management User Guide.
//
// To use this operation, you
// must have permission to perform the s3:PutJobTagging action. Related actions
// include:
//
// * CreatJob
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateJob.html)
//
// *
// GetJobTagging
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetJobTagging.html)
//
// *
// DeleteJobTagging
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteJobTagging.html)
func (c *Client) PutJobTagging(ctx context.Context, params *PutJobTaggingInput, optFns ...func(*Options)) (*PutJobTaggingOutput, error) {
	if params == nil {
		params = &PutJobTaggingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutJobTagging", params, optFns, addOperationPutJobTaggingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutJobTaggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutJobTaggingInput struct {

	// The AWS account ID associated with the S3 Batch Operations job.
	//
	// This member is required.
	AccountId *string

	// The ID for the S3 Batch Operations job whose tags you want to replace.
	//
	// This member is required.
	JobId *string

	// The set of tags to associate with the S3 Batch Operations job.
	//
	// This member is required.
	Tags []types.S3Tag
}

type PutJobTaggingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutJobTaggingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutJobTagging{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutJobTagging{}, middleware.After)
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
	if err = addEndpointPrefix_opPutJobTaggingMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutJobTaggingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutJobTagging(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addPutJobTaggingUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opPutJobTaggingMiddleware struct {
}

func (*endpointPrefix_opPutJobTaggingMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opPutJobTaggingMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*PutJobTaggingInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opPutJobTaggingMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opPutJobTaggingMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opPutJobTagging(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutJobTagging",
	}
}

func copyPutJobTaggingInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*PutJobTaggingInput)
	if !ok {
		return nil, fmt.Errorf("expect *PutJobTaggingInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func backFillPutJobTaggingAccountID(input interface{}, v string) error {
	in := input.(*PutJobTaggingInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addPutJobTaggingUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyPutJobTaggingInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseDualstack:            options.UseDualstack,
		UseARNRegion:            options.UseARNRegion,
	})
}
