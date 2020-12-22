// Code generated by smithy-go-codegen DO NOT EDIT.

package personalizeevents

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/personalizeevents/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpPutEvents struct {
}

func (*awsRestjson1_serializeOpPutEvents) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpPutEvents) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*PutEventsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/events")
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

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentPutEventsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsPutEventsInput(v *PutEventsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentPutEventsInput(v *PutEventsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.EventList != nil {
		ok := object.Key("eventList")
		if err := awsRestjson1_serializeDocumentEventList(v.EventList, ok); err != nil {
			return err
		}
	}

	if v.SessionId != nil {
		ok := object.Key("sessionId")
		ok.String(*v.SessionId)
	}

	if v.TrackingId != nil {
		ok := object.Key("trackingId")
		ok.String(*v.TrackingId)
	}

	if v.UserId != nil {
		ok := object.Key("userId")
		ok.String(*v.UserId)
	}

	return nil
}

type awsRestjson1_serializeOpPutItems struct {
}

func (*awsRestjson1_serializeOpPutItems) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpPutItems) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*PutItemsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/items")
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

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentPutItemsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsPutItemsInput(v *PutItemsInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentPutItemsInput(v *PutItemsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DatasetArn != nil {
		ok := object.Key("datasetArn")
		ok.String(*v.DatasetArn)
	}

	if v.Items != nil {
		ok := object.Key("items")
		if err := awsRestjson1_serializeDocumentItemList(v.Items, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpPutUsers struct {
}

func (*awsRestjson1_serializeOpPutUsers) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpPutUsers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*PutUsersInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/users")
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

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentPutUsersInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsPutUsersInput(v *PutUsersInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentPutUsersInput(v *PutUsersInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DatasetArn != nil {
		ok := object.Key("datasetArn")
		ok.String(*v.DatasetArn)
	}

	if v.Users != nil {
		ok := object.Key("users")
		if err := awsRestjson1_serializeDocumentUserList(v.Users, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentEvent(v *types.Event, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.EventId != nil {
		ok := object.Key("eventId")
		ok.String(*v.EventId)
	}

	if v.EventType != nil {
		ok := object.Key("eventType")
		ok.String(*v.EventType)
	}

	if v.EventValue != nil {
		ok := object.Key("eventValue")
		ok.Float(*v.EventValue)
	}

	if v.Impression != nil {
		ok := object.Key("impression")
		if err := awsRestjson1_serializeDocumentImpression(v.Impression, ok); err != nil {
			return err
		}
	}

	if v.ItemId != nil {
		ok := object.Key("itemId")
		ok.String(*v.ItemId)
	}

	if v.Properties != nil {
		ok := object.Key("properties")
		ok.String(*v.Properties)
	}

	if v.RecommendationId != nil {
		ok := object.Key("recommendationId")
		ok.String(*v.RecommendationId)
	}

	if v.SentAt != nil {
		ok := object.Key("sentAt")
		ok.Double(smithytime.FormatEpochSeconds(*v.SentAt))
	}

	return nil
}

func awsRestjson1_serializeDocumentEventList(v []types.Event, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentEvent(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentImpression(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentItem(v *types.Item, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ItemId != nil {
		ok := object.Key("itemId")
		ok.String(*v.ItemId)
	}

	if v.Properties != nil {
		ok := object.Key("properties")
		ok.String(*v.Properties)
	}

	return nil
}

func awsRestjson1_serializeDocumentItemList(v []types.Item, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentItem(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentUser(v *types.User, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Properties != nil {
		ok := object.Key("properties")
		ok.String(*v.Properties)
	}

	if v.UserId != nil {
		ok := object.Key("userId")
		ok.String(*v.UserId)
	}

	return nil
}

func awsRestjson1_serializeDocumentUserList(v []types.User, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentUser(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}
