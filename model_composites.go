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

// checks if the Composites type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Composites{}

// Composites struct for Composites
type Composites struct {
	Realm       []string             `json:"realm,omitempty"`
	Client      *map[string][]string `json:"client,omitempty"`
	Application *map[string][]string `json:"application,omitempty"`
}

// NewComposites instantiates a new Composites object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComposites() *Composites {
	this := Composites{}
	return &this
}

// NewCompositesWithDefaults instantiates a new Composites object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompositesWithDefaults() *Composites {
	this := Composites{}
	return &this
}

// GetRealm returns the Realm field value if set, zero value otherwise.
func (o *Composites) GetRealm() []string {
	if o == nil || IsNil(o.Realm) {
		var ret []string
		return ret
	}
	return o.Realm
}

// GetRealmOk returns a tuple with the Realm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Composites) GetRealmOk() ([]string, bool) {
	if o == nil || IsNil(o.Realm) {
		return nil, false
	}
	return o.Realm, true
}

// HasRealm returns a boolean if a field has been set.
func (o *Composites) HasRealm() bool {
	if o != nil && !IsNil(o.Realm) {
		return true
	}

	return false
}

// SetRealm gets a reference to the given []string and assigns it to the Realm field.
func (o *Composites) SetRealm(v []string) {
	o.Realm = v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *Composites) GetClient() map[string][]string {
	if o == nil || IsNil(o.Client) {
		var ret map[string][]string
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Composites) GetClientOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *Composites) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given map[string][]string and assigns it to the Client field.
func (o *Composites) SetClient(v map[string][]string) {
	o.Client = &v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *Composites) GetApplication() map[string][]string {
	if o == nil || IsNil(o.Application) {
		var ret map[string][]string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Composites) GetApplicationOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *Composites) HasApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given map[string][]string and assigns it to the Application field.
func (o *Composites) SetApplication(v map[string][]string) {
	o.Application = &v
}

func (o Composites) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Composites) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Realm) {
		toSerialize["realm"] = o.Realm
	}
	if !IsNil(o.Client) {
		toSerialize["client"] = o.Client
	}
	if !IsNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	return toSerialize, nil
}

type NullableComposites struct {
	value *Composites
	isSet bool
}

func (v NullableComposites) Get() *Composites {
	return v.value
}

func (v *NullableComposites) Set(val *Composites) {
	v.value = val
	v.isSet = true
}

func (v NullableComposites) IsSet() bool {
	return v.isSet
}

func (v *NullableComposites) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComposites(val *Composites) *NullableComposites {
	return &NullableComposites{value: val, isSet: true}
}

func (v NullableComposites) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComposites) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
