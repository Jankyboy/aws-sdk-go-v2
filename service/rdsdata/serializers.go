// Code generated by smithy-go-codegen DO NOT EDIT.

package rdsdata

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/rdsdata/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpBatchExecuteStatement struct {
}

func (*awsRestjson1_serializeOpBatchExecuteStatement) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpBatchExecuteStatement) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*BatchExecuteStatementInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/BatchExecute")
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
	if err := awsRestjson1_serializeOpDocumentBatchExecuteStatementInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsBatchExecuteStatementInput(v *BatchExecuteStatementInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentBatchExecuteStatementInput(v *BatchExecuteStatementInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Database != nil {
		ok := object.Key("database")
		ok.String(*v.Database)
	}

	if v.ParameterSets != nil {
		ok := object.Key("parameterSets")
		if err := awsRestjson1_serializeDocumentSqlParameterSets(v.ParameterSets, ok); err != nil {
			return err
		}
	}

	if v.ResourceArn != nil {
		ok := object.Key("resourceArn")
		ok.String(*v.ResourceArn)
	}

	if v.Schema != nil {
		ok := object.Key("schema")
		ok.String(*v.Schema)
	}

	if v.SecretArn != nil {
		ok := object.Key("secretArn")
		ok.String(*v.SecretArn)
	}

	if v.Sql != nil {
		ok := object.Key("sql")
		ok.String(*v.Sql)
	}

	if v.TransactionId != nil {
		ok := object.Key("transactionId")
		ok.String(*v.TransactionId)
	}

	return nil
}

type awsRestjson1_serializeOpBeginTransaction struct {
}

func (*awsRestjson1_serializeOpBeginTransaction) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpBeginTransaction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*BeginTransactionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/BeginTransaction")
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
	if err := awsRestjson1_serializeOpDocumentBeginTransactionInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsBeginTransactionInput(v *BeginTransactionInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentBeginTransactionInput(v *BeginTransactionInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Database != nil {
		ok := object.Key("database")
		ok.String(*v.Database)
	}

	if v.ResourceArn != nil {
		ok := object.Key("resourceArn")
		ok.String(*v.ResourceArn)
	}

	if v.Schema != nil {
		ok := object.Key("schema")
		ok.String(*v.Schema)
	}

	if v.SecretArn != nil {
		ok := object.Key("secretArn")
		ok.String(*v.SecretArn)
	}

	return nil
}

type awsRestjson1_serializeOpCommitTransaction struct {
}

func (*awsRestjson1_serializeOpCommitTransaction) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpCommitTransaction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CommitTransactionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/CommitTransaction")
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
	if err := awsRestjson1_serializeOpDocumentCommitTransactionInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsCommitTransactionInput(v *CommitTransactionInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentCommitTransactionInput(v *CommitTransactionInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceArn != nil {
		ok := object.Key("resourceArn")
		ok.String(*v.ResourceArn)
	}

	if v.SecretArn != nil {
		ok := object.Key("secretArn")
		ok.String(*v.SecretArn)
	}

	if v.TransactionId != nil {
		ok := object.Key("transactionId")
		ok.String(*v.TransactionId)
	}

	return nil
}

type awsRestjson1_serializeOpExecuteSql struct {
}

func (*awsRestjson1_serializeOpExecuteSql) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpExecuteSql) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ExecuteSqlInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/ExecuteSql")
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
	if err := awsRestjson1_serializeOpDocumentExecuteSqlInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsExecuteSqlInput(v *ExecuteSqlInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentExecuteSqlInput(v *ExecuteSqlInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AwsSecretStoreArn != nil {
		ok := object.Key("awsSecretStoreArn")
		ok.String(*v.AwsSecretStoreArn)
	}

	if v.Database != nil {
		ok := object.Key("database")
		ok.String(*v.Database)
	}

	if v.DbClusterOrInstanceArn != nil {
		ok := object.Key("dbClusterOrInstanceArn")
		ok.String(*v.DbClusterOrInstanceArn)
	}

	if v.Schema != nil {
		ok := object.Key("schema")
		ok.String(*v.Schema)
	}

	if v.SqlStatements != nil {
		ok := object.Key("sqlStatements")
		ok.String(*v.SqlStatements)
	}

	return nil
}

type awsRestjson1_serializeOpExecuteStatement struct {
}

func (*awsRestjson1_serializeOpExecuteStatement) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpExecuteStatement) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ExecuteStatementInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/Execute")
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
	if err := awsRestjson1_serializeOpDocumentExecuteStatementInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsExecuteStatementInput(v *ExecuteStatementInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentExecuteStatementInput(v *ExecuteStatementInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ContinueAfterTimeout {
		ok := object.Key("continueAfterTimeout")
		ok.Boolean(v.ContinueAfterTimeout)
	}

	if v.Database != nil {
		ok := object.Key("database")
		ok.String(*v.Database)
	}

	if v.IncludeResultMetadata {
		ok := object.Key("includeResultMetadata")
		ok.Boolean(v.IncludeResultMetadata)
	}

	if v.Parameters != nil {
		ok := object.Key("parameters")
		if err := awsRestjson1_serializeDocumentSqlParametersList(v.Parameters, ok); err != nil {
			return err
		}
	}

	if v.ResourceArn != nil {
		ok := object.Key("resourceArn")
		ok.String(*v.ResourceArn)
	}

	if v.ResultSetOptions != nil {
		ok := object.Key("resultSetOptions")
		if err := awsRestjson1_serializeDocumentResultSetOptions(v.ResultSetOptions, ok); err != nil {
			return err
		}
	}

	if v.Schema != nil {
		ok := object.Key("schema")
		ok.String(*v.Schema)
	}

	if v.SecretArn != nil {
		ok := object.Key("secretArn")
		ok.String(*v.SecretArn)
	}

	if v.Sql != nil {
		ok := object.Key("sql")
		ok.String(*v.Sql)
	}

	if v.TransactionId != nil {
		ok := object.Key("transactionId")
		ok.String(*v.TransactionId)
	}

	return nil
}

type awsRestjson1_serializeOpRollbackTransaction struct {
}

func (*awsRestjson1_serializeOpRollbackTransaction) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpRollbackTransaction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*RollbackTransactionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/RollbackTransaction")
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
	if err := awsRestjson1_serializeOpDocumentRollbackTransactionInput(input, jsonEncoder.Value); err != nil {
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
func awsRestjson1_serializeOpHttpBindingsRollbackTransactionInput(v *RollbackTransactionInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	return nil
}

func awsRestjson1_serializeOpDocumentRollbackTransactionInput(v *RollbackTransactionInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ResourceArn != nil {
		ok := object.Key("resourceArn")
		ok.String(*v.ResourceArn)
	}

	if v.SecretArn != nil {
		ok := object.Key("secretArn")
		ok.String(*v.SecretArn)
	}

	if v.TransactionId != nil {
		ok := object.Key("transactionId")
		ok.String(*v.TransactionId)
	}

	return nil
}

func awsRestjson1_serializeDocumentArrayOfArray(v []types.ArrayValue, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			continue
		}
		if err := awsRestjson1_serializeDocumentArrayValue(v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentArrayValue(v types.ArrayValue, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	switch uv := v.(type) {
	case *types.ArrayValueMemberArrayValues:
		av := object.Key("arrayValues")
		if err := awsRestjson1_serializeDocumentArrayOfArray(uv.Value, av); err != nil {
			return err
		}

	case *types.ArrayValueMemberBooleanValues:
		av := object.Key("booleanValues")
		if err := awsRestjson1_serializeDocumentBooleanArray(uv.Value, av); err != nil {
			return err
		}

	case *types.ArrayValueMemberDoubleValues:
		av := object.Key("doubleValues")
		if err := awsRestjson1_serializeDocumentDoubleArray(uv.Value, av); err != nil {
			return err
		}

	case *types.ArrayValueMemberLongValues:
		av := object.Key("longValues")
		if err := awsRestjson1_serializeDocumentLongArray(uv.Value, av); err != nil {
			return err
		}

	case *types.ArrayValueMemberStringValues:
		av := object.Key("stringValues")
		if err := awsRestjson1_serializeDocumentStringArray(uv.Value, av); err != nil {
			return err
		}

	default:
		return fmt.Errorf("attempted to serialize unknown member type %T for union %T", uv, v)

	}
	return nil
}

func awsRestjson1_serializeDocumentBooleanArray(v []bool, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.Boolean(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentDoubleArray(v []float64, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.Double(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentField(v types.Field, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	switch uv := v.(type) {
	case *types.FieldMemberArrayValue:
		av := object.Key("arrayValue")
		if err := awsRestjson1_serializeDocumentArrayValue(uv.Value, av); err != nil {
			return err
		}

	case *types.FieldMemberBlobValue:
		av := object.Key("blobValue")
		av.Base64EncodeBytes(uv.Value)

	case *types.FieldMemberBooleanValue:
		av := object.Key("booleanValue")
		av.Boolean(uv.Value)

	case *types.FieldMemberDoubleValue:
		av := object.Key("doubleValue")
		av.Double(uv.Value)

	case *types.FieldMemberIsNull:
		av := object.Key("isNull")
		av.Boolean(uv.Value)

	case *types.FieldMemberLongValue:
		av := object.Key("longValue")
		av.Long(uv.Value)

	case *types.FieldMemberStringValue:
		av := object.Key("stringValue")
		av.String(uv.Value)

	default:
		return fmt.Errorf("attempted to serialize unknown member type %T for union %T", uv, v)

	}
	return nil
}

func awsRestjson1_serializeDocumentLongArray(v []int64, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.Long(v[i])
	}
	return nil
}

func awsRestjson1_serializeDocumentResultSetOptions(v *types.ResultSetOptions, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.DecimalReturnType) > 0 {
		ok := object.Key("decimalReturnType")
		ok.String(string(v.DecimalReturnType))
	}

	return nil
}

func awsRestjson1_serializeDocumentSqlParameter(v *types.SqlParameter, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Name != nil {
		ok := object.Key("name")
		ok.String(*v.Name)
	}

	if len(v.TypeHint) > 0 {
		ok := object.Key("typeHint")
		ok.String(string(v.TypeHint))
	}

	if v.Value != nil {
		ok := object.Key("value")
		if err := awsRestjson1_serializeDocumentField(v.Value, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentSqlParameterSets(v [][]types.SqlParameter, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			continue
		}
		if err := awsRestjson1_serializeDocumentSqlParametersList(v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentSqlParametersList(v []types.SqlParameter, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsRestjson1_serializeDocumentSqlParameter(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsRestjson1_serializeDocumentStringArray(v []string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(v[i])
	}
	return nil
}
