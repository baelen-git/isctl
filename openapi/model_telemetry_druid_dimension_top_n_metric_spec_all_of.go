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

// TelemetryDruidDimensionTopNMetricSpecAllOf struct for TelemetryDruidDimensionTopNMetricSpecAllOf
type TelemetryDruidDimensionTopNMetricSpecAllOf struct {
	// Specifies the sorting order. It can be one of the following values. \"lexicographic\", \"alphanumeric\", \"numeric\", \"strlen\". * lexicographic - Sorts values by converting Strings to their UTF-8 byte array representations and comparing lexicographically, byte-by-byte. * alphanumeric - Suitable for strings with both numeric and non-numeric content, e.g. \"file12 sorts after file2\". See https://github.com/amjjd/java-alphanum for more details on how this ordering sorts values. This ordering is not suitable for numbers with decimal points or negative numbers. * numeric - Sorts values as numbers, supports integers and floating point values. Negative values are supported. This sorting order will try to parse all string values as numbers. Unparseable values are treated as nulls, and nulls precede numbers. When comparing two unparseable values (e.g., \"hello\" and \"world\"), this ordering will sort by comparing the unparsed strings lexicographically. * strlen - Sorts values by the their string lengths. When there is a tie, this comparator falls back to using the String compareTo method. * version - Sorts values as versions, e.g. \"10.0 sorts after 9.0\", \"1.0.0-SNAPSHOT sorts after 1.0.0\".
	Ordering *string `json:"ordering,omitempty" yaml:"ordering,omitempty"`
	// The starting point of the sort. For example, if a previousStop value is 'b', all values before 'b' are discarded. This field can be used to paginate through all the dimension values.
	PreviousStop *string `json:"previousStop,omitempty" yaml:"previousStop,omitempty"`
}

// NewTelemetryDruidDimensionTopNMetricSpecAllOf instantiates a new TelemetryDruidDimensionTopNMetricSpecAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidDimensionTopNMetricSpecAllOf() *TelemetryDruidDimensionTopNMetricSpecAllOf {
	this := TelemetryDruidDimensionTopNMetricSpecAllOf{}
	var ordering string = "lexicographic"
	this.Ordering = &ordering
	return &this
}

// NewTelemetryDruidDimensionTopNMetricSpecAllOfWithDefaults instantiates a new TelemetryDruidDimensionTopNMetricSpecAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidDimensionTopNMetricSpecAllOfWithDefaults() *TelemetryDruidDimensionTopNMetricSpecAllOf {
	this := TelemetryDruidDimensionTopNMetricSpecAllOf{}
	var ordering string = "lexicographic"
	this.Ordering = &ordering
	return &this
}

// GetOrdering returns the Ordering field value if set, zero value otherwise.
func (o *TelemetryDruidDimensionTopNMetricSpecAllOf) GetOrdering() string {
	if o == nil || o.Ordering == nil {
		var ret string
		return ret
	}
	return *o.Ordering
}

// GetOrderingOk returns a tuple with the Ordering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidDimensionTopNMetricSpecAllOf) GetOrderingOk() (*string, bool) {
	if o == nil || o.Ordering == nil {
		return nil, false
	}
	return o.Ordering, true
}

// HasOrdering returns a boolean if a field has been set.
func (o *TelemetryDruidDimensionTopNMetricSpecAllOf) HasOrdering() bool {
	if o != nil && o.Ordering != nil {
		return true
	}

	return false
}

// SetOrdering gets a reference to the given string and assigns it to the Ordering field.
func (o *TelemetryDruidDimensionTopNMetricSpecAllOf) SetOrdering(v string) {
	o.Ordering = &v
}

// GetPreviousStop returns the PreviousStop field value if set, zero value otherwise.
func (o *TelemetryDruidDimensionTopNMetricSpecAllOf) GetPreviousStop() string {
	if o == nil || o.PreviousStop == nil {
		var ret string
		return ret
	}
	return *o.PreviousStop
}

// GetPreviousStopOk returns a tuple with the PreviousStop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidDimensionTopNMetricSpecAllOf) GetPreviousStopOk() (*string, bool) {
	if o == nil || o.PreviousStop == nil {
		return nil, false
	}
	return o.PreviousStop, true
}

// HasPreviousStop returns a boolean if a field has been set.
func (o *TelemetryDruidDimensionTopNMetricSpecAllOf) HasPreviousStop() bool {
	if o != nil && o.PreviousStop != nil {
		return true
	}

	return false
}

// SetPreviousStop gets a reference to the given string and assigns it to the PreviousStop field.
func (o *TelemetryDruidDimensionTopNMetricSpecAllOf) SetPreviousStop(v string) {
	o.PreviousStop = &v
}

func (o TelemetryDruidDimensionTopNMetricSpecAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ordering != nil {
		toSerialize["ordering"] = o.Ordering
	}
	if o.PreviousStop != nil {
		toSerialize["previousStop"] = o.PreviousStop
	}
	return json.Marshal(toSerialize)
}

type NullableTelemetryDruidDimensionTopNMetricSpecAllOf struct {
	value *TelemetryDruidDimensionTopNMetricSpecAllOf
	isSet bool
}

func (v NullableTelemetryDruidDimensionTopNMetricSpecAllOf) Get() *TelemetryDruidDimensionTopNMetricSpecAllOf {
	return v.value
}

func (v *NullableTelemetryDruidDimensionTopNMetricSpecAllOf) Set(val *TelemetryDruidDimensionTopNMetricSpecAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidDimensionTopNMetricSpecAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidDimensionTopNMetricSpecAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidDimensionTopNMetricSpecAllOf(val *TelemetryDruidDimensionTopNMetricSpecAllOf) *NullableTelemetryDruidDimensionTopNMetricSpecAllOf {
	return &NullableTelemetryDruidDimensionTopNMetricSpecAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidDimensionTopNMetricSpecAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidDimensionTopNMetricSpecAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
