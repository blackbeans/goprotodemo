// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

package code_blackbeans_com

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type User struct {
	Uid              *int32  `protobuf:"varint,1,req,name=uid" json:"uid,omitempty"`
	Name             *string `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *User) Reset()         { *this = User{} }
func (this *User) String() string { return proto.CompactTextString(this) }
func (*User) ProtoMessage()       {}

func (this *User) GetUid() int32 {
	if this != nil && this.Uid != nil {
		return *this.Uid
	}
	return 0
}

func (this *User) GetName() string {
	if this != nil && this.Name != nil {
		return *this.Name
	}
	return ""
}

func init() {
}