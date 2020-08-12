/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-31T04:35:53Z.
 *
 * API version: 1.0.9-2110
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/* Customised for isctl */
package openapi

import (
	"encoding/json"
)

// NiatelemetryNiaInventoryAllOf Definition of the list of properties defined in 'niatelemetry.NiaInventory', excluding properties defined in parent classes.
type NiatelemetryNiaInventoryAllOf struct {
	// CPU usage of device being inventoried. This determines the percentage of CPU resources used.
	Cpu *float32 `json:"Cpu,omitempty" yaml:"Cpu,omitempty"`
	// Last crash reset reason of device being inventoried. This determines the last reason for a device's restart due to crash of the system.
	CrashResetLogs *string `json:"CrashResetLogs,omitempty" yaml:"CrashResetLogs,omitempty"`
	// Name of device being inventoried. The name the user assigns to the device is inventoried here.
	DeviceName *string `json:"DeviceName,omitempty" yaml:"DeviceName,omitempty"`
	// Type of device being inventoried. This determines whether the device is a controller, leaf or spine.
	DeviceType *string               `json:"DeviceType,omitempty" yaml:"DeviceType,omitempty"`
	Disk       *NiatelemetryDiskinfo `json:"Disk,omitempty" yaml:"Disk,omitempty"`
	// Scale of fabric extendors utilized.
	FexCount *int64 `json:"FexCount,omitempty" yaml:"FexCount,omitempty"`
	// Last log in time device being inventoried. This determines the last login time on the device.
	LogInTime *string `json:"LogInTime,omitempty" yaml:"LogInTime,omitempty"`
	// Last log out time of device being inventoried. This determines the last logout time on the device.
	LogOutTime *string `json:"LogOutTime,omitempty" yaml:"LogOutTime,omitempty"`
	// Number of Macsec configured interfaces on a TOR.
	MacSecCount *int64 `json:"MacSecCount,omitempty" yaml:"MacSecCount,omitempty"`
	// Number of Macsec configured interfaces on a Spine.
	MacSecFabCount *int64 `json:"MacSecFabCount,omitempty" yaml:"MacSecFabCount,omitempty"`
	// Number of total Macsec configured interfaces for all nodes.
	MacsecTotalCount *int64 `json:"MacsecTotalCount,omitempty" yaml:"MacsecTotalCount,omitempty"`
	// Memory usage of device being inventoried. This determines the percentage of memory resources used.
	Memory *int64 `json:"Memory,omitempty" yaml:"Memory,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty" yaml:"RecordType,omitempty"`
	// Version of record being pushed. This determines what was the API version for data available from the device.
	RecordVersion *string `json:"RecordVersion,omitempty" yaml:"RecordVersion,omitempty"`
	// Total nuumber of v4 and v6 routes per node.
	RoutePrefixCount *int64 `json:"RoutePrefixCount,omitempty" yaml:"RoutePrefixCount,omitempty"`
	// Scale of v4 routes per node.
	RoutePrefixV4Count *int64 `json:"RoutePrefixV4Count,omitempty" yaml:"RoutePrefixV4Count,omitempty"`
	// Scale of v6 routes per node.
	RoutePrefixV6Count *int64 `json:"RoutePrefixV6Count,omitempty" yaml:"RoutePrefixV6Count,omitempty"`
	// Serial number of device being invetoried. The serial number is unique per device and will be used as the key.
	Serial *string `json:"Serial,omitempty" yaml:"Serial,omitempty"`
	// Last software downloaded of device being inventoried. This determines if software download API was used.
	SoftwareDownload *string `json:"SoftwareDownload,omitempty" yaml:"SoftwareDownload,omitempty"`
	// Software version of device being inventoried. The various software version values for each device are available on cisco.com.
	Version          *string                                  `json:"Version,omitempty" yaml:"Version,omitempty"`
	LicenseState     *NiatelemetryNiaLicenseStateRelationship `json:"LicenseState,omitempty" yaml:"LicenseState,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship     `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewNiatelemetryNiaInventoryAllOf instantiates a new NiatelemetryNiaInventoryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryNiaInventoryAllOf() *NiatelemetryNiaInventoryAllOf {
	this := NiatelemetryNiaInventoryAllOf{}
	return &this
}

// NewNiatelemetryNiaInventoryAllOfWithDefaults instantiates a new NiatelemetryNiaInventoryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryNiaInventoryAllOfWithDefaults() *NiatelemetryNiaInventoryAllOf {
	this := NiatelemetryNiaInventoryAllOf{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryAllOf) GetCpu() float32 {
	if o == nil || o.Cpu == nil {
		var ret float32
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryAllOf) GetCpuOk() (*float32, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryAllOf) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given float32 and assigns it to the Cpu field.
func (o *NiatelemetryNiaInventoryAllOf) SetCpu(v float32) {
	o.Cpu = &v
}

// GetCrashResetLogs returns the CrashResetLogs field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryAllOf) GetCrashResetLogs() string {
	if o == nil || o.CrashResetLogs == nil {
		var ret string
		return ret
	}
	return *o.CrashResetLogs
}

// GetCrashResetLogsOk returns a tuple with the CrashResetLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryAllOf) GetCrashResetLogsOk() (*string, bool) {
	if o == nil || o.CrashResetLogs == nil {
		return nil, false
	}
	return o.CrashResetLogs, true
}

// HasCrashResetLogs returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryAllOf) HasCrashResetLogs() bool {
	if o != nil && o.CrashResetLogs != nil {
		return true
	}

	return false
}

// SetCrashResetLogs gets a reference to the given string and assigns it to the CrashResetLogs field.
func (o *NiatelemetryNiaInventoryAllOf) SetCrashResetLogs(v string) {
	o.CrashResetLogs = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryAllOf) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryAllOf) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryAllOf) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *NiatelemetryNiaInventoryAllOf) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryAllOf) GetDeviceType() string {
	if o == nil || o.DeviceType == nil {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryAllOf) GetDeviceTypeOk() (*string, bool) {
	if o == nil || o.DeviceType == nil {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryAllOf) HasDeviceType() bool {
	if o != nil && o.DeviceType != nil {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *NiatelemetryNiaInventoryAllOf) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryAllOf) GetDisk() NiatelemetryDiskinfo {
	if o == nil || o.Disk == nil {
		var ret NiatelemetryDiskinfo
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryAllOf) GetDiskOk() (*NiatelemetryDiskinfo, bool) {
	if o == nil || o.Disk == nil {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryAllOf) HasDisk() bool {
	if o != nil && o.Disk != nil {
		return true
	}

	return false
}

// SetDisk gets a reference to the given NiatelemetryDiskinfo and assigns it to the Disk field.
func (o *NiatelemetryNiaInventoryAllOf) SetDisk(v NiatelemetryDiskinfo) {
	o.Disk = &v
}

// GetFexCount returns the FexCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryAllOf) GetFexCount() int64 {
	if o == nil || o.FexCount == nil {
		var ret int64
		return ret
	}
	return *o.FexCount
}

// GetFexCountOk returns a tuple with the FexCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryAllOf) GetFexCountOk() (*int64, bool) {
	if o == nil || o.FexCount == nil {
		return nil, false
	}
	return o.FexCount, true
}

// HasFexCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryAllOf) HasFexCount() bool {
	if o != nil && o.FexCount != nil {
		return true
	}

	return false
}

// SetFexCount gets a reference to the given int64 and assigns it to the FexCount field.
func (o *NiatelemetryNiaInventoryAllOf) SetFexCount(v int64) {
	o.FexCount = &v
}

// GetLogInTime returns the LogInTime field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryAllOf) GetLogInTime() string {
	if o == nil || o.LogInTime == nil {
		var ret string
		return ret
	}
	return *o.LogInTime
}

// GetLogInTimeOk returns a tuple with the LogInTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryAllOf) GetLogInTimeOk() (*string, bool) {
	if o == nil || o.LogInTime == nil {
		return nil, false
	}
	return o.LogInTime, true
}

// HasLogInTime returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryAllOf) HasLogInTime() bool {
	if o != nil && o.LogInTime != nil {
		return true
	}

	return false
}

// SetLogInTime gets a reference to the given string and assigns it to the LogInTime field.
func (o *NiatelemetryNiaInventoryAllOf) SetLogInTime(v string) {
	o.LogInTime = &v
}

// GetLogOutTime returns the LogOutTime field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryAllOf) GetLogOutTime() string {
	if o == nil || o.LogOutTime == nil {
		var ret string
		return ret
	}
	return *o.LogOutTime
}

// GetLogOutTimeOk returns a tuple with the LogOutTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryAllOf) GetLogOutTimeOk() (*string, bool) {
	if o == nil || o.LogOutTime == nil {
		return nil, false
	}
	return o.LogOutTime, true
}

// HasLogOutTime returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryAllOf) HasLogOutTime() bool {
	if o != nil && o.LogOutTime != nil {
		return true
	}

	return false
}

// SetLogOutTime gets a reference to the given string and assigns it to the LogOutTime field.
func (o *NiatelemetryNiaInventoryAllOf) SetLogOutTime(v string) {
	o.LogOutTime = &v
}

// GetMacSecCount returns the MacSecCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryAllOf) GetMacSecCount() int64 {
	if o == nil || o.MacSecCount == nil {
		var ret int64
		return ret
	}
	return *o.MacSecCount
}

// GetMacSecCountOk returns a tuple with the MacSecCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryAllOf) GetMacSecCountOk() (*int64, bool) {
	if o == nil || o.MacSecCount == nil {
		return nil, false
	}
	return o.MacSecCount, true
}

// HasMacSecCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryAllOf) HasMacSecCount() bool {
	if o != nil && o.MacSecCount != nil {
		return true
	}

	return false
}

// SetMacSecCount gets a reference to the given int64 and assigns it to the MacSecCount field.
func (o *NiatelemetryNiaInventoryAllOf) SetMacSecCount(v int64) {
	o.MacSecCount = &v
}

// GetMacSecFabCount returns the MacSecFabCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryAllOf) GetMacSecFabCount() int64 {
	if o == nil || o.MacSecFabCount == nil {
		var ret int64
		return ret
	}
	return *o.MacSecFabCount
}

// GetMacSecFabCountOk returns a tuple with the MacSecFabCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryAllOf) GetMacSecFabCountOk() (*int64, bool) {
	if o == nil || o.MacSecFabCount == nil {
		return nil, false
	}
	return o.MacSecFabCount, true
}

// HasMacSecFabCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryAllOf) HasMacSecFabCount() bool {
	if o != nil && o.MacSecFabCount != nil {
		return true
	}

	return false
}

// SetMacSecFabCount gets a reference to the given int64 and assigns it to the MacSecFabCount field.
func (o *NiatelemetryNiaInventoryAllOf) SetMacSecFabCount(v int64) {
	o.MacSecFabCount = &v
}

// GetMacsecTotalCount returns the MacsecTotalCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryAllOf) GetMacsecTotalCount() int64 {
	if o == nil || o.MacsecTotalCount == nil {
		var ret int64
		return ret
	}
	return *o.MacsecTotalCount
}

// GetMacsecTotalCountOk returns a tuple with the MacsecTotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryAllOf) GetMacsecTotalCountOk() (*int64, bool) {
	if o == nil || o.MacsecTotalCount == nil {
		return nil, false
	}
	return o.MacsecTotalCount, true
}

// HasMacsecTotalCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryAllOf) HasMacsecTotalCount() bool {
	if o != nil && o.MacsecTotalCount != nil {
		return true
	}

	return false
}

// SetMacsecTotalCount gets a reference to the given int64 and assigns it to the MacsecTotalCount field.
func (o *NiatelemetryNiaInventoryAllOf) SetMacsecTotalCount(v int64) {
	o.MacsecTotalCount = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryAllOf) GetMemory() int64 {
	if o == nil || o.Memory == nil {
		var ret int64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryAllOf) GetMemoryOk() (*int64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryAllOf) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int64 and assigns it to the Memory field.
func (o *NiatelemetryNiaInventoryAllOf) SetMemory(v int64) {
	o.Memory = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryAllOf) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryAllOf) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryNiaInventoryAllOf) SetRecordType(v string) {
	o.RecordType = &v
}

// GetRecordVersion returns the RecordVersion field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryAllOf) GetRecordVersion() string {
	if o == nil || o.RecordVersion == nil {
		var ret string
		return ret
	}
	return *o.RecordVersion
}

// GetRecordVersionOk returns a tuple with the RecordVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryAllOf) GetRecordVersionOk() (*string, bool) {
	if o == nil || o.RecordVersion == nil {
		return nil, false
	}
	return o.RecordVersion, true
}

// HasRecordVersion returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryAllOf) HasRecordVersion() bool {
	if o != nil && o.RecordVersion != nil {
		return true
	}

	return false
}

// SetRecordVersion gets a reference to the given string and assigns it to the RecordVersion field.
func (o *NiatelemetryNiaInventoryAllOf) SetRecordVersion(v string) {
	o.RecordVersion = &v
}

// GetRoutePrefixCount returns the RoutePrefixCount field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryAllOf) GetRoutePrefixCount() int64 {
	if o == nil || o.RoutePrefixCount == nil {
		var ret int64
		return ret
	}
	return *o.RoutePrefixCount
}

// GetRoutePrefixCountOk returns a tuple with the RoutePrefixCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryAllOf) GetRoutePrefixCountOk() (*int64, bool) {
	if o == nil || o.RoutePrefixCount == nil {
		return nil, false
	}
	return o.RoutePrefixCount, true
}

// HasRoutePrefixCount returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryAllOf) HasRoutePrefixCount() bool {
	if o != nil && o.RoutePrefixCount != nil {
		return true
	}

	return false
}

// SetRoutePrefixCount gets a reference to the given int64 and assigns it to the RoutePrefixCount field.
func (o *NiatelemetryNiaInventoryAllOf) SetRoutePrefixCount(v int64) {
	o.RoutePrefixCount = &v
}

// GetRoutePrefixV4Count returns the RoutePrefixV4Count field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryAllOf) GetRoutePrefixV4Count() int64 {
	if o == nil || o.RoutePrefixV4Count == nil {
		var ret int64
		return ret
	}
	return *o.RoutePrefixV4Count
}

// GetRoutePrefixV4CountOk returns a tuple with the RoutePrefixV4Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryAllOf) GetRoutePrefixV4CountOk() (*int64, bool) {
	if o == nil || o.RoutePrefixV4Count == nil {
		return nil, false
	}
	return o.RoutePrefixV4Count, true
}

// HasRoutePrefixV4Count returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryAllOf) HasRoutePrefixV4Count() bool {
	if o != nil && o.RoutePrefixV4Count != nil {
		return true
	}

	return false
}

// SetRoutePrefixV4Count gets a reference to the given int64 and assigns it to the RoutePrefixV4Count field.
func (o *NiatelemetryNiaInventoryAllOf) SetRoutePrefixV4Count(v int64) {
	o.RoutePrefixV4Count = &v
}

// GetRoutePrefixV6Count returns the RoutePrefixV6Count field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryAllOf) GetRoutePrefixV6Count() int64 {
	if o == nil || o.RoutePrefixV6Count == nil {
		var ret int64
		return ret
	}
	return *o.RoutePrefixV6Count
}

// GetRoutePrefixV6CountOk returns a tuple with the RoutePrefixV6Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryAllOf) GetRoutePrefixV6CountOk() (*int64, bool) {
	if o == nil || o.RoutePrefixV6Count == nil {
		return nil, false
	}
	return o.RoutePrefixV6Count, true
}

// HasRoutePrefixV6Count returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryAllOf) HasRoutePrefixV6Count() bool {
	if o != nil && o.RoutePrefixV6Count != nil {
		return true
	}

	return false
}

// SetRoutePrefixV6Count gets a reference to the given int64 and assigns it to the RoutePrefixV6Count field.
func (o *NiatelemetryNiaInventoryAllOf) SetRoutePrefixV6Count(v int64) {
	o.RoutePrefixV6Count = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryAllOf) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryAllOf) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryAllOf) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *NiatelemetryNiaInventoryAllOf) SetSerial(v string) {
	o.Serial = &v
}

// GetSoftwareDownload returns the SoftwareDownload field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryAllOf) GetSoftwareDownload() string {
	if o == nil || o.SoftwareDownload == nil {
		var ret string
		return ret
	}
	return *o.SoftwareDownload
}

// GetSoftwareDownloadOk returns a tuple with the SoftwareDownload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryAllOf) GetSoftwareDownloadOk() (*string, bool) {
	if o == nil || o.SoftwareDownload == nil {
		return nil, false
	}
	return o.SoftwareDownload, true
}

// HasSoftwareDownload returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryAllOf) HasSoftwareDownload() bool {
	if o != nil && o.SoftwareDownload != nil {
		return true
	}

	return false
}

// SetSoftwareDownload gets a reference to the given string and assigns it to the SoftwareDownload field.
func (o *NiatelemetryNiaInventoryAllOf) SetSoftwareDownload(v string) {
	o.SoftwareDownload = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *NiatelemetryNiaInventoryAllOf) SetVersion(v string) {
	o.Version = &v
}

// GetLicenseState returns the LicenseState field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryAllOf) GetLicenseState() NiatelemetryNiaLicenseStateRelationship {
	if o == nil || o.LicenseState == nil {
		var ret NiatelemetryNiaLicenseStateRelationship
		return ret
	}
	return *o.LicenseState
}

// GetLicenseStateOk returns a tuple with the LicenseState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryAllOf) GetLicenseStateOk() (*NiatelemetryNiaLicenseStateRelationship, bool) {
	if o == nil || o.LicenseState == nil {
		return nil, false
	}
	return o.LicenseState, true
}

// HasLicenseState returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryAllOf) HasLicenseState() bool {
	if o != nil && o.LicenseState != nil {
		return true
	}

	return false
}

// SetLicenseState gets a reference to the given NiatelemetryNiaLicenseStateRelationship and assigns it to the LicenseState field.
func (o *NiatelemetryNiaInventoryAllOf) SetLicenseState(v NiatelemetryNiaLicenseStateRelationship) {
	o.LicenseState = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryNiaInventoryAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryNiaInventoryAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryNiaInventoryAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryNiaInventoryAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryNiaInventoryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cpu != nil {
		toSerialize["Cpu"] = o.Cpu
	}
	if o.CrashResetLogs != nil {
		toSerialize["CrashResetLogs"] = o.CrashResetLogs
	}
	if o.DeviceName != nil {
		toSerialize["DeviceName"] = o.DeviceName
	}
	if o.DeviceType != nil {
		toSerialize["DeviceType"] = o.DeviceType
	}
	if o.Disk != nil {
		toSerialize["Disk"] = o.Disk
	}
	if o.FexCount != nil {
		toSerialize["FexCount"] = o.FexCount
	}
	if o.LogInTime != nil {
		toSerialize["LogInTime"] = o.LogInTime
	}
	if o.LogOutTime != nil {
		toSerialize["LogOutTime"] = o.LogOutTime
	}
	if o.MacSecCount != nil {
		toSerialize["MacSecCount"] = o.MacSecCount
	}
	if o.MacSecFabCount != nil {
		toSerialize["MacSecFabCount"] = o.MacSecFabCount
	}
	if o.MacsecTotalCount != nil {
		toSerialize["MacsecTotalCount"] = o.MacsecTotalCount
	}
	if o.Memory != nil {
		toSerialize["Memory"] = o.Memory
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.RecordVersion != nil {
		toSerialize["RecordVersion"] = o.RecordVersion
	}
	if o.RoutePrefixCount != nil {
		toSerialize["RoutePrefixCount"] = o.RoutePrefixCount
	}
	if o.RoutePrefixV4Count != nil {
		toSerialize["RoutePrefixV4Count"] = o.RoutePrefixV4Count
	}
	if o.RoutePrefixV6Count != nil {
		toSerialize["RoutePrefixV6Count"] = o.RoutePrefixV6Count
	}
	if o.Serial != nil {
		toSerialize["Serial"] = o.Serial
	}
	if o.SoftwareDownload != nil {
		toSerialize["SoftwareDownload"] = o.SoftwareDownload
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.LicenseState != nil {
		toSerialize["LicenseState"] = o.LicenseState
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableNiatelemetryNiaInventoryAllOf struct {
	value *NiatelemetryNiaInventoryAllOf
	isSet bool
}

func (v NullableNiatelemetryNiaInventoryAllOf) Get() *NiatelemetryNiaInventoryAllOf {
	return v.value
}

func (v *NullableNiatelemetryNiaInventoryAllOf) Set(val *NiatelemetryNiaInventoryAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryNiaInventoryAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryNiaInventoryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryNiaInventoryAllOf(val *NiatelemetryNiaInventoryAllOf) *NullableNiatelemetryNiaInventoryAllOf {
	return &NullableNiatelemetryNiaInventoryAllOf{value: val, isSet: true}
}

func (v NullableNiatelemetryNiaInventoryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryNiaInventoryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
