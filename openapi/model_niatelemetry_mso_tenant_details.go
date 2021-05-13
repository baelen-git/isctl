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

// NiatelemetryMsoTenantDetails Details of tenant in Multi-Site Orchestrator.
type NiatelemetryMsoTenantDetails struct {
	MoBaseMo `yaml:"MoBaseMo,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// Number of schemas assigned to each tenant in Multi-Site Orchestrator.
	NumberOfSchemasAssignedPerTenantInMso *int64 `json:"NumberOfSchemasAssignedPerTenantInMso,omitempty" yaml:"NumberOfSchemasAssignedPerTenantInMso,omitempty"`
	// Number of sites each tenant is deployed to.
	SitesEachTenantIsDeployedToInMso *int64 `json:"SitesEachTenantIsDeployedToInMso,omitempty" yaml:"SitesEachTenantIsDeployedToInMso,omitempty"`
	// ID of tenant in Multi-Site Orchestrator.
	TenantId *string `json:"TenantId,omitempty" yaml:"TenantId,omitempty"`
	// Name of the tenant in Multi-Site Orchestrator.
	TenantName       *string                              `json:"TenantName,omitempty" yaml:"TenantName,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty" yaml:"RegisteredDevice,omitempty"`
}

// NewNiatelemetryMsoTenantDetails instantiates a new NiatelemetryMsoTenantDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiatelemetryMsoTenantDetails(classId string, objectType string) *NiatelemetryMsoTenantDetails {
	this := NiatelemetryMsoTenantDetails{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewNiatelemetryMsoTenantDetailsWithDefaults instantiates a new NiatelemetryMsoTenantDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiatelemetryMsoTenantDetailsWithDefaults() *NiatelemetryMsoTenantDetails {
	this := NiatelemetryMsoTenantDetails{}
	var classId string = "niatelemetry.MsoTenantDetails"
	this.ClassId = classId
	var objectType string = "niatelemetry.MsoTenantDetails"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *NiatelemetryMsoTenantDetails) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoTenantDetails) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *NiatelemetryMsoTenantDetails) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *NiatelemetryMsoTenantDetails) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoTenantDetails) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *NiatelemetryMsoTenantDetails) SetObjectType(v string) {
	o.ObjectType = v
}

// GetNumberOfSchemasAssignedPerTenantInMso returns the NumberOfSchemasAssignedPerTenantInMso field value if set, zero value otherwise.
func (o *NiatelemetryMsoTenantDetails) GetNumberOfSchemasAssignedPerTenantInMso() int64 {
	if o == nil || o.NumberOfSchemasAssignedPerTenantInMso == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfSchemasAssignedPerTenantInMso
}

// GetNumberOfSchemasAssignedPerTenantInMsoOk returns a tuple with the NumberOfSchemasAssignedPerTenantInMso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoTenantDetails) GetNumberOfSchemasAssignedPerTenantInMsoOk() (*int64, bool) {
	if o == nil || o.NumberOfSchemasAssignedPerTenantInMso == nil {
		return nil, false
	}
	return o.NumberOfSchemasAssignedPerTenantInMso, true
}

// HasNumberOfSchemasAssignedPerTenantInMso returns a boolean if a field has been set.
func (o *NiatelemetryMsoTenantDetails) HasNumberOfSchemasAssignedPerTenantInMso() bool {
	if o != nil && o.NumberOfSchemasAssignedPerTenantInMso != nil {
		return true
	}

	return false
}

// SetNumberOfSchemasAssignedPerTenantInMso gets a reference to the given int64 and assigns it to the NumberOfSchemasAssignedPerTenantInMso field.
func (o *NiatelemetryMsoTenantDetails) SetNumberOfSchemasAssignedPerTenantInMso(v int64) {
	o.NumberOfSchemasAssignedPerTenantInMso = &v
}

// GetSitesEachTenantIsDeployedToInMso returns the SitesEachTenantIsDeployedToInMso field value if set, zero value otherwise.
func (o *NiatelemetryMsoTenantDetails) GetSitesEachTenantIsDeployedToInMso() int64 {
	if o == nil || o.SitesEachTenantIsDeployedToInMso == nil {
		var ret int64
		return ret
	}
	return *o.SitesEachTenantIsDeployedToInMso
}

// GetSitesEachTenantIsDeployedToInMsoOk returns a tuple with the SitesEachTenantIsDeployedToInMso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoTenantDetails) GetSitesEachTenantIsDeployedToInMsoOk() (*int64, bool) {
	if o == nil || o.SitesEachTenantIsDeployedToInMso == nil {
		return nil, false
	}
	return o.SitesEachTenantIsDeployedToInMso, true
}

// HasSitesEachTenantIsDeployedToInMso returns a boolean if a field has been set.
func (o *NiatelemetryMsoTenantDetails) HasSitesEachTenantIsDeployedToInMso() bool {
	if o != nil && o.SitesEachTenantIsDeployedToInMso != nil {
		return true
	}

	return false
}

// SetSitesEachTenantIsDeployedToInMso gets a reference to the given int64 and assigns it to the SitesEachTenantIsDeployedToInMso field.
func (o *NiatelemetryMsoTenantDetails) SetSitesEachTenantIsDeployedToInMso(v int64) {
	o.SitesEachTenantIsDeployedToInMso = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *NiatelemetryMsoTenantDetails) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoTenantDetails) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *NiatelemetryMsoTenantDetails) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *NiatelemetryMsoTenantDetails) SetTenantId(v string) {
	o.TenantId = &v
}

// GetTenantName returns the TenantName field value if set, zero value otherwise.
func (o *NiatelemetryMsoTenantDetails) GetTenantName() string {
	if o == nil || o.TenantName == nil {
		var ret string
		return ret
	}
	return *o.TenantName
}

// GetTenantNameOk returns a tuple with the TenantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoTenantDetails) GetTenantNameOk() (*string, bool) {
	if o == nil || o.TenantName == nil {
		return nil, false
	}
	return o.TenantName, true
}

// HasTenantName returns a boolean if a field has been set.
func (o *NiatelemetryMsoTenantDetails) HasTenantName() bool {
	if o != nil && o.TenantName != nil {
		return true
	}

	return false
}

// SetTenantName gets a reference to the given string and assigns it to the TenantName field.
func (o *NiatelemetryMsoTenantDetails) SetTenantName(v string) {
	o.TenantName = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *NiatelemetryMsoTenantDetails) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NiatelemetryMsoTenantDetails) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *NiatelemetryMsoTenantDetails) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *NiatelemetryMsoTenantDetails) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o NiatelemetryMsoTenantDetails) MarshalJSON() ([]byte, error) {
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
	if o.NumberOfSchemasAssignedPerTenantInMso != nil {
		toSerialize["NumberOfSchemasAssignedPerTenantInMso"] = o.NumberOfSchemasAssignedPerTenantInMso
	}
	if o.SitesEachTenantIsDeployedToInMso != nil {
		toSerialize["SitesEachTenantIsDeployedToInMso"] = o.SitesEachTenantIsDeployedToInMso
	}
	if o.TenantId != nil {
		toSerialize["TenantId"] = o.TenantId
	}
	if o.TenantName != nil {
		toSerialize["TenantName"] = o.TenantName
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	return json.Marshal(toSerialize)
}

type NullableNiatelemetryMsoTenantDetails struct {
	value *NiatelemetryMsoTenantDetails
	isSet bool
}

func (v NullableNiatelemetryMsoTenantDetails) Get() *NiatelemetryMsoTenantDetails {
	return v.value
}

func (v *NullableNiatelemetryMsoTenantDetails) Set(val *NiatelemetryMsoTenantDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableNiatelemetryMsoTenantDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableNiatelemetryMsoTenantDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiatelemetryMsoTenantDetails(val *NiatelemetryMsoTenantDetails) *NullableNiatelemetryMsoTenantDetails {
	return &NullableNiatelemetryMsoTenantDetails{value: val, isSet: true}
}

func (v NullableNiatelemetryMsoTenantDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiatelemetryMsoTenantDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
