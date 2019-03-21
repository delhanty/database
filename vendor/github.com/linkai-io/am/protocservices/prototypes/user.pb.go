// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: prototypes/user.proto

package prototypes // import "github.com/linkai-io/am/protocservices/prototypes"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type UserContext struct {
	TraceID              string   `protobuf:"bytes,1,opt,name=TraceID,proto3" json:"TraceID,omitempty"`
	OrgID                int32    `protobuf:"varint,2,opt,name=OrgID,proto3" json:"OrgID,omitempty"`
	UserID               int32    `protobuf:"varint,3,opt,name=UserID,proto3" json:"UserID,omitempty"`
	IPAddress            string   `protobuf:"bytes,4,opt,name=IPAddress,proto3" json:"IPAddress,omitempty"`
	Roles                []string `protobuf:"bytes,5,rep,name=Roles" json:"Roles,omitempty"`
	OrgCID               string   `protobuf:"bytes,6,opt,name=OrgCID,proto3" json:"OrgCID,omitempty"`
	UserCID              string   `protobuf:"bytes,7,opt,name=UserCID,proto3" json:"UserCID,omitempty"`
	SubscriptionID       int32    `protobuf:"varint,8,opt,name=SubscriptionID,proto3" json:"SubscriptionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserContext) Reset()         { *m = UserContext{} }
func (m *UserContext) String() string { return proto.CompactTextString(m) }
func (*UserContext) ProtoMessage()    {}
func (*UserContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_058fb3eca9f54508, []int{0}
}
func (m *UserContext) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserContext.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UserContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserContext.Merge(dst, src)
}
func (m *UserContext) XXX_Size() int {
	return m.Size()
}
func (m *UserContext) XXX_DiscardUnknown() {
	xxx_messageInfo_UserContext.DiscardUnknown(m)
}

var xxx_messageInfo_UserContext proto.InternalMessageInfo

func (m *UserContext) GetTraceID() string {
	if m != nil {
		return m.TraceID
	}
	return ""
}

func (m *UserContext) GetOrgID() int32 {
	if m != nil {
		return m.OrgID
	}
	return 0
}

func (m *UserContext) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *UserContext) GetIPAddress() string {
	if m != nil {
		return m.IPAddress
	}
	return ""
}

func (m *UserContext) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *UserContext) GetOrgCID() string {
	if m != nil {
		return m.OrgCID
	}
	return ""
}

func (m *UserContext) GetUserCID() string {
	if m != nil {
		return m.UserCID
	}
	return ""
}

func (m *UserContext) GetSubscriptionID() int32 {
	if m != nil {
		return m.SubscriptionID
	}
	return 0
}

type User struct {
	OrgID                int32    `protobuf:"varint,1,opt,name=OrgID,proto3" json:"OrgID,omitempty"`
	OrgCID               string   `protobuf:"bytes,2,opt,name=OrgCID,proto3" json:"OrgCID,omitempty"`
	UserCID              string   `protobuf:"bytes,3,opt,name=UserCID,proto3" json:"UserCID,omitempty"`
	UserID               int32    `protobuf:"varint,4,opt,name=UserID,proto3" json:"UserID,omitempty"`
	UserEmail            string   `protobuf:"bytes,5,opt,name=UserEmail,proto3" json:"UserEmail,omitempty"`
	FirstName            string   `protobuf:"bytes,6,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,7,opt,name=LastName,proto3" json:"LastName,omitempty"`
	StatusID             int32    `protobuf:"varint,8,opt,name=StatusID,proto3" json:"StatusID,omitempty"`
	CreationTime         int64    `protobuf:"varint,9,opt,name=CreationTime,proto3" json:"CreationTime,omitempty"`
	Deleted              bool     `protobuf:"varint,10,opt,name=Deleted,proto3" json:"Deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_058fb3eca9f54508, []int{1}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_User.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return m.Size()
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetOrgID() int32 {
	if m != nil {
		return m.OrgID
	}
	return 0
}

func (m *User) GetOrgCID() string {
	if m != nil {
		return m.OrgCID
	}
	return ""
}

func (m *User) GetUserCID() string {
	if m != nil {
		return m.UserCID
	}
	return ""
}

func (m *User) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *User) GetUserEmail() string {
	if m != nil {
		return m.UserEmail
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetStatusID() int32 {
	if m != nil {
		return m.StatusID
	}
	return 0
}

func (m *User) GetCreationTime() int64 {
	if m != nil {
		return m.CreationTime
	}
	return 0
}

func (m *User) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

type UserFilter struct {
	Start                int32       `protobuf:"varint,1,opt,name=Start,proto3" json:"Start,omitempty"`
	Limit                int32       `protobuf:"varint,2,opt,name=Limit,proto3" json:"Limit,omitempty"`
	OrgID                int32       `protobuf:"varint,3,opt,name=OrgID,proto3" json:"OrgID,omitempty"`
	Filters              *FilterType `protobuf:"bytes,4,opt,name=Filters" json:"Filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UserFilter) Reset()         { *m = UserFilter{} }
func (m *UserFilter) String() string { return proto.CompactTextString(m) }
func (*UserFilter) ProtoMessage()    {}
func (*UserFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_058fb3eca9f54508, []int{2}
}
func (m *UserFilter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserFilter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *UserFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserFilter.Merge(dst, src)
}
func (m *UserFilter) XXX_Size() int {
	return m.Size()
}
func (m *UserFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_UserFilter.DiscardUnknown(m)
}

var xxx_messageInfo_UserFilter proto.InternalMessageInfo

func (m *UserFilter) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *UserFilter) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *UserFilter) GetOrgID() int32 {
	if m != nil {
		return m.OrgID
	}
	return 0
}

func (m *UserFilter) GetFilters() *FilterType {
	if m != nil {
		return m.Filters
	}
	return nil
}

func init() {
	proto.RegisterType((*UserContext)(nil), "UserContext")
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterType((*UserFilter)(nil), "UserFilter")
}
func (m *UserContext) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserContext) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.TraceID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.TraceID)))
		i += copy(dAtA[i:], m.TraceID)
	}
	if m.OrgID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.OrgID))
	}
	if m.UserID != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.UserID))
	}
	if len(m.IPAddress) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.IPAddress)))
		i += copy(dAtA[i:], m.IPAddress)
	}
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			dAtA[i] = 0x2a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.OrgCID) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.OrgCID)))
		i += copy(dAtA[i:], m.OrgCID)
	}
	if len(m.UserCID) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.UserCID)))
		i += copy(dAtA[i:], m.UserCID)
	}
	if m.SubscriptionID != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.SubscriptionID))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.OrgID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.OrgID))
	}
	if len(m.OrgCID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.OrgCID)))
		i += copy(dAtA[i:], m.OrgCID)
	}
	if len(m.UserCID) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.UserCID)))
		i += copy(dAtA[i:], m.UserCID)
	}
	if m.UserID != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.UserID))
	}
	if len(m.UserEmail) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.UserEmail)))
		i += copy(dAtA[i:], m.UserEmail)
	}
	if len(m.FirstName) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.FirstName)))
		i += copy(dAtA[i:], m.FirstName)
	}
	if len(m.LastName) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.LastName)))
		i += copy(dAtA[i:], m.LastName)
	}
	if m.StatusID != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.StatusID))
	}
	if m.CreationTime != 0 {
		dAtA[i] = 0x48
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.CreationTime))
	}
	if m.Deleted {
		dAtA[i] = 0x50
		i++
		if m.Deleted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *UserFilter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserFilter) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Start != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.Start))
	}
	if m.Limit != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.Limit))
	}
	if m.OrgID != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.OrgID))
	}
	if m.Filters != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.Filters.Size()))
		n1, err := m.Filters.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintUser(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *UserContext) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TraceID)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.OrgID != 0 {
		n += 1 + sovUser(uint64(m.OrgID))
	}
	if m.UserID != 0 {
		n += 1 + sovUser(uint64(m.UserID))
	}
	l = len(m.IPAddress)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if len(m.Roles) > 0 {
		for _, s := range m.Roles {
			l = len(s)
			n += 1 + l + sovUser(uint64(l))
		}
	}
	l = len(m.OrgCID)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.UserCID)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.SubscriptionID != 0 {
		n += 1 + sovUser(uint64(m.SubscriptionID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *User) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OrgID != 0 {
		n += 1 + sovUser(uint64(m.OrgID))
	}
	l = len(m.OrgCID)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.UserCID)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.UserID != 0 {
		n += 1 + sovUser(uint64(m.UserID))
	}
	l = len(m.UserEmail)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.FirstName)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.LastName)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.StatusID != 0 {
		n += 1 + sovUser(uint64(m.StatusID))
	}
	if m.CreationTime != 0 {
		n += 1 + sovUser(uint64(m.CreationTime))
	}
	if m.Deleted {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UserFilter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Start != 0 {
		n += 1 + sovUser(uint64(m.Start))
	}
	if m.Limit != 0 {
		n += 1 + sovUser(uint64(m.Limit))
	}
	if m.OrgID != 0 {
		n += 1 + sovUser(uint64(m.OrgID))
	}
	if m.Filters != nil {
		l = m.Filters.Size()
		n += 1 + l + sovUser(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovUser(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozUser(x uint64) (n int) {
	return sovUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserContext) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: UserContext: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserContext: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TraceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TraceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrgID", wireType)
			}
			m.OrgID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrgID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			m.UserID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IPAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IPAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Roles = append(m.Roles, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrgCID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrgCID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserCID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserCID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubscriptionID", wireType)
			}
			m.SubscriptionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubscriptionID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
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
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrgID", wireType)
			}
			m.OrgID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrgID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrgCID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrgCID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserCID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserCID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			m.UserID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserEmail", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserEmail = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FirstName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusID", wireType)
			}
			m.StatusID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StatusID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationTime", wireType)
			}
			m.CreationTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreationTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deleted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Deleted = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
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
func (m *UserFilter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: UserFilter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserFilter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			m.Start = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Start |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrgID", wireType)
			}
			m.OrgID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrgID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Filters == nil {
				m.Filters = &FilterType{}
			}
			if err := m.Filters.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
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
func skipUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
				return 0, ErrInvalidLengthUser
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowUser
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
				next, err := skipUser(dAtA[start:])
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
	ErrInvalidLengthUser = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUser   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("prototypes/user.proto", fileDescriptor_user_058fb3eca9f54508) }

var fileDescriptor_user_058fb3eca9f54508 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcd, 0xaa, 0xd3, 0x40,
	0x14, 0xc7, 0x9d, 0xa6, 0x69, 0x9b, 0xa9, 0xb8, 0x08, 0x2a, 0xc3, 0x55, 0x42, 0x08, 0x28, 0xd9,
	0xd8, 0xa0, 0x3e, 0x81, 0x26, 0x5e, 0x08, 0x5c, 0xbc, 0x92, 0xd6, 0x8d, 0xbb, 0x69, 0x7a, 0xac,
	0x83, 0x49, 0x13, 0x66, 0x26, 0x62, 0xdf, 0xa4, 0x8f, 0xe4, 0xd2, 0x47, 0x90, 0xf6, 0x45, 0x64,
	0x66, 0xf2, 0xd1, 0x16, 0xee, 0x6e, 0x7e, 0xff, 0xc3, 0xf9, 0xf8, 0x9f, 0x39, 0xf8, 0x59, 0xcd,
	0x2b, 0x59, 0xc9, 0x7d, 0x0d, 0x22, 0x6a, 0x04, 0xf0, 0x85, 0xe6, 0x9b, 0x17, 0x67, 0xf2, 0x77,
	0x56, 0x48, 0xe0, 0xea, 0x6d, 0x82, 0xc1, 0x09, 0xe1, 0xf9, 0x57, 0x01, 0x3c, 0xae, 0x76, 0x12,
	0x7e, 0x4b, 0x97, 0xe0, 0xe9, 0x8a, 0xd3, 0x1c, 0xd2, 0x84, 0x20, 0x1f, 0x85, 0x4e, 0xd6, 0xa1,
	0xfb, 0x14, 0xdb, 0xf7, 0x7c, 0x9b, 0x26, 0x64, 0xe4, 0xa3, 0xd0, 0xce, 0x0c, 0xb8, 0xcf, 0xf1,
	0x44, 0xa5, 0xa7, 0x09, 0xb1, 0xb4, 0xdc, 0x92, 0xfb, 0x12, 0x3b, 0xe9, 0x97, 0x0f, 0x9b, 0x0d,
	0x07, 0x21, 0xc8, 0x58, 0x57, 0x1a, 0x04, 0x55, 0x2b, 0xab, 0x0a, 0x10, 0xc4, 0xf6, 0xad, 0xd0,
	0xc9, 0x0c, 0xa8, 0x5a, 0xf7, 0x7c, 0x1b, 0xa7, 0x09, 0x99, 0xe8, 0x84, 0x96, 0xd4, 0x4c, 0x7a,
	0xc4, 0x34, 0x21, 0x53, 0x33, 0x53, 0x8b, 0xee, 0x6b, 0xfc, 0x64, 0xd9, 0xac, 0x45, 0xce, 0x59,
	0x2d, 0x59, 0xb5, 0x4b, 0x13, 0x32, 0xd3, 0x53, 0x5c, 0xa9, 0xc1, 0x61, 0x84, 0xc7, 0x2a, 0x67,
	0x30, 0x81, 0xae, 0x4c, 0xb4, 0x8d, 0x47, 0x0f, 0x35, 0xb6, 0x2e, 0x1b, 0x0f, 0xb6, 0xc7, 0xd7,
	0xb6, 0xd5, 0xeb, 0x53, 0x49, 0x59, 0x41, 0x6c, 0x63, 0xbb, 0x17, 0x54, 0xf4, 0x96, 0x71, 0x21,
	0x3f, 0xd3, 0x12, 0x5a, 0x8f, 0x83, 0xe0, 0xde, 0xe0, 0xd9, 0x1d, 0x6d, 0x83, 0xc6, 0x67, 0xcf,
	0x2a, 0xb6, 0x94, 0x54, 0x36, 0xa2, 0xb7, 0xd8, 0xb3, 0x1b, 0xe0, 0xc7, 0x31, 0x07, 0xaa, 0xac,
	0xae, 0x58, 0x09, 0xc4, 0xf1, 0x51, 0x68, 0x65, 0x17, 0x9a, 0x72, 0x92, 0x40, 0x01, 0x12, 0x36,
	0x04, 0xfb, 0x28, 0x9c, 0x65, 0x1d, 0x06, 0x0d, 0xc6, 0x6a, 0xc0, 0x5b, 0x7d, 0x18, 0x6a, 0x3f,
	0x4b, 0x49, 0xb9, 0xec, 0xf6, 0xa3, 0x41, 0xa9, 0x77, 0xac, 0x64, 0xb2, 0xfb, 0x7a, 0x0d, 0xc3,
	0x2e, 0xad, 0xf3, 0x5d, 0xbe, 0xc2, 0x53, 0x53, 0xcb, 0x7c, 0xfb, 0xfc, 0xdd, 0x7c, 0x61, 0x78,
	0xb5, 0xaf, 0x21, 0xeb, 0x62, 0x1f, 0xe3, 0x3f, 0x47, 0x0f, 0xfd, 0x3d, 0x7a, 0xe8, 0xdf, 0xd1,
	0x43, 0x87, 0x93, 0xf7, 0xe8, 0xdb, 0xdb, 0x2d, 0x93, 0x3f, 0x9a, 0xf5, 0x22, 0xaf, 0xca, 0xa8,
	0x60, 0xbb, 0x9f, 0x94, 0xbd, 0x61, 0x55, 0x44, 0xcb, 0x48, 0x1f, 0x6a, 0x2e, 0x80, 0xff, 0x62,
	0x39, 0x88, 0x68, 0xb8, 0xe6, 0xf5, 0x44, 0xbf, 0xdf, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xd2,
	0x44, 0x74, 0x72, 0xf9, 0x02, 0x00, 0x00,
}
