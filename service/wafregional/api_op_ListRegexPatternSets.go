// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListRegexPatternSetsInput struct {
	_ struct{} `type:"structure"`

	// Specifies the number of RegexPatternSet objects that you want AWS WAF to
	// return for this request. If you have more RegexPatternSet objects than the
	// number you specify for Limit, the response includes a NextMarker value that
	// you can use to get another batch of RegexPatternSet objects.
	Limit *int64 `type:"integer"`

	// If you specify a value for Limit and you have more RegexPatternSet objects
	// than the value of Limit, AWS WAF returns a NextMarker value in the response
	// that allows you to list another group of RegexPatternSet objects. For the
	// second and subsequent ListRegexPatternSets requests, specify the value of
	// NextMarker from the previous response to get information about another batch
	// of RegexPatternSet objects.
	NextMarker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListRegexPatternSetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRegexPatternSetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListRegexPatternSetsInput"}
	if s.NextMarker != nil && len(*s.NextMarker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextMarker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListRegexPatternSetsOutput struct {
	_ struct{} `type:"structure"`

	// If you have more RegexPatternSet objects than the number that you specified
	// for Limit in the request, the response includes a NextMarker value. To list
	// more RegexPatternSet objects, submit another ListRegexPatternSets request,
	// and specify the NextMarker value from the response in the NextMarker value
	// in the next request.
	NextMarker *string `min:"1" type:"string"`

	// An array of RegexPatternSetSummary objects.
	RegexPatternSets []RegexPatternSetSummary `type:"list"`
}

// String returns the string representation
func (s ListRegexPatternSetsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListRegexPatternSets = "ListRegexPatternSets"

// ListRegexPatternSetsRequest returns a request value for making API operation for
// AWS WAF Regional.
//
//
// This is AWS WAF Classic documentation. For more information, see AWS WAF
// Classic (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the AWS
// WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
// With the latest version, AWS WAF has a single set of endpoints for regional
// and global use.
//
// Returns an array of RegexPatternSetSummary objects.
//
//    // Example sending a request using ListRegexPatternSetsRequest.
//    req := client.ListRegexPatternSetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/ListRegexPatternSets
func (c *Client) ListRegexPatternSetsRequest(input *ListRegexPatternSetsInput) ListRegexPatternSetsRequest {
	op := &aws.Operation{
		Name:       opListRegexPatternSets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListRegexPatternSetsInput{}
	}

	req := c.newRequest(op, input, &ListRegexPatternSetsOutput{})

	return ListRegexPatternSetsRequest{Request: req, Input: input, Copy: c.ListRegexPatternSetsRequest}
}

// ListRegexPatternSetsRequest is the request type for the
// ListRegexPatternSets API operation.
type ListRegexPatternSetsRequest struct {
	*aws.Request
	Input *ListRegexPatternSetsInput
	Copy  func(*ListRegexPatternSetsInput) ListRegexPatternSetsRequest
}

// Send marshals and sends the ListRegexPatternSets API request.
func (r ListRegexPatternSetsRequest) Send(ctx context.Context) (*ListRegexPatternSetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListRegexPatternSetsResponse{
		ListRegexPatternSetsOutput: r.Request.Data.(*ListRegexPatternSetsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListRegexPatternSetsResponse is the response type for the
// ListRegexPatternSets API operation.
type ListRegexPatternSetsResponse struct {
	*ListRegexPatternSetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListRegexPatternSets request.
func (r *ListRegexPatternSetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
