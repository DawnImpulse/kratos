/*
Ory Identities API

This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more.

API version:
Contact: office@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the UpdateLoginFlowWithWebAuthnMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateLoginFlowWithWebAuthnMethod{}

// UpdateLoginFlowWithWebAuthnMethod Update Login Flow with WebAuthn Method
type UpdateLoginFlowWithWebAuthnMethod struct {
	// Sending the anti-csrf token is only required for browser login flows.
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Identifier is the email or username of the user trying to log in.
	Identifier string `json:"identifier"`
	// Method should be set to \"webAuthn\" when logging in using the WebAuthn strategy.
	Method string `json:"method"`
	// Login a WebAuthn Security Key  This must contain the ID of the WebAuthN connection.
	WebauthnLogin *string `json:"webauthn_login,omitempty"`
}

type _UpdateLoginFlowWithWebAuthnMethod UpdateLoginFlowWithWebAuthnMethod

// NewUpdateLoginFlowWithWebAuthnMethod instantiates a new UpdateLoginFlowWithWebAuthnMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLoginFlowWithWebAuthnMethod(identifier string, method string) *UpdateLoginFlowWithWebAuthnMethod {
	this := UpdateLoginFlowWithWebAuthnMethod{}
	this.Identifier = identifier
	this.Method = method
	return &this
}

// NewUpdateLoginFlowWithWebAuthnMethodWithDefaults instantiates a new UpdateLoginFlowWithWebAuthnMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLoginFlowWithWebAuthnMethodWithDefaults() *UpdateLoginFlowWithWebAuthnMethod {
	this := UpdateLoginFlowWithWebAuthnMethod{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *UpdateLoginFlowWithWebAuthnMethod) GetCsrfToken() string {
	if o == nil || IsNil(o.CsrfToken) {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoginFlowWithWebAuthnMethod) GetCsrfTokenOk() (*string, bool) {
	if o == nil || IsNil(o.CsrfToken) {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *UpdateLoginFlowWithWebAuthnMethod) HasCsrfToken() bool {
	if o != nil && !IsNil(o.CsrfToken) {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *UpdateLoginFlowWithWebAuthnMethod) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetIdentifier returns the Identifier field value
func (o *UpdateLoginFlowWithWebAuthnMethod) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *UpdateLoginFlowWithWebAuthnMethod) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *UpdateLoginFlowWithWebAuthnMethod) SetIdentifier(v string) {
	o.Identifier = v
}

// GetMethod returns the Method field value
func (o *UpdateLoginFlowWithWebAuthnMethod) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *UpdateLoginFlowWithWebAuthnMethod) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *UpdateLoginFlowWithWebAuthnMethod) SetMethod(v string) {
	o.Method = v
}

// GetWebauthnLogin returns the WebauthnLogin field value if set, zero value otherwise.
func (o *UpdateLoginFlowWithWebAuthnMethod) GetWebauthnLogin() string {
	if o == nil || IsNil(o.WebauthnLogin) {
		var ret string
		return ret
	}
	return *o.WebauthnLogin
}

// GetWebauthnLoginOk returns a tuple with the WebauthnLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoginFlowWithWebAuthnMethod) GetWebauthnLoginOk() (*string, bool) {
	if o == nil || IsNil(o.WebauthnLogin) {
		return nil, false
	}
	return o.WebauthnLogin, true
}

// HasWebauthnLogin returns a boolean if a field has been set.
func (o *UpdateLoginFlowWithWebAuthnMethod) HasWebauthnLogin() bool {
	if o != nil && !IsNil(o.WebauthnLogin) {
		return true
	}

	return false
}

// SetWebauthnLogin gets a reference to the given string and assigns it to the WebauthnLogin field.
func (o *UpdateLoginFlowWithWebAuthnMethod) SetWebauthnLogin(v string) {
	o.WebauthnLogin = &v
}

func (o UpdateLoginFlowWithWebAuthnMethod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateLoginFlowWithWebAuthnMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CsrfToken) {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	toSerialize["identifier"] = o.Identifier
	toSerialize["method"] = o.Method
	if !IsNil(o.WebauthnLogin) {
		toSerialize["webauthn_login"] = o.WebauthnLogin
	}
	return toSerialize, nil
}

func (o *UpdateLoginFlowWithWebAuthnMethod) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"identifier",
		"method",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateLoginFlowWithWebAuthnMethod := _UpdateLoginFlowWithWebAuthnMethod{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateLoginFlowWithWebAuthnMethod)

	if err != nil {
		return err
	}

	*o = UpdateLoginFlowWithWebAuthnMethod(varUpdateLoginFlowWithWebAuthnMethod)

	return err
}

type NullableUpdateLoginFlowWithWebAuthnMethod struct {
	value *UpdateLoginFlowWithWebAuthnMethod
	isSet bool
}

func (v NullableUpdateLoginFlowWithWebAuthnMethod) Get() *UpdateLoginFlowWithWebAuthnMethod {
	return v.value
}

func (v *NullableUpdateLoginFlowWithWebAuthnMethod) Set(val *UpdateLoginFlowWithWebAuthnMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLoginFlowWithWebAuthnMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLoginFlowWithWebAuthnMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLoginFlowWithWebAuthnMethod(val *UpdateLoginFlowWithWebAuthnMethod) *NullableUpdateLoginFlowWithWebAuthnMethod {
	return &NullableUpdateLoginFlowWithWebAuthnMethod{value: val, isSet: true}
}

func (v NullableUpdateLoginFlowWithWebAuthnMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLoginFlowWithWebAuthnMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
