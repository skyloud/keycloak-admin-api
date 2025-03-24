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

// checks if the UPGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UPGroup{}

// UPGroup struct for UPGroup
type UPGroup struct {
	Name               *string                           `json:"name,omitempty"`
	DisplayHeader      *string                           `json:"displayHeader,omitempty"`
	DisplayDescription *string                           `json:"displayDescription,omitempty"`
	Annotations        map[string]map[string]interface{} `json:"annotations,omitempty"`
}

// NewUPGroup instantiates a new UPGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUPGroup() *UPGroup {
	this := UPGroup{}
	return &this
}

// NewUPGroupWithDefaults instantiates a new UPGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUPGroupWithDefaults() *UPGroup {
	this := UPGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UPGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UPGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UPGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UPGroup) SetName(v string) {
	o.Name = &v
}

// GetDisplayHeader returns the DisplayHeader field value if set, zero value otherwise.
func (o *UPGroup) GetDisplayHeader() string {
	if o == nil || IsNil(o.DisplayHeader) {
		var ret string
		return ret
	}
	return *o.DisplayHeader
}

// GetDisplayHeaderOk returns a tuple with the DisplayHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UPGroup) GetDisplayHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayHeader) {
		return nil, false
	}
	return o.DisplayHeader, true
}

// HasDisplayHeader returns a boolean if a field has been set.
func (o *UPGroup) HasDisplayHeader() bool {
	if o != nil && !IsNil(o.DisplayHeader) {
		return true
	}

	return false
}

// SetDisplayHeader gets a reference to the given string and assigns it to the DisplayHeader field.
func (o *UPGroup) SetDisplayHeader(v string) {
	o.DisplayHeader = &v
}

// GetDisplayDescription returns the DisplayDescription field value if set, zero value otherwise.
func (o *UPGroup) GetDisplayDescription() string {
	if o == nil || IsNil(o.DisplayDescription) {
		var ret string
		return ret
	}
	return *o.DisplayDescription
}

// GetDisplayDescriptionOk returns a tuple with the DisplayDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UPGroup) GetDisplayDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayDescription) {
		return nil, false
	}
	return o.DisplayDescription, true
}

// HasDisplayDescription returns a boolean if a field has been set.
func (o *UPGroup) HasDisplayDescription() bool {
	if o != nil && !IsNil(o.DisplayDescription) {
		return true
	}

	return false
}

// SetDisplayDescription gets a reference to the given string and assigns it to the DisplayDescription field.
func (o *UPGroup) SetDisplayDescription(v string) {
	o.DisplayDescription = &v
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *UPGroup) GetAnnotations() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Annotations) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UPGroup) GetAnnotationsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Annotations) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *UPGroup) HasAnnotations() bool {
	if o != nil && !IsNil(o.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]map[string]interface{} and assigns it to the Annotations field.
func (o *UPGroup) SetAnnotations(v map[string]map[string]interface{}) {
	o.Annotations = v
}

func (o UPGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UPGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DisplayHeader) {
		toSerialize["displayHeader"] = o.DisplayHeader
	}
	if !IsNil(o.DisplayDescription) {
		toSerialize["displayDescription"] = o.DisplayDescription
	}
	if !IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	return toSerialize, nil
}

type NullableUPGroup struct {
	value *UPGroup
	isSet bool
}

func (v NullableUPGroup) Get() *UPGroup {
	return v.value
}

func (v *NullableUPGroup) Set(val *UPGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableUPGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableUPGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUPGroup(val *UPGroup) *NullableUPGroup {
	return &NullableUPGroup{value: val, isSet: true}
}

func (v NullableUPGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUPGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
