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

// checks if the ContinueWithSetOrySessionToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContinueWithSetOrySessionToken{}

// ContinueWithSetOrySessionToken Indicates that a session was issued, and the application should use this token for authenticated requests
type ContinueWithSetOrySessionToken struct {
	// action will always be \"set_ory_session_token\"
	Action string `json:"action"`
	// Token is the token of the session
	OrySessionToken *string `json:"ory_session_token,omitempty"`
}

type _ContinueWithSetOrySessionToken ContinueWithSetOrySessionToken

// NewContinueWithSetOrySessionToken instantiates a new ContinueWithSetOrySessionToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContinueWithSetOrySessionToken(action string) *ContinueWithSetOrySessionToken {
	this := ContinueWithSetOrySessionToken{}
	this.Action = action
	return &this
}

// NewContinueWithSetOrySessionTokenWithDefaults instantiates a new ContinueWithSetOrySessionToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContinueWithSetOrySessionTokenWithDefaults() *ContinueWithSetOrySessionToken {
	this := ContinueWithSetOrySessionToken{}
	return &this
}

// GetAction returns the Action field value
func (o *ContinueWithSetOrySessionToken) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ContinueWithSetOrySessionToken) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ContinueWithSetOrySessionToken) SetAction(v string) {
	o.Action = v
}

// GetOrySessionToken returns the OrySessionToken field value if set, zero value otherwise.
func (o *ContinueWithSetOrySessionToken) GetOrySessionToken() string {
	if o == nil || IsNil(o.OrySessionToken) {
		var ret string
		return ret
	}
	return *o.OrySessionToken
}

// GetOrySessionTokenOk returns a tuple with the OrySessionToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContinueWithSetOrySessionToken) GetOrySessionTokenOk() (*string, bool) {
	if o == nil || IsNil(o.OrySessionToken) {
		return nil, false
	}
	return o.OrySessionToken, true
}

// HasOrySessionToken returns a boolean if a field has been set.
func (o *ContinueWithSetOrySessionToken) HasOrySessionToken() bool {
	if o != nil && !IsNil(o.OrySessionToken) {
		return true
	}

	return false
}

// SetOrySessionToken gets a reference to the given string and assigns it to the OrySessionToken field.
func (o *ContinueWithSetOrySessionToken) SetOrySessionToken(v string) {
	o.OrySessionToken = &v
}

func (o ContinueWithSetOrySessionToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContinueWithSetOrySessionToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	if !IsNil(o.OrySessionToken) {
		toSerialize["ory_session_token"] = o.OrySessionToken
	}
	return toSerialize, nil
}

func (o *ContinueWithSetOrySessionToken) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
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

	varContinueWithSetOrySessionToken := _ContinueWithSetOrySessionToken{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContinueWithSetOrySessionToken)

	if err != nil {
		return err
	}

	*o = ContinueWithSetOrySessionToken(varContinueWithSetOrySessionToken)

	return err
}

type NullableContinueWithSetOrySessionToken struct {
	value *ContinueWithSetOrySessionToken
	isSet bool
}

func (v NullableContinueWithSetOrySessionToken) Get() *ContinueWithSetOrySessionToken {
	return v.value
}

func (v *NullableContinueWithSetOrySessionToken) Set(val *ContinueWithSetOrySessionToken) {
	v.value = val
	v.isSet = true
}

func (v NullableContinueWithSetOrySessionToken) IsSet() bool {
	return v.isSet
}

func (v *NullableContinueWithSetOrySessionToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContinueWithSetOrySessionToken(val *ContinueWithSetOrySessionToken) *NullableContinueWithSetOrySessionToken {
	return &NullableContinueWithSetOrySessionToken{value: val, isSet: true}
}

func (v NullableContinueWithSetOrySessionToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContinueWithSetOrySessionToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
