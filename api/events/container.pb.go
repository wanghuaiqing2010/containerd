// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/containerd/containerd/api/events/container.proto

/*
	Package events is a generated protocol buffer package.

	It is generated from these files:
		github.com/containerd/containerd/api/events/container.proto
		github.com/containerd/containerd/api/events/content.proto
		github.com/containerd/containerd/api/events/image.proto
		github.com/containerd/containerd/api/events/namespace.proto
		github.com/containerd/containerd/api/events/snapshot.proto
		github.com/containerd/containerd/api/events/task.proto

	It has these top-level messages:
		ContainerCreate
		ContainerUpdate
		ContainerDelete
		ContentDelete
		ImageCreate
		ImageUpdate
		ImageDelete
		NamespaceCreate
		NamespaceUpdate
		NamespaceDelete
		SnapshotPrepare
		SnapshotCommit
		SnapshotRemove
		TaskCreate
		TaskStart
		TaskDelete
		TaskIO
		TaskExit
		TaskOOM
		TaskExecAdded
		TaskExecStarted
		TaskPaused
		TaskResumed
		TaskCheckpointed
*/
package events

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import google_protobuf1 "github.com/gogo/protobuf/types"
import _ "github.com/containerd/containerd/protobuf/plugin"

import github_com_containerd_typeurl "github.com/containerd/typeurl"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ContainerCreate struct {
	ID      string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Image   string                   `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Runtime *ContainerCreate_Runtime `protobuf:"bytes,3,opt,name=runtime" json:"runtime,omitempty"`
}

func (m *ContainerCreate) Reset()                    { *m = ContainerCreate{} }
func (*ContainerCreate) ProtoMessage()               {}
func (*ContainerCreate) Descriptor() ([]byte, []int) { return fileDescriptorContainer, []int{0} }

type ContainerCreate_Runtime struct {
	Name    string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Options *google_protobuf1.Any `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
}

func (m *ContainerCreate_Runtime) Reset()      { *m = ContainerCreate_Runtime{} }
func (*ContainerCreate_Runtime) ProtoMessage() {}
func (*ContainerCreate_Runtime) Descriptor() ([]byte, []int) {
	return fileDescriptorContainer, []int{0, 0}
}

type ContainerUpdate struct {
	ID          string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Image       string            `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Labels      map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SnapshotKey string            `protobuf:"bytes,4,opt,name=snapshot_key,json=snapshotKey,proto3" json:"snapshot_key,omitempty"`
}

func (m *ContainerUpdate) Reset()                    { *m = ContainerUpdate{} }
func (*ContainerUpdate) ProtoMessage()               {}
func (*ContainerUpdate) Descriptor() ([]byte, []int) { return fileDescriptorContainer, []int{1} }

type ContainerDelete struct {
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *ContainerDelete) Reset()                    { *m = ContainerDelete{} }
func (*ContainerDelete) ProtoMessage()               {}
func (*ContainerDelete) Descriptor() ([]byte, []int) { return fileDescriptorContainer, []int{2} }

func init() {
	proto.RegisterType((*ContainerCreate)(nil), "containerd.events.ContainerCreate")
	proto.RegisterType((*ContainerCreate_Runtime)(nil), "containerd.events.ContainerCreate.Runtime")
	proto.RegisterType((*ContainerUpdate)(nil), "containerd.events.ContainerUpdate")
	proto.RegisterType((*ContainerDelete)(nil), "containerd.events.ContainerDelete")
}

// Field returns the value for the given fieldpath as a string, if defined.
// If the value is not defined, the second value will be false.
func (m *ContainerCreate) Field(fieldpath []string) (string, bool) {
	if len(fieldpath) == 0 {
		return "", false
	}

	switch fieldpath[0] {
	case "id":
		return string(m.ID), len(m.ID) > 0
	case "image":
		return string(m.Image), len(m.Image) > 0
	case "runtime":
		// NOTE(stevvooe): This is probably not correct in many cases.
		// We assume that the target message also implements the Field
		// method, which isn't likely true in a lot of cases.
		//
		// If you have a broken build and have found this comment,
		// you may be closer to a solution.
		if m.Runtime == nil {
			return "", false
		}

		return m.Runtime.Field(fieldpath[1:])
	}
	return "", false
}

// Field returns the value for the given fieldpath as a string, if defined.
// If the value is not defined, the second value will be false.
func (m *ContainerCreate_Runtime) Field(fieldpath []string) (string, bool) {
	if len(fieldpath) == 0 {
		return "", false
	}

	switch fieldpath[0] {
	case "name":
		return string(m.Name), len(m.Name) > 0
	case "options":
		decoded, err := github_com_containerd_typeurl.UnmarshalAny(m.Options)
		if err != nil {
			return "", false
		}

		adaptor, ok := decoded.(interface {
			Field([]string) (string, bool)
		})
		if !ok {
			return "", false
		}
		return adaptor.Field(fieldpath[1:])
	}
	return "", false
}

// Field returns the value for the given fieldpath as a string, if defined.
// If the value is not defined, the second value will be false.
func (m *ContainerUpdate) Field(fieldpath []string) (string, bool) {
	if len(fieldpath) == 0 {
		return "", false
	}

	switch fieldpath[0] {
	case "id":
		return string(m.ID), len(m.ID) > 0
	case "image":
		return string(m.Image), len(m.Image) > 0
	case "labels":
		// Labels fields have been special-cased by name. If this breaks,
		// add better special casing to fieldpath plugin.
		if len(m.Labels) == 0 {
			return "", false
		}
		value, ok := m.Labels[strings.Join(fieldpath[1:], ".")]
		return value, ok
	case "snapshot_key":
		return string(m.SnapshotKey), len(m.SnapshotKey) > 0
	}
	return "", false
}

// Field returns the value for the given fieldpath as a string, if defined.
// If the value is not defined, the second value will be false.
func (m *ContainerDelete) Field(fieldpath []string) (string, bool) {
	if len(fieldpath) == 0 {
		return "", false
	}

	switch fieldpath[0] {
	case "id":
		return string(m.ID), len(m.ID) > 0
	}
	return "", false
}
func (m *ContainerCreate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContainerCreate) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if len(m.Image) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.Image)))
		i += copy(dAtA[i:], m.Image)
	}
	if m.Runtime != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintContainer(dAtA, i, uint64(m.Runtime.Size()))
		n1, err := m.Runtime.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *ContainerCreate_Runtime) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContainerCreate_Runtime) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Options != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintContainer(dAtA, i, uint64(m.Options.Size()))
		n2, err := m.Options.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *ContainerUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContainerUpdate) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if len(m.Image) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.Image)))
		i += copy(dAtA[i:], m.Image)
	}
	if len(m.Labels) > 0 {
		for k, _ := range m.Labels {
			dAtA[i] = 0x1a
			i++
			v := m.Labels[k]
			mapSize := 1 + len(k) + sovContainer(uint64(len(k))) + 1 + len(v) + sovContainer(uint64(len(v)))
			i = encodeVarintContainer(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintContainer(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintContainer(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if len(m.SnapshotKey) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.SnapshotKey)))
		i += copy(dAtA[i:], m.SnapshotKey)
	}
	return i, nil
}

func (m *ContainerDelete) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContainerDelete) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	return i, nil
}

func encodeVarintContainer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ContainerCreate) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	l = len(m.Image)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	if m.Runtime != nil {
		l = m.Runtime.Size()
		n += 1 + l + sovContainer(uint64(l))
	}
	return n
}

func (m *ContainerCreate_Runtime) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	if m.Options != nil {
		l = m.Options.Size()
		n += 1 + l + sovContainer(uint64(l))
	}
	return n
}

func (m *ContainerUpdate) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	l = len(m.Image)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovContainer(uint64(len(k))) + 1 + len(v) + sovContainer(uint64(len(v)))
			n += mapEntrySize + 1 + sovContainer(uint64(mapEntrySize))
		}
	}
	l = len(m.SnapshotKey)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	return n
}

func (m *ContainerDelete) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	return n
}

func sovContainer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozContainer(x uint64) (n int) {
	return sovContainer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ContainerCreate) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ContainerCreate{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Image:` + fmt.Sprintf("%v", this.Image) + `,`,
		`Runtime:` + strings.Replace(fmt.Sprintf("%v", this.Runtime), "ContainerCreate_Runtime", "ContainerCreate_Runtime", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ContainerCreate_Runtime) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ContainerCreate_Runtime{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Options:` + strings.Replace(fmt.Sprintf("%v", this.Options), "Any", "google_protobuf1.Any", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ContainerUpdate) String() string {
	if this == nil {
		return "nil"
	}
	keysForLabels := make([]string, 0, len(this.Labels))
	for k, _ := range this.Labels {
		keysForLabels = append(keysForLabels, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForLabels)
	mapStringForLabels := "map[string]string{"
	for _, k := range keysForLabels {
		mapStringForLabels += fmt.Sprintf("%v: %v,", k, this.Labels[k])
	}
	mapStringForLabels += "}"
	s := strings.Join([]string{`&ContainerUpdate{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Image:` + fmt.Sprintf("%v", this.Image) + `,`,
		`Labels:` + mapStringForLabels + `,`,
		`SnapshotKey:` + fmt.Sprintf("%v", this.SnapshotKey) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ContainerDelete) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ContainerDelete{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringContainer(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ContainerCreate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContainer
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
			return fmt.Errorf("proto: ContainerCreate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContainerCreate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Image = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Runtime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Runtime == nil {
				m.Runtime = &ContainerCreate_Runtime{}
			}
			if err := m.Runtime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContainer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContainer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ContainerCreate_Runtime) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContainer
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
			return fmt.Errorf("proto: Runtime: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Runtime: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Options == nil {
				m.Options = &google_protobuf1.Any{}
			}
			if err := m.Options.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContainer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContainer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ContainerUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContainer
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
			return fmt.Errorf("proto: ContainerUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContainerUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Image = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowContainer
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowContainer
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthContainer
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowContainer
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthContainer
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipContainer(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthContainer
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Labels[mapkey] = mapvalue
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnapshotKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SnapshotKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContainer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContainer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ContainerDelete) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContainer
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
			return fmt.Errorf("proto: ContainerDelete: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContainerDelete: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContainer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContainer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipContainer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContainer
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
					return 0, ErrIntOverflowContainer
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
					return 0, ErrIntOverflowContainer
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
				return 0, ErrInvalidLengthContainer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowContainer
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
				next, err := skipContainer(dAtA[start:])
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
	ErrInvalidLengthContainer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContainer   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/containerd/containerd/api/events/container.proto", fileDescriptorContainer)
}

var fileDescriptorContainer = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0x3b, 0x49, 0x6d, 0x71, 0x22, 0xa8, 0x43, 0x91, 0x98, 0x45, 0xac, 0x5d, 0x55, 0x17,
	0x13, 0x88, 0x1b, 0xb5, 0x1b, 0x6d, 0xab, 0x20, 0x2a, 0x48, 0xc0, 0x8d, 0x1b, 0x99, 0x34, 0xd3,
	0x74, 0x30, 0x99, 0x09, 0xc9, 0xa4, 0x90, 0x9d, 0x8f, 0xe2, 0xe3, 0x74, 0xe9, 0xc2, 0x85, 0x2b,
	0xb1, 0x01, 0xdf, 0xc0, 0x07, 0x90, 0xcc, 0x34, 0x6d, 0xf0, 0x72, 0xef, 0xa5, 0xab, 0x9c, 0x3f,
	0xdf, 0x97, 0x73, 0x7e, 0x87, 0x81, 0xb3, 0x98, 0xc9, 0x4d, 0x19, 0xe2, 0x95, 0x48, 0xbd, 0x95,
	0xe0, 0x92, 0x30, 0x4e, 0xf3, 0xa8, 0x1b, 0x92, 0x8c, 0x79, 0x74, 0x4b, 0xb9, 0x2c, 0x4e, 0x55,
	0x9c, 0xe5, 0x42, 0x0a, 0x74, 0xf7, 0x24, 0xc3, 0x5a, 0xe2, 0x8c, 0x62, 0x11, 0x0b, 0xd5, 0xf5,
	0x9a, 0x48, 0x0b, 0x9d, 0xfb, 0xb1, 0x10, 0x71, 0x42, 0x3d, 0x95, 0x85, 0xe5, 0xda, 0x23, 0xbc,
	0x3a, 0xb4, 0x5e, 0x5c, 0xbb, 0xc0, 0xd1, 0x94, 0x25, 0x65, 0xcc, 0xb8, 0xb7, 0x66, 0x34, 0x89,
	0x32, 0x22, 0x37, 0xfa, 0x0f, 0x93, 0x1f, 0x00, 0xde, 0x5e, 0xb4, 0xf2, 0x45, 0x4e, 0x89, 0xa4,
	0xe8, 0x1e, 0x34, 0x58, 0x64, 0x83, 0x31, 0x98, 0xde, 0x9c, 0x0f, 0xea, 0x5f, 0x0f, 0x8c, 0x37,
	0xcb, 0xc0, 0x60, 0x11, 0x1a, 0xc1, 0x1b, 0x2c, 0x25, 0x31, 0xb5, 0x8d, 0xa6, 0x15, 0xe8, 0x04,
	0x2d, 0xe1, 0x30, 0x2f, 0xb9, 0x64, 0x29, 0xb5, 0xcd, 0x31, 0x98, 0x5a, 0xfe, 0x63, 0x7c, 0x81,
	0x0c, 0xff, 0x37, 0x02, 0x07, 0xda, 0x11, 0xb4, 0x56, 0xe7, 0x3d, 0x1c, 0x1e, 0x6a, 0x08, 0xc1,
	0x3e, 0x27, 0x29, 0xd5, 0x0b, 0x04, 0x2a, 0x46, 0x18, 0x0e, 0x45, 0x26, 0x99, 0xe0, 0x85, 0x1a,
	0x6e, 0xf9, 0x23, 0xac, 0xaf, 0x82, 0x5b, 0x40, 0xfc, 0x92, 0x57, 0x41, 0x2b, 0x9a, 0xfc, 0xe9,
	0x62, 0x7d, 0xcc, 0xa2, 0xf3, 0xb1, 0x5e, 0xc3, 0x41, 0x42, 0x42, 0x9a, 0x14, 0xb6, 0x39, 0x36,
	0xa7, 0x96, 0x8f, 0xaf, 0xa2, 0xd2, 0x13, 0xf0, 0x3b, 0x65, 0x78, 0xc5, 0x65, 0x5e, 0x05, 0x07,
	0x37, 0x7a, 0x08, 0x6f, 0x15, 0x9c, 0x64, 0xc5, 0x46, 0xc8, 0xcf, 0x5f, 0x68, 0x65, 0xf7, 0xd5,
	0x10, 0xab, 0xad, 0xbd, 0xa5, 0x95, 0xf3, 0x0c, 0x5a, 0x1d, 0x27, 0xba, 0x03, 0xcd, 0x46, 0xa8,
	0xf1, 0x9b, 0xb0, 0xd9, 0x70, 0x4b, 0x92, 0xf2, 0xb8, 0xa1, 0x4a, 0x9e, 0x1b, 0x4f, 0xc1, 0xe4,
	0x51, 0x07, 0x73, 0x49, 0x13, 0x7a, 0x39, 0xe6, 0xfc, 0xc3, 0x6e, 0xef, 0xf6, 0x7e, 0xee, 0xdd,
	0xde, 0xd7, 0xda, 0x05, 0xbb, 0xda, 0x05, 0xdf, 0x6b, 0x17, 0xfc, 0xae, 0x5d, 0xf0, 0xed, 0xaf,
	0x0b, 0x3e, 0xf9, 0x67, 0x3c, 0xe5, 0x99, 0xfe, 0x84, 0x03, 0x75, 0xfb, 0x27, 0xff, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xc4, 0x03, 0x7d, 0xc8, 0x07, 0x03, 0x00, 0x00,
}