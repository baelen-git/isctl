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

// ComputePersistentMemoryOperation The operation that can be performed on the Persistent Memory Modules on the servers.
type ComputePersistentMemoryOperation struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
	// Administrative actions that can be performed on the Persistent Memory Modules. * `None` - No action on the selected Persistent Memory Modules. * `SecureErase` - Secure Erase action on the selected Persistent Memory Modules. * `Unlock` - Unlock action on the selected Persistent Memory Modules.
	AdminAction *string `json:"AdminAction,omitempty" yaml:"AdminAction,omitempty"`
	// Indicates whether the value of the 'securePassphrase' property has been set.
	IsSecurePassphraseSet *bool                            `json:"IsSecurePassphraseSet,omitempty" yaml:"IsSecurePassphraseSet,omitempty"`
	Modules               *[]ComputePersistentMemoryModule `json:"Modules,omitempty" yaml:"Modules,omitempty"`
	// Secure passphrase of the Persistent Memory Modules of the server.
	SecurePassphrase *string `json:"SecurePassphrase,omitempty" yaml:"SecurePassphrase,omitempty"`
}

// NewComputePersistentMemoryOperation instantiates a new ComputePersistentMemoryOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputePersistentMemoryOperation() *ComputePersistentMemoryOperation {
	this := ComputePersistentMemoryOperation{}
	var adminAction string = "None"
	this.AdminAction = &adminAction
	return &this
}

// NewComputePersistentMemoryOperationWithDefaults instantiates a new ComputePersistentMemoryOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputePersistentMemoryOperationWithDefaults() *ComputePersistentMemoryOperation {
	this := ComputePersistentMemoryOperation{}
	var adminAction string = "None"
	this.AdminAction = &adminAction
	return &this
}

// GetAdminAction returns the AdminAction field value if set, zero value otherwise.
func (o *ComputePersistentMemoryOperation) GetAdminAction() string {
	if o == nil || o.AdminAction == nil {
		var ret string
		return ret
	}
	return *o.AdminAction
}

// GetAdminActionOk returns a tuple with the AdminAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePersistentMemoryOperation) GetAdminActionOk() (*string, bool) {
	if o == nil || o.AdminAction == nil {
		return nil, false
	}
	return o.AdminAction, true
}

// HasAdminAction returns a boolean if a field has been set.
func (o *ComputePersistentMemoryOperation) HasAdminAction() bool {
	if o != nil && o.AdminAction != nil {
		return true
	}

	return false
}

// SetAdminAction gets a reference to the given string and assigns it to the AdminAction field.
func (o *ComputePersistentMemoryOperation) SetAdminAction(v string) {
	o.AdminAction = &v
}

// GetIsSecurePassphraseSet returns the IsSecurePassphraseSet field value if set, zero value otherwise.
func (o *ComputePersistentMemoryOperation) GetIsSecurePassphraseSet() bool {
	if o == nil || o.IsSecurePassphraseSet == nil {
		var ret bool
		return ret
	}
	return *o.IsSecurePassphraseSet
}

// GetIsSecurePassphraseSetOk returns a tuple with the IsSecurePassphraseSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePersistentMemoryOperation) GetIsSecurePassphraseSetOk() (*bool, bool) {
	if o == nil || o.IsSecurePassphraseSet == nil {
		return nil, false
	}
	return o.IsSecurePassphraseSet, true
}

// HasIsSecurePassphraseSet returns a boolean if a field has been set.
func (o *ComputePersistentMemoryOperation) HasIsSecurePassphraseSet() bool {
	if o != nil && o.IsSecurePassphraseSet != nil {
		return true
	}

	return false
}

// SetIsSecurePassphraseSet gets a reference to the given bool and assigns it to the IsSecurePassphraseSet field.
func (o *ComputePersistentMemoryOperation) SetIsSecurePassphraseSet(v bool) {
	o.IsSecurePassphraseSet = &v
}

// GetModules returns the Modules field value if set, zero value otherwise.
func (o *ComputePersistentMemoryOperation) GetModules() []ComputePersistentMemoryModule {
	if o == nil || o.Modules == nil {
		var ret []ComputePersistentMemoryModule
		return ret
	}
	return *o.Modules
}

// GetModulesOk returns a tuple with the Modules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePersistentMemoryOperation) GetModulesOk() (*[]ComputePersistentMemoryModule, bool) {
	if o == nil || o.Modules == nil {
		return nil, false
	}
	return o.Modules, true
}

// HasModules returns a boolean if a field has been set.
func (o *ComputePersistentMemoryOperation) HasModules() bool {
	if o != nil && o.Modules != nil {
		return true
	}

	return false
}

// SetModules gets a reference to the given []ComputePersistentMemoryModule and assigns it to the Modules field.
func (o *ComputePersistentMemoryOperation) SetModules(v []ComputePersistentMemoryModule) {
	o.Modules = &v
}

// GetSecurePassphrase returns the SecurePassphrase field value if set, zero value otherwise.
func (o *ComputePersistentMemoryOperation) GetSecurePassphrase() string {
	if o == nil || o.SecurePassphrase == nil {
		var ret string
		return ret
	}
	return *o.SecurePassphrase
}

// GetSecurePassphraseOk returns a tuple with the SecurePassphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePersistentMemoryOperation) GetSecurePassphraseOk() (*string, bool) {
	if o == nil || o.SecurePassphrase == nil {
		return nil, false
	}
	return o.SecurePassphrase, true
}

// HasSecurePassphrase returns a boolean if a field has been set.
func (o *ComputePersistentMemoryOperation) HasSecurePassphrase() bool {
	if o != nil && o.SecurePassphrase != nil {
		return true
	}

	return false
}

// SetSecurePassphrase gets a reference to the given string and assigns it to the SecurePassphrase field.
func (o *ComputePersistentMemoryOperation) SetSecurePassphrase(v string) {
	o.SecurePassphrase = &v
}

func (o ComputePersistentMemoryOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if o.AdminAction != nil {
		toSerialize["AdminAction"] = o.AdminAction
	}
	if o.IsSecurePassphraseSet != nil {
		toSerialize["IsSecurePassphraseSet"] = o.IsSecurePassphraseSet
	}
	if o.Modules != nil {
		toSerialize["Modules"] = o.Modules
	}
	if o.SecurePassphrase != nil {
		toSerialize["SecurePassphrase"] = o.SecurePassphrase
	}
	return json.Marshal(toSerialize)
}

type NullableComputePersistentMemoryOperation struct {
	value *ComputePersistentMemoryOperation
	isSet bool
}

func (v NullableComputePersistentMemoryOperation) Get() *ComputePersistentMemoryOperation {
	return v.value
}

func (v *NullableComputePersistentMemoryOperation) Set(val *ComputePersistentMemoryOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableComputePersistentMemoryOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableComputePersistentMemoryOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputePersistentMemoryOperation(val *ComputePersistentMemoryOperation) *NullableComputePersistentMemoryOperation {
	return &NullableComputePersistentMemoryOperation{value: val, isSet: true}
}

func (v NullableComputePersistentMemoryOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputePersistentMemoryOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
