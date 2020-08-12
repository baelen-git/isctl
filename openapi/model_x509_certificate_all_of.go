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
	"time"
)

// X509CertificateAllOf Definition of the list of properties defined in 'x509.Certificate', excluding properties defined in parent classes.
type X509CertificateAllOf struct {
	Issuer *PkixDistinguishedName `json:"Issuer,omitempty" yaml:"Issuer,omitempty"`
	// The date on which the certificate's validity period ends.
	NotAfter *time.Time `json:"NotAfter,omitempty" yaml:"NotAfter,omitempty"`
	// The date on which the certificate's validity period begins.
	NotBefore *time.Time `json:"NotBefore,omitempty" yaml:"NotBefore,omitempty"`
	// The base64 encoded certificate in PEM format.
	PemCertificate *string `json:"PemCertificate,omitempty" yaml:"PemCertificate,omitempty"`
	// The computed SHA-256 fingerprint of the certificate. Equivalent to 'openssl x509 -fingerprint -sha256'.
	Sha256Fingerprint *string `json:"Sha256Fingerprint,omitempty" yaml:"Sha256Fingerprint,omitempty"`
	// Signature algorithm, as specified in [RFC 5280](https://tools.ietf.org/html/rfc5280).
	SignatureAlgorithm *string                `json:"SignatureAlgorithm,omitempty" yaml:"SignatureAlgorithm,omitempty"`
	Subject            *PkixDistinguishedName `json:"Subject,omitempty" yaml:"Subject,omitempty"`
}

// NewX509CertificateAllOf instantiates a new X509CertificateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewX509CertificateAllOf() *X509CertificateAllOf {
	this := X509CertificateAllOf{}
	return &this
}

// NewX509CertificateAllOfWithDefaults instantiates a new X509CertificateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewX509CertificateAllOfWithDefaults() *X509CertificateAllOf {
	this := X509CertificateAllOf{}
	return &this
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *X509CertificateAllOf) GetIssuer() PkixDistinguishedName {
	if o == nil || o.Issuer == nil {
		var ret PkixDistinguishedName
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509CertificateAllOf) GetIssuerOk() (*PkixDistinguishedName, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *X509CertificateAllOf) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given PkixDistinguishedName and assigns it to the Issuer field.
func (o *X509CertificateAllOf) SetIssuer(v PkixDistinguishedName) {
	o.Issuer = &v
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise.
func (o *X509CertificateAllOf) GetNotAfter() time.Time {
	if o == nil || o.NotAfter == nil {
		var ret time.Time
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509CertificateAllOf) GetNotAfterOk() (*time.Time, bool) {
	if o == nil || o.NotAfter == nil {
		return nil, false
	}
	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *X509CertificateAllOf) HasNotAfter() bool {
	if o != nil && o.NotAfter != nil {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given time.Time and assigns it to the NotAfter field.
func (o *X509CertificateAllOf) SetNotAfter(v time.Time) {
	o.NotAfter = &v
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise.
func (o *X509CertificateAllOf) GetNotBefore() time.Time {
	if o == nil || o.NotBefore == nil {
		var ret time.Time
		return ret
	}
	return *o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509CertificateAllOf) GetNotBeforeOk() (*time.Time, bool) {
	if o == nil || o.NotBefore == nil {
		return nil, false
	}
	return o.NotBefore, true
}

// HasNotBefore returns a boolean if a field has been set.
func (o *X509CertificateAllOf) HasNotBefore() bool {
	if o != nil && o.NotBefore != nil {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given time.Time and assigns it to the NotBefore field.
func (o *X509CertificateAllOf) SetNotBefore(v time.Time) {
	o.NotBefore = &v
}

// GetPemCertificate returns the PemCertificate field value if set, zero value otherwise.
func (o *X509CertificateAllOf) GetPemCertificate() string {
	if o == nil || o.PemCertificate == nil {
		var ret string
		return ret
	}
	return *o.PemCertificate
}

// GetPemCertificateOk returns a tuple with the PemCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509CertificateAllOf) GetPemCertificateOk() (*string, bool) {
	if o == nil || o.PemCertificate == nil {
		return nil, false
	}
	return o.PemCertificate, true
}

// HasPemCertificate returns a boolean if a field has been set.
func (o *X509CertificateAllOf) HasPemCertificate() bool {
	if o != nil && o.PemCertificate != nil {
		return true
	}

	return false
}

// SetPemCertificate gets a reference to the given string and assigns it to the PemCertificate field.
func (o *X509CertificateAllOf) SetPemCertificate(v string) {
	o.PemCertificate = &v
}

// GetSha256Fingerprint returns the Sha256Fingerprint field value if set, zero value otherwise.
func (o *X509CertificateAllOf) GetSha256Fingerprint() string {
	if o == nil || o.Sha256Fingerprint == nil {
		var ret string
		return ret
	}
	return *o.Sha256Fingerprint
}

// GetSha256FingerprintOk returns a tuple with the Sha256Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509CertificateAllOf) GetSha256FingerprintOk() (*string, bool) {
	if o == nil || o.Sha256Fingerprint == nil {
		return nil, false
	}
	return o.Sha256Fingerprint, true
}

// HasSha256Fingerprint returns a boolean if a field has been set.
func (o *X509CertificateAllOf) HasSha256Fingerprint() bool {
	if o != nil && o.Sha256Fingerprint != nil {
		return true
	}

	return false
}

// SetSha256Fingerprint gets a reference to the given string and assigns it to the Sha256Fingerprint field.
func (o *X509CertificateAllOf) SetSha256Fingerprint(v string) {
	o.Sha256Fingerprint = &v
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value if set, zero value otherwise.
func (o *X509CertificateAllOf) GetSignatureAlgorithm() string {
	if o == nil || o.SignatureAlgorithm == nil {
		var ret string
		return ret
	}
	return *o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509CertificateAllOf) GetSignatureAlgorithmOk() (*string, bool) {
	if o == nil || o.SignatureAlgorithm == nil {
		return nil, false
	}
	return o.SignatureAlgorithm, true
}

// HasSignatureAlgorithm returns a boolean if a field has been set.
func (o *X509CertificateAllOf) HasSignatureAlgorithm() bool {
	if o != nil && o.SignatureAlgorithm != nil {
		return true
	}

	return false
}

// SetSignatureAlgorithm gets a reference to the given string and assigns it to the SignatureAlgorithm field.
func (o *X509CertificateAllOf) SetSignatureAlgorithm(v string) {
	o.SignatureAlgorithm = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *X509CertificateAllOf) GetSubject() PkixDistinguishedName {
	if o == nil || o.Subject == nil {
		var ret PkixDistinguishedName
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509CertificateAllOf) GetSubjectOk() (*PkixDistinguishedName, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *X509CertificateAllOf) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given PkixDistinguishedName and assigns it to the Subject field.
func (o *X509CertificateAllOf) SetSubject(v PkixDistinguishedName) {
	o.Subject = &v
}

func (o X509CertificateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Issuer != nil {
		toSerialize["Issuer"] = o.Issuer
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
	if o.Subject != nil {
		toSerialize["Subject"] = o.Subject
	}
	return json.Marshal(toSerialize)
}

type NullableX509CertificateAllOf struct {
	value *X509CertificateAllOf
	isSet bool
}

func (v NullableX509CertificateAllOf) Get() *X509CertificateAllOf {
	return v.value
}

func (v *NullableX509CertificateAllOf) Set(val *X509CertificateAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableX509CertificateAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableX509CertificateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableX509CertificateAllOf(val *X509CertificateAllOf) *NullableX509CertificateAllOf {
	return &NullableX509CertificateAllOf{value: val, isSet: true}
}

func (v NullableX509CertificateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableX509CertificateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
