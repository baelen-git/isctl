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

// UcsdUcsdRestoreParameters Restore Configuration Parameters for UCS Director restore workflow.
type UcsdUcsdRestoreParameters struct {
	RecoveryConfigParams `yaml:"RecoveryConfigParams,inline"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty" yaml:"IsPasswordSet,omitempty"`
	// The complete location of the path on the server. The location should be specified in the following format- hostname-or-ipv4address<:port>/absolute-file-path.
	Location *string `json:"Location,omitempty" yaml:"Location,omitempty"`
	// The password of the target backup server. Only required if the target server is accessed using SFTP or SCP protocol.
	Password *string `json:"Password,omitempty" yaml:"Password,omitempty"`
	// The protocol used to backup the UCS Director.
	Protocol *string `json:"Protocol,omitempty" yaml:"Protocol,omitempty"`
	// Decides whether UCS Director property files should also be restored.
	RestoreConfigurationFiles *bool `json:"RestoreConfigurationFiles,omitempty" yaml:"RestoreConfigurationFiles,omitempty"`
	// Decides whether license should also be restored.
	RestoreLicense *bool `json:"RestoreLicense,omitempty" yaml:"RestoreLicense,omitempty"`
	// The username of the target backup server. Only required if the target server is accessed using SFTP or SCP protocol.
	Username *string `json:"Username,omitempty" yaml:"Username,omitempty"`
}

// NewUcsdUcsdRestoreParameters instantiates a new UcsdUcsdRestoreParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUcsdUcsdRestoreParameters() *UcsdUcsdRestoreParameters {
	this := UcsdUcsdRestoreParameters{}
	return &this
}

// NewUcsdUcsdRestoreParametersWithDefaults instantiates a new UcsdUcsdRestoreParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUcsdUcsdRestoreParametersWithDefaults() *UcsdUcsdRestoreParameters {
	this := UcsdUcsdRestoreParameters{}
	return &this
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *UcsdUcsdRestoreParameters) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParameters) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *UcsdUcsdRestoreParameters) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *UcsdUcsdRestoreParameters) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *UcsdUcsdRestoreParameters) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParameters) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *UcsdUcsdRestoreParameters) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *UcsdUcsdRestoreParameters) SetLocation(v string) {
	o.Location = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UcsdUcsdRestoreParameters) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParameters) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UcsdUcsdRestoreParameters) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UcsdUcsdRestoreParameters) SetPassword(v string) {
	o.Password = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *UcsdUcsdRestoreParameters) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParameters) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *UcsdUcsdRestoreParameters) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *UcsdUcsdRestoreParameters) SetProtocol(v string) {
	o.Protocol = &v
}

// GetRestoreConfigurationFiles returns the RestoreConfigurationFiles field value if set, zero value otherwise.
func (o *UcsdUcsdRestoreParameters) GetRestoreConfigurationFiles() bool {
	if o == nil || o.RestoreConfigurationFiles == nil {
		var ret bool
		return ret
	}
	return *o.RestoreConfigurationFiles
}

// GetRestoreConfigurationFilesOk returns a tuple with the RestoreConfigurationFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParameters) GetRestoreConfigurationFilesOk() (*bool, bool) {
	if o == nil || o.RestoreConfigurationFiles == nil {
		return nil, false
	}
	return o.RestoreConfigurationFiles, true
}

// HasRestoreConfigurationFiles returns a boolean if a field has been set.
func (o *UcsdUcsdRestoreParameters) HasRestoreConfigurationFiles() bool {
	if o != nil && o.RestoreConfigurationFiles != nil {
		return true
	}

	return false
}

// SetRestoreConfigurationFiles gets a reference to the given bool and assigns it to the RestoreConfigurationFiles field.
func (o *UcsdUcsdRestoreParameters) SetRestoreConfigurationFiles(v bool) {
	o.RestoreConfigurationFiles = &v
}

// GetRestoreLicense returns the RestoreLicense field value if set, zero value otherwise.
func (o *UcsdUcsdRestoreParameters) GetRestoreLicense() bool {
	if o == nil || o.RestoreLicense == nil {
		var ret bool
		return ret
	}
	return *o.RestoreLicense
}

// GetRestoreLicenseOk returns a tuple with the RestoreLicense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParameters) GetRestoreLicenseOk() (*bool, bool) {
	if o == nil || o.RestoreLicense == nil {
		return nil, false
	}
	return o.RestoreLicense, true
}

// HasRestoreLicense returns a boolean if a field has been set.
func (o *UcsdUcsdRestoreParameters) HasRestoreLicense() bool {
	if o != nil && o.RestoreLicense != nil {
		return true
	}

	return false
}

// SetRestoreLicense gets a reference to the given bool and assigns it to the RestoreLicense field.
func (o *UcsdUcsdRestoreParameters) SetRestoreLicense(v bool) {
	o.RestoreLicense = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UcsdUcsdRestoreParameters) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdUcsdRestoreParameters) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UcsdUcsdRestoreParameters) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UcsdUcsdRestoreParameters) SetUsername(v string) {
	o.Username = &v
}

func (o UcsdUcsdRestoreParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedRecoveryConfigParams, errRecoveryConfigParams := json.Marshal(o.RecoveryConfigParams)
	if errRecoveryConfigParams != nil {
		return []byte{}, errRecoveryConfigParams
	}
	errRecoveryConfigParams = json.Unmarshal([]byte(serializedRecoveryConfigParams), &toSerialize)
	if errRecoveryConfigParams != nil {
		return []byte{}, errRecoveryConfigParams
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.Location != nil {
		toSerialize["Location"] = o.Location
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	if o.RestoreConfigurationFiles != nil {
		toSerialize["RestoreConfigurationFiles"] = o.RestoreConfigurationFiles
	}
	if o.RestoreLicense != nil {
		toSerialize["RestoreLicense"] = o.RestoreLicense
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableUcsdUcsdRestoreParameters struct {
	value *UcsdUcsdRestoreParameters
	isSet bool
}

func (v NullableUcsdUcsdRestoreParameters) Get() *UcsdUcsdRestoreParameters {
	return v.value
}

func (v *NullableUcsdUcsdRestoreParameters) Set(val *UcsdUcsdRestoreParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableUcsdUcsdRestoreParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableUcsdUcsdRestoreParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUcsdUcsdRestoreParameters(val *UcsdUcsdRestoreParameters) *NullableUcsdUcsdRestoreParameters {
	return &NullableUcsdUcsdRestoreParameters{value: val, isSet: true}
}

func (v NullableUcsdUcsdRestoreParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUcsdUcsdRestoreParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
