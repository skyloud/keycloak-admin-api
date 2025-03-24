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

// checks if the RoleRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleRepresentation{}

// RoleRepresentation struct for RoleRepresentation
type RoleRepresentation struct {
	Id                 *string              `json:"id,omitempty"`
	Name               *string              `json:"name,omitempty"`
	Description        *string              `json:"description,omitempty"`
	ScopeParamRequired *bool                `json:"scopeParamRequired,omitempty"`
	Composite          *bool                `json:"composite,omitempty"`
	Composites         *Composites          `json:"composites,omitempty"`
	ClientRole         *bool                `json:"clientRole,omitempty"`
	ContainerId        *string              `json:"containerId,omitempty"`
	Attributes         *map[string][]string `json:"attributes,omitempty"`
}

// NewRoleRepresentation instantiates a new RoleRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleRepresentation() *RoleRepresentation {
	this := RoleRepresentation{}
	return &this
}

// NewRoleRepresentationWithDefaults instantiates a new RoleRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleRepresentationWithDefaults() *RoleRepresentation {
	this := RoleRepresentation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RoleRepresentation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleRepresentation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RoleRepresentation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RoleRepresentation) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoleRepresentation) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleRepresentation) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoleRepresentation) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoleRepresentation) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RoleRepresentation) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleRepresentation) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RoleRepresentation) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RoleRepresentation) SetDescription(v string) {
	o.Description = &v
}

// GetScopeParamRequired returns the ScopeParamRequired field value if set, zero value otherwise.
func (o *RoleRepresentation) GetScopeParamRequired() bool {
	if o == nil || IsNil(o.ScopeParamRequired) {
		var ret bool
		return ret
	}
	return *o.ScopeParamRequired
}

// GetScopeParamRequiredOk returns a tuple with the ScopeParamRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleRepresentation) GetScopeParamRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.ScopeParamRequired) {
		return nil, false
	}
	return o.ScopeParamRequired, true
}

// HasScopeParamRequired returns a boolean if a field has been set.
func (o *RoleRepresentation) HasScopeParamRequired() bool {
	if o != nil && !IsNil(o.ScopeParamRequired) {
		return true
	}

	return false
}

// SetScopeParamRequired gets a reference to the given bool and assigns it to the ScopeParamRequired field.
func (o *RoleRepresentation) SetScopeParamRequired(v bool) {
	o.ScopeParamRequired = &v
}

// GetComposite returns the Composite field value if set, zero value otherwise.
func (o *RoleRepresentation) GetComposite() bool {
	if o == nil || IsNil(o.Composite) {
		var ret bool
		return ret
	}
	return *o.Composite
}

// GetCompositeOk returns a tuple with the Composite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleRepresentation) GetCompositeOk() (*bool, bool) {
	if o == nil || IsNil(o.Composite) {
		return nil, false
	}
	return o.Composite, true
}

// HasComposite returns a boolean if a field has been set.
func (o *RoleRepresentation) HasComposite() bool {
	if o != nil && !IsNil(o.Composite) {
		return true
	}

	return false
}

// SetComposite gets a reference to the given bool and assigns it to the Composite field.
func (o *RoleRepresentation) SetComposite(v bool) {
	o.Composite = &v
}

// GetComposites returns the Composites field value if set, zero value otherwise.
func (o *RoleRepresentation) GetComposites() Composites {
	if o == nil || IsNil(o.Composites) {
		var ret Composites
		return ret
	}
	return *o.Composites
}

// GetCompositesOk returns a tuple with the Composites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleRepresentation) GetCompositesOk() (*Composites, bool) {
	if o == nil || IsNil(o.Composites) {
		return nil, false
	}
	return o.Composites, true
}

// HasComposites returns a boolean if a field has been set.
func (o *RoleRepresentation) HasComposites() bool {
	if o != nil && !IsNil(o.Composites) {
		return true
	}

	return false
}

// SetComposites gets a reference to the given Composites and assigns it to the Composites field.
func (o *RoleRepresentation) SetComposites(v Composites) {
	o.Composites = &v
}

// GetClientRole returns the ClientRole field value if set, zero value otherwise.
func (o *RoleRepresentation) GetClientRole() bool {
	if o == nil || IsNil(o.ClientRole) {
		var ret bool
		return ret
	}
	return *o.ClientRole
}

// GetClientRoleOk returns a tuple with the ClientRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleRepresentation) GetClientRoleOk() (*bool, bool) {
	if o == nil || IsNil(o.ClientRole) {
		return nil, false
	}
	return o.ClientRole, true
}

// HasClientRole returns a boolean if a field has been set.
func (o *RoleRepresentation) HasClientRole() bool {
	if o != nil && !IsNil(o.ClientRole) {
		return true
	}

	return false
}

// SetClientRole gets a reference to the given bool and assigns it to the ClientRole field.
func (o *RoleRepresentation) SetClientRole(v bool) {
	o.ClientRole = &v
}

// GetContainerId returns the ContainerId field value if set, zero value otherwise.
func (o *RoleRepresentation) GetContainerId() string {
	if o == nil || IsNil(o.ContainerId) {
		var ret string
		return ret
	}
	return *o.ContainerId
}

// GetContainerIdOk returns a tuple with the ContainerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleRepresentation) GetContainerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerId) {
		return nil, false
	}
	return o.ContainerId, true
}

// HasContainerId returns a boolean if a field has been set.
func (o *RoleRepresentation) HasContainerId() bool {
	if o != nil && !IsNil(o.ContainerId) {
		return true
	}

	return false
}

// SetContainerId gets a reference to the given string and assigns it to the ContainerId field.
func (o *RoleRepresentation) SetContainerId(v string) {
	o.ContainerId = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *RoleRepresentation) GetAttributes() map[string][]string {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string][]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleRepresentation) GetAttributesOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *RoleRepresentation) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string][]string and assigns it to the Attributes field.
func (o *RoleRepresentation) SetAttributes(v map[string][]string) {
	o.Attributes = &v
}

func (o RoleRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ScopeParamRequired) {
		toSerialize["scopeParamRequired"] = o.ScopeParamRequired
	}
	if !IsNil(o.Composite) {
		toSerialize["composite"] = o.Composite
	}
	if !IsNil(o.Composites) {
		toSerialize["composites"] = o.Composites
	}
	if !IsNil(o.ClientRole) {
		toSerialize["clientRole"] = o.ClientRole
	}
	if !IsNil(o.ContainerId) {
		toSerialize["containerId"] = o.ContainerId
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableRoleRepresentation struct {
	value *RoleRepresentation
	isSet bool
}

func (v NullableRoleRepresentation) Get() *RoleRepresentation {
	return v.value
}

func (v *NullableRoleRepresentation) Set(val *RoleRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleRepresentation(val *RoleRepresentation) *NullableRoleRepresentation {
	return &NullableRoleRepresentation{value: val, isSet: true}
}

func (v NullableRoleRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
