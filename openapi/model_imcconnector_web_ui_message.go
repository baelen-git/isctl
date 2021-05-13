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

// ImcconnectorWebUiMessage Message contains the request to send to the platforms management entities HTTP server. Returns the HTTP response body from the platform, or error if plugin encounters error executing the request.
type ImcconnectorWebUiMessage struct {
	ConnectorAuthMessage `yaml:"ConnectorAuthMessage,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType" yaml:"ObjectType"`
	// The body content of the UI HTTP request to send to the BMC platform.
	WebUiRequest *string `json:"WebUiRequest,omitempty" yaml:"WebUiRequest,omitempty"`
}

// NewImcconnectorWebUiMessage instantiates a new ImcconnectorWebUiMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImcconnectorWebUiMessage(classId string, objectType string) *ImcconnectorWebUiMessage {
	this := ImcconnectorWebUiMessage{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewImcconnectorWebUiMessageWithDefaults instantiates a new ImcconnectorWebUiMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImcconnectorWebUiMessageWithDefaults() *ImcconnectorWebUiMessage {
	this := ImcconnectorWebUiMessage{}
	var classId string = "imcconnector.WebUiMessage"
	this.ClassId = classId
	var objectType string = "imcconnector.WebUiMessage"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ImcconnectorWebUiMessage) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ImcconnectorWebUiMessage) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ImcconnectorWebUiMessage) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ImcconnectorWebUiMessage) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ImcconnectorWebUiMessage) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ImcconnectorWebUiMessage) SetObjectType(v string) {
	o.ObjectType = v
}

// GetWebUiRequest returns the WebUiRequest field value if set, zero value otherwise.
func (o *ImcconnectorWebUiMessage) GetWebUiRequest() string {
	if o == nil || o.WebUiRequest == nil {
		var ret string
		return ret
	}
	return *o.WebUiRequest
}

// GetWebUiRequestOk returns a tuple with the WebUiRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImcconnectorWebUiMessage) GetWebUiRequestOk() (*string, bool) {
	if o == nil || o.WebUiRequest == nil {
		return nil, false
	}
	return o.WebUiRequest, true
}

// HasWebUiRequest returns a boolean if a field has been set.
func (o *ImcconnectorWebUiMessage) HasWebUiRequest() bool {
	if o != nil && o.WebUiRequest != nil {
		return true
	}

	return false
}

// SetWebUiRequest gets a reference to the given string and assigns it to the WebUiRequest field.
func (o *ImcconnectorWebUiMessage) SetWebUiRequest(v string) {
	o.WebUiRequest = &v
}

func (o ImcconnectorWebUiMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectorAuthMessage, errConnectorAuthMessage := json.Marshal(o.ConnectorAuthMessage)
	if errConnectorAuthMessage != nil {
		return []byte{}, errConnectorAuthMessage
	}
	errConnectorAuthMessage = json.Unmarshal([]byte(serializedConnectorAuthMessage), &toSerialize)
	if errConnectorAuthMessage != nil {
		return []byte{}, errConnectorAuthMessage
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.WebUiRequest != nil {
		toSerialize["WebUiRequest"] = o.WebUiRequest
	}
	return json.Marshal(toSerialize)
}

type NullableImcconnectorWebUiMessage struct {
	value *ImcconnectorWebUiMessage
	isSet bool
}

func (v NullableImcconnectorWebUiMessage) Get() *ImcconnectorWebUiMessage {
	return v.value
}

func (v *NullableImcconnectorWebUiMessage) Set(val *ImcconnectorWebUiMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableImcconnectorWebUiMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableImcconnectorWebUiMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImcconnectorWebUiMessage(val *ImcconnectorWebUiMessage) *NullableImcconnectorWebUiMessage {
	return &NullableImcconnectorWebUiMessage{value: val, isSet: true}
}

func (v NullableImcconnectorWebUiMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImcconnectorWebUiMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
