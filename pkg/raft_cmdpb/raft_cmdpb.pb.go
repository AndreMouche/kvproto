// Code generated by protoc-gen-go.
// source: raft_cmdpb.proto
// DO NOT EDIT!

/*
Package raft_cmdpb is a generated protocol buffer package.

It is generated from these files:
	raft_cmdpb.proto

It has these top-level messages:
	GetRequest
	GetResponse
	SeekRequest
	SeekResponse
	PutRequest
	PutResponse
	DeleteRequest
	DeleteResponse
	Request
	Response
	ChangePeerRequest
	ChangePeerResponse
	SplitRequest
	SplitResponse
	CompactLogRequest
	CompactLogResponse
	AdminRequest
	AdminResponse
	RegionLeaderRequest
	RegionLeaderResponse
	RegionDetailRequest
	RegionDetailResponse
	StatusRequest
	StatusResponse
	RaftRequestHeader
	RaftResponseHeader
	RaftCommandRequest
	RaftCommandResponse
*/
package raft_cmdpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import metapb "github.com/pingcap/kvproto/pkg/metapb"
import errorpb "github.com/pingcap/kvproto/pkg/errorpb"
import raftpb "github.com/pingcap/kvproto/pkg/raftpb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type CommandType int32

const (
	CommandType_Invalid CommandType = 0
	CommandType_Get     CommandType = 1
	CommandType_Seek    CommandType = 2
	CommandType_Put     CommandType = 3
	CommandType_Delete  CommandType = 4
)

var CommandType_name = map[int32]string{
	0: "Invalid",
	1: "Get",
	2: "Seek",
	3: "Put",
	4: "Delete",
}
var CommandType_value = map[string]int32{
	"Invalid": 0,
	"Get":     1,
	"Seek":    2,
	"Put":     3,
	"Delete":  4,
}

func (x CommandType) Enum() *CommandType {
	p := new(CommandType)
	*p = x
	return p
}
func (x CommandType) String() string {
	return proto.EnumName(CommandType_name, int32(x))
}
func (x *CommandType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CommandType_value, data, "CommandType")
	if err != nil {
		return err
	}
	*x = CommandType(value)
	return nil
}
func (CommandType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AdminCommandType int32

const (
	AdminCommandType_InvalidAdmin AdminCommandType = 0
	AdminCommandType_ChangePeer   AdminCommandType = 1
	AdminCommandType_Split        AdminCommandType = 2
	AdminCommandType_CompactLog   AdminCommandType = 3
)

var AdminCommandType_name = map[int32]string{
	0: "InvalidAdmin",
	1: "ChangePeer",
	2: "Split",
	3: "CompactLog",
}
var AdminCommandType_value = map[string]int32{
	"InvalidAdmin": 0,
	"ChangePeer":   1,
	"Split":        2,
	"CompactLog":   3,
}

func (x AdminCommandType) Enum() *AdminCommandType {
	p := new(AdminCommandType)
	*p = x
	return p
}
func (x AdminCommandType) String() string {
	return proto.EnumName(AdminCommandType_name, int32(x))
}
func (x *AdminCommandType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AdminCommandType_value, data, "AdminCommandType")
	if err != nil {
		return err
	}
	*x = AdminCommandType(value)
	return nil
}
func (AdminCommandType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type StatusCommandType int32

const (
	StatusCommandType_InvalidStatus StatusCommandType = 0
	StatusCommandType_RegionLeader  StatusCommandType = 1
)

var StatusCommandType_name = map[int32]string{
	0: "InvalidStatus",
	1: "RegionLeader",
}
var StatusCommandType_value = map[string]int32{
	"InvalidStatus": 0,
	"RegionLeader":  1,
}

func (x StatusCommandType) Enum() *StatusCommandType {
	p := new(StatusCommandType)
	*p = x
	return p
}
func (x StatusCommandType) String() string {
	return proto.EnumName(StatusCommandType_name, int32(x))
}
func (x *StatusCommandType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(StatusCommandType_value, data, "StatusCommandType")
	if err != nil {
		return err
	}
	*x = StatusCommandType(value)
	return nil
}
func (StatusCommandType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type GetRequest struct {
	Key              []byte `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type GetResponse struct {
	Value            []byte `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type SeekRequest struct {
	Key              []byte `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SeekRequest) Reset()                    { *m = SeekRequest{} }
func (m *SeekRequest) String() string            { return proto.CompactTextString(m) }
func (*SeekRequest) ProtoMessage()               {}
func (*SeekRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SeekRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type SeekResponse struct {
	Key              []byte `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value            []byte `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SeekResponse) Reset()                    { *m = SeekResponse{} }
func (m *SeekResponse) String() string            { return proto.CompactTextString(m) }
func (*SeekResponse) ProtoMessage()               {}
func (*SeekResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SeekResponse) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *SeekResponse) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type PutRequest struct {
	Key              []byte `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value            []byte `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PutRequest) Reset()                    { *m = PutRequest{} }
func (m *PutRequest) String() string            { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()               {}
func (*PutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PutRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *PutRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type PutResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *PutResponse) Reset()                    { *m = PutResponse{} }
func (m *PutResponse) String() string            { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()               {}
func (*PutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type DeleteRequest struct {
	Key              []byte `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type DeleteResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type Request struct {
	CmdType          *CommandType   `protobuf:"varint,1,opt,name=cmd_type,enum=raft_cmdpb.CommandType" json:"cmd_type,omitempty"`
	Get              *GetRequest    `protobuf:"bytes,2,opt,name=get" json:"get,omitempty"`
	Seek             *SeekRequest   `protobuf:"bytes,3,opt,name=seek" json:"seek,omitempty"`
	Put              *PutRequest    `protobuf:"bytes,4,opt,name=put" json:"put,omitempty"`
	Delete           *DeleteRequest `protobuf:"bytes,5,opt,name=delete" json:"delete,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Request) GetCmdType() CommandType {
	if m != nil && m.CmdType != nil {
		return *m.CmdType
	}
	return CommandType_Invalid
}

func (m *Request) GetGet() *GetRequest {
	if m != nil {
		return m.Get
	}
	return nil
}

func (m *Request) GetSeek() *SeekRequest {
	if m != nil {
		return m.Seek
	}
	return nil
}

func (m *Request) GetPut() *PutRequest {
	if m != nil {
		return m.Put
	}
	return nil
}

func (m *Request) GetDelete() *DeleteRequest {
	if m != nil {
		return m.Delete
	}
	return nil
}

type Response struct {
	CmdType          *CommandType    `protobuf:"varint,1,opt,name=cmd_type,enum=raft_cmdpb.CommandType" json:"cmd_type,omitempty"`
	Get              *GetResponse    `protobuf:"bytes,2,opt,name=get" json:"get,omitempty"`
	Seek             *SeekResponse   `protobuf:"bytes,3,opt,name=seek" json:"seek,omitempty"`
	Put              *PutResponse    `protobuf:"bytes,4,opt,name=put" json:"put,omitempty"`
	Delete           *DeleteResponse `protobuf:"bytes,5,opt,name=delete" json:"delete,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Response) GetCmdType() CommandType {
	if m != nil && m.CmdType != nil {
		return *m.CmdType
	}
	return CommandType_Invalid
}

func (m *Response) GetGet() *GetResponse {
	if m != nil {
		return m.Get
	}
	return nil
}

func (m *Response) GetSeek() *SeekResponse {
	if m != nil {
		return m.Seek
	}
	return nil
}

func (m *Response) GetPut() *PutResponse {
	if m != nil {
		return m.Put
	}
	return nil
}

func (m *Response) GetDelete() *DeleteResponse {
	if m != nil {
		return m.Delete
	}
	return nil
}

type ChangePeerRequest struct {
	ChangeType       *raftpb.ConfChangeType `protobuf:"varint,1,opt,name=change_type,enum=raftpb.ConfChangeType" json:"change_type,omitempty"`
	Peer             *metapb.Peer           `protobuf:"bytes,2,opt,name=peer" json:"peer,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ChangePeerRequest) Reset()                    { *m = ChangePeerRequest{} }
func (m *ChangePeerRequest) String() string            { return proto.CompactTextString(m) }
func (*ChangePeerRequest) ProtoMessage()               {}
func (*ChangePeerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ChangePeerRequest) GetChangeType() raftpb.ConfChangeType {
	if m != nil && m.ChangeType != nil {
		return *m.ChangeType
	}
	return raftpb.ConfChangeType_AddNode
}

func (m *ChangePeerRequest) GetPeer() *metapb.Peer {
	if m != nil {
		return m.Peer
	}
	return nil
}

type ChangePeerResponse struct {
	Region           *metapb.Region `protobuf:"bytes,1,opt,name=region" json:"region,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *ChangePeerResponse) Reset()                    { *m = ChangePeerResponse{} }
func (m *ChangePeerResponse) String() string            { return proto.CompactTextString(m) }
func (*ChangePeerResponse) ProtoMessage()               {}
func (*ChangePeerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ChangePeerResponse) GetRegion() *metapb.Region {
	if m != nil {
		return m.Region
	}
	return nil
}

type SplitRequest struct {
	// The split_key must be in the been splitting region.
	// If the split_key is none, we will choose a proper key
	// to split the region in half.
	SplitKey []byte `protobuf:"bytes,1,opt,name=split_key" json:"split_key,omitempty"`
	// We split the region into two, first uses the origin
	// parent region id, and the second uses the new_region_id.
	// We must guarantee that the new_region_id is global unique.
	NewRegionId *uint64 `protobuf:"varint,2,opt,name=new_region_id" json:"new_region_id,omitempty"`
	// The peer ids for the new split region.
	NewPeerIds       []uint64 `protobuf:"varint,3,rep,name=new_peer_ids" json:"new_peer_ids,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SplitRequest) Reset()                    { *m = SplitRequest{} }
func (m *SplitRequest) String() string            { return proto.CompactTextString(m) }
func (*SplitRequest) ProtoMessage()               {}
func (*SplitRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *SplitRequest) GetSplitKey() []byte {
	if m != nil {
		return m.SplitKey
	}
	return nil
}

func (m *SplitRequest) GetNewRegionId() uint64 {
	if m != nil && m.NewRegionId != nil {
		return *m.NewRegionId
	}
	return 0
}

func (m *SplitRequest) GetNewPeerIds() []uint64 {
	if m != nil {
		return m.NewPeerIds
	}
	return nil
}

type SplitResponse struct {
	Left             *metapb.Region `protobuf:"bytes,1,opt,name=left" json:"left,omitempty"`
	Right            *metapb.Region `protobuf:"bytes,2,opt,name=right" json:"right,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *SplitResponse) Reset()                    { *m = SplitResponse{} }
func (m *SplitResponse) String() string            { return proto.CompactTextString(m) }
func (*SplitResponse) ProtoMessage()               {}
func (*SplitResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *SplitResponse) GetLeft() *metapb.Region {
	if m != nil {
		return m.Left
	}
	return nil
}

func (m *SplitResponse) GetRight() *metapb.Region {
	if m != nil {
		return m.Right
	}
	return nil
}

type CompactLogRequest struct {
	CompactIndex     *uint64 `protobuf:"varint,1,opt,name=compact_index" json:"compact_index,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CompactLogRequest) Reset()                    { *m = CompactLogRequest{} }
func (m *CompactLogRequest) String() string            { return proto.CompactTextString(m) }
func (*CompactLogRequest) ProtoMessage()               {}
func (*CompactLogRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *CompactLogRequest) GetCompactIndex() uint64 {
	if m != nil && m.CompactIndex != nil {
		return *m.CompactIndex
	}
	return 0
}

type CompactLogResponse struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CompactLogResponse) Reset()                    { *m = CompactLogResponse{} }
func (m *CompactLogResponse) String() string            { return proto.CompactTextString(m) }
func (*CompactLogResponse) ProtoMessage()               {}
func (*CompactLogResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

type AdminRequest struct {
	CmdType          *AdminCommandType  `protobuf:"varint,1,opt,name=cmd_type,enum=raft_cmdpb.AdminCommandType" json:"cmd_type,omitempty"`
	ChangePeer       *ChangePeerRequest `protobuf:"bytes,2,opt,name=change_peer" json:"change_peer,omitempty"`
	Split            *SplitRequest      `protobuf:"bytes,3,opt,name=split" json:"split,omitempty"`
	CompactLog       *CompactLogRequest `protobuf:"bytes,4,opt,name=compact_log" json:"compact_log,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *AdminRequest) Reset()                    { *m = AdminRequest{} }
func (m *AdminRequest) String() string            { return proto.CompactTextString(m) }
func (*AdminRequest) ProtoMessage()               {}
func (*AdminRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *AdminRequest) GetCmdType() AdminCommandType {
	if m != nil && m.CmdType != nil {
		return *m.CmdType
	}
	return AdminCommandType_InvalidAdmin
}

func (m *AdminRequest) GetChangePeer() *ChangePeerRequest {
	if m != nil {
		return m.ChangePeer
	}
	return nil
}

func (m *AdminRequest) GetSplit() *SplitRequest {
	if m != nil {
		return m.Split
	}
	return nil
}

func (m *AdminRequest) GetCompactLog() *CompactLogRequest {
	if m != nil {
		return m.CompactLog
	}
	return nil
}

type AdminResponse struct {
	CmdType          *AdminCommandType   `protobuf:"varint,1,opt,name=cmd_type,enum=raft_cmdpb.AdminCommandType" json:"cmd_type,omitempty"`
	ChangePeer       *ChangePeerResponse `protobuf:"bytes,2,opt,name=change_peer" json:"change_peer,omitempty"`
	Split            *SplitResponse      `protobuf:"bytes,3,opt,name=split" json:"split,omitempty"`
	CompactLog       *CompactLogResponse `protobuf:"bytes,4,opt,name=compact_log" json:"compact_log,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *AdminResponse) Reset()                    { *m = AdminResponse{} }
func (m *AdminResponse) String() string            { return proto.CompactTextString(m) }
func (*AdminResponse) ProtoMessage()               {}
func (*AdminResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *AdminResponse) GetCmdType() AdminCommandType {
	if m != nil && m.CmdType != nil {
		return *m.CmdType
	}
	return AdminCommandType_InvalidAdmin
}

func (m *AdminResponse) GetChangePeer() *ChangePeerResponse {
	if m != nil {
		return m.ChangePeer
	}
	return nil
}

func (m *AdminResponse) GetSplit() *SplitResponse {
	if m != nil {
		return m.Split
	}
	return nil
}

func (m *AdminResponse) GetCompactLog() *CompactLogResponse {
	if m != nil {
		return m.CompactLog
	}
	return nil
}

// For get the leader of the region.
type RegionLeaderRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RegionLeaderRequest) Reset()                    { *m = RegionLeaderRequest{} }
func (m *RegionLeaderRequest) String() string            { return proto.CompactTextString(m) }
func (*RegionLeaderRequest) ProtoMessage()               {}
func (*RegionLeaderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

type RegionLeaderResponse struct {
	Leader           *metapb.Peer `protobuf:"bytes,1,opt,name=leader" json:"leader,omitempty"`
	CurrentTerm      *uint64      `protobuf:"varint,2,opt,name=current_term" json:"current_term,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *RegionLeaderResponse) Reset()                    { *m = RegionLeaderResponse{} }
func (m *RegionLeaderResponse) String() string            { return proto.CompactTextString(m) }
func (*RegionLeaderResponse) ProtoMessage()               {}
func (*RegionLeaderResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *RegionLeaderResponse) GetLeader() *metapb.Peer {
	if m != nil {
		return m.Leader
	}
	return nil
}

func (m *RegionLeaderResponse) GetCurrentTerm() uint64 {
	if m != nil && m.CurrentTerm != nil {
		return *m.CurrentTerm
	}
	return 0
}

// For getting more information of the region.
// We add some admin operations (ChangePeer, Split...) into the pb job list,
// then pd server will peek the first one, handle it and then pop it from the job lib.
// But sometimes, the pd server may crash before popping. When another pd server
// starts and finds the job is running but not finished, it will first check whether
// the raft server already has handled this job.
// E,g, for ChangePeer, if we add Peer10 into region1 and find region1 has already had
// Peer10, we can think this ChangePeer is finished, and can pop this job from job list
// directly.
type RegionDetailRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RegionDetailRequest) Reset()                    { *m = RegionDetailRequest{} }
func (m *RegionDetailRequest) String() string            { return proto.CompactTextString(m) }
func (*RegionDetailRequest) ProtoMessage()               {}
func (*RegionDetailRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

type RegionDetailResponse struct {
	Region           *metapb.Region `protobuf:"bytes,1,opt,name=region" json:"region,omitempty"`
	Leader           *metapb.Peer   `protobuf:"bytes,2,opt,name=leader" json:"leader,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *RegionDetailResponse) Reset()                    { *m = RegionDetailResponse{} }
func (m *RegionDetailResponse) String() string            { return proto.CompactTextString(m) }
func (*RegionDetailResponse) ProtoMessage()               {}
func (*RegionDetailResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *RegionDetailResponse) GetRegion() *metapb.Region {
	if m != nil {
		return m.Region
	}
	return nil
}

func (m *RegionDetailResponse) GetLeader() *metapb.Peer {
	if m != nil {
		return m.Leader
	}
	return nil
}

type StatusRequest struct {
	CmdType          *StatusCommandType   `protobuf:"varint,1,opt,name=cmd_type,enum=raft_cmdpb.StatusCommandType" json:"cmd_type,omitempty"`
	RegionLeader     *RegionLeaderRequest `protobuf:"bytes,2,opt,name=region_leader" json:"region_leader,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *StatusRequest) Reset()                    { *m = StatusRequest{} }
func (m *StatusRequest) String() string            { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()               {}
func (*StatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

func (m *StatusRequest) GetCmdType() StatusCommandType {
	if m != nil && m.CmdType != nil {
		return *m.CmdType
	}
	return StatusCommandType_InvalidStatus
}

func (m *StatusRequest) GetRegionLeader() *RegionLeaderRequest {
	if m != nil {
		return m.RegionLeader
	}
	return nil
}

type StatusResponse struct {
	CmdType          *StatusCommandType    `protobuf:"varint,1,opt,name=cmd_type,enum=raft_cmdpb.StatusCommandType" json:"cmd_type,omitempty"`
	RegionLeader     *RegionLeaderResponse `protobuf:"bytes,2,opt,name=region_leader" json:"region_leader,omitempty"`
	CurrentTerm      *uint64               `protobuf:"varint,3,opt,name=current_term" json:"current_term,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{23} }

func (m *StatusResponse) GetCmdType() StatusCommandType {
	if m != nil && m.CmdType != nil {
		return *m.CmdType
	}
	return StatusCommandType_InvalidStatus
}

func (m *StatusResponse) GetRegionLeader() *RegionLeaderResponse {
	if m != nil {
		return m.RegionLeader
	}
	return nil
}

func (m *StatusResponse) GetCurrentTerm() uint64 {
	if m != nil && m.CurrentTerm != nil {
		return *m.CurrentTerm
	}
	return 0
}

type RaftRequestHeader struct {
	RegionId *uint64      `protobuf:"varint,1,opt,name=region_id" json:"region_id,omitempty"`
	Peer     *metapb.Peer `protobuf:"bytes,2,opt,name=peer" json:"peer,omitempty"`
	// true for read linearization
	ReadQuorum *bool `protobuf:"varint,3,opt,name=read_quorum" json:"read_quorum,omitempty"`
	// 16 bytes, to distinguish request.
	Uuid             []byte `protobuf:"bytes,4,opt,name=uuid" json:"uuid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RaftRequestHeader) Reset()                    { *m = RaftRequestHeader{} }
func (m *RaftRequestHeader) String() string            { return proto.CompactTextString(m) }
func (*RaftRequestHeader) ProtoMessage()               {}
func (*RaftRequestHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{24} }

func (m *RaftRequestHeader) GetRegionId() uint64 {
	if m != nil && m.RegionId != nil {
		return *m.RegionId
	}
	return 0
}

func (m *RaftRequestHeader) GetPeer() *metapb.Peer {
	if m != nil {
		return m.Peer
	}
	return nil
}

func (m *RaftRequestHeader) GetReadQuorum() bool {
	if m != nil && m.ReadQuorum != nil {
		return *m.ReadQuorum
	}
	return false
}

func (m *RaftRequestHeader) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

type RaftResponseHeader struct {
	Error            *errorpb.Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Uuid             []byte         `protobuf:"bytes,2,opt,name=uuid" json:"uuid,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *RaftResponseHeader) Reset()                    { *m = RaftResponseHeader{} }
func (m *RaftResponseHeader) String() string            { return proto.CompactTextString(m) }
func (*RaftResponseHeader) ProtoMessage()               {}
func (*RaftResponseHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{25} }

func (m *RaftResponseHeader) GetError() *errorpb.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *RaftResponseHeader) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

type RaftCommandRequest struct {
	Header *RaftRequestHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// We can't enclose normal requests and administrator request
	// at same time.
	Requests         []*Request     `protobuf:"bytes,2,rep,name=requests" json:"requests,omitempty"`
	AdminRequest     *AdminRequest  `protobuf:"bytes,3,opt,name=admin_request" json:"admin_request,omitempty"`
	StatusRequest    *StatusRequest `protobuf:"bytes,4,opt,name=status_request" json:"status_request,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *RaftCommandRequest) Reset()                    { *m = RaftCommandRequest{} }
func (m *RaftCommandRequest) String() string            { return proto.CompactTextString(m) }
func (*RaftCommandRequest) ProtoMessage()               {}
func (*RaftCommandRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{26} }

func (m *RaftCommandRequest) GetHeader() *RaftRequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *RaftCommandRequest) GetRequests() []*Request {
	if m != nil {
		return m.Requests
	}
	return nil
}

func (m *RaftCommandRequest) GetAdminRequest() *AdminRequest {
	if m != nil {
		return m.AdminRequest
	}
	return nil
}

func (m *RaftCommandRequest) GetStatusRequest() *StatusRequest {
	if m != nil {
		return m.StatusRequest
	}
	return nil
}

type RaftCommandResponse struct {
	Header           *RaftResponseHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Responses        []*Response         `protobuf:"bytes,2,rep,name=responses" json:"responses,omitempty"`
	AdminResponse    *AdminResponse      `protobuf:"bytes,3,opt,name=admin_response" json:"admin_response,omitempty"`
	StatusResponse   *StatusResponse     `protobuf:"bytes,4,opt,name=status_response" json:"status_response,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *RaftCommandResponse) Reset()                    { *m = RaftCommandResponse{} }
func (m *RaftCommandResponse) String() string            { return proto.CompactTextString(m) }
func (*RaftCommandResponse) ProtoMessage()               {}
func (*RaftCommandResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{27} }

func (m *RaftCommandResponse) GetHeader() *RaftResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *RaftCommandResponse) GetResponses() []*Response {
	if m != nil {
		return m.Responses
	}
	return nil
}

func (m *RaftCommandResponse) GetAdminResponse() *AdminResponse {
	if m != nil {
		return m.AdminResponse
	}
	return nil
}

func (m *RaftCommandResponse) GetStatusResponse() *StatusResponse {
	if m != nil {
		return m.StatusResponse
	}
	return nil
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "raft_cmdpb.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "raft_cmdpb.GetResponse")
	proto.RegisterType((*SeekRequest)(nil), "raft_cmdpb.SeekRequest")
	proto.RegisterType((*SeekResponse)(nil), "raft_cmdpb.SeekResponse")
	proto.RegisterType((*PutRequest)(nil), "raft_cmdpb.PutRequest")
	proto.RegisterType((*PutResponse)(nil), "raft_cmdpb.PutResponse")
	proto.RegisterType((*DeleteRequest)(nil), "raft_cmdpb.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "raft_cmdpb.DeleteResponse")
	proto.RegisterType((*Request)(nil), "raft_cmdpb.Request")
	proto.RegisterType((*Response)(nil), "raft_cmdpb.Response")
	proto.RegisterType((*ChangePeerRequest)(nil), "raft_cmdpb.ChangePeerRequest")
	proto.RegisterType((*ChangePeerResponse)(nil), "raft_cmdpb.ChangePeerResponse")
	proto.RegisterType((*SplitRequest)(nil), "raft_cmdpb.SplitRequest")
	proto.RegisterType((*SplitResponse)(nil), "raft_cmdpb.SplitResponse")
	proto.RegisterType((*CompactLogRequest)(nil), "raft_cmdpb.CompactLogRequest")
	proto.RegisterType((*CompactLogResponse)(nil), "raft_cmdpb.CompactLogResponse")
	proto.RegisterType((*AdminRequest)(nil), "raft_cmdpb.AdminRequest")
	proto.RegisterType((*AdminResponse)(nil), "raft_cmdpb.AdminResponse")
	proto.RegisterType((*RegionLeaderRequest)(nil), "raft_cmdpb.RegionLeaderRequest")
	proto.RegisterType((*RegionLeaderResponse)(nil), "raft_cmdpb.RegionLeaderResponse")
	proto.RegisterType((*RegionDetailRequest)(nil), "raft_cmdpb.RegionDetailRequest")
	proto.RegisterType((*RegionDetailResponse)(nil), "raft_cmdpb.RegionDetailResponse")
	proto.RegisterType((*StatusRequest)(nil), "raft_cmdpb.StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "raft_cmdpb.StatusResponse")
	proto.RegisterType((*RaftRequestHeader)(nil), "raft_cmdpb.RaftRequestHeader")
	proto.RegisterType((*RaftResponseHeader)(nil), "raft_cmdpb.RaftResponseHeader")
	proto.RegisterType((*RaftCommandRequest)(nil), "raft_cmdpb.RaftCommandRequest")
	proto.RegisterType((*RaftCommandResponse)(nil), "raft_cmdpb.RaftCommandResponse")
	proto.RegisterEnum("raft_cmdpb.CommandType", CommandType_name, CommandType_value)
	proto.RegisterEnum("raft_cmdpb.AdminCommandType", AdminCommandType_name, AdminCommandType_value)
	proto.RegisterEnum("raft_cmdpb.StatusCommandType", StatusCommandType_name, StatusCommandType_value)
}

var fileDescriptor0 = []byte{
	// 951 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0x6d, 0x6f, 0xe3, 0x44,
	0x10, 0xc6, 0x67, 0x27, 0xcd, 0x8d, 0xed, 0xe0, 0xb8, 0x39, 0x9a, 0x46, 0x6d, 0x38, 0x99, 0x3b,
	0x2e, 0x17, 0x44, 0x10, 0x29, 0x02, 0xbe, 0x1e, 0x77, 0x88, 0x17, 0x55, 0x50, 0x95, 0x7e, 0x44,
	0xb2, 0x4c, 0xbc, 0x4d, 0xad, 0x26, 0xb6, 0xeb, 0x17, 0x68, 0xff, 0x05, 0xbf, 0x0c, 0x90, 0x10,
	0xfc, 0x1e, 0xd6, 0xbb, 0xb3, 0xce, 0x6e, 0xec, 0xa2, 0xeb, 0xa7, 0x28, 0xb3, 0x33, 0xcf, 0xcc,
	0xf3, 0xcc, 0xec, 0xac, 0xc1, 0xc9, 0x82, 0xcb, 0xc2, 0x5f, 0x6e, 0xc2, 0xf4, 0x97, 0x79, 0x9a,
	0x25, 0x45, 0xe2, 0xc2, 0xd6, 0x32, 0xb6, 0x36, 0xa4, 0x08, 0xc4, 0xc9, 0xd8, 0x26, 0x59, 0x96,
	0x64, 0xf5, 0x5f, 0xab, 0x72, 0x14, 0xff, 0xbc, 0x43, 0x80, 0x6f, 0x48, 0x71, 0x4e, 0x6e, 0x4a,
	0x92, 0x17, 0xae, 0x09, 0xfa, 0x35, 0xb9, 0x1b, 0x69, 0x4f, 0xb5, 0xa9, 0xe5, 0x1d, 0x81, 0xc9,
	0x8e, 0xf2, 0x34, 0x89, 0x73, 0xe2, 0xda, 0xd0, 0xf9, 0x35, 0x58, 0x97, 0x04, 0x4f, 0xc7, 0x60,
	0xfe, 0x44, 0xc8, 0x75, 0x6b, 0xe4, 0x0c, 0x2c, 0x7e, 0x86, 0xa1, 0xf2, 0xe1, 0x16, 0xe7, 0x11,
	0xf3, 0x9d, 0x02, 0x9c, 0x95, 0xad, 0x05, 0xec, 0x7a, 0xda, 0x60, 0x32, 0x4f, 0x0e, 0x4a, 0xcb,
	0xb3, 0xdf, 0x90, 0x35, 0x29, 0x48, 0x6b, 0x09, 0x0e, 0xf4, 0xc5, 0x29, 0xfa, 0xff, 0xa5, 0xc1,
	0x9e, 0x70, 0x7d, 0x09, 0x3d, 0xaa, 0x94, 0x5f, 0xdc, 0xa5, 0x9c, 0x4e, 0x7f, 0x71, 0x30, 0x97,
	0x14, 0x7d, 0x9d, 0x6c, 0x36, 0x41, 0x1c, 0x5e, 0xd0, 0x63, 0xf7, 0x03, 0xd0, 0x57, 0xa4, 0x60,
	0x25, 0x98, 0x8b, 0xf7, 0x64, 0x2f, 0x49, 0xb7, 0xe7, 0x60, 0xe4, 0x94, 0xf0, 0x48, 0x67, 0x5e,
	0x0a, 0x96, 0x2c, 0x12, 0xc5, 0x4a, 0xcb, 0x62, 0x64, 0x34, 0xb1, 0x24, 0x09, 0x5e, 0x42, 0x37,
	0x64, 0x95, 0x8f, 0x3a, 0xcc, 0xef, 0x50, 0xf6, 0x53, 0x18, 0x7b, 0xff, 0x68, 0xd0, 0xab, 0x45,
	0x7e, 0x00, 0xa7, 0x67, 0x32, 0xa7, 0x83, 0x06, 0x27, 0x04, 0xfc, 0x50, 0x21, 0x35, 0x6a, 0x92,
	0x42, 0xbf, 0x67, 0x32, 0xab, 0x83, 0x06, 0x2b, 0xf4, 0x9a, 0xed, 0xd0, 0x1a, 0xb7, 0xd1, 0xc2,
	0x56, 0xfd, 0x0c, 0x83, 0xd7, 0x57, 0x41, 0xbc, 0x22, 0x67, 0x84, 0x64, 0x42, 0x97, 0x8f, 0xc0,
	0x5c, 0x32, 0xa3, 0x4c, 0x91, 0x8b, 0xc8, 0xe8, 0xc5, 0x97, 0x3c, 0x86, 0x31, 0x1c, 0x83, 0x91,
	0xd2, 0x58, 0xa4, 0x68, 0xcd, 0xf1, 0x42, 0x54, 0x78, 0xde, 0x67, 0xe0, 0xca, 0xe8, 0x58, 0xdf,
	0x04, 0xba, 0x19, 0x59, 0x45, 0x49, 0xcc, 0x90, 0xcd, 0x45, 0x5f, 0xc4, 0x9c, 0x33, 0xab, 0xf7,
	0x03, 0x9d, 0xe9, 0x74, 0x1d, 0xd5, 0x6d, 0x1a, 0xc0, 0xe3, 0xbc, 0xfa, 0xef, 0x6f, 0xe7, 0xf5,
	0x09, 0xd8, 0x31, 0xf9, 0xcd, 0xe7, 0x30, 0x7e, 0x14, 0xb2, 0xec, 0x86, 0x3b, 0x04, 0xab, 0x32,
	0x57, 0xf5, 0x50, 0x63, 0x4e, 0xf5, 0xd4, 0xa7, 0x86, 0x77, 0x0a, 0x36, 0xe2, 0x61, 0x01, 0x47,
	0x60, 0xac, 0xc9, 0x65, 0xd1, 0x9e, 0xde, 0x3d, 0x86, 0x4e, 0x16, 0xad, 0xae, 0x44, 0xd3, 0x76,
	0xab, 0x9b, 0x51, 0xc5, 0x92, 0x4d, 0x1a, 0x2c, 0x8b, 0xd3, 0x64, 0x25, 0x4a, 0xa4, 0xf5, 0x2c,
	0xb9, 0xd1, 0x8f, 0xe2, 0x90, 0xdc, 0x32, 0x68, 0xc3, 0x1b, 0x52, 0xfe, 0x92, 0x2f, 0x6a, 0xfe,
	0xa7, 0x06, 0xd6, 0xab, 0x70, 0x13, 0xc5, 0x22, 0x7a, 0xde, 0x98, 0xa7, 0x23, 0xb9, 0x65, 0xcc,
	0x57, 0x1e, 0xaa, 0x45, 0xdd, 0x1f, 0x49, 0xf9, 0x63, 0x65, 0x04, 0x1b, 0x3d, 0x7d, 0x01, 0x1d,
	0x26, 0x62, 0xeb, 0x8c, 0xc9, 0x6a, 0x57, 0xe0, 0x48, 0x65, 0x9d, 0xac, 0x70, 0xd6, 0x8e, 0x77,
	0xe6, 0x5b, 0xa5, 0xef, 0xfd, 0xad, 0x81, 0x8d, 0x8c, 0x50, 0xe2, 0x87, 0x52, 0x3a, 0x69, 0xa3,
	0x34, 0xb9, 0x8f, 0x12, 0x26, 0x99, 0xaa, 0x9c, 0x0e, 0x5b, 0x38, 0xa1, 0xe7, 0x49, 0x1b, 0xa9,
	0xc9, 0x7d, 0xa4, 0xb0, 0x4f, 0x4f, 0x60, 0x9f, 0xf7, 0xfc, 0x94, 0x04, 0x61, 0xad, 0xa4, 0xf7,
	0x3d, 0x0c, 0x55, 0x73, 0x3d, 0x55, 0xdd, 0x35, 0xb3, 0xe0, 0x5c, 0x29, 0x57, 0xa1, 0x1a, 0xcd,
	0x65, 0x99, 0x65, 0x24, 0x2e, 0xfc, 0x82, 0x64, 0x1b, 0x3e, 0xb0, 0xdb, 0x14, 0x6f, 0xa8, 0x6b,
	0xb4, 0x16, 0x29, 0x2e, 0x44, 0x0a, 0x61, 0x7e, 0xbb, 0x9b, 0x23, 0x95, 0xd0, 0x76, 0x1b, 0x6f,
	0xe9, 0x3d, 0x28, 0x82, 0xa2, 0xcc, 0x45, 0xab, 0x3f, 0x69, 0x34, 0x49, 0xe9, 0x33, 0x77, 0x96,
	0xbb, 0xf4, 0x39, 0xd8, 0x78, 0xe5, 0x94, 0x34, 0xef, 0xcb, 0x51, 0x6d, 0x92, 0xfd, 0xae, 0x41,
	0x5f, 0xa4, 0x46, 0x2a, 0x0f, 0xce, 0xfd, 0x45, 0x7b, 0xee, 0xa7, 0xf7, 0xe7, 0xc6, 0x4c, 0xbb,
	0xca, 0xeb, 0x4c, 0x79, 0x02, 0x83, 0x73, 0x1a, 0x88, 0x15, 0x7e, 0xcb, 0x42, 0xaa, 0x4d, 0xb3,
	0x5d, 0x29, 0xec, 0x0a, 0xff, 0xdf, 0x7a, 0x73, 0xf7, 0xc1, 0xcc, 0x68, 0xa0, 0x7f, 0x53, 0x26,
	0x59, 0xc9, 0x81, 0x7b, 0xae, 0x05, 0x46, 0x59, 0xd2, 0x70, 0x83, 0x3d, 0x8e, 0xaf, 0xc0, 0xe5,
	0x69, 0x78, 0x31, 0x98, 0x87, 0xae, 0x18, 0xf6, 0xa5, 0x50, 0xb7, 0x51, 0x7c, 0x37, 0x7c, 0x5d,
	0xfd, 0xd6, 0x10, 0xfc, 0x31, 0xfe, 0x43, 0xe3, 0x18, 0x28, 0x86, 0x68, 0xde, 0xc7, 0xd0, 0xbd,
	0x92, 0xc7, 0x4d, 0x91, 0xaf, 0x49, 0xed, 0x39, 0xf4, 0x32, 0x6e, 0xc8, 0x29, 0xae, 0x4e, 0x03,
	0xf6, 0x55, 0xe5, 0xc4, 0x48, 0xd8, 0x41, 0x75, 0x37, 0x7d, 0x74, 0x6e, 0x5b, 0x17, 0xca, 0xee,
	0xfa, 0x14, 0xfa, 0x39, 0xeb, 0x55, 0x1d, 0x61, 0xb4, 0x5c, 0x46, 0x79, 0xec, 0xbc, 0x7f, 0x35,
	0x3a, 0xf5, 0x32, 0xa1, 0x7a, 0x67, 0xa8, 0x8c, 0x26, 0x4d, 0x46, 0x8a, 0x8a, 0x2f, 0xaa, 0x6e,
	0x71, 0x8b, 0xe0, 0x34, 0x54, 0x39, 0x21, 0x30, 0xad, 0x51, 0x90, 0xe2, 0x96, 0xb6, 0x85, 0xa1,
	0xee, 0xaf, 0x13, 0x78, 0xb7, 0xa6, 0x85, 0x31, 0x46, 0xf3, 0x31, 0x55, 0x67, 0x7a, 0xf6, 0x15,
	0x98, 0xf2, 0xc4, 0x9a, 0xb0, 0xf7, 0x5d, 0x4c, 0x3f, 0xab, 0xa2, 0xd0, 0x79, 0xc7, 0xdd, 0x03,
	0x9d, 0xbe, 0xf8, 0x8e, 0xe6, 0xf6, 0xc0, 0xa8, 0xde, 0x74, 0xe7, 0x51, 0x65, 0xa2, 0xcf, 0xb6,
	0xa3, 0xbb, 0x00, 0x5d, 0xfe, 0x2c, 0x3b, 0xc6, 0xec, 0x47, 0x70, 0x1a, 0xcb, 0xd1, 0x01, 0x0b,
	0x81, 0xd8, 0x11, 0x45, 0xeb, 0x03, 0x6c, 0xf7, 0x21, 0x05, 0x7d, 0x0c, 0x1d, 0xb6, 0xf0, 0x28,
	0x6a, 0x75, 0x54, 0xef, 0x32, 0x47, 0x9f, 0x7d, 0x09, 0x83, 0xe6, 0x65, 0x1a, 0x80, 0x8d, 0x88,
	0xfc, 0x8c, 0x42, 0xd2, 0x24, 0xf2, 0xf5, 0x71, 0xb4, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x51,
	0x1f, 0xe4, 0xb7, 0xfa, 0x0a, 0x00, 0x00,
}
