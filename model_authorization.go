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

// checks if the Authorization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Authorization{}

// Authorization struct for Authorization
type Authorization struct {
	Permissions []Permission `json:"permissions,omitempty"`
}

// NewAuthorization instantiates a new Authorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorization() *Authorization {
	this := Authorization{}
	return &this
}

// NewAuthorizationWithDefaults instantiates a new Authorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationWithDefaults() *Authorization {
	this := Authorization{}
	return &this
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *Authorization) GetPermissions() []Permission {
	if o == nil || IsNil(o.Permissions) {
		var ret []Permission
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Authorization) GetPermissionsOk() ([]Permission, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *Authorization) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []Permission and assigns it to the Permissions field.
func (o *Authorization) SetPermissions(v []Permission) {
	o.Permissions = v
}

func (o Authorization) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Authorization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	return toSerialize, nil
}

type NullableAuthorization struct {
	value *Authorization
	isSet bool
}

func (v NullableAuthorization) Get() *Authorization {
	return v.value
}

func (v *NullableAuthorization) Set(val *Authorization) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorization(val *Authorization) *NullableAuthorization {
	return &NullableAuthorization{value: val, isSet: true}
}

func (v NullableAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
