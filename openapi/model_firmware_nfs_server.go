/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-04-28T13:03:38Z.
 *
 * API version: 1.0.9-4267
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/* Customised for isctl */
package openapi

import (
	"encoding/json"
)

// FirmwareNfsServer Network share file server.
type FirmwareNfsServer struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// The location to the image file. The accepted format is IP-or-hostname/folder1/folder2/.../imageFile.
	FileLocation *string `json:"FileLocation,omitempty" yaml:"FileLocation,omitempty"`
	// Mount option as configured on the NFS Server. For example:nolock.
	MountOptions *string `json:"MountOptions,omitempty" yaml:"MountOptions,omitempty"`
	// Filename of the image in the remote share location. For example:ucs-c220m5-huu-3.1.2c.iso.
	RemoteFile *string `json:"RemoteFile,omitempty" yaml:"RemoteFile,omitempty"`
	// NFS Server Hostname or IP Address. For example:NFS-server-hostname or 10.10.8.7. The remote share server should have network connectivity with the CIMC of the selected devices for a successful upgrade.
	RemoteIp *string `json:"RemoteIp,omitempty" yaml:"RemoteIp,omitempty"`
	// Directory where the image is stored. For example:/share/subfolder.
	RemoteShare *string `json:"RemoteShare,omitempty" yaml:"RemoteShare,omitempty"`
}

// NewFirmwareNfsServer instantiates a new FirmwareNfsServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareNfsServer(classId string, objectType string) *FirmwareNfsServer {
	this := FirmwareNfsServer{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFirmwareNfsServerWithDefaults instantiates a new FirmwareNfsServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareNfsServerWithDefaults() *FirmwareNfsServer {
	this := FirmwareNfsServer{}
	var classId string = "firmware.NfsServer"
	this.ClassId = classId
	var objectType string = "firmware.NfsServer"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareNfsServer) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareNfsServer) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareNfsServer) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareNfsServer) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareNfsServer) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareNfsServer) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFileLocation returns the FileLocation field value if set, zero value otherwise.
func (o *FirmwareNfsServer) GetFileLocation() string {
	if o == nil || o.FileLocation == nil {
		var ret string
		return ret
	}
	return *o.FileLocation
}

// GetFileLocationOk returns a tuple with the FileLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareNfsServer) GetFileLocationOk() (*string, bool) {
	if o == nil || o.FileLocation == nil {
		return nil, false
	}
	return o.FileLocation, true
}

// HasFileLocation returns a boolean if a field has been set.
func (o *FirmwareNfsServer) HasFileLocation() bool {
	if o != nil && o.FileLocation != nil {
		return true
	}

	return false
}

// SetFileLocation gets a reference to the given string and assigns it to the FileLocation field.
func (o *FirmwareNfsServer) SetFileLocation(v string) {
	o.FileLocation = &v
}

// GetMountOptions returns the MountOptions field value if set, zero value otherwise.
func (o *FirmwareNfsServer) GetMountOptions() string {
	if o == nil || o.MountOptions == nil {
		var ret string
		return ret
	}
	return *o.MountOptions
}

// GetMountOptionsOk returns a tuple with the MountOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareNfsServer) GetMountOptionsOk() (*string, bool) {
	if o == nil || o.MountOptions == nil {
		return nil, false
	}
	return o.MountOptions, true
}

// HasMountOptions returns a boolean if a field has been set.
func (o *FirmwareNfsServer) HasMountOptions() bool {
	if o != nil && o.MountOptions != nil {
		return true
	}

	return false
}

// SetMountOptions gets a reference to the given string and assigns it to the MountOptions field.
func (o *FirmwareNfsServer) SetMountOptions(v string) {
	o.MountOptions = &v
}

// GetRemoteFile returns the RemoteFile field value if set, zero value otherwise.
func (o *FirmwareNfsServer) GetRemoteFile() string {
	if o == nil || o.RemoteFile == nil {
		var ret string
		return ret
	}
	return *o.RemoteFile
}

// GetRemoteFileOk returns a tuple with the RemoteFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareNfsServer) GetRemoteFileOk() (*string, bool) {
	if o == nil || o.RemoteFile == nil {
		return nil, false
	}
	return o.RemoteFile, true
}

// HasRemoteFile returns a boolean if a field has been set.
func (o *FirmwareNfsServer) HasRemoteFile() bool {
	if o != nil && o.RemoteFile != nil {
		return true
	}

	return false
}

// SetRemoteFile gets a reference to the given string and assigns it to the RemoteFile field.
func (o *FirmwareNfsServer) SetRemoteFile(v string) {
	o.RemoteFile = &v
}

// GetRemoteIp returns the RemoteIp field value if set, zero value otherwise.
func (o *FirmwareNfsServer) GetRemoteIp() string {
	if o == nil || o.RemoteIp == nil {
		var ret string
		return ret
	}
	return *o.RemoteIp
}

// GetRemoteIpOk returns a tuple with the RemoteIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareNfsServer) GetRemoteIpOk() (*string, bool) {
	if o == nil || o.RemoteIp == nil {
		return nil, false
	}
	return o.RemoteIp, true
}

// HasRemoteIp returns a boolean if a field has been set.
func (o *FirmwareNfsServer) HasRemoteIp() bool {
	if o != nil && o.RemoteIp != nil {
		return true
	}

	return false
}

// SetRemoteIp gets a reference to the given string and assigns it to the RemoteIp field.
func (o *FirmwareNfsServer) SetRemoteIp(v string) {
	o.RemoteIp = &v
}

// GetRemoteShare returns the RemoteShare field value if set, zero value otherwise.
func (o *FirmwareNfsServer) GetRemoteShare() string {
	if o == nil || o.RemoteShare == nil {
		var ret string
		return ret
	}
	return *o.RemoteShare
}

// GetRemoteShareOk returns a tuple with the RemoteShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareNfsServer) GetRemoteShareOk() (*string, bool) {
	if o == nil || o.RemoteShare == nil {
		return nil, false
	}
	return o.RemoteShare, true
}

// HasRemoteShare returns a boolean if a field has been set.
func (o *FirmwareNfsServer) HasRemoteShare() bool {
	if o != nil && o.RemoteShare != nil {
		return true
	}

	return false
}

// SetRemoteShare gets a reference to the given string and assigns it to the RemoteShare field.
func (o *FirmwareNfsServer) SetRemoteShare(v string) {
	o.RemoteShare = &v
}

func (o FirmwareNfsServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.FileLocation != nil {
		toSerialize["FileLocation"] = o.FileLocation
	}
	if o.MountOptions != nil {
		toSerialize["MountOptions"] = o.MountOptions
	}
	if o.RemoteFile != nil {
		toSerialize["RemoteFile"] = o.RemoteFile
	}
	if o.RemoteIp != nil {
		toSerialize["RemoteIp"] = o.RemoteIp
	}
	if o.RemoteShare != nil {
		toSerialize["RemoteShare"] = o.RemoteShare
	}
	return json.Marshal(toSerialize)
}

type NullableFirmwareNfsServer struct {
	value *FirmwareNfsServer
	isSet bool
}

func (v NullableFirmwareNfsServer) Get() *FirmwareNfsServer {
	return v.value
}

func (v *NullableFirmwareNfsServer) Set(val *FirmwareNfsServer) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareNfsServer) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareNfsServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareNfsServer(val *FirmwareNfsServer) *NullableFirmwareNfsServer {
	return &NullableFirmwareNfsServer{value: val, isSet: true}
}

func (v NullableFirmwareNfsServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareNfsServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
