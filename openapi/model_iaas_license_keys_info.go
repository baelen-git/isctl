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

// IaasLicenseKeysInfo License list with the utilization info for UCSD.
type IaasLicenseKeysInfo struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
	// Number of licenses available for the UCSD PID (Product ID).
	Count *int64 `json:"Count,omitempty" yaml:"Count,omitempty"`
	// Expiration date for the license.
	ExpirationDate *string `json:"ExpirationDate,omitempty" yaml:"ExpirationDate,omitempty"`
	// UCS Director Unique license ID.
	LicenseId *string `json:"LicenseId,omitempty" yaml:"LicenseId,omitempty"`
	// PID (Product ID) for UCSD License.
	Pid *string `json:"Pid,omitempty" yaml:"Pid,omitempty"`
}

// NewIaasLicenseKeysInfo instantiates a new IaasLicenseKeysInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIaasLicenseKeysInfo() *IaasLicenseKeysInfo {
	this := IaasLicenseKeysInfo{}
	return &this
}

// NewIaasLicenseKeysInfoWithDefaults instantiates a new IaasLicenseKeysInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIaasLicenseKeysInfoWithDefaults() *IaasLicenseKeysInfo {
	this := IaasLicenseKeysInfo{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *IaasLicenseKeysInfo) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasLicenseKeysInfo) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *IaasLicenseKeysInfo) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *IaasLicenseKeysInfo) SetCount(v int64) {
	o.Count = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *IaasLicenseKeysInfo) GetExpirationDate() string {
	if o == nil || o.ExpirationDate == nil {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasLicenseKeysInfo) GetExpirationDateOk() (*string, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *IaasLicenseKeysInfo) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate != nil {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *IaasLicenseKeysInfo) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetLicenseId returns the LicenseId field value if set, zero value otherwise.
func (o *IaasLicenseKeysInfo) GetLicenseId() string {
	if o == nil || o.LicenseId == nil {
		var ret string
		return ret
	}
	return *o.LicenseId
}

// GetLicenseIdOk returns a tuple with the LicenseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasLicenseKeysInfo) GetLicenseIdOk() (*string, bool) {
	if o == nil || o.LicenseId == nil {
		return nil, false
	}
	return o.LicenseId, true
}

// HasLicenseId returns a boolean if a field has been set.
func (o *IaasLicenseKeysInfo) HasLicenseId() bool {
	if o != nil && o.LicenseId != nil {
		return true
	}

	return false
}

// SetLicenseId gets a reference to the given string and assigns it to the LicenseId field.
func (o *IaasLicenseKeysInfo) SetLicenseId(v string) {
	o.LicenseId = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *IaasLicenseKeysInfo) GetPid() string {
	if o == nil || o.Pid == nil {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IaasLicenseKeysInfo) GetPidOk() (*string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *IaasLicenseKeysInfo) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *IaasLicenseKeysInfo) SetPid(v string) {
	o.Pid = &v
}

func (o IaasLicenseKeysInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.Count != nil {
		toSerialize["Count"] = o.Count
	}
	if o.ExpirationDate != nil {
		toSerialize["ExpirationDate"] = o.ExpirationDate
	}
	if o.LicenseId != nil {
		toSerialize["LicenseId"] = o.LicenseId
	}
	if o.Pid != nil {
		toSerialize["Pid"] = o.Pid
	}
	return json.Marshal(toSerialize)
}

type NullableIaasLicenseKeysInfo struct {
	value *IaasLicenseKeysInfo
	isSet bool
}

func (v NullableIaasLicenseKeysInfo) Get() *IaasLicenseKeysInfo {
	return v.value
}

func (v *NullableIaasLicenseKeysInfo) Set(val *IaasLicenseKeysInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIaasLicenseKeysInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIaasLicenseKeysInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIaasLicenseKeysInfo(val *IaasLicenseKeysInfo) *NullableIaasLicenseKeysInfo {
	return &NullableIaasLicenseKeysInfo{value: val, isSet: true}
}

func (v NullableIaasLicenseKeysInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIaasLicenseKeysInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
