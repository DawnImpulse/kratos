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

// checks if the UiNodeScriptAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UiNodeScriptAttributes{}

// UiNodeScriptAttributes struct for UiNodeScriptAttributes
type UiNodeScriptAttributes struct {
	// The script async type
	Async bool `json:"async"`
	// The script cross origin policy
	Crossorigin string `json:"crossorigin"`
	// A unique identifier
	Id string `json:"id"`
	// The script's integrity hash
	Integrity string `json:"integrity"`
	// NodeType represents this node's types. It is a mirror of `node.type` and is primarily used to allow compatibility with OpenAPI 3.0. In this struct it technically always is \"script\".
	NodeType string `json:"node_type"`
	// Nonce for CSP  A nonce you may want to use to improve your Content Security Policy. You do not have to use this value but if you want to improve your CSP policies you may use it. You can also choose to use your own nonce value!
	Nonce string `json:"nonce"`
	// The script referrer policy
	Referrerpolicy string `json:"referrerpolicy"`
	// The script source
	Src string `json:"src"`
	// The script MIME type
	Type string `json:"type"`
}

type _UiNodeScriptAttributes UiNodeScriptAttributes

// NewUiNodeScriptAttributes instantiates a new UiNodeScriptAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiNodeScriptAttributes(async bool, crossorigin string, id string, integrity string, nodeType string, nonce string, referrerpolicy string, src string, type_ string) *UiNodeScriptAttributes {
	this := UiNodeScriptAttributes{}
	this.Async = async
	this.Crossorigin = crossorigin
	this.Id = id
	this.Integrity = integrity
	this.NodeType = nodeType
	this.Nonce = nonce
	this.Referrerpolicy = referrerpolicy
	this.Src = src
	this.Type = type_
	return &this
}

// NewUiNodeScriptAttributesWithDefaults instantiates a new UiNodeScriptAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiNodeScriptAttributesWithDefaults() *UiNodeScriptAttributes {
	this := UiNodeScriptAttributes{}
	return &this
}

// GetAsync returns the Async field value
func (o *UiNodeScriptAttributes) GetAsync() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Async
}

// GetAsyncOk returns a tuple with the Async field value
// and a boolean to check if the value has been set.
func (o *UiNodeScriptAttributes) GetAsyncOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Async, true
}

// SetAsync sets field value
func (o *UiNodeScriptAttributes) SetAsync(v bool) {
	o.Async = v
}

// GetCrossorigin returns the Crossorigin field value
func (o *UiNodeScriptAttributes) GetCrossorigin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Crossorigin
}

// GetCrossoriginOk returns a tuple with the Crossorigin field value
// and a boolean to check if the value has been set.
func (o *UiNodeScriptAttributes) GetCrossoriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Crossorigin, true
}

// SetCrossorigin sets field value
func (o *UiNodeScriptAttributes) SetCrossorigin(v string) {
	o.Crossorigin = v
}

// GetId returns the Id field value
func (o *UiNodeScriptAttributes) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UiNodeScriptAttributes) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UiNodeScriptAttributes) SetId(v string) {
	o.Id = v
}

// GetIntegrity returns the Integrity field value
func (o *UiNodeScriptAttributes) GetIntegrity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Integrity
}

// GetIntegrityOk returns a tuple with the Integrity field value
// and a boolean to check if the value has been set.
func (o *UiNodeScriptAttributes) GetIntegrityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Integrity, true
}

// SetIntegrity sets field value
func (o *UiNodeScriptAttributes) SetIntegrity(v string) {
	o.Integrity = v
}

// GetNodeType returns the NodeType field value
func (o *UiNodeScriptAttributes) GetNodeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *UiNodeScriptAttributes) GetNodeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value
func (o *UiNodeScriptAttributes) SetNodeType(v string) {
	o.NodeType = v
}

// GetNonce returns the Nonce field value
func (o *UiNodeScriptAttributes) GetNonce() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *UiNodeScriptAttributes) GetNonceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *UiNodeScriptAttributes) SetNonce(v string) {
	o.Nonce = v
}

// GetReferrerpolicy returns the Referrerpolicy field value
func (o *UiNodeScriptAttributes) GetReferrerpolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Referrerpolicy
}

// GetReferrerpolicyOk returns a tuple with the Referrerpolicy field value
// and a boolean to check if the value has been set.
func (o *UiNodeScriptAttributes) GetReferrerpolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Referrerpolicy, true
}

// SetReferrerpolicy sets field value
func (o *UiNodeScriptAttributes) SetReferrerpolicy(v string) {
	o.Referrerpolicy = v
}

// GetSrc returns the Src field value
func (o *UiNodeScriptAttributes) GetSrc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Src
}

// GetSrcOk returns a tuple with the Src field value
// and a boolean to check if the value has been set.
func (o *UiNodeScriptAttributes) GetSrcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Src, true
}

// SetSrc sets field value
func (o *UiNodeScriptAttributes) SetSrc(v string) {
	o.Src = v
}

// GetType returns the Type field value
func (o *UiNodeScriptAttributes) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UiNodeScriptAttributes) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UiNodeScriptAttributes) SetType(v string) {
	o.Type = v
}

func (o UiNodeScriptAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UiNodeScriptAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["async"] = o.Async
	toSerialize["crossorigin"] = o.Crossorigin
	toSerialize["id"] = o.Id
	toSerialize["integrity"] = o.Integrity
	toSerialize["node_type"] = o.NodeType
	toSerialize["nonce"] = o.Nonce
	toSerialize["referrerpolicy"] = o.Referrerpolicy
	toSerialize["src"] = o.Src
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *UiNodeScriptAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"async",
		"crossorigin",
		"id",
		"integrity",
		"node_type",
		"nonce",
		"referrerpolicy",
		"src",
		"type",
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

	varUiNodeScriptAttributes := _UiNodeScriptAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUiNodeScriptAttributes)

	if err != nil {
		return err
	}

	*o = UiNodeScriptAttributes(varUiNodeScriptAttributes)

	return err
}

type NullableUiNodeScriptAttributes struct {
	value *UiNodeScriptAttributes
	isSet bool
}

func (v NullableUiNodeScriptAttributes) Get() *UiNodeScriptAttributes {
	return v.value
}

func (v *NullableUiNodeScriptAttributes) Set(val *UiNodeScriptAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableUiNodeScriptAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableUiNodeScriptAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiNodeScriptAttributes(val *UiNodeScriptAttributes) *NullableUiNodeScriptAttributes {
	return &NullableUiNodeScriptAttributes{value: val, isSet: true}
}

func (v NullableUiNodeScriptAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiNodeScriptAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
