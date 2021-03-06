// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package detective

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteGraphInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the behavior graph to disable.
	//
	// GraphArn is a required field
	GraphArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteGraphInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteGraphInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteGraphInput"}

	if s.GraphArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("GraphArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteGraphInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GraphArn != nil {
		v := *s.GraphArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GraphArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteGraphOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteGraphOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteGraphOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteGraph = "DeleteGraph"

// DeleteGraphRequest returns a request value for making API operation for
// Amazon Detective.
//
// Disables the specified behavior graph and queues it to be deleted. This operation
// removes the graph from each member account's list of behavior graphs.
//
// DeleteGraph can only be called by the master account for a behavior graph.
//
//    // Example sending a request using DeleteGraphRequest.
//    req := client.DeleteGraphRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/detective-2018-10-26/DeleteGraph
func (c *Client) DeleteGraphRequest(input *DeleteGraphInput) DeleteGraphRequest {
	op := &aws.Operation{
		Name:       opDeleteGraph,
		HTTPMethod: "POST",
		HTTPPath:   "/graph/removal",
	}

	if input == nil {
		input = &DeleteGraphInput{}
	}

	req := c.newRequest(op, input, &DeleteGraphOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteGraphRequest{Request: req, Input: input, Copy: c.DeleteGraphRequest}
}

// DeleteGraphRequest is the request type for the
// DeleteGraph API operation.
type DeleteGraphRequest struct {
	*aws.Request
	Input *DeleteGraphInput
	Copy  func(*DeleteGraphInput) DeleteGraphRequest
}

// Send marshals and sends the DeleteGraph API request.
func (r DeleteGraphRequest) Send(ctx context.Context) (*DeleteGraphResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteGraphResponse{
		DeleteGraphOutput: r.Request.Data.(*DeleteGraphOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteGraphResponse is the response type for the
// DeleteGraph API operation.
type DeleteGraphResponse struct {
	*DeleteGraphOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteGraph request.
func (r *DeleteGraphResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
