/*
Keycloak Admin REST API

This is a REST API reference for the Keycloak Admin REST API.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UPConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UPConfig{}

// UPConfig struct for UPConfig
type UPConfig struct {
	Attributes []UPAttribute `json:"attributes,omitempty"`
	Groups     []UPGroup     `json:"groups,omitempty"`
}

// NewUPConfig instantiates a new UPConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUPConfig() *UPConfig {
	this := UPConfig{}
	return &this
}

// NewUPConfigWithDefaults instantiates a new UPConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUPConfigWithDefaults() *UPConfig {
	this := UPConfig{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UPConfig) GetAttributes() []UPAttribute {
	if o == nil || IsNil(o.Attributes) {
		var ret []UPAttribute
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UPConfig) GetAttributesOk() ([]UPAttribute, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UPConfig) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []UPAttribute and assigns it to the Attributes field.
func (o *UPConfig) SetAttributes(v []UPAttribute) {
	o.Attributes = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *UPConfig) GetGroups() []UPGroup {
	if o == nil || IsNil(o.Groups) {
		var ret []UPGroup
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UPConfig) GetGroupsOk() ([]UPGroup, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *UPConfig) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []UPGroup and assigns it to the Groups field.
func (o *UPConfig) SetGroups(v []UPGroup) {
	o.Groups = v
}

func (o UPConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UPConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	return toSerialize, nil
}

type NullableUPConfig struct {
	value *UPConfig
	isSet bool
}

func (v NullableUPConfig) Get() *UPConfig {
	return v.value
}

func (v *NullableUPConfig) Set(val *UPConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableUPConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableUPConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUPConfig(val *UPConfig) *NullableUPConfig {
	return &NullableUPConfig{value: val, isSet: true}
}

func (v NullableUPConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUPConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
