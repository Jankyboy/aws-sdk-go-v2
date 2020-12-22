// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

type ComplexNestedErrorData struct {
	Foo *string
}

// A union with a representative set of types for members.
//
// The following types satisfy this interface:
//  MyUnionMemberStringValue
//  MyUnionMemberBooleanValue
//  MyUnionMemberNumberValue
//  MyUnionMemberBlobValue
//  MyUnionMemberTimestampValue
//  MyUnionMemberEnumValue
//  MyUnionMemberListValue
//  MyUnionMemberMapValue
//  MyUnionMemberStructureValue
type MyUnion interface {
	isMyUnion()
}

type MyUnionMemberStringValue struct {
	Value string
}

func (*MyUnionMemberStringValue) isMyUnion() {}

type MyUnionMemberBooleanValue struct {
	Value bool
}

func (*MyUnionMemberBooleanValue) isMyUnion() {}

type MyUnionMemberNumberValue struct {
	Value int32
}

func (*MyUnionMemberNumberValue) isMyUnion() {}

type MyUnionMemberBlobValue struct {
	Value []byte
}

func (*MyUnionMemberBlobValue) isMyUnion() {}

type MyUnionMemberTimestampValue struct {
	Value time.Time
}

func (*MyUnionMemberTimestampValue) isMyUnion() {}

type MyUnionMemberEnumValue struct {
	Value FooEnum
}

func (*MyUnionMemberEnumValue) isMyUnion() {}

type MyUnionMemberListValue struct {
	Value []string
}

func (*MyUnionMemberListValue) isMyUnion() {}

type MyUnionMemberMapValue struct {
	Value map[string]string
}

func (*MyUnionMemberMapValue) isMyUnion() {}

type MyUnionMemberStructureValue struct {
	Value GreetingStruct
}

func (*MyUnionMemberStructureValue) isMyUnion() {}

type GreetingStruct struct {
	Hi *string
}

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte
}

func (*UnknownUnionMember) isMyUnion() {}
