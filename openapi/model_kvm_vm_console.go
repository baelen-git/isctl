/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-04-28T13:03:38Z.
 *
 * API version: 1.0.9-4267
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"reflect"
	"strings"
)

// KvmVmConsole API to launch the virtual machine console.
type KvmVmConsole struct {
	TunnelingTunnel
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The URL of the KVM MUX to connect to.
	KvmMuxUrl            *string                                  `json:"KvmMuxUrl,omitempty"`
	VirtualMachine       *HyperflexHxapVirtualMachineRelationship `json:"VirtualMachine,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KvmVmConsole KvmVmConsole

// NewKvmVmConsole instantiates a new KvmVmConsole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvmVmConsole(classId string, objectType string) *KvmVmConsole {
	this := KvmVmConsole{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKvmVmConsoleWithDefaults instantiates a new KvmVmConsole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvmVmConsoleWithDefaults() *KvmVmConsole {
	this := KvmVmConsole{}
	var classId string = "kvm.VmConsole"
	this.ClassId = classId
	var objectType string = "kvm.VmConsole"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KvmVmConsole) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KvmVmConsole) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KvmVmConsole) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KvmVmConsole) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KvmVmConsole) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KvmVmConsole) SetObjectType(v string) {
	o.ObjectType = v
}

// GetKvmMuxUrl returns the KvmMuxUrl field value if set, zero value otherwise.
func (o *KvmVmConsole) GetKvmMuxUrl() string {
	if o == nil || o.KvmMuxUrl == nil {
		var ret string
		return ret
	}
	return *o.KvmMuxUrl
}

// GetKvmMuxUrlOk returns a tuple with the KvmMuxUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmVmConsole) GetKvmMuxUrlOk() (*string, bool) {
	if o == nil || o.KvmMuxUrl == nil {
		return nil, false
	}
	return o.KvmMuxUrl, true
}

// HasKvmMuxUrl returns a boolean if a field has been set.
func (o *KvmVmConsole) HasKvmMuxUrl() bool {
	if o != nil && o.KvmMuxUrl != nil {
		return true
	}

	return false
}

// SetKvmMuxUrl gets a reference to the given string and assigns it to the KvmMuxUrl field.
func (o *KvmVmConsole) SetKvmMuxUrl(v string) {
	o.KvmMuxUrl = &v
}

// GetVirtualMachine returns the VirtualMachine field value if set, zero value otherwise.
func (o *KvmVmConsole) GetVirtualMachine() HyperflexHxapVirtualMachineRelationship {
	if o == nil || o.VirtualMachine == nil {
		var ret HyperflexHxapVirtualMachineRelationship
		return ret
	}
	return *o.VirtualMachine
}

// GetVirtualMachineOk returns a tuple with the VirtualMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KvmVmConsole) GetVirtualMachineOk() (*HyperflexHxapVirtualMachineRelationship, bool) {
	if o == nil || o.VirtualMachine == nil {
		return nil, false
	}
	return o.VirtualMachine, true
}

// HasVirtualMachine returns a boolean if a field has been set.
func (o *KvmVmConsole) HasVirtualMachine() bool {
	if o != nil && o.VirtualMachine != nil {
		return true
	}

	return false
}

// SetVirtualMachine gets a reference to the given HyperflexHxapVirtualMachineRelationship and assigns it to the VirtualMachine field.
func (o *KvmVmConsole) SetVirtualMachine(v HyperflexHxapVirtualMachineRelationship) {
	o.VirtualMachine = &v
}

func (o KvmVmConsole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedTunnelingTunnel, errTunnelingTunnel := json.Marshal(o.TunnelingTunnel)
	if errTunnelingTunnel != nil {
		return []byte{}, errTunnelingTunnel
	}
	errTunnelingTunnel = json.Unmarshal([]byte(serializedTunnelingTunnel), &toSerialize)
	if errTunnelingTunnel != nil {
		return []byte{}, errTunnelingTunnel
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.KvmMuxUrl != nil {
		toSerialize["KvmMuxUrl"] = o.KvmMuxUrl
	}
	if o.VirtualMachine != nil {
		toSerialize["VirtualMachine"] = o.VirtualMachine
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KvmVmConsole) UnmarshalJSON(bytes []byte) (err error) {
	type KvmVmConsoleWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The URL of the KVM MUX to connect to.
		KvmMuxUrl      *string                                  `json:"KvmMuxUrl,omitempty"`
		VirtualMachine *HyperflexHxapVirtualMachineRelationship `json:"VirtualMachine,omitempty"`
	}

	varKvmVmConsoleWithoutEmbeddedStruct := KvmVmConsoleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKvmVmConsoleWithoutEmbeddedStruct)
	if err == nil {
		varKvmVmConsole := _KvmVmConsole{}
		varKvmVmConsole.ClassId = varKvmVmConsoleWithoutEmbeddedStruct.ClassId
		varKvmVmConsole.ObjectType = varKvmVmConsoleWithoutEmbeddedStruct.ObjectType
		varKvmVmConsole.KvmMuxUrl = varKvmVmConsoleWithoutEmbeddedStruct.KvmMuxUrl
		varKvmVmConsole.VirtualMachine = varKvmVmConsoleWithoutEmbeddedStruct.VirtualMachine
		*o = KvmVmConsole(varKvmVmConsole)
	} else {
		return err
	}

	varKvmVmConsole := _KvmVmConsole{}

	err = json.Unmarshal(bytes, &varKvmVmConsole)
	if err == nil {
		o.TunnelingTunnel = varKvmVmConsole.TunnelingTunnel
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "KvmMuxUrl")
		delete(additionalProperties, "VirtualMachine")

		// remove fields from embedded structs
		reflectTunnelingTunnel := reflect.ValueOf(o.TunnelingTunnel)
		for i := 0; i < reflectTunnelingTunnel.Type().NumField(); i++ {
			t := reflectTunnelingTunnel.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKvmVmConsole struct {
	value *KvmVmConsole
	isSet bool
}

func (v NullableKvmVmConsole) Get() *KvmVmConsole {
	return v.value
}

func (v *NullableKvmVmConsole) Set(val *KvmVmConsole) {
	v.value = val
	v.isSet = true
}

func (v NullableKvmVmConsole) IsSet() bool {
	return v.isSet
}

func (v *NullableKvmVmConsole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKvmVmConsole(val *KvmVmConsole) *NullableKvmVmConsole {
	return &NullableKvmVmConsole{value: val, isSet: true}
}

func (v NullableKvmVmConsole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKvmVmConsole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
