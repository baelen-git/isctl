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

// KubernetesPodStatusAllOf Definition of the list of properties defined in 'kubernetes.PodStatus', excluding properties defined in parent classes.
type KubernetesPodStatusAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// The IP of the host that the pod is running on.
	HostIp *string `json:"HostIp,omitempty" yaml:"HostIp,omitempty"`
	// The phase of a Pod is a simple, high-level summary of where the Pod is in its lifecycle.
	Phase *string `json:"Phase,omitempty" yaml:"Phase,omitempty"`
	// The IP of the pod. The IP is allocated by the Pod CIDR from the kubernetes cluster itself.
	PodIp *string `json:"PodIp,omitempty" yaml:"PodIp,omitempty"`
	// The Quality of Service (QOS) classification assigned to the pod based on resource requirements.
	QosClass *string `json:"QosClass,omitempty" yaml:"QosClass,omitempty"`
	// The time that the pod was started.
	StartTime *string `json:"StartTime,omitempty" yaml:"StartTime,omitempty"`
}

// NewKubernetesPodStatusAllOf instantiates a new KubernetesPodStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesPodStatusAllOf(classId string, objectType string) *KubernetesPodStatusAllOf {
	this := KubernetesPodStatusAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesPodStatusAllOfWithDefaults instantiates a new KubernetesPodStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesPodStatusAllOfWithDefaults() *KubernetesPodStatusAllOf {
	this := KubernetesPodStatusAllOf{}
	var classId string = "kubernetes.PodStatus"
	this.ClassId = classId
	var objectType string = "kubernetes.PodStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesPodStatusAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesPodStatusAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesPodStatusAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesPodStatusAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesPodStatusAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesPodStatusAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHostIp returns the HostIp field value if set, zero value otherwise.
func (o *KubernetesPodStatusAllOf) GetHostIp() string {
	if o == nil || o.HostIp == nil {
		var ret string
		return ret
	}
	return *o.HostIp
}

// GetHostIpOk returns a tuple with the HostIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesPodStatusAllOf) GetHostIpOk() (*string, bool) {
	if o == nil || o.HostIp == nil {
		return nil, false
	}
	return o.HostIp, true
}

// HasHostIp returns a boolean if a field has been set.
func (o *KubernetesPodStatusAllOf) HasHostIp() bool {
	if o != nil && o.HostIp != nil {
		return true
	}

	return false
}

// SetHostIp gets a reference to the given string and assigns it to the HostIp field.
func (o *KubernetesPodStatusAllOf) SetHostIp(v string) {
	o.HostIp = &v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *KubernetesPodStatusAllOf) GetPhase() string {
	if o == nil || o.Phase == nil {
		var ret string
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesPodStatusAllOf) GetPhaseOk() (*string, bool) {
	if o == nil || o.Phase == nil {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *KubernetesPodStatusAllOf) HasPhase() bool {
	if o != nil && o.Phase != nil {
		return true
	}

	return false
}

// SetPhase gets a reference to the given string and assigns it to the Phase field.
func (o *KubernetesPodStatusAllOf) SetPhase(v string) {
	o.Phase = &v
}

// GetPodIp returns the PodIp field value if set, zero value otherwise.
func (o *KubernetesPodStatusAllOf) GetPodIp() string {
	if o == nil || o.PodIp == nil {
		var ret string
		return ret
	}
	return *o.PodIp
}

// GetPodIpOk returns a tuple with the PodIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesPodStatusAllOf) GetPodIpOk() (*string, bool) {
	if o == nil || o.PodIp == nil {
		return nil, false
	}
	return o.PodIp, true
}

// HasPodIp returns a boolean if a field has been set.
func (o *KubernetesPodStatusAllOf) HasPodIp() bool {
	if o != nil && o.PodIp != nil {
		return true
	}

	return false
}

// SetPodIp gets a reference to the given string and assigns it to the PodIp field.
func (o *KubernetesPodStatusAllOf) SetPodIp(v string) {
	o.PodIp = &v
}

// GetQosClass returns the QosClass field value if set, zero value otherwise.
func (o *KubernetesPodStatusAllOf) GetQosClass() string {
	if o == nil || o.QosClass == nil {
		var ret string
		return ret
	}
	return *o.QosClass
}

// GetQosClassOk returns a tuple with the QosClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesPodStatusAllOf) GetQosClassOk() (*string, bool) {
	if o == nil || o.QosClass == nil {
		return nil, false
	}
	return o.QosClass, true
}

// HasQosClass returns a boolean if a field has been set.
func (o *KubernetesPodStatusAllOf) HasQosClass() bool {
	if o != nil && o.QosClass != nil {
		return true
	}

	return false
}

// SetQosClass gets a reference to the given string and assigns it to the QosClass field.
func (o *KubernetesPodStatusAllOf) SetQosClass(v string) {
	o.QosClass = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *KubernetesPodStatusAllOf) GetStartTime() string {
	if o == nil || o.StartTime == nil {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesPodStatusAllOf) GetStartTimeOk() (*string, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *KubernetesPodStatusAllOf) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *KubernetesPodStatusAllOf) SetStartTime(v string) {
	o.StartTime = &v
}

func (o KubernetesPodStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.HostIp != nil {
		toSerialize["HostIp"] = o.HostIp
	}
	if o.Phase != nil {
		toSerialize["Phase"] = o.Phase
	}
	if o.PodIp != nil {
		toSerialize["PodIp"] = o.PodIp
	}
	if o.QosClass != nil {
		toSerialize["QosClass"] = o.QosClass
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	return json.Marshal(toSerialize)
}

type NullableKubernetesPodStatusAllOf struct {
	value *KubernetesPodStatusAllOf
	isSet bool
}

func (v NullableKubernetesPodStatusAllOf) Get() *KubernetesPodStatusAllOf {
	return v.value
}

func (v *NullableKubernetesPodStatusAllOf) Set(val *KubernetesPodStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesPodStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesPodStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesPodStatusAllOf(val *KubernetesPodStatusAllOf) *NullableKubernetesPodStatusAllOf {
	return &NullableKubernetesPodStatusAllOf{value: val, isSet: true}
}

func (v NullableKubernetesPodStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesPodStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
