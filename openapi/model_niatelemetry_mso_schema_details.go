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

// NiatelemetryMsoSchemaDetails Details of schema in Multi-Site Orchestrator.
type NiatelemetryMsoSchemaDetails struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Number of policy objects per schema.
	NumberOfPolicyObjectsPerSchema *int64 `json:"NumberOfPolicyObjectsPerSchema,omitempty" yaml:"NumberOfPolicyObjectsPerSchema,omitempty"`
	// Number of tenants assigned per schema in Multi-Site Orchestrator.
	NumberOfTemplatesPerSchema *int64 `json:"NumberOfTemplatesPerSchema,omitempty" yaml:"NumberOfTemplatesPerSchema,omitempty"`
	// Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected.
	RecordType *string `json:"RecordType,omitempty" yaml:"RecordType,omitempty"`
	// Schema ID in Multi-Site Orchestrator.
	SchemaId *string `json:"SchemaId,omitempty" yaml:"SchemaId,omitempty"`
	// Schema name in Multi-Site Orchestrator.
	SchemaName       *string                              `json:"SchemaName,omitempty" yaml:"SchemaName,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewNiatelemetryMsoSchemaDetails instantiates a new NiatelemetryMsoSchemaDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryMsoSchemaDetails(classId string, objectType string) *NiatelemetryMsoSchemaDetails {
	this := NiatelemetryMsoSchemaDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryMsoSchemaDetailsWithDefaults instantiates a new NiatelemetryMsoSchemaDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryMsoSchemaDetailsWithDefaults() *NiatelemetryMsoSchemaDetails {
	this := NiatelemetryMsoSchemaDetails{}
	var classId string = "niatelemetry.MsoSchemaDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.MsoSchemaDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryMsoSchemaDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSchemaDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryMsoSchemaDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryMsoSchemaDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSchemaDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryMsoSchemaDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetNumberOfPolicyObjectsPerSchema returns the NumberOfPolicyObjectsPerSchema field value if set, zero value otherwise.
func (o *NiatelemetryMsoSchemaDetails) GetNumberOfPolicyObjectsPerSchema() int64 {
	if o == nil || o.NumberOfPolicyObjectsPerSchema == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfPolicyObjectsPerSchema
}

// GetNumberOfPolicyObjectsPerSchemaOk returns a tuple with the NumberOfPolicyObjectsPerSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSchemaDetails) GetNumberOfPolicyObjectsPerSchemaOk() (*int64, bool) {
	if o == nil || o.NumberOfPolicyObjectsPerSchema == nil {
		return nil, false
	}
	return o.NumberOfPolicyObjectsPerSchema, true
}

// HasNumberOfPolicyObjectsPerSchema returns a boolean if a field has been set.
func (o *NiatelemetryMsoSchemaDetails) HasNumberOfPolicyObjectsPerSchema() bool {
	if o != nil && o.NumberOfPolicyObjectsPerSchema != nil {
		return true
	}

	return false
}

// SetNumberOfPolicyObjectsPerSchema gets a reference to the given int64 and assigns it to the NumberOfPolicyObjectsPerSchema field.
func (o *NiatelemetryMsoSchemaDetails) SetNumberOfPolicyObjectsPerSchema(v int64) {
	o.NumberOfPolicyObjectsPerSchema = &v
}

// GetNumberOfTemplatesPerSchema returns the NumberOfTemplatesPerSchema field value if set, zero value otherwise.
func (o *NiatelemetryMsoSchemaDetails) GetNumberOfTemplatesPerSchema() int64 {
	if o == nil || o.NumberOfTemplatesPerSchema == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfTemplatesPerSchema
}

// GetNumberOfTemplatesPerSchemaOk returns a tuple with the NumberOfTemplatesPerSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSchemaDetails) GetNumberOfTemplatesPerSchemaOk() (*int64, bool) {
	if o == nil || o.NumberOfTemplatesPerSchema == nil {
		return nil, false
	}
	return o.NumberOfTemplatesPerSchema, true
}

// HasNumberOfTemplatesPerSchema returns a boolean if a field has been set.
func (o *NiatelemetryMsoSchemaDetails) HasNumberOfTemplatesPerSchema() bool {
	if o != nil && o.NumberOfTemplatesPerSchema != nil {
		return true
	}

	return false
}

// SetNumberOfTemplatesPerSchema gets a reference to the given int64 and assigns it to the NumberOfTemplatesPerSchema field.
func (o *NiatelemetryMsoSchemaDetails) SetNumberOfTemplatesPerSchema(v int64) {
	o.NumberOfTemplatesPerSchema = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *NiatelemetryMsoSchemaDetails) GetRecordType() string {
	if o == nil || o.RecordType == nil {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSchemaDetails) GetRecordTypeOk() (*string, bool) {
	if o == nil || o.RecordType == nil {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *NiatelemetryMsoSchemaDetails) HasRecordType() bool {
	if o != nil && o.RecordType != nil {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *NiatelemetryMsoSchemaDetails) SetRecordType(v string) {
	o.RecordType = &v
}

// GetSchemaId returns the SchemaId field value if set, zero value otherwise.
func (o *NiatelemetryMsoSchemaDetails) GetSchemaId() string {
	if o == nil || o.SchemaId == nil {
		var ret string
		return ret
	}
	return *o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSchemaDetails) GetSchemaIdOk() (*string, bool) {
	if o == nil || o.SchemaId == nil {
		return nil, false
	}
	return o.SchemaId, true
}

// HasSchemaId returns a boolean if a field has been set.
func (o *NiatelemetryMsoSchemaDetails) HasSchemaId() bool {
	if o != nil && o.SchemaId != nil {
		return true
	}

	return false
}

// SetSchemaId gets a reference to the given string and assigns it to the SchemaId field.
func (o *NiatelemetryMsoSchemaDetails) SetSchemaId(v string) {
	o.SchemaId = &v
}

// GetSchemaName returns the SchemaName field value if set, zero value otherwise.
func (o *NiatelemetryMsoSchemaDetails) GetSchemaName() string {
	if o == nil || o.SchemaName == nil {
		var ret string
		return ret
	}
	return *o.SchemaName
}

// GetSchemaNameOk returns a tuple with the SchemaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSchemaDetails) GetSchemaNameOk() (*string, bool) {
	if o == nil || o.SchemaName == nil {
		return nil, false
	}
	return o.SchemaName, true
}

// HasSchemaName returns a boolean if a field has been set.
func (o *NiatelemetryMsoSchemaDetails) HasSchemaName() bool {
	if o != nil && o.SchemaName != nil {
		return true
	}

	return false
}

// SetSchemaName gets a reference to the given string and assigns it to the SchemaName field.
func (o *NiatelemetryMsoSchemaDetails) SetSchemaName(v string) {
	o.SchemaName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryMsoSchemaDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoSchemaDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryMsoSchemaDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryMsoSchemaDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryMsoSchemaDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.NumberOfPolicyObjectsPerSchema != nil {
		toSerialize["NumberOfPolicyObjectsPerSchema"] = o.NumberOfPolicyObjectsPerSchema
	}
	if o.NumberOfTemplatesPerSchema != nil {
		toSerialize["NumberOfTemplatesPerSchema"] = o.NumberOfTemplatesPerSchema
	}
	if o.RecordType != nil {
		toSerialize["RecordType"] = o.RecordType
	}
	if o.SchemaId != nil {
		toSerialize["SchemaId"] = o.SchemaId
	}
	if o.SchemaName != nil {
		toSerialize["SchemaName"] = o.SchemaName
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableNiatelemetryMsoSchemaDetails struct {
	value *NiatelemetryMsoSchemaDetails
	isSet bool
}

func (v NullableNiatelemetryMsoSchemaDetails) Get() *NiatelemetryMsoSchemaDetails {
	return v.value
}

func (v *NullableNiatelemetryMsoSchemaDetails) Set(val *NiatelemetryMsoSchemaDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryMsoSchemaDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryMsoSchemaDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryMsoSchemaDetails(val *NiatelemetryMsoSchemaDetails) *NullableNiatelemetryMsoSchemaDetails {
	return &NullableNiatelemetryMsoSchemaDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryMsoSchemaDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryMsoSchemaDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
