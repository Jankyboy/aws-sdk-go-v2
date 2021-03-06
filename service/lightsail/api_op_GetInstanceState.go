// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetInstanceStateInput struct {
	_ struct{} `type:"structure"`

	// The name of the instance to get state information about.
	//
	// InstanceName is a required field
	InstanceName *string `locationName:"instanceName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetInstanceStateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetInstanceStateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetInstanceStateInput"}

	if s.InstanceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetInstanceStateOutput struct {
	_ struct{} `type:"structure"`

	// The state of the instance.
	State *InstanceState `locationName:"state" type:"structure"`
}

// String returns the string representation
func (s GetInstanceStateOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetInstanceState = "GetInstanceState"

// GetInstanceStateRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns the state of a specific instance. Works on one instance at a time.
//
//    // Example sending a request using GetInstanceStateRequest.
//    req := client.GetInstanceStateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetInstanceState
func (c *Client) GetInstanceStateRequest(input *GetInstanceStateInput) GetInstanceStateRequest {
	op := &aws.Operation{
		Name:       opGetInstanceState,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetInstanceStateInput{}
	}

	req := c.newRequest(op, input, &GetInstanceStateOutput{})

	return GetInstanceStateRequest{Request: req, Input: input, Copy: c.GetInstanceStateRequest}
}

// GetInstanceStateRequest is the request type for the
// GetInstanceState API operation.
type GetInstanceStateRequest struct {
	*aws.Request
	Input *GetInstanceStateInput
	Copy  func(*GetInstanceStateInput) GetInstanceStateRequest
}

// Send marshals and sends the GetInstanceState API request.
func (r GetInstanceStateRequest) Send(ctx context.Context) (*GetInstanceStateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetInstanceStateResponse{
		GetInstanceStateOutput: r.Request.Data.(*GetInstanceStateOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetInstanceStateResponse is the response type for the
// GetInstanceState API operation.
type GetInstanceStateResponse struct {
	*GetInstanceStateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetInstanceState request.
func (r *GetInstanceStateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
