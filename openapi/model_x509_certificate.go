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
	"time"
)

// X509Certificate The representation of an X.509 certificate.
type X509Certificate struct {
	MoBaseComplexType `yaml:"MoBaseComplexType,inline"`
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId" yaml:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string                        `json:"ObjectType" yaml:"ObjectType"`
	Issuer     NullablePkixDistinguishedName `json:"Issuer,omitempty" yaml:"Issuer,omitempty"`
	// The date on which the certificate's validity period ends.
	NotAfter *time.Time `json:"NotAfter,omitempty" yaml:"NotAfter,omitempty"`
	// The date on which the certificate's validity period begins.
	NotBefore *time.Time `json:"NotBefore,omitempty" yaml:"NotBefore,omitempty"`
	// The base64 encoded certificate in PEM format.
	PemCertificate *string `json:"PemCertificate,omitempty" yaml:"PemCertificate,omitempty"`
	// The computed SHA-256 fingerprint of the certificate. Equivalent to 'openssl x509 -fingerprint -sha256'.
	Sha256Fingerprint *string `json:"Sha256Fingerprint,omitempty" yaml:"Sha256Fingerprint,omitempty"`
	// Signature algorithm, as specified in [RFC 5280](https://tools.ietf.org/html/rfc5280).
	SignatureAlgorithm *string                       `json:"SignatureAlgorithm,omitempty" yaml:"SignatureAlgorithm,omitempty"`
	Subject            NullablePkixDistinguishedName `json:"Subject,omitempty" yaml:"Subject,omitempty"`
}

// NewX509Certificate instantiates a new X509Certificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewX509Certificate(classId string, objectType string) *X509Certificate {
	this := X509Certificate{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewX509CertificateWithDefaults instantiates a new X509Certificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewX509CertificateWithDefaults() *X509Certificate {
	this := X509Certificate{}
	var classId string = "x509.Certificate"
	this.ClassId = classId
	var objectType string = "x509.Certificate"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *X509Certificate) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *X509Certificate) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *X509Certificate) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *X509Certificate) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *X509Certificate) GetIssuer() PkixDistinguishedName {
	if o == nil || o.Issuer.Get() == nil {
		var ret PkixDistinguishedName
		return ret
	}
	return *o.Issuer.Get()
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *X509Certificate) GetIssuerOk() (*PkixDistinguishedName, bool) {
	if o == nil {
		return nil, false
	}
	return o.Issuer.Get(), o.Issuer.IsSet()
}

// HasIssuer returns a boolean if a field has been set.
func (o *X509Certificate) HasIssuer() bool {
	if o != nil && o.Issuer.IsSet() {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given NullablePkixDistinguishedName and assigns it to the Issuer field.
func (o *X509Certificate) SetIssuer(v PkixDistinguishedName) {
	o.Issuer.Set(&v)
}

// SetIssuerNil sets the value for Issuer to be an explicit nil
func (o *X509Certificate) SetIssuerNil() {
	o.Issuer.Set(nil)
}

// UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
func (o *X509Certificate) UnsetIssuer() {
	o.Issuer.Unset()
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise.
func (o *X509Certificate) GetNotAfter() time.Time {
	if o == nil || o.NotAfter == nil {
		var ret time.Time
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetNotAfterOk() (*time.Time, bool) {
	if o == nil || o.NotAfter == nil {
		return nil, false
	}
	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *X509Certificate) HasNotAfter() bool {
	if o != nil && o.NotAfter != nil {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given time.Time and assigns it to the NotAfter field.
func (o *X509Certificate) SetNotAfter(v time.Time) {
	o.NotAfter = &v
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise.
func (o *X509Certificate) GetNotBefore() time.Time {
	if o == nil || o.NotBefore == nil {
		var ret time.Time
		return ret
	}
	return *o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetNotBeforeOk() (*time.Time, bool) {
	if o == nil || o.NotBefore == nil {
		return nil, false
	}
	return o.NotBefore, true
}

// HasNotBefore returns a boolean if a field has been set.
func (o *X509Certificate) HasNotBefore() bool {
	if o != nil && o.NotBefore != nil {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given time.Time and assigns it to the NotBefore field.
func (o *X509Certificate) SetNotBefore(v time.Time) {
	o.NotBefore = &v
}

// GetPemCertificate returns the PemCertificate field value if set, zero value otherwise.
func (o *X509Certificate) GetPemCertificate() string {
	if o == nil || o.PemCertificate == nil {
		var ret string
		return ret
	}
	return *o.PemCertificate
}

// GetPemCertificateOk returns a tuple with the PemCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetPemCertificateOk() (*string, bool) {
	if o == nil || o.PemCertificate == nil {
		return nil, false
	}
	return o.PemCertificate, true
}

// HasPemCertificate returns a boolean if a field has been set.
func (o *X509Certificate) HasPemCertificate() bool {
	if o != nil && o.PemCertificate != nil {
		return true
	}

	return false
}

// SetPemCertificate gets a reference to the given string and assigns it to the PemCertificate field.
func (o *X509Certificate) SetPemCertificate(v string) {
	o.PemCertificate = &v
}

// GetSha256Fingerprint returns the Sha256Fingerprint field value if set, zero value otherwise.
func (o *X509Certificate) GetSha256Fingerprint() string {
	if o == nil || o.Sha256Fingerprint == nil {
		var ret string
		return ret
	}
	return *o.Sha256Fingerprint
}

// GetSha256FingerprintOk returns a tuple with the Sha256Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetSha256FingerprintOk() (*string, bool) {
	if o == nil || o.Sha256Fingerprint == nil {
		return nil, false
	}
	return o.Sha256Fingerprint, true
}

// HasSha256Fingerprint returns a boolean if a field has been set.
func (o *X509Certificate) HasSha256Fingerprint() bool {
	if o != nil && o.Sha256Fingerprint != nil {
		return true
	}

	return false
}

// SetSha256Fingerprint gets a reference to the given string and assigns it to the Sha256Fingerprint field.
func (o *X509Certificate) SetSha256Fingerprint(v string) {
	o.Sha256Fingerprint = &v
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value if set, zero value otherwise.
func (o *X509Certificate) GetSignatureAlgorithm() string {
	if o == nil || o.SignatureAlgorithm == nil {
		var ret string
		return ret
	}
	return *o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509Certificate) GetSignatureAlgorithmOk() (*string, bool) {
	if o == nil || o.SignatureAlgorithm == nil {
		return nil, false
	}
	return o.SignatureAlgorithm, true
}

// HasSignatureAlgorithm returns a boolean if a field has been set.
func (o *X509Certificate) HasSignatureAlgorithm() bool {
	if o != nil && o.SignatureAlgorithm != nil {
		return true
	}

	return false
}

// SetSignatureAlgorithm gets a reference to the given string and assigns it to the SignatureAlgorithm field.
func (o *X509Certificate) SetSignatureAlgorithm(v string) {
	o.SignatureAlgorithm = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *X509Certificate) GetSubject() PkixDistinguishedName {
	if o == nil || o.Subject.Get() == nil {
		var ret PkixDistinguishedName
		return ret
	}
	return *o.Subject.Get()
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *X509Certificate) GetSubjectOk() (*PkixDistinguishedName, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subject.Get(), o.Subject.IsSet()
}

// HasSubject returns a boolean if a field has been set.
func (o *X509Certificate) HasSubject() bool {
	if o != nil && o.Subject.IsSet() {
		return true
	}

	return false
}

// SetSubject gets a reference to the given NullablePkixDistinguishedName and assigns it to the Subject field.
func (o *X509Certificate) SetSubject(v PkixDistinguishedName) {
	o.Subject.Set(&v)
}

// SetSubjectNil sets the value for Subject to be an explicit nil
func (o *X509Certificate) SetSubjectNil() {
	o.Subject.Set(nil)
}

// UnsetSubject ensures that no value is present for Subject, not even an explicit nil
func (o *X509Certificate) UnsetSubject() {
	o.Subject.Unset()
}

func (o X509Certificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Issuer.IsSet() {
		toSerialize["Issuer"] = o.Issuer.Get()
	}
	if o.NotAfter != nil {
		toSerialize["NotAfter"] = o.NotAfter
	}
	if o.NotBefore != nil {
		toSerialize["NotBefore"] = o.NotBefore
	}
	if o.PemCertificate != nil {
		toSerialize["PemCertificate"] = o.PemCertificate
	}
	if o.Sha256Fingerprint != nil {
		toSerialize["Sha256Fingerprint"] = o.Sha256Fingerprint
	}
	if o.SignatureAlgorithm != nil {
		toSerialize["SignatureAlgorithm"] = o.SignatureAlgorithm
	}
	if o.Subject.IsSet() {
		toSerialize["Subject"] = o.Subject.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableX509Certificate struct {
	value *X509Certificate
	isSet bool
}

func (v NullableX509Certificate) Get() *X509Certificate {
	return v.value
}

func (v *NullableX509Certificate) Set(val *X509Certificate) {
	v.value = val
	v.isSet = true
}

func (v NullableX509Certificate) IsSet() bool {
	return v.isSet
}

func (v *NullableX509Certificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableX509Certificate(val *X509Certificate) *NullableX509Certificate {
	return &NullableX509Certificate{value: val, isSet: true}
}

func (v NullableX509Certificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableX509Certificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
