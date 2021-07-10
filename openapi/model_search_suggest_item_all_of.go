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
)

// SearchSuggestItemAllOf Definition of the list of properties defined in 'search.SuggestItem', excluding properties defined in parent classes.
type SearchSuggestItemAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Flag for returning complete objects that matched the global search criteria.
	CompleteMo *bool `json:"CompleteMo,omitempty"`
	// Additional filter parameters for global search. You can also specify OData select fields here. Maximum Query Length is limited to 10000.
	Rawquery *string `json:"Rawquery,omitempty"`
	// Starting offset for the results to be returned from external search engine.
	Skip *int64 `json:"Skip,omitempty"`
	// Main search term used for global search across all Managed Objects that has search enabled. Search Term can be up to 200 characters long.
	SuggestTerm *string `json:"SuggestTerm,omitempty"`
	// Maximum number of results to be returned from external search engine.
	Top *int64 `json:"Top,omitempty"`
	// Object type filter of a Managed Object. Search will be restricted only on the specified object types.  Do not provide IndexMoTypes filter in the rawquery, if you specify values in this field.
	Type                 *string `json:"Type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchSuggestItemAllOf SearchSuggestItemAllOf

// NewSearchSuggestItemAllOf instantiates a new SearchSuggestItemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchSuggestItemAllOf(classId string, objectType string) *SearchSuggestItemAllOf {
	this := SearchSuggestItemAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSearchSuggestItemAllOfWithDefaults instantiates a new SearchSuggestItemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchSuggestItemAllOfWithDefaults() *SearchSuggestItemAllOf {
	this := SearchSuggestItemAllOf{}
	var classId string = "search.SuggestItem"
	this.ClassId = classId
	var objectType string = "search.SuggestItem"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SearchSuggestItemAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SearchSuggestItemAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SearchSuggestItemAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SearchSuggestItemAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SearchSuggestItemAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SearchSuggestItemAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCompleteMo returns the CompleteMo field value if set, zero value otherwise.
func (o *SearchSuggestItemAllOf) GetCompleteMo() bool {
	if o == nil || o.CompleteMo == nil {
		var ret bool
		return ret
	}
	return *o.CompleteMo
}

// GetCompleteMoOk returns a tuple with the CompleteMo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSuggestItemAllOf) GetCompleteMoOk() (*bool, bool) {
	if o == nil || o.CompleteMo == nil {
		return nil, false
	}
	return o.CompleteMo, true
}

// HasCompleteMo returns a boolean if a field has been set.
func (o *SearchSuggestItemAllOf) HasCompleteMo() bool {
	if o != nil && o.CompleteMo != nil {
		return true
	}

	return false
}

// SetCompleteMo gets a reference to the given bool and assigns it to the CompleteMo field.
func (o *SearchSuggestItemAllOf) SetCompleteMo(v bool) {
	o.CompleteMo = &v
}

// GetRawquery returns the Rawquery field value if set, zero value otherwise.
func (o *SearchSuggestItemAllOf) GetRawquery() string {
	if o == nil || o.Rawquery == nil {
		var ret string
		return ret
	}
	return *o.Rawquery
}

// GetRawqueryOk returns a tuple with the Rawquery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSuggestItemAllOf) GetRawqueryOk() (*string, bool) {
	if o == nil || o.Rawquery == nil {
		return nil, false
	}
	return o.Rawquery, true
}

// HasRawquery returns a boolean if a field has been set.
func (o *SearchSuggestItemAllOf) HasRawquery() bool {
	if o != nil && o.Rawquery != nil {
		return true
	}

	return false
}

// SetRawquery gets a reference to the given string and assigns it to the Rawquery field.
func (o *SearchSuggestItemAllOf) SetRawquery(v string) {
	o.Rawquery = &v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *SearchSuggestItemAllOf) GetSkip() int64 {
	if o == nil || o.Skip == nil {
		var ret int64
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSuggestItemAllOf) GetSkipOk() (*int64, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *SearchSuggestItemAllOf) HasSkip() bool {
	if o != nil && o.Skip != nil {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int64 and assigns it to the Skip field.
func (o *SearchSuggestItemAllOf) SetSkip(v int64) {
	o.Skip = &v
}

// GetSuggestTerm returns the SuggestTerm field value if set, zero value otherwise.
func (o *SearchSuggestItemAllOf) GetSuggestTerm() string {
	if o == nil || o.SuggestTerm == nil {
		var ret string
		return ret
	}
	return *o.SuggestTerm
}

// GetSuggestTermOk returns a tuple with the SuggestTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSuggestItemAllOf) GetSuggestTermOk() (*string, bool) {
	if o == nil || o.SuggestTerm == nil {
		return nil, false
	}
	return o.SuggestTerm, true
}

// HasSuggestTerm returns a boolean if a field has been set.
func (o *SearchSuggestItemAllOf) HasSuggestTerm() bool {
	if o != nil && o.SuggestTerm != nil {
		return true
	}

	return false
}

// SetSuggestTerm gets a reference to the given string and assigns it to the SuggestTerm field.
func (o *SearchSuggestItemAllOf) SetSuggestTerm(v string) {
	o.SuggestTerm = &v
}

// GetTop returns the Top field value if set, zero value otherwise.
func (o *SearchSuggestItemAllOf) GetTop() int64 {
	if o == nil || o.Top == nil {
		var ret int64
		return ret
	}
	return *o.Top
}

// GetTopOk returns a tuple with the Top field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSuggestItemAllOf) GetTopOk() (*int64, bool) {
	if o == nil || o.Top == nil {
		return nil, false
	}
	return o.Top, true
}

// HasTop returns a boolean if a field has been set.
func (o *SearchSuggestItemAllOf) HasTop() bool {
	if o != nil && o.Top != nil {
		return true
	}

	return false
}

// SetTop gets a reference to the given int64 and assigns it to the Top field.
func (o *SearchSuggestItemAllOf) SetTop(v int64) {
	o.Top = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SearchSuggestItemAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSuggestItemAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SearchSuggestItemAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SearchSuggestItemAllOf) SetType(v string) {
	o.Type = &v
}

func (o SearchSuggestItemAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.CompleteMo != nil {
		toSerialize["CompleteMo"] = o.CompleteMo
	}
	if o.Rawquery != nil {
		toSerialize["Rawquery"] = o.Rawquery
	}
	if o.Skip != nil {
		toSerialize["Skip"] = o.Skip
	}
	if o.SuggestTerm != nil {
		toSerialize["SuggestTerm"] = o.SuggestTerm
	}
	if o.Top != nil {
		toSerialize["Top"] = o.Top
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SearchSuggestItemAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSearchSuggestItemAllOf := _SearchSuggestItemAllOf{}

	if err = json.Unmarshal(bytes, &varSearchSuggestItemAllOf); err == nil {
		*o = SearchSuggestItemAllOf(varSearchSuggestItemAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "CompleteMo")
		delete(additionalProperties, "Rawquery")
		delete(additionalProperties, "Skip")
		delete(additionalProperties, "SuggestTerm")
		delete(additionalProperties, "Top")
		delete(additionalProperties, "Type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchSuggestItemAllOf struct {
	value *SearchSuggestItemAllOf
	isSet bool
}

func (v NullableSearchSuggestItemAllOf) Get() *SearchSuggestItemAllOf {
	return v.value
}

func (v *NullableSearchSuggestItemAllOf) Set(val *SearchSuggestItemAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSuggestItemAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSuggestItemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSuggestItemAllOf(val *SearchSuggestItemAllOf) *NullableSearchSuggestItemAllOf {
	return &NullableSearchSuggestItemAllOf{value: val, isSet: true}
}

func (v NullableSearchSuggestItemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSuggestItemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
