// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteAttendeeInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime SDK attendee ID.
	//
	// AttendeeId is a required field
	AttendeeId *string `location:"uri" locationName:"attendeeId" type:"string" required:"true"`

	// The Amazon Chime SDK meeting ID.
	//
	// MeetingId is a required field
	MeetingId *string `location:"uri" locationName:"meetingId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAttendeeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAttendeeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAttendeeInput"}

	if s.AttendeeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AttendeeId"))
	}

	if s.MeetingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeetingId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAttendeeInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AttendeeId != nil {
		v := *s.AttendeeId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "attendeeId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MeetingId != nil {
		v := *s.MeetingId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meetingId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteAttendeeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAttendeeOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAttendeeOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteAttendee = "DeleteAttendee"

// DeleteAttendeeRequest returns a request value for making API operation for
// Amazon Chime.
//
// Deletes an attendee from the specified Amazon Chime SDK meeting and deletes
// their JoinToken. Attendees are automatically deleted when a Amazon Chime
// SDK meeting is deleted. For more information about the Amazon Chime SDK,
// see Using the Amazon Chime SDK (https://docs.aws.amazon.com/chime/latest/dg/meetings-sdk.html)
// in the Amazon Chime Developer Guide.
//
//    // Example sending a request using DeleteAttendeeRequest.
//    req := client.DeleteAttendeeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/DeleteAttendee
func (c *Client) DeleteAttendeeRequest(input *DeleteAttendeeInput) DeleteAttendeeRequest {
	op := &aws.Operation{
		Name:       opDeleteAttendee,
		HTTPMethod: "DELETE",
		HTTPPath:   "/meetings/{meetingId}/attendees/{attendeeId}",
	}

	if input == nil {
		input = &DeleteAttendeeInput{}
	}

	req := c.newRequest(op, input, &DeleteAttendeeOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteAttendeeRequest{Request: req, Input: input, Copy: c.DeleteAttendeeRequest}
}

// DeleteAttendeeRequest is the request type for the
// DeleteAttendee API operation.
type DeleteAttendeeRequest struct {
	*aws.Request
	Input *DeleteAttendeeInput
	Copy  func(*DeleteAttendeeInput) DeleteAttendeeRequest
}

// Send marshals and sends the DeleteAttendee API request.
func (r DeleteAttendeeRequest) Send(ctx context.Context) (*DeleteAttendeeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAttendeeResponse{
		DeleteAttendeeOutput: r.Request.Data.(*DeleteAttendeeOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAttendeeResponse is the response type for the
// DeleteAttendee API operation.
type DeleteAttendeeResponse struct {
	*DeleteAttendeeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAttendee request.
func (r *DeleteAttendeeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
