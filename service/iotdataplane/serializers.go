// Code generated by smithy-go-codegen DO NOT EDIT.

package iotdataplane

import (
	"bytes"
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpDeleteThingShadow struct {
}

func (*awsRestjson1_serializeOpDeleteThingShadow) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDeleteThingShadow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DeleteThingShadowInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/things/{thingName}/shadow")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "DELETE"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsDeleteThingShadowInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsDeleteThingShadowInput(v *DeleteThingShadowInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ShadowName != nil {
		encoder.SetQuery("name").String(*v.ShadowName)
	}

	if v.ThingName == nil || len(*v.ThingName) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member thingName must not be empty")}
	}
	if v.ThingName != nil {
		if err := encoder.SetURI("thingName").String(*v.ThingName); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpGetThingShadow struct {
}

func (*awsRestjson1_serializeOpGetThingShadow) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetThingShadow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetThingShadowInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/things/{thingName}/shadow")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsGetThingShadowInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetThingShadowInput(v *GetThingShadowInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ShadowName != nil {
		encoder.SetQuery("name").String(*v.ShadowName)
	}

	if v.ThingName == nil || len(*v.ThingName) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member thingName must not be empty")}
	}
	if v.ThingName != nil {
		if err := encoder.SetURI("thingName").String(*v.ThingName); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpListNamedShadowsForThing struct {
}

func (*awsRestjson1_serializeOpListNamedShadowsForThing) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpListNamedShadowsForThing) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListNamedShadowsForThingInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/api/things/shadow/ListNamedShadowsForThing/{thingName}")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsListNamedShadowsForThingInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsListNamedShadowsForThingInput(v *ListNamedShadowsForThingInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.NextToken != nil {
		encoder.SetQuery("nextToken").String(*v.NextToken)
	}

	if v.PageSize != nil {
		encoder.SetQuery("pageSize").Integer(*v.PageSize)
	}

	if v.ThingName == nil || len(*v.ThingName) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member thingName must not be empty")}
	}
	if v.ThingName != nil {
		if err := encoder.SetURI("thingName").String(*v.ThingName); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpPublish struct {
}

func (*awsRestjson1_serializeOpPublish) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpPublish) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*PublishInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/topics/{topic}")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsPublishInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if input.Payload != nil {
		if !restEncoder.HasHeader("Content-Type") {
			restEncoder.SetHeader("Content-Type").String("application/octet-stream")
		}

		payload := bytes.NewReader(input.Payload)
		if request, err = request.SetStream(payload); err != nil {
			return out, metadata, &smithy.SerializationError{Err: err}
		}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsPublishInput(v *PublishInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.Qos != 0 {
		encoder.SetQuery("qos").Integer(v.Qos)
	}

	if v.Topic == nil || len(*v.Topic) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member topic must not be empty")}
	}
	if v.Topic != nil {
		if err := encoder.SetURI("topic").String(*v.Topic); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpUpdateThingShadow struct {
}

func (*awsRestjson1_serializeOpUpdateThingShadow) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpUpdateThingShadow) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdateThingShadowInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/things/{thingName}/shadow")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsUpdateThingShadowInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if input.Payload != nil {
		if !restEncoder.HasHeader("Content-Type") {
			restEncoder.SetHeader("Content-Type").String("application/octet-stream")
		}

		payload := bytes.NewReader(input.Payload)
		if request, err = request.SetStream(payload); err != nil {
			return out, metadata, &smithy.SerializationError{Err: err}
		}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsUpdateThingShadowInput(v *UpdateThingShadowInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ShadowName != nil {
		encoder.SetQuery("name").String(*v.ShadowName)
	}

	if v.ThingName == nil || len(*v.ThingName) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member thingName must not be empty")}
	}
	if v.ThingName != nil {
		if err := encoder.SetURI("thingName").String(*v.ThingName); err != nil {
			return err
		}
	}

	return nil
}
