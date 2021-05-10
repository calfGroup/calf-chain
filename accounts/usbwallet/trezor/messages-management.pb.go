

package trezor

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)


var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf


const _ = proto.ProtoPackageIsVersion3 


type ApplySettings_PassphraseSourceType int32

const (
	ApplySettings_ASK    ApplySettings_PassphraseSourceType = 0
	ApplySettings_DEVICE ApplySettings_PassphraseSourceType = 1
	ApplySettings_HOST   ApplySettings_PassphraseSourceType = 2
)

var ApplySettings_PassphraseSourceType_name = map[int32]string{
	0: "ASK",
	1: "DEVICE",
	2: "HOST",
}

var ApplySettings_PassphraseSourceType_value = map[string]int32{
	"ASK":    0,
	"DEVICE": 1,
	"HOST":   2,
}

func (x ApplySettings_PassphraseSourceType) Enum() *ApplySettings_PassphraseSourceType {
	p := new(ApplySettings_PassphraseSourceType)
	*p = x
	return p
}

func (x ApplySettings_PassphraseSourceType) String() string {
	return proto.EnumName(ApplySettings_PassphraseSourceType_name, int32(x))
}

func (x *ApplySettings_PassphraseSourceType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ApplySettings_PassphraseSourceType_value, data, "ApplySettings_PassphraseSourceType")
	if err != nil {
		return err
	}
	*x = ApplySettings_PassphraseSourceType(value)
	return nil
}

func (ApplySettings_PassphraseSourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{4, 0}
}


type RecoveryDevice_RecoveryDeviceType int32

const (

	RecoveryDevice_RecoveryDeviceType_ScrambledWords RecoveryDevice_RecoveryDeviceType = 0
	RecoveryDevice_RecoveryDeviceType_Matrix         RecoveryDevice_RecoveryDeviceType = 1
)

var RecoveryDevice_RecoveryDeviceType_name = map[int32]string{
	0: "RecoveryDeviceType_ScrambledWords",
	1: "RecoveryDeviceType_Matrix",
}

var RecoveryDevice_RecoveryDeviceType_value = map[string]int32{
	"RecoveryDeviceType_ScrambledWords": 0,
	"RecoveryDeviceType_Matrix":         1,
}

func (x RecoveryDevice_RecoveryDeviceType) Enum() *RecoveryDevice_RecoveryDeviceType {
	p := new(RecoveryDevice_RecoveryDeviceType)
	*p = x
	return p
}

func (x RecoveryDevice_RecoveryDeviceType) String() string {
	return proto.EnumName(RecoveryDevice_RecoveryDeviceType_name, int32(x))
}

func (x *RecoveryDevice_RecoveryDeviceType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RecoveryDevice_RecoveryDeviceType_value, data, "RecoveryDevice_RecoveryDeviceType")
	if err != nil {
		return err
	}
	*x = RecoveryDevice_RecoveryDeviceType(value)
	return nil
}

func (RecoveryDevice_RecoveryDeviceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{17, 0}
}


type WordRequest_WordRequestType int32

const (
	WordRequest_WordRequestType_Plain   WordRequest_WordRequestType = 0
	WordRequest_WordRequestType_Matrix9 WordRequest_WordRequestType = 1
	WordRequest_WordRequestType_Matrix6 WordRequest_WordRequestType = 2
)

var WordRequest_WordRequestType_name = map[int32]string{
	0: "WordRequestType_Plain",
	1: "WordRequestType_Matrix9",
	2: "WordRequestType_Matrix6",
}

var WordRequest_WordRequestType_value = map[string]int32{
	"WordRequestType_Plain":   0,
	"WordRequestType_Matrix9": 1,
	"WordRequestType_Matrix6": 2,
}

func (x WordRequest_WordRequestType) Enum() *WordRequest_WordRequestType {
	p := new(WordRequest_WordRequestType)
	*p = x
	return p
}

func (x WordRequest_WordRequestType) String() string {
	return proto.EnumName(WordRequest_WordRequestType_name, int32(x))
}

func (x *WordRequest_WordRequestType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WordRequest_WordRequestType_value, data, "WordRequest_WordRequestType")
	if err != nil {
		return err
	}
	*x = WordRequest_WordRequestType(value)
	return nil
}

func (WordRequest_WordRequestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{18, 0}
}


type Initialize struct {
	State                []byte   `protobuf:"bytes,1,opt,name=state" json:"state,omitempty"`
	SkipPassphrase       *bool    `protobuf:"varint,2,opt,name=skip_passphrase,json=skipPassphrase" json:"skip_passphrase,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Initialize) Reset()         { *m = Initialize{} }
func (m *Initialize) String() string { return proto.CompactTextString(m) }
func (*Initialize) ProtoMessage()    {}
func (*Initialize) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{0}
}

func (m *Initialize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Initialize.Unmarshal(m, b)
}
func (m *Initialize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Initialize.Marshal(b, m, deterministic)
}
func (m *Initialize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Initialize.Merge(m, src)
}
func (m *Initialize) XXX_Size() int {
	return xxx_messageInfo_Initialize.Size(m)
}
func (m *Initialize) XXX_DiscardUnknown() {
	xxx_messageInfo_Initialize.DiscardUnknown(m)
}

var xxx_messageInfo_Initialize proto.InternalMessageInfo

func (m *Initialize) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *Initialize) GetSkipPassphrase() bool {
	if m != nil && m.SkipPassphrase != nil {
		return *m.SkipPassphrase
	}
	return false
}

type GetFeatures struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeatures) Reset()         { *m = GetFeatures{} }
func (m *GetFeatures) String() string { return proto.CompactTextString(m) }
func (*GetFeatures) ProtoMessage()    {}
func (*GetFeatures) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{1}
}

func (m *GetFeatures) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeatures.Unmarshal(m, b)
}
func (m *GetFeatures) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeatures.Marshal(b, m, deterministic)
}
func (m *GetFeatures) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeatures.Merge(m, src)
}
func (m *GetFeatures) XXX_Size() int {
	return xxx_messageInfo_GetFeatures.Size(m)
}
func (m *GetFeatures) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeatures.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeatures proto.InternalMessageInfo


type Features struct {
	Vendor               *string  `protobuf:"bytes,1,opt,name=vendor" json:"vendor,omitempty"`
	MajorVersion         *uint32  `protobuf:"varint,2,opt,name=major_version,json=majorVersion" json:"major_version,omitempty"`
	MinorVersion         *uint32  `protobuf:"varint,3,opt,name=minor_version,json=minorVersion" json:"minor_version,omitempty"`
	PatchVersion         *uint32  `protobuf:"varint,4,opt,name=patch_version,json=patchVersion" json:"patch_version,omitempty"`
	BootloaderMode       *bool    `protobuf:"varint,5,opt,name=bootloader_mode,json=bootloaderMode" json:"bootloader_mode,omitempty"`
	DeviceId             *string  `protobuf:"bytes,6,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	PinProtection        *bool    `protobuf:"varint,7,opt,name=pin_protection,json=pinProtection" json:"pin_protection,omitempty"`
	PassphraseProtection *bool    `protobuf:"varint,8,opt,name=passphrase_protection,json=passphraseProtection" json:"passphrase_protection,omitempty"`
	Language             *string  `protobuf:"bytes,9,opt,name=language" json:"language,omitempty"`
	Label                *string  `protobuf:"bytes,10,opt,name=label" json:"label,omitempty"`
	Initialized          *bool    `protobuf:"varint,12,opt,name=initialized" json:"initialized,omitempty"`
	Revision             []byte   `protobuf:"bytes,13,opt,name=revision" json:"revision,omitempty"`
	BootloaderHash       []byte   `protobuf:"bytes,14,opt,name=bootloader_hash,json=bootloaderHash" json:"bootloader_hash,omitempty"`
	Imported             *bool    `protobuf:"varint,15,opt,name=imported" json:"imported,omitempty"`
	PinCached            *bool    `protobuf:"varint,16,opt,name=pin_cached,json=pinCached" json:"pin_cached,omitempty"`
	PassphraseCached     *bool    `protobuf:"varint,17,opt,name=passphrase_cached,json=passphraseCached" json:"passphrase_cached,omitempty"`
	FirmwarePresent      *bool    `protobuf:"varint,18,opt,name=firmware_present,json=firmwarePresent" json:"firmware_present,omitempty"`
	NeedsBackup          *bool    `protobuf:"varint,19,opt,name=needs_backup,json=needsBackup" json:"needs_backup,omitempty"`
	Flags                *uint32  `protobuf:"varint,20,opt,name=flags" json:"flags,omitempty"`
	Model                *string  `protobuf:"bytes,21,opt,name=model" json:"model,omitempty"`
	FwMajor              *uint32  `protobuf:"varint,22,opt,name=fw_major,json=fwMajor" json:"fw_major,omitempty"`
	FwMinor              *uint32  `protobuf:"varint,23,opt,name=fw_minor,json=fwMinor" json:"fw_minor,omitempty"`
	FwPatch              *uint32  `protobuf:"varint,24,opt,name=fw_patch,json=fwPatch" json:"fw_patch,omitempty"`
	FwVendor             *string  `protobuf:"bytes,25,opt,name=fw_vendor,json=fwVendor" json:"fw_vendor,omitempty"`
	FwVendorKeys         []byte   `protobuf:"bytes,26,opt,name=fw_vendor_keys,json=fwVendorKeys" json:"fw_vendor_keys,omitempty"`
	UnfinishedBackup     *bool    `protobuf:"varint,27,opt,name=unfinished_backup,json=unfinishedBackup" json:"unfinished_backup,omitempty"`
	NoBackup             *bool    `protobuf:"varint,28,opt,name=no_backup,json=noBackup" json:"no_backup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Features) Reset()         { *m = Features{} }
func (m *Features) String() string { return proto.CompactTextString(m) }
func (*Features) ProtoMessage()    {}
func (*Features) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{2}
}

func (m *Features) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Features.Unmarshal(m, b)
}
func (m *Features) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Features.Marshal(b, m, deterministic)
}
func (m *Features) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Features.Merge(m, src)
}
func (m *Features) XXX_Size() int {
	return xxx_messageInfo_Features.Size(m)
}
func (m *Features) XXX_DiscardUnknown() {
	xxx_messageInfo_Features.DiscardUnknown(m)
}

var xxx_messageInfo_Features proto.InternalMessageInfo

func (m *Features) GetVendor() string {
	if m != nil && m.Vendor != nil {
		return *m.Vendor
	}
	return ""
}

func (m *Features) GetMajorVersion() uint32 {
	if m != nil && m.MajorVersion != nil {
		return *m.MajorVersion
	}
	return 0
}

func (m *Features) GetMinorVersion() uint32 {
	if m != nil && m.MinorVersion != nil {
		return *m.MinorVersion
	}
	return 0
}

func (m *Features) GetPatchVersion() uint32 {
	if m != nil && m.PatchVersion != nil {
		return *m.PatchVersion
	}
	return 0
}

func (m *Features) GetBootloaderMode() bool {
	if m != nil && m.BootloaderMode != nil {
		return *m.BootloaderMode
	}
	return false
}

func (m *Features) GetDeviceId() string {
	if m != nil && m.DeviceId != nil {
		return *m.DeviceId
	}
	return ""
}

func (m *Features) GetPinProtection() bool {
	if m != nil && m.PinProtection != nil {
		return *m.PinProtection
	}
	return false
}

func (m *Features) GetPassphraseProtection() bool {
	if m != nil && m.PassphraseProtection != nil {
		return *m.PassphraseProtection
	}
	return false
}

func (m *Features) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

func (m *Features) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *Features) GetInitialized() bool {
	if m != nil && m.Initialized != nil {
		return *m.Initialized
	}
	return false
}

func (m *Features) GetRevision() []byte {
	if m != nil {
		return m.Revision
	}
	return nil
}

func (m *Features) GetBootloaderHash() []byte {
	if m != nil {
		return m.BootloaderHash
	}
	return nil
}

func (m *Features) GetImported() bool {
	if m != nil && m.Imported != nil {
		return *m.Imported
	}
	return false
}

func (m *Features) GetPinCached() bool {
	if m != nil && m.PinCached != nil {
		return *m.PinCached
	}
	return false
}

func (m *Features) GetPassphraseCached() bool {
	if m != nil && m.PassphraseCached != nil {
		return *m.PassphraseCached
	}
	return false
}

func (m *Features) GetFirmwarePresent() bool {
	if m != nil && m.FirmwarePresent != nil {
		return *m.FirmwarePresent
	}
	return false
}

func (m *Features) GetNeedsBackup() bool {
	if m != nil && m.NeedsBackup != nil {
		return *m.NeedsBackup
	}
	return false
}

func (m *Features) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func (m *Features) GetModel() string {
	if m != nil && m.Model != nil {
		return *m.Model
	}
	return ""
}

func (m *Features) GetFwMajor() uint32 {
	if m != nil && m.FwMajor != nil {
		return *m.FwMajor
	}
	return 0
}

func (m *Features) GetFwMinor() uint32 {
	if m != nil && m.FwMinor != nil {
		return *m.FwMinor
	}
	return 0
}

func (m *Features) GetFwPatch() uint32 {
	if m != nil && m.FwPatch != nil {
		return *m.FwPatch
	}
	return 0
}

func (m *Features) GetFwVendor() string {
	if m != nil && m.FwVendor != nil {
		return *m.FwVendor
	}
	return ""
}

func (m *Features) GetFwVendorKeys() []byte {
	if m != nil {
		return m.FwVendorKeys
	}
	return nil
}

func (m *Features) GetUnfinishedBackup() bool {
	if m != nil && m.UnfinishedBackup != nil {
		return *m.UnfinishedBackup
	}
	return false
}

func (m *Features) GetNoBackup() bool {
	if m != nil && m.NoBackup != nil {
		return *m.NoBackup
	}
	return false
}

//*
// Request: clear session (removes cached PIN, passphrase, etc).
// @start
// @next Success
type ClearSession struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClearSession) Reset()         { *m = ClearSession{} }
func (m *ClearSession) String() string { return proto.CompactTextString(m) }
func (*ClearSession) ProtoMessage()    {}
func (*ClearSession) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{3}
}

func (m *ClearSession) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClearSession.Unmarshal(m, b)
}
func (m *ClearSession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClearSession.Marshal(b, m, deterministic)
}
func (m *ClearSession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClearSession.Merge(m, src)
}
func (m *ClearSession) XXX_Size() int {
	return xxx_messageInfo_ClearSession.Size(m)
}
func (m *ClearSession) XXX_DiscardUnknown() {
	xxx_messageInfo_ClearSession.DiscardUnknown(m)
}

var xxx_messageInfo_ClearSession proto.InternalMessageInfo


type ApplySettings struct {
	Language             *string                             `protobuf:"bytes,1,opt,name=language" json:"language,omitempty"`
	Label                *string                             `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
	UsePassphrase        *bool                               `protobuf:"varint,3,opt,name=use_passphrase,json=usePassphrase" json:"use_passphrase,omitempty"`
	Homescreen           []byte                              `protobuf:"bytes,4,opt,name=homescreen" json:"homescreen,omitempty"`
	PassphraseSource     *ApplySettings_PassphraseSourceType `protobuf:"varint,5,opt,name=passphrase_source,json=passphraseSource,enum=hw.trezor.messages.management.ApplySettings_PassphraseSourceType" json:"passphrase_source,omitempty"`
	AutoLockDelayMs      *uint32                             `protobuf:"varint,6,opt,name=auto_lock_delay_ms,json=autoLockDelayMs" json:"auto_lock_delay_ms,omitempty"`
	DisplayRotation      *uint32                             `protobuf:"varint,7,opt,name=display_rotation,json=displayRotation" json:"display_rotation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *ApplySettings) Reset()         { *m = ApplySettings{} }
func (m *ApplySettings) String() string { return proto.CompactTextString(m) }
func (*ApplySettings) ProtoMessage()    {}
func (*ApplySettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{4}
}

func (m *ApplySettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplySettings.Unmarshal(m, b)
}
func (m *ApplySettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplySettings.Marshal(b, m, deterministic)
}
func (m *ApplySettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplySettings.Merge(m, src)
}
func (m *ApplySettings) XXX_Size() int {
	return xxx_messageInfo_ApplySettings.Size(m)
}
func (m *ApplySettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplySettings.DiscardUnknown(m)
}

var xxx_messageInfo_ApplySettings proto.InternalMessageInfo

func (m *ApplySettings) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

func (m *ApplySettings) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *ApplySettings) GetUsePassphrase() bool {
	if m != nil && m.UsePassphrase != nil {
		return *m.UsePassphrase
	}
	return false
}

func (m *ApplySettings) GetHomescreen() []byte {
	if m != nil {
		return m.Homescreen
	}
	return nil
}

func (m *ApplySettings) GetPassphraseSource() ApplySettings_PassphraseSourceType {
	if m != nil && m.PassphraseSource != nil {
		return *m.PassphraseSource
	}
	return ApplySettings_ASK
}

func (m *ApplySettings) GetAutoLockDelayMs() uint32 {
	if m != nil && m.AutoLockDelayMs != nil {
		return *m.AutoLockDelayMs
	}
	return 0
}

func (m *ApplySettings) GetDisplayRotation() uint32 {
	if m != nil && m.DisplayRotation != nil {
		return *m.DisplayRotation
	}
	return 0
}


type ApplyFlags struct {
	Flags                *uint32  `protobuf:"varint,1,opt,name=flags" json:"flags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyFlags) Reset()         { *m = ApplyFlags{} }
func (m *ApplyFlags) String() string { return proto.CompactTextString(m) }
func (*ApplyFlags) ProtoMessage()    {}
func (*ApplyFlags) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{5}
}

func (m *ApplyFlags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyFlags.Unmarshal(m, b)
}
func (m *ApplyFlags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyFlags.Marshal(b, m, deterministic)
}
func (m *ApplyFlags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyFlags.Merge(m, src)
}
func (m *ApplyFlags) XXX_Size() int {
	return xxx_messageInfo_ApplyFlags.Size(m)
}
func (m *ApplyFlags) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyFlags.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyFlags proto.InternalMessageInfo

func (m *ApplyFlags) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}


type ChangePin struct {
	Remove               *bool    `protobuf:"varint,1,opt,name=remove" json:"remove,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangePin) Reset()         { *m = ChangePin{} }
func (m *ChangePin) String() string { return proto.CompactTextString(m) }
func (*ChangePin) ProtoMessage()    {}
func (*ChangePin) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{6}
}

func (m *ChangePin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangePin.Unmarshal(m, b)
}
func (m *ChangePin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangePin.Marshal(b, m, deterministic)
}
func (m *ChangePin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangePin.Merge(m, src)
}
func (m *ChangePin) XXX_Size() int {
	return xxx_messageInfo_ChangePin.Size(m)
}
func (m *ChangePin) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangePin.DiscardUnknown(m)
}

var xxx_messageInfo_ChangePin proto.InternalMessageInfo

func (m *ChangePin) GetRemove() bool {
	if m != nil && m.Remove != nil {
		return *m.Remove
	}
	return false
}


type Ping struct {
	Message              *string  `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	ButtonProtection     *bool    `protobuf:"varint,2,opt,name=button_protection,json=buttonProtection" json:"button_protection,omitempty"`
	PinProtection        *bool    `protobuf:"varint,3,opt,name=pin_protection,json=pinProtection" json:"pin_protection,omitempty"`
	PassphraseProtection *bool    `protobuf:"varint,4,opt,name=passphrase_protection,json=passphraseProtection" json:"passphrase_protection,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{7}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *Ping) GetButtonProtection() bool {
	if m != nil && m.ButtonProtection != nil {
		return *m.ButtonProtection
	}
	return false
}

func (m *Ping) GetPinProtection() bool {
	if m != nil && m.PinProtection != nil {
		return *m.PinProtection
	}
	return false
}

func (m *Ping) GetPassphraseProtection() bool {
	if m != nil && m.PassphraseProtection != nil {
		return *m.PassphraseProtection
	}
	return false
}


type Cancel struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cancel) Reset()         { *m = Cancel{} }
func (m *Cancel) String() string { return proto.CompactTextString(m) }
func (*Cancel) ProtoMessage()    {}
func (*Cancel) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{8}
}

func (m *Cancel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cancel.Unmarshal(m, b)
}
func (m *Cancel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cancel.Marshal(b, m, deterministic)
}
func (m *Cancel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cancel.Merge(m, src)
}
func (m *Cancel) XXX_Size() int {
	return xxx_messageInfo_Cancel.Size(m)
}
func (m *Cancel) XXX_DiscardUnknown() {
	xxx_messageInfo_Cancel.DiscardUnknown(m)
}

var xxx_messageInfo_Cancel proto.InternalMessageInfo


type GetEntropy struct {
	Size                 *uint32  `protobuf:"varint,1,req,name=size" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEntropy) Reset()         { *m = GetEntropy{} }
func (m *GetEntropy) String() string { return proto.CompactTextString(m) }
func (*GetEntropy) ProtoMessage()    {}
func (*GetEntropy) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{9}
}

func (m *GetEntropy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEntropy.Unmarshal(m, b)
}
func (m *GetEntropy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEntropy.Marshal(b, m, deterministic)
}
func (m *GetEntropy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEntropy.Merge(m, src)
}
func (m *GetEntropy) XXX_Size() int {
	return xxx_messageInfo_GetEntropy.Size(m)
}
func (m *GetEntropy) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEntropy.DiscardUnknown(m)
}

var xxx_messageInfo_GetEntropy proto.InternalMessageInfo

func (m *GetEntropy) GetSize() uint32 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}


type Entropy struct {
	Entropy              []byte   `protobuf:"bytes,1,req,name=entropy" json:"entropy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entropy) Reset()         { *m = Entropy{} }
func (m *Entropy) String() string { return proto.CompactTextString(m) }
func (*Entropy) ProtoMessage()    {}
func (*Entropy) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{10}
}

func (m *Entropy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entropy.Unmarshal(m, b)
}
func (m *Entropy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entropy.Marshal(b, m, deterministic)
}
func (m *Entropy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entropy.Merge(m, src)
}
func (m *Entropy) XXX_Size() int {
	return xxx_messageInfo_Entropy.Size(m)
}
func (m *Entropy) XXX_DiscardUnknown() {
	xxx_messageInfo_Entropy.DiscardUnknown(m)
}

var xxx_messageInfo_Entropy proto.InternalMessageInfo

func (m *Entropy) GetEntropy() []byte {
	if m != nil {
		return m.Entropy
	}
	return nil
}

type WipeDevice struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WipeDevice) Reset()         { *m = WipeDevice{} }
func (m *WipeDevice) String() string { return proto.CompactTextString(m) }
func (*WipeDevice) ProtoMessage()    {}
func (*WipeDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{11}
}

func (m *WipeDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WipeDevice.Unmarshal(m, b)
}
func (m *WipeDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WipeDevice.Marshal(b, m, deterministic)
}
func (m *WipeDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WipeDevice.Merge(m, src)
}
func (m *WipeDevice) XXX_Size() int {
	return xxx_messageInfo_WipeDevice.Size(m)
}
func (m *WipeDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_WipeDevice.DiscardUnknown(m)
}

var xxx_messageInfo_WipeDevice proto.InternalMessageInfo


type LoadDevice struct {
	Mnemonic             *string     `protobuf:"bytes,1,opt,name=mnemonic" json:"mnemonic,omitempty"`
	Node                 *HDNodeType `protobuf:"bytes,2,opt,name=node" json:"node,omitempty"`
	Pin                  *string     `protobuf:"bytes,3,opt,name=pin" json:"pin,omitempty"`
	PassphraseProtection *bool       `protobuf:"varint,4,opt,name=passphrase_protection,json=passphraseProtection" json:"passphrase_protection,omitempty"`
	Language             *string     `protobuf:"bytes,5,opt,name=language,def=english" json:"language,omitempty"`
	Label                *string     `protobuf:"bytes,6,opt,name=label" json:"label,omitempty"`
	SkipChecksum         *bool       `protobuf:"varint,7,opt,name=skip_checksum,json=skipChecksum" json:"skip_checksum,omitempty"`
	U2FCounter           *uint32     `protobuf:"varint,8,opt,name=u2f_counter,json=u2fCounter" json:"u2f_counter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *LoadDevice) Reset()         { *m = LoadDevice{} }
func (m *LoadDevice) String() string { return proto.CompactTextString(m) }
func (*LoadDevice) ProtoMessage()    {}
func (*LoadDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{12}
}

func (m *LoadDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadDevice.Unmarshal(m, b)
}
func (m *LoadDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadDevice.Marshal(b, m, deterministic)
}
func (m *LoadDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadDevice.Merge(m, src)
}
func (m *LoadDevice) XXX_Size() int {
	return xxx_messageInfo_LoadDevice.Size(m)
}
func (m *LoadDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadDevice.DiscardUnknown(m)
}

var xxx_messageInfo_LoadDevice proto.InternalMessageInfo

const Default_LoadDevice_Language string = "english"

func (m *LoadDevice) GetMnemonic() string {
	if m != nil && m.Mnemonic != nil {
		return *m.Mnemonic
	}
	return ""
}

func (m *LoadDevice) GetNode() *HDNodeType {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *LoadDevice) GetPin() string {
	if m != nil && m.Pin != nil {
		return *m.Pin
	}
	return ""
}

func (m *LoadDevice) GetPassphraseProtection() bool {
	if m != nil && m.PassphraseProtection != nil {
		return *m.PassphraseProtection
	}
	return false
}

func (m *LoadDevice) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return Default_LoadDevice_Language
}

func (m *LoadDevice) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *LoadDevice) GetSkipChecksum() bool {
	if m != nil && m.SkipChecksum != nil {
		return *m.SkipChecksum
	}
	return false
}

func (m *LoadDevice) GetU2FCounter() uint32 {
	if m != nil && m.U2FCounter != nil {
		return *m.U2FCounter
	}
	return 0
}


type ResetDevice struct {
	DisplayRandom        *bool    `protobuf:"varint,1,opt,name=display_random,json=displayRandom" json:"display_random,omitempty"`
	Strength             *uint32  `protobuf:"varint,2,opt,name=strength,def=256" json:"strength,omitempty"`
	PassphraseProtection *bool    `protobuf:"varint,3,opt,name=passphrase_protection,json=passphraseProtection" json:"passphrase_protection,omitempty"`
	PinProtection        *bool    `protobuf:"varint,4,opt,name=pin_protection,json=pinProtection" json:"pin_protection,omitempty"`
	Language             *string  `protobuf:"bytes,5,opt,name=language,def=english" json:"language,omitempty"`
	Label                *string  `protobuf:"bytes,6,opt,name=label" json:"label,omitempty"`
	U2FCounter           *uint32  `protobuf:"varint,7,opt,name=u2f_counter,json=u2fCounter" json:"u2f_counter,omitempty"`
	SkipBackup           *bool    `protobuf:"varint,8,opt,name=skip_backup,json=skipBackup" json:"skip_backup,omitempty"`
	NoBackup             *bool    `protobuf:"varint,9,opt,name=no_backup,json=noBackup" json:"no_backup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetDevice) Reset()         { *m = ResetDevice{} }
func (m *ResetDevice) String() string { return proto.CompactTextString(m) }
func (*ResetDevice) ProtoMessage()    {}
func (*ResetDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{13}
}

func (m *ResetDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetDevice.Unmarshal(m, b)
}
func (m *ResetDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetDevice.Marshal(b, m, deterministic)
}
func (m *ResetDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetDevice.Merge(m, src)
}
func (m *ResetDevice) XXX_Size() int {
	return xxx_messageInfo_ResetDevice.Size(m)
}
func (m *ResetDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetDevice.DiscardUnknown(m)
}

var xxx_messageInfo_ResetDevice proto.InternalMessageInfo

const Default_ResetDevice_Strength uint32 = 256
const Default_ResetDevice_Language string = "english"

func (m *ResetDevice) GetDisplayRandom() bool {
	if m != nil && m.DisplayRandom != nil {
		return *m.DisplayRandom
	}
	return false
}

func (m *ResetDevice) GetStrength() uint32 {
	if m != nil && m.Strength != nil {
		return *m.Strength
	}
	return Default_ResetDevice_Strength
}

func (m *ResetDevice) GetPassphraseProtection() bool {
	if m != nil && m.PassphraseProtection != nil {
		return *m.PassphraseProtection
	}
	return false
}

func (m *ResetDevice) GetPinProtection() bool {
	if m != nil && m.PinProtection != nil {
		return *m.PinProtection
	}
	return false
}

func (m *ResetDevice) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return Default_ResetDevice_Language
}

func (m *ResetDevice) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *ResetDevice) GetU2FCounter() uint32 {
	if m != nil && m.U2FCounter != nil {
		return *m.U2FCounter
	}
	return 0
}

func (m *ResetDevice) GetSkipBackup() bool {
	if m != nil && m.SkipBackup != nil {
		return *m.SkipBackup
	}
	return false
}

func (m *ResetDevice) GetNoBackup() bool {
	if m != nil && m.NoBackup != nil {
		return *m.NoBackup
	}
	return false
}


type BackupDevice struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BackupDevice) Reset()         { *m = BackupDevice{} }
func (m *BackupDevice) String() string { return proto.CompactTextString(m) }
func (*BackupDevice) ProtoMessage()    {}
func (*BackupDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{14}
}

func (m *BackupDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackupDevice.Unmarshal(m, b)
}
func (m *BackupDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackupDevice.Marshal(b, m, deterministic)
}
func (m *BackupDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackupDevice.Merge(m, src)
}
func (m *BackupDevice) XXX_Size() int {
	return xxx_messageInfo_BackupDevice.Size(m)
}
func (m *BackupDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_BackupDevice.DiscardUnknown(m)
}

var xxx_messageInfo_BackupDevice proto.InternalMessageInfo


type EntropyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntropyRequest) Reset()         { *m = EntropyRequest{} }
func (m *EntropyRequest) String() string { return proto.CompactTextString(m) }
func (*EntropyRequest) ProtoMessage()    {}
func (*EntropyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{15}
}

func (m *EntropyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntropyRequest.Unmarshal(m, b)
}
func (m *EntropyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntropyRequest.Marshal(b, m, deterministic)
}
func (m *EntropyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntropyRequest.Merge(m, src)
}
func (m *EntropyRequest) XXX_Size() int {
	return xxx_messageInfo_EntropyRequest.Size(m)
}
func (m *EntropyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EntropyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EntropyRequest proto.InternalMessageInfo


type EntropyAck struct {
	Entropy              []byte   `protobuf:"bytes,1,opt,name=entropy" json:"entropy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntropyAck) Reset()         { *m = EntropyAck{} }
func (m *EntropyAck) String() string { return proto.CompactTextString(m) }
func (*EntropyAck) ProtoMessage()    {}
func (*EntropyAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{16}
}

func (m *EntropyAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntropyAck.Unmarshal(m, b)
}
func (m *EntropyAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntropyAck.Marshal(b, m, deterministic)
}
func (m *EntropyAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntropyAck.Merge(m, src)
}
func (m *EntropyAck) XXX_Size() int {
	return xxx_messageInfo_EntropyAck.Size(m)
}
func (m *EntropyAck) XXX_DiscardUnknown() {
	xxx_messageInfo_EntropyAck.DiscardUnknown(m)
}

var xxx_messageInfo_EntropyAck proto.InternalMessageInfo

func (m *EntropyAck) GetEntropy() []byte {
	if m != nil {
		return m.Entropy
	}
	return nil
}


type RecoveryDevice struct {
	WordCount            *uint32 `protobuf:"varint,1,opt,name=word_count,json=wordCount" json:"word_count,omitempty"`
	PassphraseProtection *bool   `protobuf:"varint,2,opt,name=passphrase_protection,json=passphraseProtection" json:"passphrase_protection,omitempty"`
	PinProtection        *bool   `protobuf:"varint,3,opt,name=pin_protection,json=pinProtection" json:"pin_protection,omitempty"`
	Language             *string `protobuf:"bytes,4,opt,name=language,def=english" json:"language,omitempty"`
	Label                *string `protobuf:"bytes,5,opt,name=label" json:"label,omitempty"`
	EnforceWordlist      *bool   `protobuf:"varint,6,opt,name=enforce_wordlist,json=enforceWordlist" json:"enforce_wordlist,omitempty"`

	Type                 *RecoveryDevice_RecoveryDeviceType `protobuf:"varint,8,opt,name=type,enum=hw.trezor.messages.management.RecoveryDevice_RecoveryDeviceType" json:"type,omitempty"`
	U2FCounter           *uint32                            `protobuf:"varint,9,opt,name=u2f_counter,json=u2fCounter" json:"u2f_counter,omitempty"`
	DryRun               *bool                              `protobuf:"varint,10,opt,name=dry_run,json=dryRun" json:"dry_run,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *RecoveryDevice) Reset()         { *m = RecoveryDevice{} }
func (m *RecoveryDevice) String() string { return proto.CompactTextString(m) }
func (*RecoveryDevice) ProtoMessage()    {}
func (*RecoveryDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{17}
}

func (m *RecoveryDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecoveryDevice.Unmarshal(m, b)
}
func (m *RecoveryDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecoveryDevice.Marshal(b, m, deterministic)
}
func (m *RecoveryDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecoveryDevice.Merge(m, src)
}
func (m *RecoveryDevice) XXX_Size() int {
	return xxx_messageInfo_RecoveryDevice.Size(m)
}
func (m *RecoveryDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_RecoveryDevice.DiscardUnknown(m)
}

var xxx_messageInfo_RecoveryDevice proto.InternalMessageInfo

const Default_RecoveryDevice_Language string = "english"

func (m *RecoveryDevice) GetWordCount() uint32 {
	if m != nil && m.WordCount != nil {
		return *m.WordCount
	}
	return 0
}

func (m *RecoveryDevice) GetPassphraseProtection() bool {
	if m != nil && m.PassphraseProtection != nil {
		return *m.PassphraseProtection
	}
	return false
}

func (m *RecoveryDevice) GetPinProtection() bool {
	if m != nil && m.PinProtection != nil {
		return *m.PinProtection
	}
	return false
}

func (m *RecoveryDevice) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return Default_RecoveryDevice_Language
}

func (m *RecoveryDevice) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *RecoveryDevice) GetEnforceWordlist() bool {
	if m != nil && m.EnforceWordlist != nil {
		return *m.EnforceWordlist
	}
	return false
}

func (m *RecoveryDevice) GetType() RecoveryDevice_RecoveryDeviceType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return RecoveryDevice_RecoveryDeviceType_ScrambledWords
}

func (m *RecoveryDevice) GetU2FCounter() uint32 {
	if m != nil && m.U2FCounter != nil {
		return *m.U2FCounter
	}
	return 0
}

func (m *RecoveryDevice) GetDryRun() bool {
	if m != nil && m.DryRun != nil {
		return *m.DryRun
	}
	return false
}


type WordRequest struct {
	Type                 *WordRequest_WordRequestType `protobuf:"varint,1,opt,name=type,enum=hw.trezor.messages.management.WordRequest_WordRequestType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *WordRequest) Reset()         { *m = WordRequest{} }
func (m *WordRequest) String() string { return proto.CompactTextString(m) }
func (*WordRequest) ProtoMessage()    {}
func (*WordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{18}
}

func (m *WordRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WordRequest.Unmarshal(m, b)
}
func (m *WordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WordRequest.Marshal(b, m, deterministic)
}
func (m *WordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WordRequest.Merge(m, src)
}
func (m *WordRequest) XXX_Size() int {
	return xxx_messageInfo_WordRequest.Size(m)
}
func (m *WordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WordRequest proto.InternalMessageInfo

func (m *WordRequest) GetType() WordRequest_WordRequestType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return WordRequest_WordRequestType_Plain
}


type WordAck struct {
	Word                 *string  `protobuf:"bytes,1,req,name=word" json:"word,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WordAck) Reset()         { *m = WordAck{} }
func (m *WordAck) String() string { return proto.CompactTextString(m) }
func (*WordAck) ProtoMessage()    {}
func (*WordAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{19}
}

func (m *WordAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WordAck.Unmarshal(m, b)
}
func (m *WordAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WordAck.Marshal(b, m, deterministic)
}
func (m *WordAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WordAck.Merge(m, src)
}
func (m *WordAck) XXX_Size() int {
	return xxx_messageInfo_WordAck.Size(m)
}
func (m *WordAck) XXX_DiscardUnknown() {
	xxx_messageInfo_WordAck.DiscardUnknown(m)
}

var xxx_messageInfo_WordAck proto.InternalMessageInfo

func (m *WordAck) GetWord() string {
	if m != nil && m.Word != nil {
		return *m.Word
	}
	return ""
}

type SetU2FCounter struct {
	U2FCounter           *uint32  `protobuf:"varint,1,opt,name=u2f_counter,json=u2fCounter" json:"u2f_counter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetU2FCounter) Reset()         { *m = SetU2FCounter{} }
func (m *SetU2FCounter) String() string { return proto.CompactTextString(m) }
func (*SetU2FCounter) ProtoMessage()    {}
func (*SetU2FCounter) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c720c20d27aa029, []int{20}
}

func (m *SetU2FCounter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetU2FCounter.Unmarshal(m, b)
}
func (m *SetU2FCounter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetU2FCounter.Marshal(b, m, deterministic)
}
func (m *SetU2FCounter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetU2FCounter.Merge(m, src)
}
func (m *SetU2FCounter) XXX_Size() int {
	return xxx_messageInfo_SetU2FCounter.Size(m)
}
func (m *SetU2FCounter) XXX_DiscardUnknown() {
	xxx_messageInfo_SetU2FCounter.DiscardUnknown(m)
}

var xxx_messageInfo_SetU2FCounter proto.InternalMessageInfo

func (m *SetU2FCounter) GetU2FCounter() uint32 {
	if m != nil && m.U2FCounter != nil {
		return *m.U2FCounter
	}
	return 0
}

func init() {
	proto.RegisterEnum("hw.trezor.messages.management.ApplySettings_PassphraseSourceType", ApplySettings_PassphraseSourceType_name, ApplySettings_PassphraseSourceType_value)
	proto.RegisterEnum("hw.trezor.messages.management.RecoveryDevice_RecoveryDeviceType", RecoveryDevice_RecoveryDeviceType_name, RecoveryDevice_RecoveryDeviceType_value)
	proto.RegisterEnum("hw.trezor.messages.management.WordRequest_WordRequestType", WordRequest_WordRequestType_name, WordRequest_WordRequestType_value)
	proto.RegisterType((*Initialize)(nil), "hw.trezor.messages.management.Initialize")
	proto.RegisterType((*GetFeatures)(nil), "hw.trezor.messages.management.GetFeatures")
	proto.RegisterType((*Features)(nil), "hw.trezor.messages.management.Features")
	proto.RegisterType((*ClearSession)(nil), "hw.trezor.messages.management.ClearSession")
	proto.RegisterType((*ApplySettings)(nil), "hw.trezor.messages.management.ApplySettings")
	proto.RegisterType((*ApplyFlags)(nil), "hw.trezor.messages.management.ApplyFlags")
	proto.RegisterType((*ChangePin)(nil), "hw.trezor.messages.management.ChangePin")
	proto.RegisterType((*Ping)(nil), "hw.trezor.messages.management.Ping")
	proto.RegisterType((*Cancel)(nil), "hw.trezor.messages.management.Cancel")
	proto.RegisterType((*GetEntropy)(nil), "hw.trezor.messages.management.GetEntropy")
	proto.RegisterType((*Entropy)(nil), "hw.trezor.messages.management.Entropy")
	proto.RegisterType((*WipeDevice)(nil), "hw.trezor.messages.management.WipeDevice")
	proto.RegisterType((*LoadDevice)(nil), "hw.trezor.messages.management.LoadDevice")
	proto.RegisterType((*ResetDevice)(nil), "hw.trezor.messages.management.ResetDevice")
	proto.RegisterType((*BackupDevice)(nil), "hw.trezor.messages.management.BackupDevice")
	proto.RegisterType((*EntropyRequest)(nil), "hw.trezor.messages.management.EntropyRequest")
	proto.RegisterType((*EntropyAck)(nil), "hw.trezor.messages.management.EntropyAck")
	proto.RegisterType((*RecoveryDevice)(nil), "hw.trezor.messages.management.RecoveryDevice")
	proto.RegisterType((*WordRequest)(nil), "hw.trezor.messages.management.WordRequest")
	proto.RegisterType((*WordAck)(nil), "hw.trezor.messages.management.WordAck")
	proto.RegisterType((*SetU2FCounter)(nil), "hw.trezor.messages.management.SetU2FCounter")
}

func init() { proto.RegisterFile("messages-management.proto", fileDescriptor_0c720c20d27aa029) }

var fileDescriptor_0c720c20d27aa029 = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x57, 0xdd, 0x6e, 0xdb, 0xc8,
	0x15, 0x8e, 0x7e, 0x62, 0x49, 0xc7, 0xfa, 0xcb, 0xd4, 0x8e, 0xe9, 0xb8, 0x6e, 0x1c, 0xba, 0x6e,
	0x12, 0x04, 0x15, 0x0a, 0x17, 0x09, 0x90, 0x5c, 0x14, 0x75, 0xec, 0xfc, 0x21, 0x71, 0x6a, 0xd0,
	0x6e, 0x02, 0xf4, 0x86, 0x18, 0x91, 0x47, 0xd2, 0xd4, 0xe4, 0x0c, 0xcb, 0x19, 0xda, 0x55, 0x5e,
	0x60, 0x6f, 0xf6, 0x45, 0x16, 0xfb, 0x1c, 0x7b, 0xb5, 0xcf, 0xb0, 0xef, 0xb2, 0x98, 0x19, 0x52,
	0xa2, 0x65, 0x3b, 0x46, 0x76, 0xef, 0xe6, 0x7c, 0xe7, 0xe3, 0x68, 0xce, 0x77, 0xbe, 0x39, 0x63,
	0xc3, 0x7a, 0x8c, 0x52, 0xd2, 0x31, 0xca, 0xbf, 0xc6, 0x94, 0xd3, 0x31, 0xc6, 0xc8, 0xd5, 0x20,
	0x49, 0x85, 0x12, 0x64, 0x73, 0x72, 0x3e, 0x50, 0x29, 0x7e, 0x11, 0xe9, 0xa0, 0x20, 0x0d, 0xe6,
	0xa4, 0x7b, 0xab, 0xb3, 0x2f, 0x03, 0x11, 0xc7, 0x82, 0xdb, 0xaf, 0xdc, 0xf7, 0x00, 0xef, 0x38,
	0x53, 0x8c, 0x46, 0xec, 0x0b, 0x92, 0x15, 0xb8, 0x2d, 0x15, 0x55, 0xe8, 0x54, 0xb6, 0x2a, 0x8f,
	0xda, 0x9e, 0x0d, 0xc8, 0x43, 0xe8, 0xc9, 0x53, 0x96, 0xf8, 0x09, 0x95, 0x32, 0x99, 0xa4, 0x54,
	0xa2, 0x53, 0xdd, 0xaa, 0x3c, 0x6a, 0x7a, 0x5d, 0x0d, 0x1f, 0xcd, 0x50, 0xb7, 0x03, 0xcb, 0x6f,
	0x50, 0xbd, 0x46, 0xaa, 0xb2, 0x14, 0xa5, 0xfb, 0x7d, 0x03, 0x9a, 0x45, 0x40, 0xee, 0xc2, 0xd2,
	0x19, 0xf2, 0x50, 0xa4, 0x66, 0xef, 0x96, 0x97, 0x47, 0x64, 0x1b, 0x3a, 0x31, 0xfd, 0xaf, 0x48,
	0xfd, 0x33, 0x4c, 0x25, 0x13, 0xdc, 0x6c, 0xdd, 0xf1, 0xda, 0x06, 0xfc, 0x64, 0x31, 0x43, 0x62,
	0xbc, 0x44, 0xaa, 0xe5, 0x24, 0x0d, 0x96, 0x48, 0x09, 0x55, 0xc1, 0x64, 0x46, 0xaa, 0x5b, 0x92,
	0x01, 0x0b, 0xd2, 0x43, 0xe8, 0x0d, 0x85, 0x50, 0x91, 0xa0, 0x21, 0xa6, 0x7e, 0x2c, 0x42, 0x74,
	0x6e, 0xdb, 0x5a, 0xe6, 0xf0, 0xa1, 0x08, 0x91, 0x6c, 0x40, 0x2b, 0xc4, 0x33, 0x16, 0xa0, 0xcf,
	0x42, 0x67, 0xc9, 0x1c, 0xb9, 0x69, 0x81, 0x77, 0x21, 0xd9, 0x81, 0x6e, 0xc2, 0xb8, 0xaf, 0x25,
	0xc4, 0x40, 0xe9, 0xdf, 0x6a, 0x98, 0x4d, 0x3a, 0x09, 0xe3, 0x47, 0x33, 0x90, 0xfc, 0x1d, 0x56,
	0xe7, 0x9a, 0x95, 0xd9, 0x4d, 0xc3, 0x5e, 0x99, 0x27, 0x4b, 0x1f, 0xdd, 0x83, 0x66, 0x44, 0xf9,
	0x38, 0xa3, 0x63, 0x74, 0x5a, 0xf6, 0x77, 0x8b, 0x58, 0xf7, 0x27, 0xa2, 0x43, 0x8c, 0x1c, 0x30,
	0x09, 0x1b, 0x90, 0x2d, 0x58, 0x66, 0xb3, 0x1e, 0x86, 0x4e, 0xdb, 0x6c, 0x5e, 0x86, 0xf4, 0x9e,
	0x29, 0x9e, 0x31, 0xa3, 0x4a, 0xc7, 0xb4, 0x76, 0x16, 0x2f, 0x28, 0x32, 0xa1, 0x72, 0xe2, 0x74,
	0x0d, 0xa5, 0xa4, 0xc8, 0x5b, 0x2a, 0x27, 0x7a, 0x13, 0x16, 0x27, 0x22, 0x55, 0x18, 0x3a, 0x3d,
	0xf3, 0x1b, 0xb3, 0x98, 0x6c, 0x02, 0x68, 0x41, 0x02, 0x1a, 0x4c, 0x30, 0x74, 0xfa, 0x26, 0xdb,
	0x4a, 0x18, 0xdf, 0x37, 0x00, 0x79, 0x02, 0x77, 0x4a, 0x42, 0xe4, 0xac, 0x3b, 0x86, 0xd5, 0x9f,
	0x27, 0x72, 0xf2, 0x63, 0xe8, 0x8f, 0x58, 0x1a, 0x9f, 0xd3, 0x54, 0x6b, 0x86, 0x12, 0xb9, 0x72,
	0x88, 0xe1, 0xf6, 0x0a, 0xfc, 0xc8, 0xc2, 0xe4, 0x01, 0xb4, 0x39, 0x62, 0x28, 0xfd, 0x21, 0x0d,
	0x4e, 0xb3, 0xc4, 0xf9, 0x83, 0x2d, 0xdd, 0x60, 0x2f, 0x0d, 0xa4, 0x25, 0x1b, 0x45, 0x74, 0x2c,
	0x9d, 0x15, 0xe3, 0x06, 0x1b, 0x68, 0x54, 0xf7, 0x3e, 0x72, 0x56, 0xad, 0x90, 0x26, 0x20, 0xeb,
	0xd0, 0x1c, 0x9d, 0xfb, 0xc6, 0x79, 0xce, 0x5d, 0x43, 0x6f, 0x8c, 0xce, 0x0f, 0x75, 0x58, 0xa4,
	0xb4, 0xdf, 0x9c, 0xb5, 0x59, 0x4a, 0x87, 0x79, 0xca, 0xb8, 0xcc, 0x71, 0x8a, 0xd4, 0x91, 0x0e,
	0xb5, 0x89, 0x46, 0xe7, 0x7e, 0xee, 0xfb, 0x75, 0xdb, 0xcc, 0xd1, 0xf9, 0x27, 0xeb, 0xfc, 0x3f,
	0x43, 0x77, 0x96, 0xf4, 0x4f, 0x71, 0x2a, 0x9d, 0x7b, 0x46, 0xf7, 0x76, 0xc1, 0x78, 0x8f, 0x53,
	0xa9, 0xa5, 0xcb, 0xf8, 0x88, 0x71, 0x26, 0x27, 0x18, 0x16, 0x75, 0x6e, 0x58, 0xe9, 0xe6, 0x89,
	0xbc, 0xd8, 0x0d, 0x68, 0x71, 0x51, 0x90, 0xfe, 0x68, 0x7b, 0xc4, 0x85, 0x4d, 0xba, 0x5d, 0x68,
	0xef, 0x47, 0x48, 0xd3, 0x63, 0x94, 0xba, 0xf1, 0xee, 0x77, 0x35, 0xe8, 0xec, 0x25, 0x49, 0x34,
	0x3d, 0x46, 0xa5, 0x18, 0x1f, 0xcb, 0x0b, 0xd6, 0xab, 0x5c, 0x67, 0xbd, 0x6a, 0xd9, 0x7a, 0x3b,
	0xd0, 0xcd, 0xb4, 0xb5, 0xe7, 0x93, 0xa1, 0x66, 0x2f, 0x42, 0x26, 0x71, 0x3e, 0x18, 0xc8, 0x9f,
	0x00, 0x26, 0x22, 0x46, 0x19, 0xa4, 0x88, 0xf6, 0x5e, 0xb6, 0xbd, 0x12, 0x42, 0xf8, 0x05, 0x7f,
	0x48, 0x91, 0xa5, 0x81, 0xbd, 0x97, 0xdd, 0xdd, 0xbd, 0xc1, 0x57, 0xe7, 0xda, 0xe0, 0x42, 0x05,
	0x83, 0xf9, 0x6f, 0x1e, 0x9b, 0x4d, 0x4e, 0xa6, 0x09, 0x96, 0x2d, 0x66, 0x51, 0xf2, 0x04, 0x08,
	0xcd, 0x94, 0xf0, 0x23, 0x11, 0x9c, 0xfa, 0x21, 0x46, 0x74, 0xea, 0xc7, 0xd2, 0xdc, 0xf2, 0x8e,
	0xd7, 0xd3, 0x99, 0x0f, 0x22, 0x38, 0x3d, 0xd0, 0xf8, 0xa1, 0xd4, 0x7e, 0x0c, 0x99, 0x4c, 0x34,
	0x29, 0x15, 0x8a, 0xce, 0xae, 0x7b, 0xc7, 0xeb, 0xe5, 0xb8, 0x97, 0xc3, 0xee, 0x53, 0x58, 0xb9,
	0xea, 0x04, 0xa4, 0x01, 0xb5, 0xbd, 0xe3, 0xf7, 0xfd, 0x5b, 0x04, 0x60, 0xe9, 0xe0, 0xd5, 0xa7,
	0x77, 0xfb, 0xaf, 0xfa, 0x15, 0xd2, 0x84, 0xfa, 0xdb, 0x7f, 0x1d, 0x9f, 0xf4, 0xab, 0xae, 0x0b,
	0x60, 0xca, 0x78, 0x5d, 0x78, 0xd3, 0x3a, 0xb6, 0x52, 0x72, 0xac, 0xbb, 0x0d, 0xad, 0xfd, 0x09,
	0xe5, 0x63, 0x3c, 0x62, 0x5c, 0x0f, 0xd3, 0x14, 0x63, 0x71, 0x66, 0xdb, 0xd4, 0xf4, 0xf2, 0xc8,
	0xfd, 0xa1, 0x02, 0xf5, 0x23, 0xc6, 0xc7, 0xc4, 0x81, 0x46, 0x2e, 0x56, 0xde, 0xc8, 0x22, 0xd4,
	0x7e, 0x1a, 0x66, 0x4a, 0x89, 0x0b, 0xd3, 0xcb, 0x8e, 0xf3, 0xbe, 0x4d, 0x94, 0x66, 0xd1, 0xe5,
	0x39, 0x57, 0xfb, 0xa6, 0x39, 0x57, 0xbf, 0x7e, 0xce, 0xb9, 0x4d, 0x58, 0xda, 0xa7, 0x3c, 0xc0,
	0xc8, 0xdd, 0x02, 0x78, 0x83, 0xea, 0x15, 0x57, 0xa9, 0x48, 0xa6, 0x84, 0x40, 0x5d, 0xb2, 0x2f,
	0xfa, 0xdc, 0xd5, 0x47, 0x1d, 0xcf, 0xac, 0xdd, 0x6d, 0x68, 0x14, 0x69, 0x07, 0x1a, 0x68, 0x97,
	0x86, 0xd1, 0xf6, 0x8a, 0xd0, 0x6d, 0x03, 0x7c, 0x66, 0x09, 0x1e, 0x98, 0x21, 0xed, 0xfe, 0x58,
	0x05, 0xf8, 0x20, 0x68, 0x68, 0x43, 0x6d, 0xed, 0x98, 0x63, 0x2c, 0x38, 0x0b, 0x0a, 0x6b, 0x17,
	0x31, 0x79, 0x0e, 0x75, 0xae, 0x1f, 0x02, 0xad, 0xc2, 0xf2, 0xee, 0xce, 0x55, 0x86, 0xcb, 0xdf,
	0xcc, 0xb7, 0x07, 0x1f, 0x45, 0x68, 0x4d, 0x65, 0x3e, 0x21, 0x7d, 0xa8, 0x25, 0xcc, 0xaa, 0xd2,
	0xf2, 0xf4, 0xf2, 0x37, 0x69, 0x41, 0xb6, 0x4b, 0x17, 0x4f, 0xdb, 0xbe, 0xf5, 0xa2, 0x81, 0x7c,
	0x1c, 0x31, 0x39, 0xb9, 0xea, 0x06, 0x2e, 0x95, 0x6f, 0xe0, 0x36, 0x74, 0xcc, 0xe3, 0x1c, 0x4c,
	0x30, 0x38, 0x95, 0x59, 0x9c, 0xbf, 0x44, 0x6d, 0x0d, 0xee, 0xe7, 0x18, 0xb9, 0x0f, 0xcb, 0xd9,
	0xee, 0xc8, 0x0f, 0x44, 0xc6, 0x15, 0xa6, 0xe6, 0xf9, 0xe9, 0x78, 0x90, 0xed, 0x8e, 0xf6, 0x2d,
	0xe2, 0xfe, 0x5c, 0x85, 0x65, 0x0f, 0x25, 0xaa, 0x5c, 0xae, 0x1d, 0xe8, 0xce, 0x3c, 0x4f, 0x79,
	0x28, 0xe2, 0xdc, 0x68, 0x9d, 0xc2, 0xf1, 0x06, 0x24, 0xf7, 0xa1, 0x29, 0x55, 0x8a, 0x7c, 0xac,
	0x26, 0xf6, 0xdd, 0x7e, 0x51, 0xdb, 0x7d, 0xfa, 0xcc, 0x9b, 0x81, 0xd7, 0xab, 0x51, 0xfb, 0x8a,
	0x1a, 0x97, 0x5d, 0x57, 0xbf, 0xca, 0x75, 0xbf, 0x43, 0xb4, 0x05, 0x3d, 0x1a, 0x8b, 0x7a, 0x68,
	0x82, 0x51, 0x35, 0x1f, 0xa5, 0xf6, 0xbd, 0x06, 0x0d, 0x5d, 0x35, 0x69, 0x5b, 0x97, 0x27, 0xad,
	0x5d, 0xe5, 0x5e, 0xec, 0x43, 0x37, 0xb7, 0xaf, 0x87, 0xff, 0xcb, 0x50, 0x2a, 0xf7, 0x2f, 0x00,
	0x39, 0xb2, 0x17, 0x9c, 0x5e, 0xf4, 0x74, 0xa5, 0xec, 0xe9, 0x5f, 0x6a, 0xd0, 0xf5, 0x30, 0x10,
	0x67, 0x98, 0x4e, 0xf3, 0xd6, 0x6c, 0x02, 0x9c, 0x8b, 0x34, 0xb4, 0x87, 0xcf, 0x67, 0x44, 0x4b,
	0x23, 0xe6, 0xec, 0xd7, 0x2b, 0x5e, 0xfd, 0x26, 0xc5, 0x6b, 0x37, 0x29, 0x5e, 0xbf, 0x51, 0xf1,
	0xdb, 0x65, 0xc5, 0x1f, 0x43, 0x1f, 0xf9, 0x48, 0xa4, 0x01, 0xfa, 0xfa, 0xac, 0x11, 0x93, 0xca,
	0xb4, 0xa4, 0xe9, 0xf5, 0x72, 0xfc, 0x73, 0x0e, 0x93, 0x13, 0xa8, 0xab, 0x69, 0x82, 0x46, 0xf4,
	0xee, 0xee, 0x3f, 0x6f, 0x98, 0xff, 0x17, 0xd5, 0x59, 0x08, 0xed, 0x4d, 0xd5, 0xbb, 0x2d, 0xb6,
	0xbc, 0x75, 0xa9, 0xe5, 0x6b, 0xd0, 0x08, 0xd3, 0xa9, 0x9f, 0x66, 0xdc, 0xfc, 0x75, 0xd5, 0xf4,
	0x96, 0xc2, 0x74, 0xea, 0x65, 0xdc, 0xfd, 0x0f, 0x90, 0xcb, 0xbb, 0x92, 0x1d, 0x78, 0x70, 0x19,
	0xf5, 0x8f, 0x83, 0x94, 0xc6, 0xc3, 0x08, 0x43, 0x5d, 0x8d, 0xec, 0xdf, 0x22, 0x9b, 0xb0, 0x7e,
	0x05, 0xed, 0x90, 0xaa, 0x94, 0xfd, 0xbf, 0x5f, 0x71, 0x7f, 0xaa, 0xc0, 0xb2, 0xa6, 0xe6, 0xbe,
	0x20, 0x1f, 0xf3, 0xda, 0x2b, 0xa6, 0xf6, 0x17, 0x37, 0xd4, 0x5e, 0xfa, 0xb2, 0xbc, 0x9e, 0x57,
	0xed, 0x8e, 0xa0, 0xb7, 0x90, 0x20, 0xeb, 0xb0, 0xba, 0x00, 0xf9, 0x47, 0x11, 0x65, 0xbc, 0x7f,
	0x8b, 0x6c, 0xc0, 0xda, 0x62, 0xca, 0x9e, 0xf4, 0x79, 0xbf, 0x72, 0x7d, 0xf2, 0x59, 0xbf, 0xea,
	0x6e, 0x42, 0x43, 0x27, 0xb5, 0x99, 0x09, 0xd4, 0x75, 0x87, 0xcd, 0x74, 0x6e, 0x79, 0x66, 0xed,
	0xfe, 0x0d, 0x3a, 0xc7, 0xa8, 0xfe, 0xbd, 0xfb, 0xba, 0x74, 0xbf, 0xca, 0xdd, 0xa8, 0x2c, 0x76,
	0xe3, 0xe5, 0x3f, 0x60, 0x3b, 0x10, 0xf1, 0x40, 0x52, 0x25, 0xe4, 0x84, 0x45, 0x74, 0x28, 0x0b,
	0x21, 0x22, 0x36, 0xb4, 0xff, 0xbb, 0x0c, 0xb3, 0xd1, 0xcb, 0xb5, 0x13, 0x03, 0x1e, 0x5a, 0x71,
	0x0e, 0x67, 0xd2, 0xfc, 0x1a, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x6e, 0xfc, 0x59, 0x29, 0x0d, 0x00,
	0x00,
}
