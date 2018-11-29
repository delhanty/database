// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: coordinator/coordinatorservicer.proto

package coordinator // import "github.com/linkai-io/am/protocservices/coordinator"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import prototypes "github.com/linkai-io/am/protocservices/prototypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StartGroupRequest struct {
	UserContext          *prototypes.UserContext `protobuf:"bytes,1,opt,name=UserContext" json:"UserContext,omitempty"`
	GroupID              int32                   `protobuf:"varint,2,opt,name=GroupID,proto3" json:"GroupID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *StartGroupRequest) Reset()         { *m = StartGroupRequest{} }
func (m *StartGroupRequest) String() string { return proto.CompactTextString(m) }
func (*StartGroupRequest) ProtoMessage()    {}
func (*StartGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_coordinatorservicer_c7a0a16d6559a9fe, []int{0}
}
func (m *StartGroupRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StartGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StartGroupRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *StartGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartGroupRequest.Merge(dst, src)
}
func (m *StartGroupRequest) XXX_Size() int {
	return m.Size()
}
func (m *StartGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartGroupRequest proto.InternalMessageInfo

func (m *StartGroupRequest) GetUserContext() *prototypes.UserContext {
	if m != nil {
		return m.UserContext
	}
	return nil
}

func (m *StartGroupRequest) GetGroupID() int32 {
	if m != nil {
		return m.GroupID
	}
	return 0
}

type GroupStartedResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupStartedResponse) Reset()         { *m = GroupStartedResponse{} }
func (m *GroupStartedResponse) String() string { return proto.CompactTextString(m) }
func (*GroupStartedResponse) ProtoMessage()    {}
func (*GroupStartedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_coordinatorservicer_c7a0a16d6559a9fe, []int{1}
}
func (m *GroupStartedResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupStartedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupStartedResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GroupStartedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupStartedResponse.Merge(dst, src)
}
func (m *GroupStartedResponse) XXX_Size() int {
	return m.Size()
}
func (m *GroupStartedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupStartedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GroupStartedResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StartGroupRequest)(nil), "coordinator.StartGroupRequest")
	proto.RegisterType((*GroupStartedResponse)(nil), "coordinator.GroupStartedResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CoordinatorClient is the client API for Coordinator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoordinatorClient interface {
	StartGroup(ctx context.Context, in *StartGroupRequest, opts ...grpc.CallOption) (*GroupStartedResponse, error)
}

type coordinatorClient struct {
	cc *grpc.ClientConn
}

func NewCoordinatorClient(cc *grpc.ClientConn) CoordinatorClient {
	return &coordinatorClient{cc}
}

func (c *coordinatorClient) StartGroup(ctx context.Context, in *StartGroupRequest, opts ...grpc.CallOption) (*GroupStartedResponse, error) {
	out := new(GroupStartedResponse)
	err := c.cc.Invoke(ctx, "/coordinator.Coordinator/StartGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoordinatorServer is the server API for Coordinator service.
type CoordinatorServer interface {
	StartGroup(context.Context, *StartGroupRequest) (*GroupStartedResponse, error)
}

func RegisterCoordinatorServer(s *grpc.Server, srv CoordinatorServer) {
	s.RegisterService(&_Coordinator_serviceDesc, srv)
}

func _Coordinator_StartGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinatorServer).StartGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coordinator.Coordinator/StartGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinatorServer).StartGroup(ctx, req.(*StartGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Coordinator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "coordinator.Coordinator",
	HandlerType: (*CoordinatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartGroup",
			Handler:    _Coordinator_StartGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coordinator/coordinatorservicer.proto",
}

func (m *StartGroupRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StartGroupRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UserContext != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCoordinatorservicer(dAtA, i, uint64(m.UserContext.Size()))
		n1, err := m.UserContext.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.GroupID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintCoordinatorservicer(dAtA, i, uint64(m.GroupID))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *GroupStartedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupStartedResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintCoordinatorservicer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *StartGroupRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UserContext != nil {
		l = m.UserContext.Size()
		n += 1 + l + sovCoordinatorservicer(uint64(l))
	}
	if m.GroupID != 0 {
		n += 1 + sovCoordinatorservicer(uint64(m.GroupID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GroupStartedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCoordinatorservicer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCoordinatorservicer(x uint64) (n int) {
	return sovCoordinatorservicer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StartGroupRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoordinatorservicer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StartGroupRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StartGroupRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserContext", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinatorservicer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCoordinatorservicer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UserContext == nil {
				m.UserContext = &prototypes.UserContext{}
			}
			if err := m.UserContext.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupID", wireType)
			}
			m.GroupID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinatorservicer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCoordinatorservicer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCoordinatorservicer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GroupStartedResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoordinatorservicer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GroupStartedResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupStartedResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipCoordinatorservicer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCoordinatorservicer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCoordinatorservicer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCoordinatorservicer
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCoordinatorservicer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCoordinatorservicer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCoordinatorservicer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCoordinatorservicer
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCoordinatorservicer(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCoordinatorservicer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCoordinatorservicer   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("coordinator/coordinatorservicer.proto", fileDescriptor_coordinatorservicer_c7a0a16d6559a9fe)
}

var fileDescriptor_coordinatorservicer_c7a0a16d6559a9fe = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0xce, 0xcf, 0x2f,
	0x4a, 0xc9, 0xcc, 0x4b, 0x2c, 0xc9, 0x2f, 0xd2, 0x47, 0x62, 0x17, 0xa7, 0x16, 0x95, 0x65, 0x26,
	0xa7, 0x16, 0xe9, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x71, 0x23, 0x49, 0x49, 0x89, 0x82, 0xc5,
	0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0xf5, 0x4b, 0x8b, 0x61, 0x6a, 0x94, 0x62, 0xb9, 0x04, 0x83, 0x4b,
	0x12, 0x8b, 0x4a, 0xdc, 0x8b, 0xf2, 0x4b, 0x0b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84,
	0xf4, 0xb8, 0xb8, 0x43, 0x8b, 0x53, 0x8b, 0x9c, 0xf3, 0xf3, 0x4a, 0x52, 0x2b, 0x4a, 0x24, 0x18,
	0x15, 0x18, 0x35, 0xb8, 0x8d, 0x78, 0xf4, 0x90, 0xc4, 0x82, 0x90, 0x15, 0x08, 0x49, 0x70, 0xb1,
	0x83, 0xf5, 0x7b, 0xba, 0x48, 0x30, 0x29, 0x30, 0x6a, 0xb0, 0x06, 0xc1, 0xb8, 0x4a, 0x62, 0x5c,
	0x22, 0x60, 0x26, 0xd8, 0x8e, 0xd4, 0x94, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xa3,
	0x38, 0x2e, 0x6e, 0x67, 0x84, 0xe3, 0x84, 0xfc, 0xb9, 0xb8, 0x10, 0xae, 0x10, 0x92, 0xd3, 0x43,
	0x72, 0xb8, 0x1e, 0x86, 0xf3, 0xa4, 0x14, 0x51, 0xe4, 0xb1, 0x99, 0xef, 0xe4, 0x72, 0xe2, 0x91,
	0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xce, 0x78, 0x2c, 0xc7, 0x10, 0x65,
	0x94, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x93, 0x99, 0x97, 0x9d,
	0x98, 0xa9, 0x9b, 0x99, 0xaf, 0x9f, 0x98, 0xab, 0x0f, 0x0e, 0x87, 0x64, 0x68, 0xd0, 0x15, 0x23,
	0x07, 0x67, 0x12, 0x1b, 0x58, 0xce, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x79, 0x06, 0x55, 0x8f,
	0x70, 0x01, 0x00, 0x00,
}