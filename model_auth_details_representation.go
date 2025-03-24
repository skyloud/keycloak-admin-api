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

// checks if the AuthDetailsRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthDetailsRepresentation{}

// AuthDetailsRepresentation struct for AuthDetailsRepresentation
type AuthDetailsRepresentation struct {
	RealmId   *string `json:"realmId,omitempty"`
	ClientId  *string `json:"clientId,omitempty"`
	UserId    *string `json:"userId,omitempty"`
	IpAddress *string `json:"ipAddress,omitempty"`
}

// NewAuthDetailsRepresentation instantiates a new AuthDetailsRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthDetailsRepresentation() *AuthDetailsRepresentation {
	this := AuthDetailsRepresentation{}
	return &this
}

// NewAuthDetailsRepresentationWithDefaults instantiates a new AuthDetailsRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthDetailsRepresentationWithDefaults() *AuthDetailsRepresentation {
	this := AuthDetailsRepresentation{}
	return &this
}

// GetRealmId returns the RealmId field value if set, zero value otherwise.
func (o *AuthDetailsRepresentation) GetRealmId() string {
	if o == nil || IsNil(o.RealmId) {
		var ret string
		return ret
	}
	return *o.RealmId
}

// GetRealmIdOk returns a tuple with the RealmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthDetailsRepresentation) GetRealmIdOk() (*string, bool) {
	if o == nil || IsNil(o.RealmId) {
		return nil, false
	}
	return o.RealmId, true
}

// HasRealmId returns a boolean if a field has been set.
func (o *AuthDetailsRepresentation) HasRealmId() bool {
	if o != nil && !IsNil(o.RealmId) {
		return true
	}

	return false
}

// SetRealmId gets a reference to the given string and assigns it to the RealmId field.
func (o *AuthDetailsRepresentation) SetRealmId(v string) {
	o.RealmId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AuthDetailsRepresentation) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthDetailsRepresentation) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AuthDetailsRepresentation) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AuthDetailsRepresentation) SetClientId(v string) {
	o.ClientId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AuthDetailsRepresentation) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthDetailsRepresentation) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AuthDetailsRepresentation) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AuthDetailsRepresentation) SetUserId(v string) {
	o.UserId = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *AuthDetailsRepresentation) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthDetailsRepresentation) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *AuthDetailsRepresentation) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *AuthDetailsRepresentation) SetIpAddress(v string) {
	o.IpAddress = &v
}

func (o AuthDetailsRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthDetailsRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RealmId) {
		toSerialize["realmId"] = o.RealmId
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	return toSerialize, nil
}

type NullableAuthDetailsRepresentation struct {
	value *AuthDetailsRepresentation
	isSet bool
}

func (v NullableAuthDetailsRepresentation) Get() *AuthDetailsRepresentation {
	return v.value
}

func (v *NullableAuthDetailsRepresentation) Set(val *AuthDetailsRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthDetailsRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthDetailsRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthDetailsRepresentation(val *AuthDetailsRepresentation) *NullableAuthDetailsRepresentation {
	return &NullableAuthDetailsRepresentation{value: val, isSet: true}
}

func (v NullableAuthDetailsRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthDetailsRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
